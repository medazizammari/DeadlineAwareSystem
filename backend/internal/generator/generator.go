package generator

import (
	"time"

	"github.com/google/uuid"
	"github.com/medazizammari/real-time-deadline-aware-golang/internal/domain"
)

func Start(out chan<- domain.Event) {
	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		event := domain.Event{
			ID:         uuid.New().String(),
			CreatedAt:  time.Now(),
			DeadlineMs: 100 * time.Millisecond,
		}

		select {
		case out <- event:
		default:
			// drop event → back-pressure by design
		}
	}
}
