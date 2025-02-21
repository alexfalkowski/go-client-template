package client

import (
	"context"
	"log/slog"

	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/telemetry/logger"
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
	cmd.Start(params.Lifecycle, func(ctx context.Context) {
		params.Logger.LogAttrs(ctx, slog.LevelInfo, "awesome client", nil)
	})
}
