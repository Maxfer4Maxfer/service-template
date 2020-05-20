package storekeeper

import (
	"context"
	"time"

	"github.com/maxfer4maxfer/service-template/internal/domain"
	"github.com/maxfer4maxfer/service-template/internal/platform/correlationid"
	"github.com/maxfer4maxfer/service-template/internal/platform/errors"

	"github.com/rs/zerolog"
)

var (
	errRetrieveLogs = errors.New("cannot retrieve logs records")
	errStoreLog     = errors.New("cannot store a log")
	errDeleteLog    = errors.New("cannot delete a log")
)

// Storage declares ways how Storekeeper interact with the storage subsystem.
type Storage interface {
	Complete(context.Context, *error)
	GetLogRecords(context.Context) (domain.Log, error)
	SaveToWarehouse(context.Context, domain.LogRecord) error
	DeteteLog(context.Context, domain.LogRecordID) error
}

// StorageFactory start close transaction and return Storage instance.
type StorageFactory interface {
	ForSingleOperation() Storage
	ForComplexOperation() Storage
}

// Storekeeper goes trought all log recored and save them to a separate table.
type Storekeeper struct {
	logger       *zerolog.Logger
	timeInterval time.Duration
	sf           StorageFactory
	ctx          context.Context
	cancel       context.CancelFunc
	stopChan     chan struct{}
}

// New returns the new instance of Storekeeper.
func New(
	logger *zerolog.Logger,
	timeInterval time.Duration,
	sf StorageFactory,
) *Storekeeper {
	locallogger := logger.With().Str("component", "storekeeper").Logger()

	ctxWithCancel, cancel := context.WithCancel(context.Background())

	return &Storekeeper{
		logger:       &locallogger,
		timeInterval: timeInterval,
		sf:           sf,
		ctx:          ctxWithCancel,
		cancel:       cancel,
		stopChan:     make(chan struct{}),
	}
}

// Start starts Storekeeper.
func (sk *Storekeeper) Start() {
	sk.logger.Info().Dur("period", sk.timeInterval).Msg("starting")

	ticker := time.NewTicker(sk.timeInterval)

	go func() {
		for {
			err := sk.keeplogs()
			if err != nil {
				sk.logger.Error().Err(err).
					Msg("failure to call the gatherMiningSchedulersState method")
			}

			select {
			case <-sk.stopChan:
				ticker.Stop()
				return
			case <-ticker.C:
				continue
			}
		}
	}()
}

// Shutdown stops Storekeeper.
func (sk *Storekeeper) Shutdown() {
	sk.stopChan <- struct{}{}

	sk.logger.Info().Msg("stopped")
}

func (sk *Storekeeper) keeplogs() (err error) {
	ctx, cID := correlationid.Assign(context.Background())
	logger := sk.logger.With().
		Str("correlationID", cID).Str("method", "keeplogs").Logger()

	stg := sk.sf.ForComplexOperation()
	defer stg.Complete(ctx, &err)

	logs, err := stg.GetLogRecords(ctx)
	if err != nil {
		return errRetrieveLogs.Wrap(err)
	}

	logger.Debug().Int("len(logs)", len(logs)).Msg("")

	for _, l := range logs {
		err = stg.SaveToWarehouse(ctx, l)
		if err != nil {
			return errStoreLog.Wrap(err).Scope("log", l)
		}

		err = stg.DeteteLog(ctx, l.ID)
		if err != nil {
			return errDeleteLog.Wrap(err).Scope("log", l)
		}
	}

	return nil
}
