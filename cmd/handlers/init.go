package handlers

import "github.com/programming-is-so-funny/giti/internal/fs"

type InitCommand struct {
	InitDir string
}

func (c *InitCommand) Name() string {
	return "init"
}

func (c *InitCommand) Description() string {
	return "Initializes a new giti repository in the current directory. \n This creates the internal .giti folder used to store metadata, configuration, and object data, similar to how Git initializes its own repositories.\nIt is the first command to run when starting to track a project with giti."
}

func (c *InitCommand) Run(args []string) error {
	err := fs.InitRepo(c.InitDir)
	if err != nil {
		return err
	}
	return nil
}
