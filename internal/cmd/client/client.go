package client

import (
	"context"

	"github.com/alexfalkowski/go-service/v2/di"
	"github.com/alexfalkowski/go-service/v2/telemetry/logger"
)

// Params for config.
type Params struct {
	di.In

	Lifecycle di.Lifecycle
	Logger    *logger.Logger
}

// Start the client.
func Start(params Params) {
	params.Lifecycle.Append(di.Hook{
		OnStart: func(ctx context.Context) error {
			params.Logger.Log(ctx, logger.NewText("awesome client"))

			return nil
		},
	})
}
