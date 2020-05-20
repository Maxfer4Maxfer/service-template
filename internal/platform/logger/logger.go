package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

func defaultOptions() Options {
	return Options{
		console:   false,
		debug:     false,
		fieldname: "msg",
	}
}

// New returns a configured zerolog.
func New(opts ...Option) *zerolog.Logger {
	options := defaultOptions()

	for _, o := range opts {
		o(&options)
	}

	var logger zerolog.Logger

	zerolog.MessageFieldName = options.fieldname

	logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()

	// setup output to console
	if options.console {
		logger = logger.Output(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		})
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if options.debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	return &logger
}
