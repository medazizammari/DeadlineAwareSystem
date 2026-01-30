package transport

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/medazizammari/real-time-deadline-aware-golang/internal/domain"
)

func CreateEventHandler(in chan<- domain.Event) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		ev := domain.Event{
			ID:         uuid.New().String(),
			CreatedAt:  time.Now(),
			DeadlineMs: 100 * time.Millisecond,
		}

		// non-blocking enqueue (backpressure visible)
		select {
		case in <- ev:
			w.WriteHeader(http.StatusAccepted)
		default:
			http.Error(w, "input queue full", http.StatusTooManyRequests)
		}
	}
}
