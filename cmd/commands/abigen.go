package commands

import (
	"github.com/mitchellh/cli"
)

// VersionCommand is the command to show the version of the agent
type AbigenCommand struct {
	UI cli.Ui
}

// Help implements the cli.Command interface
func (c *AbigenCommand) Help() string {
	return `Usage: greenhouse version
  Display the Greenhouse version`
}

// Synopsis implements the cli.Command interface
func (c *AbigenCommand) Synopsis() string {
	return "Display the Greenhouse version"
}

// Run implements the cli.Command interface
func (c *AbigenCommand) Run(args []string) int {
	return 0
}
