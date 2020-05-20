package calculator

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type metricsMiddleware struct {
	calculator Calculator
	addTotal   prometheus.Counter
	subTotal   prometheus.Counter
	multTotal  prometheus.Counter
	piTotal    prometheus.Counter
}

// NewMetricsMiddleware returns new metrics wrapper for a component.
func NewMetricsMiddleware(calculator Calculator) Calculator {
	addTotal := promauto.NewCounter(prometheus.CounterOpts{
		Name: "calculator_add_total",
		Help: "The total number of all added numbers",
	})
	subTotal := promauto.NewCounter(prometheus.CounterOpts{
		Name: "calculator_sub_total",
		Help: "The sum of all sub results",
	})
	multTotal := promauto.NewCounter(prometheus.CounterOpts{
		Name: "calculator_mult_total",
		Help: "The sum number of multiply results",
	})
	piTotal := promauto.NewCounter(prometheus.CounterOpts{
		Name: "calculator_pi_total",
		Help: "The total number all calls of the pi method",
	})

	return &metricsMiddleware{
		calculator: calculator,
		addTotal:   addTotal,
		subTotal:   subTotal,
		multTotal:  multTotal,
		piTotal:    piTotal,
	}
}

// -------------- Component's methods ------------------

func (mm *metricsMiddleware) Add(
	ctx context.Context, a int, b int,
) (
	sum int, err error,
) {
	sum, err = mm.calculator.Add(ctx, a, b)
	if err == nil {
		mm.addTotal.Add(float64(sum))
	}

	return sum, err
}

func (mm *metricsMiddleware) Subtract(
	ctx context.Context, a int, b int,
) (
	sub int, err error,
) {
	sub, err = mm.calculator.Subtract(ctx, a, b)
	if err == nil {
		mm.subTotal.Add(float64(sub))
	}

	return sub, err
}

func (mm *metricsMiddleware) Multiply(
	ctx context.Context, a int, b int,
) (
	mult int, err error,
) {
	mult, err = mm.calculator.Multiply(ctx, a, b)
	if err == nil {
		mm.multTotal.Add(float64(mult))
	}

	return mult, err
}

func (mm *metricsMiddleware) Pi(
	ctx context.Context, count int,
) (
	pi string, err error,
) {
	pi, err = mm.calculator.Pi(ctx, count)
	if err == nil {
		mm.piTotal.Add(1)
	}

	return pi, err
}
