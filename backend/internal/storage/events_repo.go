package storage

import (
	"context"
	"time"

	"github.com/medazizammari/real-time-deadline-aware-golang/internal/domain"
)

type EventsRepo struct {
	pg *Postgres
}

func NewEventsRepo(pg *Postgres) *EventsRepo {
	return &EventsRepo{pg: pg}
}

func (r *EventsRepo) Insert(ctx context.Context, ev domain.Event) error {
	// Store deadline in ms to keep it simple
	deadlineMs := ev.DeadlineMs.Milliseconds()

	_, err := r.pg.DB.ExecContext(ctx, `
		INSERT INTO events (id, created_at, deadline_ms, processed_at, status)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (id) DO NOTHING
	`,
		ev.ID,
		ev.CreatedAt,
		deadlineMs,
		ev.ProcessedAt,
		ev.Status,
	)
	return err
}

func (r *EventsRepo) ListLatest(ctx context.Context, limit int) ([]domain.Event, error) {
	if limit <= 0 || limit > 500 {
		limit = 50
	}

	rows, err := r.pg.DB.QueryContext(ctx, `
		SELECT id, created_at, deadline_ms, processed_at, status
		FROM events
		ORDER BY processed_at DESC
		LIMIT $1
	`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]domain.Event, 0, limit)
	for rows.Next() {
		var (
			id         string
			createdAt  time.Time
			deadlineMs int64
			processed  time.Time
			status     string
		)

		if err := rows.Scan(&id, &createdAt, &deadlineMs, &processed, &status); err != nil {
			return nil, err
		}

		out = append(out, domain.Event{
			ID:          id,
			CreatedAt:   createdAt,
			DeadlineMs:  time.Duration(deadlineMs) * time.Millisecond,
			ProcessedAt: processed,
			Status:      status,
		})
	}

	return out, rows.Err()
}
