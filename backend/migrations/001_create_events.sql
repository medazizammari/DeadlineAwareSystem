CREATE TABLE IF NOT EXISTS events (
  id           TEXT PRIMARY KEY,
  created_at   TIMESTAMPTZ NOT NULL,
  deadline_ms  BIGINT NOT NULL,
  processed_at TIMESTAMPTZ NOT NULL,
  status       TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_events_processed_at ON events (processed_at DESC);
