package pipeline

import (
	"context"
	"net/http"

	"github.com/medazizammari/real-time-deadline-aware-golang/internal/domain"
	"github.com/medazizammari/real-time-deadline-aware-golang/internal/generator"
	"github.com/medazizammari/real-time-deadline-aware-golang/internal/processor"
	"github.com/medazizammari/real-time-deadline-aware-golang/internal/storage"
)

type Pipeline struct {
	Trigger chan struct{}
	Out     chan domain.Event
	repo    *storage.EventsRepo
}

func Start(repo *storage.EventsRepo) *Pipeline {
	trigger := make(chan struct{}, 10)
	in := make(chan domain.Event, 10)
	out := make(chan domain.Event, 10)

	go generator.Start(trigger, in)

	go func() {
		for ev := range in {
			processed := processor.Process(ev)

			// Persist (best-effort; you can log errors)
			_ = repo.Insert(context.Background(), processed)

			out <- processed
		}
	}()

	return &Pipeline{Trigger: trigger, Out: out, repo: repo}
}

// Your manual trigger handler (simple)
func (p *Pipeline) TriggerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	select {
	case p.Trigger <- struct{}{}:
		w.WriteHeader(http.StatusAccepted)
	default:
		w.WriteHeader(http.StatusTooManyRequests)
	}
}
