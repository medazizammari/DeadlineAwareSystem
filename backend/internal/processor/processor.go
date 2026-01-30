package processor

import (
	"context"
	"time"

	"github.com/medazizammari/real-time-deadline-aware-golang/internal/domain"
)

func Process(event domain.Event) domain.Event {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		event.DeadlineMs,
	)
	defer cancel()

	done := make(chan struct{})

	go func() {
		time.Sleep(50 * time.Millisecond) // simulated work
		close(done)
	}()

	select {
	case <-done:
		event.Status = "on-time"
	case <-ctx.Done():
		event.Status = "late"
	}

	event.ProcessedAt = time.Now()
	return event
}
