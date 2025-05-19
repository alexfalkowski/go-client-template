package main

import (
	"context"

	"github.com/alexfalkowski/go-client-template/internal/cmd"
	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/env"
	"github.com/alexfalkowski/go-service/os"
)

func main() {
	command().ExitOnError(context.Background())
}

func command() *sc.Command {
	fs := os.NewFS()
	command := sc.New(env.NewName(fs), env.NewVersion())

	cmd.RegisterClient(command)

	return command
}
