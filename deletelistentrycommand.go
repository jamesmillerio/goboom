package main

import (
	"fmt"
	"io"
)

//DeleteListEntryCommand facilitates the
//deletion of specific entries in the
//specified list.
type DeleteListEntryCommand struct {
	definition  string
	description string
	List        string
	Entry       string
}

//NewDeleteListEntryCommand creates the
//command used to delete a specific entry.
func NewDeleteListEntryCommand(list string, entry string) *DeleteListEntryCommand {
	command := new(DeleteListEntryCommand)

	command.definition = "goboom delete <list> <entry>"
	command.description = "Deletes the specified entry from a list."
	command.List = list
	command.Entry = entry

	return command
}

//GetDescription retrieves a description of what the command does.
func (c *DeleteListEntryCommand) GetDescription() string {
	return c.description
}

//GetDefinition retrieves definition to use for this command.
func (c *DeleteListEntryCommand) GetDefinition() string {
	return c.definition
}

//Execute runs the current command.
func (c *DeleteListEntryCommand) Execute(w io.Writer, s Storage) {

	//Delete the list
	for _, list := range s.Lists {

		delete(list, c.List)

	}

	//Persist our data back.
	s.Save()

	//Notify the user.
	fmt.Fprintf(w, "List '%v' successfully deleted!\n", c.List)

}
