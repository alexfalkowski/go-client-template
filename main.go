package main

import (
	"github.com/alexfalkowski/go-client-template/internal/cmd"
	sc "github.com/alexfalkowski/go-service/cmd"
)

func main() {
	command().ExitOnError()
}

func command() *sc.Command {
	command := sc.New(cmd.Version)

	cmd.RegisterClient(command)

	return command
}
