package calculator

import (
	"context"
	"time"

	"github.com/rs/zerolog"

	"github.com/maxfer4maxfer/service-template/internal/platform/correlationid"
	"github.com/maxfer4maxfer/service-template/internal/platform/errors"
)

var (
	errCorrelationIDEmpty = errors.New("ctx doesn't have correlationID")
)

type loggingMiddleware struct {
	calculator Calculator
	logger     *zerolog.Logger
	debug      bool
}

// NewLoggingMiddleware returns new metrics wrapper for a component.
func NewLoggingMiddleware(
	logger *zerolog.Logger, calculator Calculator, debug bool,
) Calculator {
	return &loggingMiddleware{
		calculator: calculator,
		logger:     logger,
		debug:      debug,
	}
}

func (ll *loggingMiddleware) before(
	ctx context.Context, method string,
) (
	context.Context, time.Time, string,
) {
	start := time.Now()

	cID := correlationid.Extract(ctx)
	if cID == "" {
		ll.logger.Warn().Err(errCorrelationIDEmpty).Msg("")

		ctx, cID = correlationid.Assign(ctx)
	}

	ll.logger.Info().
		Str("correlationID", cID).
		Str("method", method).
		Time("start", start).
		Msg("operation started")

	return ctx, start, cID
}

func (ll *loggingMiddleware) after(method string, start time.Time, rID string) {
	ll.logger.Info().
		Str("correlationID", rID).
		Str("method", method).
		Str("duration", time.Since(start).String()).
		Msg("operation completed")
}

// -------------- Component's methods ------------------

func (ll *loggingMiddleware) Add(
	ctx context.Context, a int, b int,
) (
	sum int, err error,
) {
	method := "Add"

	ctx, start, rID := ll.before(ctx, method)
	defer ll.after(method, start, rID)

	if ll.debug {
		ll.logger.Debug().
			Str("correlationID", rID).
			Str("method", method).
			Int("a", a).
			Int("b", b).
			Msg("incoming parameters")
	}

	sum, err = ll.calculator.Add(ctx, a, b)

	if ll.debug {
		ll.logger.Debug().
			Str("correlationID", rID).
			Str("method", method).
			Int("sum", sum).
			Interface("err", err).
			Msg("outcoming parameters")
	}

	return sum, err
}

func (ll *loggingMiddleware) Subtract(
	ctx context.Context, a int, b int,
) (
	sub int, err error,
) {
	method := "Subtract"

	ctx, start, rID := ll.before(ctx, method)
	defer ll.after(method, start, rID)

	if ll.debug {
		ll.logger.Debug().
			Str("correlationID", rID).
			Str("method", method).
			Int("a", a).
			Int("b", b).
			Msg("incoming parameters")
	}

	sub, err = ll.calculator.Subtract(ctx, a, b)

	if ll.debug {
		ll.logger.Debug().
			Str("correlationID", rID).
			Str("method", method).
			Int("sub", sub).
			Interface("err", err).
			Msg("outcoming parameters")
	}

	return sub, err
}

func (ll *loggingMiddleware) Multiply(
	ctx context.Context, a int, b int,
) (
	mult int, err error,
) {
	method := "Multiply"

	ctx, start, rID := ll.before(ctx, method)
	defer ll.after(method, start, rID)

	if ll.debug {
		ll.logger.Debug().
			Str("correlationID", rID).
			Str("method", method).
			Int("a", a).
			Int("b", b).
			Msg("incoming parameters")
	}

	mult, err = ll.calculator.Multiply(ctx, a, b)

	if ll.debug {
		ll.logger.Debug().
			Str("correlationID", rID).
			Str("method", method).
			Int("mult", mult).
			Interface("err", err).
			Msg("outcoming parameters")
	}

	return mult, err
}

func (ll *loggingMiddleware) Pi(
	ctx context.Context, count int,
) (
	pi string, err error,
) {
	method := "Pi"

	ctx, start, rID := ll.before(ctx, method)
	defer ll.after(method, start, rID)

	if ll.debug {
		ll.logger.Debug().
			Str("correlationID", rID).
			Str("method", method).
			Int("count", count).
			Msg("incoming parameters")
	}

	pi, err = ll.calculator.Pi(ctx, count)

	if ll.debug {
		ll.logger.Debug().
			Str("correlationID", rID).
			Str("method", method).
			Str("pi", pi).
			Interface("err", err).
			Msg("outcoming parameters")
	}

	return pi, err
}
