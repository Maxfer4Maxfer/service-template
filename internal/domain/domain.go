package domain

import (
	"time"

	"github.com/maxfer4maxfer/service-template/internal/platform/errors"
)

var (
	// errEmptyLog throws when a log has 0 records in it
	errEmptyLog = errors.New("the log has 0 records")

	// errEmptyLogRecord throws when a log record is empty
	errEmptyLogRecord = errors.New("the log record is empty")
)

// Log defines a log.
type Log []LogRecord

// Validate performs validation of ErrorFilter.
func (ls *Log) Validate() error {
	if len(*ls) == 0 {
		return errEmptyLog
	}

	for _, l := range *ls {
		err := l.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}

// LogRecordID is an ID of a log record.
type LogRecordID int

// OperationType represents a type of the operation.
type OperationType int

const (
	// OperationAdd - Add
	OperationAdd OperationType = 1
	// OperationSub - Subtract
	OperationSub OperationType = 2
	// OperationMultiply - Multiply
	OperationMultiply OperationType = 3
	// OperationPi - Pi
	OperationPi OperationType = 4
)

// LogRecord defines a log record.
type LogRecord struct {
	ID        LogRecordID
	Operation OperationType
	CreatedAt time.Time
	Text      string
}

// Validate performs validation of ErrorFilter.
func (lr *LogRecord) Validate() error {
	if lr.Text == "" {
		return errEmptyLogRecord.Scope("id", lr.ID)
	}

	return nil
}
