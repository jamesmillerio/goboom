package main

import (
	"fmt"
	"io"
)

//EchoCommand facilitates the echoing
//of a value to screen.
type EchoCommand struct {
	definition  string
	description string
	List        string
	Entry       string
}

//NewEchoEntryCommand creates a command that echoes
//the first entry it finds with the specified name.
func NewEchoEntryCommand(entry string) *EchoCommand {
	command := new(EchoCommand)

	command.definition = "goboom echo <entry>"
	command.description = "Echoes the specified entry to the screen."
	command.Entry = entry

	return command
}

//NewEchoSpecificEntryCommand creates a command that echoes
//the entry it finds with the specified name in the specific list.
func NewEchoSpecificEntryCommand(list string, entry string) *EchoCommand {
	command := new(EchoCommand)

	command.definition = "goboom echo <list> <entry>"
	command.description = "Echoes the entry in the specified list to the screen."
	command.List = list
	command.Entry = entry

	return command
}

//GetDescription retrieves a description of what the command does.
func (c *EchoCommand) GetDescription() string {
	return c.description
}

//GetDefinition retrieves definition to use for this command.
func (c *EchoCommand) GetDefinition() string {
	return c.definition
}

//Execute runs the current command.
func (c *EchoCommand) Execute(w io.Writer, s Storage) {

	for _, list := range s.Lists {
		for key, entries := range list {
			if key == c.List || c.List == "" {
				for _, entry := range entries {
					for key, value := range entry {
						if key == c.Entry {
							fmt.Fprint(w, value+"\n")
						}
					}
				}
			}
		}
	}
}
