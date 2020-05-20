package readycheck

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

// ReadyCheck is the service to interact with news.
type ReadyCheck struct {
	logger  *zerolog.Logger
	storage *sqlx.DB
}

// New returns new instance of ReadyCheck service.
func New(logger *zerolog.Logger, storage *sqlx.DB) *ReadyCheck {
	srv := &ReadyCheck{
		logger:  logger,
		storage: storage,
	}

	return srv
}

// Ready checks that a database is up.
func (rc *ReadyCheck) Ready(ctx context.Context) error {
	return rc.storage.PingContext(ctx)
}
