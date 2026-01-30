package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/medazizammari/real-time-deadline-aware-golang/internal/domain"
	"github.com/medazizammari/real-time-deadline-aware-golang/internal/pipeline"
	"github.com/medazizammari/real-time-deadline-aware-golang/internal/transport"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// NEW: split channels
	in := make(chan domain.Event, 10)  // raw events (generator + POST /event)
	out := make(chan domain.Event, 10) // processed events (websocket stream)

	// UPDATED: pipeline now connects in -> processing -> out
	pipeline.Start(in, out)

	mux := http.NewServeMux()
	mux.Handle("/ws", transport.Handler(out))              // websocket reads from out
	mux.Handle("/event", transport.CreateEventHandler(in)) // button injects into in

	srv := &http.Server{
		Addr:    ":8080",
		Handler: withCORS(mux),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()

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
