package client

import (
	"github.com/alexfalkowski/go-client-template/internal/config"
	"github.com/alexfalkowski/go-service/v2/module"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	module.Client,
	config.Module,
	fx.Invoke(Start),
)
