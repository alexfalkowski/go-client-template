package client

import (
	"context"

	"github.com/alexfalkowski/go-service/cmd"
	tz "github.com/alexfalkowski/go-service/telemetry/logger/zap"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Params for config.
type Params struct {
	fx.In

	Lifecycle fx.Lifecycle
	Logger    *zap.Logger
}

// Start the client.
func Start(params Params) {
	cmd.Start(params.Lifecycle, func(ctx context.Context) {
		params.Logger.Info("awesome client", tz.Meta(ctx)...)
	})
}
