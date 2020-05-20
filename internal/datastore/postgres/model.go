package postgres

import (
	"database/sql"

	"github.com/lib/pq"

	"github.com/maxfer4maxfer/service-template/internal/domain"
)

// Log is a collection of log records.
type Log []LogRecord

// Convert converts Log -> []domain.Log.
func (ls *Log) Convert() domain.Log {
	result := make([]domain.LogRecord, len(*ls))
	for i, l := range *ls {
		result[i] = l.Convert()
	}

	return result
}

// LogRecord represent table 'log'.
type LogRecord struct {
	ID            int64
	OperationType sql.NullInt64  `db:"operation_type"`
	CreatedAt     pq.NullTime    `db:"created_at"`
	LogText       sql.NullString `db:"log_text"`
}

// Convert converts LogRecord -> domain.LogRecord.
func (lr *LogRecord) Convert() domain.LogRecord {
	return domain.LogRecord{
		ID:        domain.LogRecordID(lr.ID),
		Operation: domain.OperationType(int(lr.OperationType.Int64)),
		CreatedAt: lr.CreatedAt.Time,
		Text:      lr.LogText.String,
	}
}
