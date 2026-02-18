package generator

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/medazizammari/real-time-deadline-aware-golang/internal/domain"
)

func Start(trigger <-chan struct{}, out chan<- domain.Event) {
	log.Println("GENERATOR STARTED")
	for range trigger {

		event := domain.Event{
			ID:         uuid.New().String(),
			CreatedAt:  time.Now(),
			DeadlineMs: 100 * time.Millisecond,
		}

		select {
		case out <- event:
			log.Println("GEN SENT", event.ID)
		default:
			log.Println("GEN DROPPED", event.ID)
		}
	}
}
