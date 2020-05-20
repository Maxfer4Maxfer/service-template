package calculator

import (
	"context"
	"sync"
	"time"

	"github.com/maxfer4maxfer/service-template/internal/platform/errors"
	"github.com/rs/zerolog"
	"golang.org/x/time/rate"
)

var (
	errRateLimitExceeded = errors.New(
		"rate limit exceeded too many request to method")
)

// RateLimiterConfig holds configuration for a RateLimiter.
// Description https://godoc.org/golang.org/x/time/rate#Limit
type RateLimiterConfig struct {
	Interval time.Duration //rate.Every(<<Interval>>, <<B>>)
	B        int
}

type ratelimiterMiddleware struct {
	logger     *zerolog.Logger
	calculator Calculator
	// rate limiters for each method
	rls map[string]*rate.Limiter
	mu  *sync.RWMutex
	r   rate.Limit
	b   int
}

// NewRatelimiterMiddleware returns new ratelimiter wrapper for a component.
func NewRatelimiterMiddleware(
	logger *zerolog.Logger, cfg RateLimiterConfig, calculator Calculator,
) Calculator {
	return &ratelimiterMiddleware{
		logger:     logger,
		calculator: calculator,
		rls:        make(map[string]*rate.Limiter),
		mu:         &sync.RWMutex{},
		r:          rate.Every(cfg.Interval),
		b:          cfg.B,
	}
}

// limiter returns a ratelimiter for given name
// if a ratelimiter does not exist then it will be created
// with values provided in before.
func (rl *ratelimiterMiddleware) limiter(name string) *rate.Limiter {
	rl.mu.RLock()
	limiter, ok := rl.rls[name]
	rl.mu.RUnlock()

	if !ok {
		limiter = rate.NewLimiter(rl.r, rl.b)

		rl.mu.Lock()
		rl.rls[name] = limiter
		rl.mu.Unlock()
	}

	return limiter
}

// applyLimit apply a ratelimit for a given name.
func (rl *ratelimiterMiddleware) applyLimit(name string) error {
	limiter := rl.limiter(name)
	if !limiter.Allow() {
		return errRateLimitExceeded.Method(name)
	}

	return nil
}

// -------------- Component's methods ------------------

func (rl *ratelimiterMiddleware) Add(
	ctx context.Context, a int, b int,
) (
	sum int, err error,
) {
	err = rl.applyLimit("Add")
	if err != nil {
		return
	}

	return rl.calculator.Add(ctx, a, b)
}

func (rl *ratelimiterMiddleware) Subtract(
	ctx context.Context, a int, b int,
) (
	sub int, err error,
) {
	err = rl.applyLimit("Subtract")
	if err != nil {
		return
	}

	return rl.calculator.Subtract(ctx, a, b)
}

func (rl *ratelimiterMiddleware) Multiply(
	ctx context.Context, a int, b int,
) (
	mult int, err error,
) {
	err = rl.applyLimit("Multiply")
	if err != nil {
		return
	}

	return rl.calculator.Multiply(ctx, a, b)
}

func (rl *ratelimiterMiddleware) Pi(
	ctx context.Context, count int,
) (
	pi string, err error,
) {
	err = rl.applyLimit("Pi")
	if err != nil {
		return
	}

	return rl.calculator.Pi(ctx, count)
}
