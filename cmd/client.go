package cmd

import (
	"github.com/alexfalkowski/go-client-template/client"
	"github.com/alexfalkowski/go-client-template/config"
	"github.com/alexfalkowski/go-client-template/transport"
	"github.com/alexfalkowski/go-service/debug"
	"github.com/alexfalkowski/go-service/feature"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/go-service/telemetry/metrics"
	"go.uber.org/fx"
)

// ClientOptions for cmd.
var ClientOptions = []fx.Option{
	fx.NopLogger, runtime.Module, feature.Module,
	transport.Module, debug.Module,
	telemetry.Module, metrics.Module,
	config.Module, client.Module, Module,
}
