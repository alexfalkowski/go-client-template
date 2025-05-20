package main

import (
	"context"

	"github.com/alexfalkowski/go-client-template/internal/cmd"
	"github.com/alexfalkowski/go-service/v2/cli"
)

var app = cli.NewApplication(func(command cli.Commander) {
	cmd.RegisterClient(command)
})

func main() {
	app.ExitOnError(context.Background())
}
