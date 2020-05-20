package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/maxfer4maxfer/service-template/internal/platform/config"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

// Logger holds configuration for a logger.
type Logger struct {
	Debug   bool
	Console bool
}

// Database holds configuration for a database.
type Database struct {
	DSN string
	// https://godoc.org/database/sql#DB.SetConnMaxLifetime
	ConnMaxLifetime time.Duration
	// https://godoc.org/database/sql#DB.SetMaxOpenConns
	MaxOpenConns int
	// https://godoc.org/database/sql#DB.SetMaxIdleConns
	MaxIdleConns int
}

// HTTP holds configuration for a HTTP server.
type HTTP struct {
	Address string
	// Address for prometheus and http debug
	MetricsAddress string
	// https://golang.org/pkg/net/http/#Server
	// https://medium.com/@simonfrey/go-as-in-golang-standard-net-http-config-will-break-your-production-environment-1360871cb72b
	ReadTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	MaxHeaderBytes    int
}

// GRPC holds configuration for a GRPC server.
type GRPC struct {
	Address string
}

// RateLimiter holds configuration for a RateLimiter.
// Description https://godoc.org/golang.org/x/time/rate#Limit
type RateLimiter struct {
	Interval time.Duration //rate.Every(<<Interval>>, <<B>>)
	B        int
}

// CircuitBreaker holds configuration for a CircuitBreaker.
// Description https://godoc.org/github.com/sony/gobreaker#Settings
type CircuitBreaker struct {
	MaxRequests int
	Interval    time.Duration
	Timeout     time.Duration
}

// Calculator holds configuration for a Calculator component.
type Calculator struct {
	RateLimiter    RateLimiter
	CircuitBreaker CircuitBreaker
}

// Storekeeper holds configuration for a Storekeeper schedule job.
type Storekeeper struct {
	TimeInterval time.Duration
}

// PiClient holds configuration for a SupplierClient.
type PiClient struct {
	Timeout time.Duration
	Address string
}

// Config contains all configurations of the application.
type Config struct {
	logger      *zerolog.Logger
	Logger      Logger
	Database    Database
	HTTP        HTTP
	GRPC        GRPC
	Calculator  Calculator
	Storekeeper Storekeeper
	PiClient    PiClient
}

// newConfig returns the application configuration.
// Configuration read from environment variables and from command line argumemnt.
// Command line arguments have more priority then environemt variables.
func newConfig() *Config {
	config.SetupSources("calculator")

	return &Config{
		Logger: Logger{
			Debug:   viper.GetBool("logger.debug"),
			Console: viper.GetBool("logger.console"),
		},
		Database: Database{
			DSN:             viper.GetString("database.dsn"),
			ConnMaxLifetime: viper.GetDuration("database.connmaxlifetime"),
			MaxOpenConns:    viper.GetInt("database.maxopenconns"),
			MaxIdleConns:    viper.GetInt("database.maxidleconns"),
		},
		HTTP: HTTP{
			Address:           viper.GetString("http.address"),
			MetricsAddress:    viper.GetString("http.metricsaddress"),
			ReadTimeout:       viper.GetDuration("http.readtimeout"),
			ReadHeaderTimeout: viper.GetDuration("http.readheadertimeout"),
			WriteTimeout:      viper.GetDuration("http.writetimeout"),
			IdleTimeout:       viper.GetDuration("http.idletimeout"),
			MaxHeaderBytes:    viper.GetInt("http.maxheaderbytes"),
		},
		GRPC: GRPC{
			Address: viper.GetString("grpc.address"),
		},
		Calculator: Calculator{
			RateLimiter: RateLimiter{
				Interval: viper.GetDuration("calculator.ratelimiter.interval"),
				B:        viper.GetInt("calculator.ratelimiter.b"),
			},
			CircuitBreaker: CircuitBreaker{
				MaxRequests: viper.GetInt(
					"calculator.circuitbreaker.maxrequests"),
				Interval: viper.GetDuration(
					"calculator.circuitbreaker.interval"),
				Timeout: viper.GetDuration(
					"calculator.circuitbreaker.timeout"),
			},
		},
		Storekeeper: Storekeeper{
			TimeInterval: viper.GetDuration("storekeeper.timeinterval"),
		},
		PiClient: PiClient{
			Timeout: viper.GetDuration("piclient.timeout"),
			Address: viper.GetString("piclient.address"),
		},
	}
}

// PrintConfig outputs a config to a logger.
func (cfg *Config) PrintConfig() {
	// Loger
	cfg.logger.Info().Bool("logger.debug", cfg.Logger.Debug).Msg("")
	cfg.logger.Info().Bool("logger.console", cfg.Logger.Console).Msg("")

	// Database
	cfg.logger.Info().Str("database.dsn", cfg.maskDSN(cfg.Database.DSN)).Msg("")
	cfg.logger.Info().Dur("database.connmaxlifetime", cfg.Database.ConnMaxLifetime).
		Msg("")
	cfg.logger.Info().Int("database.maxopenconns", cfg.Database.MaxOpenConns).Msg("")
	cfg.logger.Info().Int("database.maxidleconns", cfg.Database.MaxIdleConns).Msg("")

	// HTTP
	cfg.logger.Info().Str("http.address", cfg.HTTP.Address).Msg("")
	cfg.logger.Info().Str("http.metricsaddress", cfg.HTTP.MetricsAddress).Msg("")
	cfg.logger.Info().Dur("http.ReadTimeout", cfg.HTTP.ReadTimeout).Msg("")
	cfg.logger.Info().Dur("http.ReadHeaderTimeout", cfg.HTTP.ReadHeaderTimeout).
		Msg("")
	cfg.logger.Info().Dur("http.WriteTimeout", cfg.HTTP.WriteTimeout).Msg("")
	cfg.logger.Info().Dur("http.IdleTimeout", cfg.HTTP.IdleTimeout).Msg("")
	cfg.logger.Info().Int("http.MaxHeaderBytes", cfg.HTTP.MaxHeaderBytes).Msg("")

	// GRPC
	cfg.logger.Info().Str("grpc.address", cfg.GRPC.Address).Msg("")

	// Calculator
	cfg.logger.Info().Dur(
		"calculator.ratelimiter.interval", cfg.Calculator.RateLimiter.Interval,
	).Msg("")
	cfg.logger.Info().Int(
		"calculator.ratelimiter.b", cfg.Calculator.RateLimiter.B,
	).Msg("")
	cfg.logger.Info().Int(
		"calculator.circuitbreaker.maxrequests",
		cfg.Calculator.CircuitBreaker.MaxRequests,
	).Msg("")
	cfg.logger.Info().Dur(
		"calculator.circuitbreaker.interval", cfg.Calculator.CircuitBreaker.Interval,
	).Msg("")
	cfg.logger.Info().Dur(
		"calculator.circuitbreaker.timeout", cfg.Calculator.CircuitBreaker.Timeout,
	).Msg("")

	// Storekeeper
	cfg.logger.Info().Dur("storekeeper.timeinterval", cfg.Storekeeper.TimeInterval).
		Msg("")

	// PiClient
	cfg.logger.Info().Dur("piclient.timeout", cfg.PiClient.Timeout).Msg("")
	cfg.logger.Info().Str("piclient.address", cfg.PiClient.Address).Msg("")
}

// SetLogger sets logger.
func (cfg *Config) SetLogger(logger *zerolog.Logger) {
	cfg.logger = logger
}

func (cfg *Config) maskDSN(dsn string) string {
	// postgres://<USERNAME>:<PASSWORD>@<IP_ADDRESS>:<PORT>/<DB_NAME>?sslmode=disable
	const (
		numberFirstChars = 2
		numberLastChars  = 2
	)

	split := strings.Split(dsn, ":")

	user := split[1][2:]
	pass := strings.Split(split[2], "@")[0]

	maskUser := fmt.Sprintf(
		"%s%s%s",
		user[:numberFirstChars],
		strings.Repeat("*", len(user)-numberFirstChars+numberLastChars),
		user[len(user)-numberLastChars:],
	)
	maskPass := strings.Repeat("*", len(pass))

	dsn = strings.Replace(dsn, user, maskUser, 1)
	dsn = strings.Replace(dsn, pass, maskPass, 1)

	return dsn
}
