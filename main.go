package main

import (
	"github.com/alexfalkowski/go-client-template/internal/cmd"
	"github.com/alexfalkowski/go-service/v3/cli"
	"github.com/alexfalkowski/go-service/v3/context"
	"github.com/alexfalkowski/go-service/v3/os"
)

var app = cli.NewApplication(func(command cli.Commander) {
	cmd.RegisterClient(command)
})

func main() {
	os.Exit(app.RunCode(context.Background()))
}
