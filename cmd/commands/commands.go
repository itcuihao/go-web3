package commands

import (
	"os"

	"github.com/mitchellh/cli"
)

func Commands() map[string]cli.CommandFactory {
	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	return map[string]cli.CommandFactory{
		"abigen": func() (cli.Command, error) {
			return &AbigenCommand{
				UI: ui,
			}, nil
		},
	}
}
