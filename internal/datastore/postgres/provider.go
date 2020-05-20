package postgres

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/maxfer4maxfer/service-template/internal/platform/errors"
	"github.com/rs/zerolog"
)

var (
	errConnectToDatabase = errors.New("cannot connect to the database")
)

// Config holds configuration for a database.
type Config struct {
	DSN string
	// https://godoc.org/database/sql#DB.SetConnMaxLifetime
	ConnMaxLifetime time.Duration
	// https://godoc.org/database/sql#DB.SetMaxOpenConns
	MaxOpenConns int
	// https://godoc.org/database/sql#DB.SetMaxIdleConns
	MaxIdleConns int
}

// New establishes connection to a database.
func New(logger *zerolog.Logger, cfg Config) (*sqlx.DB, error) {
	dsn := cfg.DSN

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, errConnectToDatabase.Wrap(err)
	}

	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)

	return db, nil
}
