package pipeline

import (
	"github.com/medazizammari/real-time-deadline-aware-golang/internal/domain"
	"github.com/medazizammari/real-time-deadline-aware-golang/internal/generator"
	"github.com/medazizammari/real-time-deadline-aware-golang/internal/processor"
)

func Start(out chan<- domain.Event) {
	in := make(chan domain.Event, 10) // bounded buffer

	go generator.Start(in)

	go func() {
		for event := range in {
			processed := processor.Process(event)
			out <- processed
		}
	}()
}
