package main

import (
	"github.com/alexfalkowski/go-client-template/internal/cmd"
	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/env"
)

func main() {
	command().ExitOnError()
}

func command() *sc.Command {
	command := sc.New(env.NewName(), env.NewVersion())

	cmd.RegisterClient(command)

	return command
}
