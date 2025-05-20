package config

import (
	"github.com/alexfalkowski/go-service/v2/config"
)

// Config for the client.
type Config struct {
	*config.Config `yaml:",inline" json:",inline" toml:",inline"`
}

func decorateConfig(cfg *Config) *config.Config {
	return cfg.Config
}
