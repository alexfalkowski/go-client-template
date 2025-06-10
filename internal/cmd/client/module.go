package client

import (
	"github.com/alexfalkowski/go-client-template/internal/config"
	"github.com/alexfalkowski/go-service/v2/di"
	"github.com/alexfalkowski/go-service/v2/module"
)

// Module for fx.
var Module = di.Module(
	module.Client,
	config.Module,
	di.Register(Start),
)
