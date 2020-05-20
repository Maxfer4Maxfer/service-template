package http

import (
	"context"
	"net/http"
	"time"

	"github.com/rs/zerolog"
)

// Config holds configuration for a HTTP server.
type Config struct {
	Address string
	// https://golang.org/pkg/net/http/#Server
	// https://medium.com/@simonfrey/go-as-in-golang-standard-net-http-config-will-break-your-production-environment-1360871cb72b
	ReadTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	MaxHeaderBytes    int
}

// Service declares interactions with the errors service.
type Service interface {
	Add(ctx context.Context, a int, b int) (sum int, err error)
	Subtract(ctx context.Context, a int, b int) (sub int, err error)
	Multiply(ctx context.Context, a int, b int) (mult int, err error)
}

// ServerHTTP is a wraper of http.Server.
type ServerHTTP struct {
	logger  *zerolog.Logger
	cfg     Config
	service Service
	server  *http.Server
}

// New returns a HTTP server.
func New(logger *zerolog.Logger, svc Service, cfg Config) *ServerHTTP {
	s := &ServerHTTP{
		logger:  logger,
		cfg:     cfg,
		service: svc,
		server: &http.Server{
			Addr:              cfg.Address,
			ReadTimeout:       cfg.ReadTimeout,
			ReadHeaderTimeout: cfg.ReadHeaderTimeout,
			WriteTimeout:      cfg.WriteTimeout,
			IdleTimeout:       cfg.IdleTimeout,
			MaxHeaderBytes:    cfg.MaxHeaderBytes,
		},
	}

	s.routes()

	return s
}

// Start starts HTTP Server.
func (s *ServerHTTP) Start() chan error {
	serverErrors := make(chan error, 1)

	go func() {
		s.logger.Info().Msgf("Start HTTP API Listening %s", s.cfg.Address)
		serverErrors <- s.server.ListenAndServe()
	}()

	return serverErrors
}

// Shutdown stops HTTP Server.
func (s *ServerHTTP) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
