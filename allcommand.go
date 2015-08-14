package main

import (
	"fmt"
	"io"
)

//maxKeyLength defines the longest key that can be printed to the screen.
const maxKeyLength = 16

//AllCommand is a command that prints everything we have stored.
type AllCommand struct {
	definition  string
	description string
	List        string
}

//NewAllCommand creates a command that lists all
//entries we have stored within our .boom file.
func NewAllCommand() *AllCommand {
	command := new(AllCommand)

	command.definition = "goboom all"
	command.description = "Displays all entries in the .boom file."

	return command
}

//NewAllListCommand creates a command that lists all
//entries within a specific list in our .boom file.
func NewAllListCommand(list string) *AllCommand {
	command := new(AllCommand)

	command.definition = "goboom <list>"
	command.description = "Displays all entries within the specified list."
	command.List = list

	return command
}

//GetDescription retrieves a description of what the command does.
func (c *AllCommand) GetDescription() string {
	return c.description
}

//GetDefinition retrieves definition to use for this command.
func (c *AllCommand) GetDefinition() string {
	return c.definition
}

//Execute runs the current command.
func (c *AllCommand) Execute(w io.Writer, s Storage) {

	for _, list := range s.Lists {

		for entries := range list {

			if c.List == "" || c.List == entries {

				fmt.Fprintf(w, defaultIndention+"%v\n", entries)

				for _, entry := range list[entries] {

					for key, value := range entry {

						if len(key) > maxKeyLength {
							key = key[:16] + "…"
						}

						fmt.Fprintf(w, defaultIndention+defaultIndention+"%v:\t%v\n", key, value)

					}
				}
			}
		}
	}
}
