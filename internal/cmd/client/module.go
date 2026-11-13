package client

import (
	"github.com/alexfalkowski/go-client-template/internal/config"
	"github.com/alexfalkowski/go-service/v3/module"
	"github.com/alexfalkowski/go-signal/v2"
)

// Module for signal.
var Module = signal.Module(
	module.Client,
	config.Module,
	signal.Register(Start),
)
