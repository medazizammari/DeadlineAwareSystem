package main

import (
	"context"
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

	events := make(chan domain.Event, 10)

	pipeline.Start(events)

	http.Handle("/ws", transport.Handler(events))

	go http.ListenAndServe(":8080", nil)

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_ = shutdownCtx
}
