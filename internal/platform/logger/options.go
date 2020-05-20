package logger

// Option is the func interface to assign options.
type Option func(*Options)

// Options is a struct for options for logger.
type Options struct {
	console   bool
	debug     bool
	fieldname string
}

// FieldName defines the FieldName options.
func FieldName(c string) Option {
	return func(o *Options) {
		o.fieldname = c
	}
}

// ConsoleOutput defines should logger output its messages
// using the console format.
func ConsoleOutput(c bool) Option {
	return func(o *Options) {
		o.console = c
	}
}

// DebugMode switch on/off debug mode.
func DebugMode(c bool) Option {
	return func(o *Options) {
		o.debug = c
	}
}
