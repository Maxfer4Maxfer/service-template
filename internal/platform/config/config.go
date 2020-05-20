package config

import (
	"os"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/maxfer4maxfer/service-template/internal/platform/errors"
)

var (
	errConfigFlags = errors.New("cannot read config flags")
	errConfigFile  = errors.New("cannot read a config file")
)

// SetupSources defines directory with a config file.
func SetupSources(cfgName string) {
	viper.SetConfigName(getConfigName(cfgName))

	viper.AddConfigPath("./config/")

	err := viper.ReadInConfig()
	if err != nil {
		panic(errConfigFile.Wrap(err))
	}

	// Environment variables
	r := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(r)
	viper.AutomaticEnv()

	pflag.Parse()

	err = viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(errConfigFlags.Wrap(err))
	}
}

// getConfigName returns a config file name.
func getConfigName(cfgName string) string {
	cfgEnvName := os.Getenv("CONFIG_NAME")

	cfgName += "."

	if cfgEnvName != "" {
		cfgName += cfgEnvName
	} else {
		cfgName += "local"
	}

	return cfgName
}
