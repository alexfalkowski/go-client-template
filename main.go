package main

import (
	"github.com/alexfalkowski/go-client-template/internal/cmd"
	"github.com/alexfalkowski/go-service/v2/cli"
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/os"
)

var app = cli.NewApplication(func(command cli.Commander) {
	cmd.RegisterClient(command)
})

func main() {
	os.Exit(app.RunCode(context.Background()))
}
