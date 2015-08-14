package main

import (
	"fmt"
	"io"
	"os/exec"
)

//OpenCommand facilitates the opening
//of links in the user's browser.
type OpenCommand struct {
	definition  string
	description string
	List        string
	Entry       string
}

//NewOpenListCommand creates a command that will
//open all links within a list in the user's browser.
func NewOpenListCommand(list string) *OpenCommand {
	command := new(OpenCommand)

	command.definition = "goboom open <list>"
	command.description = "Opens all links in a list in the user's browser."
	command.List = list

	return command
}

//NewOpenListEntryCommand creates a command that will
//open a specific link the user's browser.
func NewOpenListEntryCommand(list string, entry string) *OpenCommand {
	command := new(OpenCommand)

	command.definition = "goboom open <list> <entry>"
	command.description = "Opens the specified entry in the user's browser."
	command.List = list
	command.Entry = entry

	return command
}

//GetDescription retrieves a description of what the command does.
func (c *OpenCommand) GetDescription() string {
	return c.description
}

//GetDefinition retrieves definition to use for this command.
func (c *OpenCommand) GetDefinition() string {
	return c.definition
}

//Execute runs the current command.
func (c *OpenCommand) Execute(w io.Writer, s Storage) {

	//If entry is empty, open all links in browser.
	entries := s.FindList(c.List)

	for _, entry := range entries {
		for key, value := range entry {
			if c.Entry == "" || key == c.Entry {
				fmt.Fprintf(w, "Opening '%v.'\n", value)
				exec.Command("open", value).Start()
			}
		}
	}
}
