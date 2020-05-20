package httpserver

import "github.com/rs/zerolog"

// Option is the func interface to assign options.
type Option func(*Options)

// Options is a struct for options for logger.
type Options struct {
	logger                 *zerolog.Logger
	address                string
	debug                  bool
	readychecker           ReadyChecker
	version                Version
	maximumprofilefraction int
}

// Logger defines the logger.
func Logger(c *zerolog.Logger) Option {
	return func(o *Options) {
		o.logger = c
	}
}

// Address defines the address of the http server.
func Address(c string) Option {
	return func(o *Options) {
		o.address = c
	}
}

// DebugMode switch on/off debug mode.
func DebugMode(c bool) Option {
	return func(o *Options) {
		o.debug = c
	}
}

// ReadyCheck setups ReadyChecker.
func ReadyCheck(c ReadyChecker) Option {
	return func(o *Options) {
		o.readychecker = c
	}
}

// VersionProvider setups Version.
func VersionProvider(c Version) Option {
	return func(o *Options) {
		o.version = c
	}
}

// MaximumProfileFraction setups maximumprofilefraction.
func MaximumProfileFraction(c int) Option {
	return func(o *Options) {
		o.maximumprofilefraction = c
	}
}
