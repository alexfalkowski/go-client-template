package cmd

import (
	"github.com/alexfalkowski/go-client-template/cmd/client"
	"github.com/alexfalkowski/go-client-template/config"
	"github.com/alexfalkowski/go-service/compress"
	"github.com/alexfalkowski/go-service/encoding"
	"github.com/alexfalkowski/go-service/feature"
	"github.com/alexfalkowski/go-service/sync"
	"github.com/alexfalkowski/go-service/telemetry"
	"go.uber.org/fx"
)

// ClientOptions for cmd.
var ClientOptions = []fx.Option{
	sync.Module, compress.Module, encoding.Module,
	feature.Module, telemetry.Module,
	config.Module, client.Module, Module,
}
