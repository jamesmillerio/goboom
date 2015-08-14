package main

import (
	"fmt"
	"io"
)

//HelpCommand prints the help
//documentation to the screen.
type HelpCommand struct {
	definition  string
	description string
	commands    []Command
}

//NewHelpEntryCommand creates a command that Helpes
//the first entry it finds with the specified name.
func NewHelpCommand(commands []Command) *HelpCommand {
	command := new(HelpCommand)

	command.definition = "goboom -h"
	command.description = "Displays help information."
	command.commands = commands

	return command
}

//GetDescription retrieves a description of what the command does.
func (c *HelpCommand) GetDescription() string {
	return c.description
}

//GetDefinition retrieves definition to use for this command.
func (c *HelpCommand) GetDefinition() string {
	return c.definition
}

//Execute runs the current command.
func (c *HelpCommand) Execute(w io.Writer, s Storage) {

	fmt.Fprintf(w, "\ngoboom command reference\n")
	fmt.Fprintf(w, "-----------------------------------\n\n")

	for _, command := range c.commands {
		if command != nil {
			fmt.Fprintf(w, "%v\t%v\n", command.GetDefinition(), command.GetDescription())
		} else {
			fmt.Fprintf(w, "\t\n")
		}
	}
}
