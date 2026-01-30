package pipeline

import (
	"github.com/medazizammari/real-time-deadline-aware-golang/internal/domain"
	"github.com/medazizammari/real-time-deadline-aware-golang/internal/generator"
	"github.com/medazizammari/real-time-deadline-aware-golang/internal/processor"
)

func Start(in chan domain.Event, out chan<- domain.Event) {
	go generator.Start(in) // generator writes into in

	go func() {
		for ev := range in { // pipeline reads from in
			processed := processor.Process(ev)
			out <- processed
		}
	}()
}
