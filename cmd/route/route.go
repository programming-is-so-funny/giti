package route

import "github.com/programming-is-so-funny/giti/internal/registry"

type ExitCode int

const (
	ExitSuccess       ExitCode = 0
	ExitUsageError    ExitCode = 1
	ExitFileNotFound  ExitCode = 2
	ExitPermission    ExitCode = 3
	ExitInternalError ExitCode = 10
)

func Route(args []string) ExitCode {
	if len(args) < 1 {
		registry.Commands[0].Run(args)
	}

	for _, c := range registry.Commands {
		if c.Name() == args[0] {
			err := c.Run(args)
			if err != nil {
				return ExitUsageError
			}
			return ExitSuccess
		}
	}
	return ExitSuccess
}
