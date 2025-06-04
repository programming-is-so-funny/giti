package handlers

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/programming-is-so-funny/giti/internal/registry"
)

type HelpCommand struct{}

func (c *HelpCommand) Name() string {
	return "help"
}

func (c *HelpCommand) Description() string {
	return "Show help for giti or a specific command"
}

func (c *HelpCommand) Run(args []string) error {
	fmt.Println("Usage:\n giti <command> [options]")
	fmt.Println("Available Commands:")

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	for _, cmd := range registry.Commands {
		fmt.Fprintf(w, "  %s\t%s\n", cmd.Name(), cmd.Description())
	}
	w.Flush()

	fmt.Println("\nUse \"giti <command> --help\" for more information about a command.")
	return nil
}
