package calculator

import (
	"context"
	"fmt"
	"time"

	"github.com/maxfer4maxfer/service-template/internal/domain"
	"github.com/maxfer4maxfer/service-template/internal/platform/errors"
	"github.com/rs/zerolog"
)

var (
	errPiAccuracyNegative = errors.New(
		"number accuracy of Pi should be more then zero")
	errReceivePiProblem = errors.New("cannot receive a pi number")
	errStorageProblem   = errors.New("cannot interact with the storage subsystem")
	errStoreLog         = errors.New("cannot store a log")
)

// Config holds configuration for the Calculator component.
type Config struct {
	RateLimiter    RateLimiterConfig
	CircuitBreaker CircuitBreakerConfig
}

// Calculator declares interactions with the calculator.
type Calculator interface {
	Add(ctx context.Context, a int, b int) (sum int, err error)
	Subtract(ctx context.Context, a int, b int) (sub int, err error)
	Multiply(ctx context.Context, a int, b int) (mult int, err error)
	Pi(ctx context.Context, count int) (pi string, err error)
}

// Storage declares methods for interactiong with storage subsystems.
type Storage interface {
	Complete(context.Context, *error)
	SaveLog(context.Context, domain.LogRecord) error
}

// StorageFactory start close transaction and return Storage instance.
type StorageFactory interface {
	ForSingleOperation() Storage
	ForComplexOperation() Storage
}

// PiGetter provides pi number.
type PiGetter interface {
	Pi(ctx context.Context, count int) (string, error)
}

// calcService is the main service of that microservice.
type calculator struct {
	logger *zerolog.Logger
	sf     StorageFactory
	pi     PiGetter
}

// New returns new instance of service.
func New(
	logger *zerolog.Logger,
	sf StorageFactory,
	pi PiGetter,
	cfg Config,
) Calculator {
	var srv Calculator

	srv = &calculator{
		logger: logger,
		sf:     sf,
		pi:     pi,
	}

	srv = NewCircuitBreakerMiddleware(logger, cfg.CircuitBreaker, srv)
	srv = NewLoggingMiddleware(logger, srv, true)
	srv = NewRatelimiterMiddleware(logger, cfg.RateLimiter, srv)
	srv = NewMetricsMiddleware(srv)

	return srv
}

// Add adds to number together.
func (cal *calculator) Add(
	ctx context.Context, a int, b int,
) (
	sum int, err error,
) {
	stg := cal.sf.ForSingleOperation()

	sum = a + b

	l := domain.LogRecord{
		Operation: domain.OperationAdd,
		CreatedAt: time.Now(),
		Text:      fmt.Sprintf("%d + %d = %d", a, b, sum),
	}

	err = stg.SaveLog(ctx, l)
	if err != nil {
		return sum, errStorageProblem.Wrap(err)
	}

	return sum, nil
}

// Subtract subtracts one number for other.
func (cal *calculator) Subtract(
	ctx context.Context, a int, b int,
) (
	sub int, err error,
) {
	stg := cal.sf.ForSingleOperation()

	sub = a - b

	lr := domain.LogRecord{
		Operation: domain.OperationSub,
		CreatedAt: time.Now(),
		Text:      fmt.Sprintf("%d - %d = %d", a, b, sub),
	}

	err = stg.SaveLog(ctx, lr)
	if err != nil {
		return sub, errStoreLog.Wrap(err)
	}

	return sub, nil
}

// Multiply multiplies too numbers.
func (cal *calculator) Multiply(
	ctx context.Context, a int, b int,
) (
	mult int, err error,
) {
	stg := cal.sf.ForSingleOperation()

	mult = a * b

	lr := domain.LogRecord{
		Operation: domain.OperationMultiply,
		CreatedAt: time.Now(),
		Text:      fmt.Sprintf("%d * %d = %d", a, b, mult),
	}

	err = stg.SaveLog(ctx, lr)
	if err != nil {
		return mult, errStoreLog.Wrap(err)
	}

	return mult, nil
}

// Pi returns pi number with a given length.
func (cal *calculator) Pi(ctx context.Context, count int) (pi string, err error) {
	if count <= 0 {
		return pi, errPiAccuracyNegative.Scope("count", count)
	}

	stg := cal.sf.ForComplexOperation()
	defer stg.Complete(ctx, &err)

	lr := domain.LogRecord{
		Operation: domain.OperationPi,
		CreatedAt: time.Now(),
		Text:      fmt.Sprintf("Start request Pi(%d)", count),
	}

	err = stg.SaveLog(ctx, lr)
	if err != nil {
		return pi, errStoreLog.Wrap(err)
	}

	pi, err = cal.pi.Pi(ctx, count)
	if err != nil {
		return pi, errReceivePiProblem.Wrap(err)
	}

	lr = domain.LogRecord{
		Operation: domain.OperationPi,
		CreatedAt: time.Now(),
		Text:      fmt.Sprintf("Finish request Pi(%d) = %s", count, pi),
	}

	err = stg.SaveLog(ctx, lr)
	if err != nil {
		return pi, errStoreLog.Wrap(err)
	}

	return pi, nil
}
