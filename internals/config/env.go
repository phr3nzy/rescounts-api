// package config stores configuration information read from envinronment
// variables and/or files
package config

import (
	"fmt"

	"github.com/phr3nzy/rescounts-api/internals/utils/validator"

	"github.com/caarlos0/env"
	v "github.com/go-playground/validator/v10"
)

// config is the application configuration
// read from ENV vars.
type config struct {
	ServiceName    string `env:"SERVICE_NAME"  															                 validate:"required"`
	AppEnv         string `env:"APP_ENV"         				envDefault:"development" 						 validate:"required"`
	HOST           string `env:"HOST"										envDefault:"localhost"							 validate:"required"`
	PORT           int    `env:"PORT"       						envDefault:"3000"                    validate:"required,integer"`
	LogLevel       string `env:"LOG_LEVEL"      				envDefault:"info"                    validate:"required"`
	DisableLogging bool   `env:"DISABLE_LOGGING" 				envDefault:"false"									 validate:"bool"`
}

// parse parses, validates and then returns the application
// configuration based on environment variables.
func parse(val *v.Validate) (*config, error) {
	cfg := &config{}

	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to parse env: %w", err)
	}

	if err := val.Struct(cfg); err != nil {
		return nil, fmt.Errorf("failed to project env on struct: %w", err)
	}

	return cfg, nil
}

// GetConfig returns the current environment config.
func GetConfig() *config {
	validate := validator.New()
	config, err := parse(validate)
	if err != nil {
		panic(err)
	}
	return config
}
