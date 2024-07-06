package cmd

import (
	"github.com/alexfalkowski/go-client-template/client"
	"github.com/alexfalkowski/go-client-template/config"
	"github.com/alexfalkowski/go-service/compress"
	"github.com/alexfalkowski/go-service/encoding"
	"github.com/alexfalkowski/go-service/feature"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/go-service/telemetry/metrics"
	"go.uber.org/fx"
)

// ClientOptions for cmd.
var ClientOptions = []fx.Option{
	feature.Module,
	compress.Module, encoding.Module,
	telemetry.Module, metrics.Module,
	config.Module, client.Module, Module,
}
