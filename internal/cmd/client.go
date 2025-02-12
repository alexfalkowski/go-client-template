package cmd

import (
	"github.com/alexfalkowski/go-client-template/internal/cmd/client"
	"github.com/alexfalkowski/go-client-template/internal/config"
	"github.com/alexfalkowski/go-service/feature"
	"github.com/alexfalkowski/go-service/module"
	"github.com/alexfalkowski/go-service/telemetry"
	"go.uber.org/fx"
)

// ClientOptions for cmd.
var ClientOptions = []fx.Option{
	module.Module, feature.Module, telemetry.Module,
	config.Module, client.Module, Module,
}
