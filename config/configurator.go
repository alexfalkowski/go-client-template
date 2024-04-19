package config

import (
	"github.com/alexfalkowski/go-client-template/health"
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/config"
)

// NewConfigurator for config.
func NewConfigurator(i *cmd.InputConfig) (config.Configurator, error) {
	c := &Config{}

	return c, i.Unmarshal(c)
}

func healthConfig(cfg config.Configurator) *health.Config {
	return cfg.(*Config).Health
}
