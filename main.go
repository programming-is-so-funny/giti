package main

import (
	"os"

	"github.com/programming-is-so-funny/giti/cmd/command"
	"github.com/programming-is-so-funny/giti/cmd/route"
)

func main() {
	command.RegisterCommands()
	code := route.Route(os.Args[1:])
	os.Exit(int(code))
}
