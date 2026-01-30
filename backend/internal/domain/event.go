package domain

import "time"

type Event struct {
	ID          string        `json:"id"`
	CreatedAt   time.Time     `json:"createdAt"`
	DeadlineMs  time.Duration `json:"deadlineMs"`
	ProcessedAt time.Time     `json:"processedAt,omitempty"`
	Status      string        `json:"status"`
}
