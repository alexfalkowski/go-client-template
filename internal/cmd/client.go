package cmd

import (
	"github.com/alexfalkowski/go-client-template/internal/cmd/client"
	"github.com/alexfalkowski/go-client-template/internal/config"
	"github.com/alexfalkowski/go-service/v2/cli"
	"github.com/alexfalkowski/go-service/v2/feature"
	"github.com/alexfalkowski/go-service/v2/module"
	"github.com/alexfalkowski/go-service/v2/telemetry"
)

// RegisterClient for cmd.
func RegisterClient(command cli.Commander) {
	cmd := command.AddClient("client", "Start client",
		module.Module, feature.Module, telemetry.Module,
		config.Module, client.Module, cli.Module,
	)
	cmd.AddInput("")
}
