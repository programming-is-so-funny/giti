package command

import (
	"github.com/programming-is-so-funny/giti/cmd/handlers"
	contract "github.com/programming-is-so-funny/giti/internal/contracts"
	"github.com/programming-is-so-funny/giti/internal/registry"
)

func RegisterCommands() {
	registry.Commands = []contract.Command{
		&handlers.HelpCommand{},
		&handlers.InitCommand{InitDir: "."},
	}
}
