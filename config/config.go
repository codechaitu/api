package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type Env struct {
	// Environment
	Env string `envconfig:"ENV" default:"dev"`

	// Port at which server should run and look for connections
	Port string `envconfig:"PORT" default:"9000"`

	// ServerDegradationTime is used to properly shutdown the server
	ServerDegradationTime time.Duration `envconfig:"SERVER_DEGRADATION_TIME" default:"10s"`
}

func ReadEnvVariables() (*Env, error) {
	var env Env
	if err := envconfig.Process("", &env); err != nil {
		return nil, errors.Wrap(err, "error reading env variables")
	}

	return &env, nil
}
