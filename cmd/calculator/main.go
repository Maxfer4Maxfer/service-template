package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/maxfer4maxfer/service-template/internal/component/calculator"
	"github.com/maxfer4maxfer/service-template/internal/component/storekeeper"
	"github.com/maxfer4maxfer/service-template/internal/datastore/postgres"
	"github.com/maxfer4maxfer/service-template/internal/endpoint/pi"
	grpcEntrypoint "github.com/maxfer4maxfer/service-template/internal/entrypoint/grpc"
	httpEntrypoint "github.com/maxfer4maxfer/service-template/internal/entrypoint/http"
	"github.com/maxfer4maxfer/service-template/internal/platform/httpserver"
	"github.com/maxfer4maxfer/service-template/internal/platform/logger"
	"github.com/maxfer4maxfer/service-template/internal/platform/readycheck"
)

func main() {
	var (
		cfg    = newConfig()
		logger = logger.New(
			logger.ConsoleOutput(cfg.Logger.Console),
			logger.DebugMode(cfg.Logger.Debug),
		)
	)

	db, err := postgres.New(logger, postgres.Config{
		DSN:             cfg.Database.DSN,
		ConnMaxLifetime: cfg.Database.ConnMaxLifetime,
		MaxOpenConns:    cfg.Database.MaxOpenConns,
		MaxIdleConns:    cfg.Database.MaxIdleConns,
	})
	if err != nil {
		logger.Error().Err(err).Msg("")

		os.Exit(1)
	}

	var (
		readychecker = readycheck.New(logger, db)

		// --------- Pi Endpoint Client ---------
		piHTTPClient = &http.Client{Timeout: cfg.PiClient.Timeout}
		piClient     = pi.New(logger, cfg.PiClient.Address, piHTTPClient)

		// --------- Calculator ---------
		calculatorSF = postgres.NewCalculatorStorageFactory(logger, db)
		calculator   = calculator.New(
			logger, calculatorSF, piClient, calculator.Config{
				RateLimiter: calculator.RateLimiterConfig{
					Interval: cfg.Calculator.RateLimiter.Interval,
					B:        cfg.Calculator.RateLimiter.B,
				},
				CircuitBreaker: calculator.CircuitBreakerConfig{
					MaxRequests: cfg.Calculator.CircuitBreaker.MaxRequests,
					Interval:    cfg.Calculator.CircuitBreaker.Interval,
					Timeout:     cfg.Calculator.CircuitBreaker.Timeout,
				},
			},
		)

		// --------- Storekeeper scheduled task ---------
		storekeeperSF = postgres.NewStorekeeperStorageFactory(logger, db)
		storekeeper   = storekeeper.New(
			logger, cfg.Storekeeper.TimeInterval, storekeeperSF,
		)

		// --------- Entrypoint ---------
		entrypointHTTP = httpEntrypoint.New(
			logger, calculator,
			httpEntrypoint.Config{
				Address:           cfg.HTTP.Address,
				ReadTimeout:       cfg.HTTP.ReadTimeout,
				ReadHeaderTimeout: cfg.HTTP.ReadHeaderTimeout,
				WriteTimeout:      cfg.HTTP.WriteTimeout,
				IdleTimeout:       cfg.HTTP.IdleTimeout,
				MaxHeaderBytes:    cfg.HTTP.MaxHeaderBytes,
			})

		entrypointGRPC = grpcEntrypoint.New(
			logger, calculator,
			grpcEntrypoint.Config{
				Address: cfg.GRPC.Address,
			})
	)

	// Print configuration
	cfg.SetLogger(logger)
	cfg.PrintConfig()

	// Default HTTP Server handle:
	// -  Prometheus metrics and go profiling (if enable)
	// -  go profiling (if enable)
	// -  health, ready and version
	defaultHTTP := httpserver.New(
		httpserver.Logger(logger),
		httpserver.Address(cfg.HTTP.MetricsAddress),
		httpserver.ReadyCheck(readychecker),
	)

	// Start each subprocess
	errEntrypointHTTP := entrypointHTTP.Start()
	errEntrypointGRPC := entrypointGRPC.Start()
	errDefaultHTTP := defaultHTTP.Start()
	storekeeper.Start()

	logger.Info().Msg("all components are running")

	// Handle stop signals from outside
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-errEntrypointGRPC:
		logger.Error().Err(err).Msg("problem with GRPC Transport")
	case err := <-errEntrypointHTTP:
		logger.Error().Err(err).Msg("problem with HTTP Transport")
	case err := <-errDefaultHTTP:
		logger.Error().Err(err).Msg("problem with default HTTP Server")
	case <-osSignals:
		logger.Info().Msg("shutdown...")
		storekeeper.Shutdown()
		defaultHTTP.Shutdown()
		entrypointGRPC.Shutdown(context.Background())

		err := entrypointHTTP.Shutdown(context.Background())
		if err != nil {
			logger.Err(err).Msg("graceful shutdown error")
			os.Exit(1)
		}
	}
}
