package calculator

import (
	"context"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/sony/gobreaker"
)

// CircuitBreakerConfig holds configuration for a CircuitBreaker.
// Description https://godoc.org/github.com/sony/gobreaker#Settings
type CircuitBreakerConfig struct {
	MaxRequests int
	Interval    time.Duration
	Timeout     time.Duration
}

type circuitbreakerMiddleware struct {
	logger     *zerolog.Logger
	calculator Calculator
	// cbs: CircuitBreakers for each method
	cbs      map[string]*gobreaker.CircuitBreaker
	mu       *sync.RWMutex
	settings gobreaker.Settings
}

// NewCircuitBreakerMiddleware returns new CircuitBreaker wrapper for Service.
func NewCircuitBreakerMiddleware(
	logger *zerolog.Logger, cfg CircuitBreakerConfig, calculator Calculator,
) Calculator {
	settings := gobreaker.Settings{
		MaxRequests: uint32(cfg.MaxRequests),
		Interval:    cfg.Interval,
		Timeout:     cfg.Timeout,
	}

	settings.OnStateChange = func(
		name string, from gobreaker.State, to gobreaker.State,
	) {
		logger.Warn().
			Str("name", name).
			Str("from", from.String()).
			Str("to", to.String()).
			Msg("cirtuit breaker has changed state")
	}

	return &circuitbreakerMiddleware{
		logger:     logger,
		calculator: calculator,
		cbs:        make(map[string]*gobreaker.CircuitBreaker), // CircuitBreakers
		mu:         &sync.RWMutex{},
		settings:   settings,
	}
}

// circuitBreaker returns a circuit breaker for given name
// if a circuit breaker does not exist then it will be created
// with values provided in before.
func (cbm *circuitbreakerMiddleware) circuitBreaker(
	name string,
) *gobreaker.CircuitBreaker {
	cbm.mu.RLock()
	cb, ok := cbm.cbs[name]
	cbm.mu.RUnlock()

	if !ok {
		settings := cbm.settings
		settings.Name = name
		cb = gobreaker.NewCircuitBreaker(cbm.settings)

		cbm.mu.Lock()
		cbm.cbs[name] = cb
		cbm.mu.Unlock()
	}

	return cb
}

// -------------- Component's methods ------------------

func (cbm *circuitbreakerMiddleware) Add(
	ctx context.Context, a int, b int,
) (
	sum int, err error,
) {
	cb := cbm.circuitBreaker("Add")

	response, err := cb.Execute(func() (interface{}, error) {
		return cbm.calculator.Add(ctx, a, b)
	})
	if err != nil {
		return sum, err
	}

	sum = response.(int)

	return sum, err
}

func (cbm *circuitbreakerMiddleware) Subtract(
	ctx context.Context, a int, b int,
) (
	sub int, err error,
) {
	cb := cbm.circuitBreaker("Subtract")

	response, err := cb.Execute(func() (interface{}, error) {
		return cbm.calculator.Subtract(ctx, a, b)
	})
	if err != nil {
		return sub, err
	}

	sub = response.(int)

	return sub, err
}

func (cbm *circuitbreakerMiddleware) Multiply(
	ctx context.Context, a int, b int,
) (
	mult int, err error,
) {
	cb := cbm.circuitBreaker("Multiply")

	response, err := cb.Execute(func() (interface{}, error) {
		return cbm.calculator.Multiply(ctx, a, b)
	})
	if err != nil {
		return mult, err
	}

	mult = response.(int)

	return mult, err
}

func (cbm *circuitbreakerMiddleware) Pi(
	ctx context.Context, count int,
) (
	pi string, err error,
) {
	cb := cbm.circuitBreaker("Multiply")

	response, err := cb.Execute(func() (interface{}, error) {
		return cbm.calculator.Pi(ctx, count)
	})
	if err != nil {
		return pi, err
	}

	pi = response.(string)

	return pi, err
}
