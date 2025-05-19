package main

import (
	"context"

	"github.com/alexfalkowski/go-client-template/internal/cmd"
	sc "github.com/alexfalkowski/go-service/cmd"
)

var app = sc.NewApplication(func(command *sc.Command) {
	cmd.RegisterClient(command)
})

func main() {
	app.ExitOnError(context.Background())
}
