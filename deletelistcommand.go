package main

import (
	"fmt"
	"io"
)

//DeleteListCommand is a struct that
//facilitates the deletion of lists.
type DeleteListCommand struct {
	definition  string
	description string
	List        string
}

//NewDeleteListCommand creates a command that
//can be used to delete the specified list.
func NewDeleteListCommand(list string) *DeleteListCommand {
	command := new(DeleteListCommand)

	command.definition = "goboom delete <list>"
	command.description = "Deletes the specified list and its entries."
	command.List = list

	return command
}

//GetDescription retrieves a description of what the command does.
func (c *DeleteListCommand) GetDescription() string {
	return c.description
}

//GetDefinition retrieves definition to use for this command.
func (c *DeleteListCommand) GetDefinition() string {
	return c.definition
}

//Execute runs the current command.
func (c *DeleteListCommand) Execute(w io.Writer, s Storage) {

	//Delete the list
	for _, list := range s.Lists {

		delete(list, c.List)

	}

	//Persist our data back.
	s.Save()

	//Notify the user.
	fmt.Fprintf(w, "List '%v' successfully deleted!\n", c.List)

}
