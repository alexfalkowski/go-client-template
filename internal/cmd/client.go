package cmd

import (
	"github.com/alexfalkowski/go-client-template/internal/cmd/client"
	"github.com/alexfalkowski/go-client-template/internal/config"
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/feature"
	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/go-service/module"
	"github.com/alexfalkowski/go-service/telemetry"
)

// RegisterClient for cmd.
func RegisterClient(command *cmd.Command) {
	flags := flags.NewFlagSet("client")
	flags.AddInput("")

	command.AddClient("client", "Start client", flags,
		module.Module, feature.Module, telemetry.Module,
		config.Module, client.Module, cmd.Module,
	)
}
