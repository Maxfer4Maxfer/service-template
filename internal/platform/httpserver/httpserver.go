package httpserver

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"runtime"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
)

// ReadyChecker checks can the service be able to handle requests.
type ReadyChecker interface {
	Ready(ctx context.Context) error
}

// Version declares the behavior of a version information container.
type Version interface {
	BuildTime() string
	BuildBranch() string
	BuildCommit() string
	BuildSummary() string
}

func defaultOptions() Options {
	logger := zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()

	const maximumprofilefraction = 16

	return Options{
		logger:                 &logger,
		address:                "localhost:8080",
		debug:                  false,
		maximumprofilefraction: maximumprofilefraction,
	}
}

// HTTPServer start and stop infra http server.
// Main perpose serves up stuff like the Prometheus metrics route,
// the Go debug and profiling routes, and so on.
type HTTPServer struct {
	logger   *zerolog.Logger
	listener net.Listener
	mux      *http.ServeMux
}

// New returns system Default HTTP Server that handles:
// -  Prometheus metrics;
// -  go profiling (if enable);
// -  health, ready and version.
func New(opts ...Option) *HTTPServer {
	options := defaultOptions()

	for _, o := range opts {
		o(&options)
	}

	listener, err := net.Listen("tcp", options.address)
	if err != nil {
		options.logger.Error().Err(err).Str("address", options.address).
			Msg("address already in use")
		os.Exit(1)
	}

	router := http.NewServeMux()

	router.Handle("/metrics", promhttp.Handler())
	router.Handle("/ready", handlerReady(options.readychecker))
	router.Handle("/health", handleHealth())
	router.Handle("/version", handleVersion(options.version))

	if options.debug {
		runtime.SetMutexProfileFraction(options.maximumprofilefraction)

		router.HandleFunc("/debug/pprof/", pprof.Index)
		router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		router.HandleFunc("/debug/pprof/profile", pprof.Profile)
		router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		router.HandleFunc("/debug/pprof/trace", pprof.Trace)
	}

	return &HTTPServer{
		logger:   options.logger,
		listener: listener,
		mux:      router,
	}
}

// Start starts HTTP Server.
func (srv *HTTPServer) Start() chan error {
	serverErrors := make(chan error, 1)

	go func() {
		srv.logger.Info().
			Msgf("Start API Listening %s", srv.listener.Addr().String())

		serverErrors <- http.Serve(srv.listener, srv.mux)
	}()

	return serverErrors
}

// Shutdown stops HTTP Server.
func (srv *HTTPServer) Shutdown() {
	err := srv.listener.Close()
	if err != nil {
		srv.logger.Error().Err(err).Msgf("cannot stop the http server")
	}
}

func handleHealth() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}

func handlerReady(rc ReadyChecker) http.HandlerFunc {
	if rc == nil {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := rc.Ready(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	})
}

func handleVersion(version Version) http.HandlerFunc {
	if version == nil {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Version is not defined")
		})
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "BuildTime = %s \n", version.BuildTime())
		fmt.Fprintf(w, "BuildBranch = %s \n", version.BuildBranch())
		fmt.Fprintf(w, "BuildCommit = %s \n", version.BuildCommit())
		fmt.Fprintf(w, "BuildSummary = %s \n", version.BuildSummary())
	})
}
