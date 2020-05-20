package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/maxfer4maxfer/service-template/internal/component/storekeeper"
	"github.com/rs/zerolog"
)

// StorekeeperStorageFactory for the Storekeeper component.
type StorekeeperStorageFactory struct {
	logger *zerolog.Logger
	db     *sqlx.DB
}

// NewStorekeeperStorageFactory returns the new StorekeeperStorageFactory instance.
func NewStorekeeperStorageFactory(
	logger *zerolog.Logger, db *sqlx.DB,
) *StorekeeperStorageFactory {
	return &StorekeeperStorageFactory{
		logger: logger,
		db:     db,
	}
}

// ForSingleOperation returns a storage instance
// for performing single operations.
func (csf *StorekeeperStorageFactory) ForSingleOperation() storekeeper.Storage {
	return &Storage{
		Logger: csf.logger,
		DB:     csf.db,
	}
}

// ForComplexOperation returns a storage instance
// for performing operations in the one transaction.
func (csf *StorekeeperStorageFactory) ForComplexOperation() storekeeper.Storage {
	tx := csf.db.MustBegin()

	return &Storage{
		Logger: csf.logger,
		DB:     tx,
	}
}
