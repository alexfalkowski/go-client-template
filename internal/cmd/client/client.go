package client

import (
	"github.com/alexfalkowski/go-service/v3/context"
	"github.com/alexfalkowski/go-service/v3/telemetry/logger"
	"github.com/alexfalkowski/go-signal/v2"
)

// Params for config.
type Params struct {
	signal.In

	Lifecycle signal.Lifecycle
	Logger    *logger.Logger
}

// Start the client.
func Start(params Params) {
	params.Lifecycle.Append(signal.Hook{
		OnStart: func(ctx context.Context) error {
			params.Logger.Log(ctx, logger.NewText("awesome client"))

			return nil
		},
	})
}
