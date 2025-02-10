package main

import (
	"github.com/alexfalkowski/go-client-template/cmd"
	sc "github.com/alexfalkowski/go-service/cmd"
)

func main() {
	command().ExitOnError()
}

func command() *sc.Command {
	c := sc.New(cmd.Version)
	cl := c.AddClient("client", "Start client", cmd.ClientOptions...)
	c.RegisterInput(cl, "")
	c.RegisterOutput(cl, "")

	return c
}
