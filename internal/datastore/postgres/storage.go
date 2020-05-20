package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/maxfer4maxfer/service-template/internal/domain"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type dbquerier interface {
	sqlx.Ext
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
}

// Storage interacts with postgres.
type Storage struct {
	Logger *zerolog.Logger
	DB     dbquerier
}

// NewStorage returns a storage instance that can be used for interaction with
// the storage subsystem.
func NewStorage(logger *zerolog.Logger, db dbquerier) *Storage {
	return &Storage{
		Logger: logger,
		DB:     db,
	}
}

// Complete extracts a transaction from the provided storage and perform Commit.
func (s *Storage) Complete(ctx context.Context, err *error) {
	tx, ok := s.DB.(interface {
		Commit() error
		Rollback() error
	})

	// we are not in a transaction
	if !ok {
		return
	}

	// something went wrong. Rollback
	if *err != nil {
		txErr := tx.Rollback()
		if txErr != nil {
			*err = errors.Wrap(*err, "Storage.Complete something went wrong")
			*err = errors.Wrap(txErr, "Storage.Complete cannon rollback the transaction")
		}

		return
	}

	// there is no previous error was occurred. Can commit
	txErr := tx.Commit()
	if txErr != nil {
		*err = errors.Wrap(txErr, "Storage.Complete cannot commit the transaction ")
	}
}

// GetLogRecords return all records from the log table.
func (s *Storage) GetLogRecords(ctx context.Context) (domain.Log, error) {
	q := `SELECT * FROM log`

	log := &Log{}
	err := s.DB.Select(log, q)

	return log.Convert(), err
}

// SaveLog saves news in 'log' table.
func (s *Storage) SaveLog(ctx context.Context, l domain.LogRecord) error {
	q := `INSERT INTO log (
	operation_type,
	created_at, 
	log_text
	) VALUES ($1, $2, $3)
	RETURNING id`

	_, err := s.DB.Exec(q,
		l.Operation,
		l.CreatedAt,
		l.Text,
	)

	return err
}

// SaveToWarehouse saves news in 'news_for_suppliers' table.
func (s *Storage) SaveToWarehouse(ctx context.Context, l domain.LogRecord) error {
	q := `INSERT INTO log_warehouse (
	operation_type,
	created_at, 
	log_text
	) VALUES ($1, $2, $3)
	RETURNING id`

	_, err := s.DB.Exec(q,
		l.Operation,
		l.CreatedAt,
		l.Text,
	)

	return err
}

// DeteteLog returns all unpublished news.
func (s *Storage) DeteteLog(ctx context.Context, id domain.LogRecordID) error {
	q := `DELETE FROM log WHERE id = $1`

	_, err := s.DB.Exec(q, id)

	return err
}
