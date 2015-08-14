package main

import (
	"fmt"
	"io"
)

//OverviewCommand prints a summary of what's stored in each list.
type OverviewCommand struct {
}

//GetDescription retrieves a description of what the command does.
func (c *OverviewCommand) GetDescription() string {
	return "Displays an overview of your lists."
}

//GetDefinition retrieves definition to use for this command.
func (c *OverviewCommand) GetDefinition() string {
	return "goboom"
}

//Execute runs the current command.
func (c *OverviewCommand) Execute(w io.Writer, s Storage) {

	for _, list := range s.Lists {

		for k := range list {
			fmt.Fprintf(w, defaultIndention+"%v (%v)\n", k, len(list[k]))
		}
	}
}
