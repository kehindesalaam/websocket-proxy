package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Specification represents structured configuration variables
type Specification struct {
	Name     string `envconfig:"SERVICE_NAME" default:"message-socket-service"`
	Debug    bool   `envconfig:"DEBUG" default:"true"`
	LogLevel string `envconfig:"LOG_LEVEL" default:"info"`
	Port     uint16 `envconfig:"WEB_PORT" default:"8000"`
	Version  string `envconfig:"VERSION" default:"dev"`
}

// LoadEnv loads config variables into Specification
func LoadEnv() (*Specification, error) {
	var conf Specification
	err := envconfig.Process("", &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
