package client

import (
	"context"

	"github.com/alexfalkowski/go-service/v2/telemetry/logger"
	"go.uber.org/fx"
)

// Params for config.
type Params struct {
	fx.In

	Lifecycle fx.Lifecycle
	Logger    *logger.Logger
}

// Start the client.
func Start(params Params) {
	params.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			params.Logger.Log(ctx, logger.NewText("awesome client"))

			return nil
		},
	})
}
