package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/maxfer4maxfer/service-template/internal/component/calculator"
	"github.com/rs/zerolog"
)

// CalculatorStorageFactory for the 'calculator' component.
type CalculatorStorageFactory struct {
	logger *zerolog.Logger
	db     *sqlx.DB
}

// NewCalculatorStorageFactory returns the new CalculatorStorageFactory instance.
func NewCalculatorStorageFactory(
	logger *zerolog.Logger, db *sqlx.DB) *CalculatorStorageFactory {
	return &CalculatorStorageFactory{
		logger: logger,
		db:     db,
	}
}

// ForSingleOperation returns a storage instance
// for performing single operations.
func (sf *CalculatorStorageFactory) ForSingleOperation() calculator.Storage {
	return NewStorage(sf.logger, sf.db)
}

// ForComplexOperation returns a storage instance
// for performing operations in the one transaction.
func (sf *CalculatorStorageFactory) ForComplexOperation() calculator.Storage {
	tx := sf.db.MustBegin()

	return NewStorage(sf.logger, tx)
}
