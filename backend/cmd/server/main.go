package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/medazizammari/real-time-deadline-aware-golang/internal/pipeline"
	"github.com/medazizammari/real-time-deadline-aware-golang/internal/storage"
	"github.com/medazizammari/real-time-deadline-aware-golang/internal/transport"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is required")
	}

	pg, err := storage.NewPostgres(dsn)
	if err != nil {
		log.Fatalf("db connect failed: %v", err)
	}
	defer pg.Close()

	repo := storage.NewEventsRepo(pg)

	// Start your pipeline in manual-trigger mode (as we did before)
	p := pipeline.Start(repo)

	http.Handle("/ws", transport.Handler(p.Out))

	// Trigger endpoint (you already have)
	http.HandleFunc("/event", p.TriggerHandler)

	// New endpoint: list latest
	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		limit := 50
		if v := r.URL.Query().Get("limit"); v != "" {
			if n, err := strconv.Atoi(v); err == nil {
				limit = n
			}
		}

		events, err := repo.ListLatest(r.Context(), limit)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		transport.WriteJSON(w, events) // helper you can add (see below)
	})

	srv := &http.Server{Addr: ":8080"}

	go func() {
		log.Println("server listening on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen failed: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_ = srv.Shutdown(shutdownCtx)
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Vite dev server
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
