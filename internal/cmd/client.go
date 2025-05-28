package cmd

import (
	"github.com/alexfalkowski/go-client-template/internal/cmd/client"
	"github.com/alexfalkowski/go-service/v2/cli"
)

// RegisterClient for cmd.
func RegisterClient(command cli.Commander) {
	cmd := command.AddClient("client", "Start client", client.Module)

	cmd.AddInput("")
}
