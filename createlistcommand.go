package main

import (
	"fmt"
	"io"
)

//CreateListCommand is a struct that
//facilitates the deletion of lists.
type CreateListCommand struct {
	definition  string
	description string
	List        string
}

//NewCreateListCommand creates a command that
//can be used to Create the specified list.
func NewCreateListCommand(list string) *CreateListCommand {
	command := new(CreateListCommand)

	command.definition = "goboom <list>"
	command.description = "Creates the specified list."
	command.List = list

	return command
}

//GetDescription retrieves a description of what the command does.
func (c *CreateListCommand) GetDescription() string {
	return c.description
}

//GetDefinition retrieves definition to use for this command.
func (c *CreateListCommand) GetDefinition() string {
	return c.definition
}

//Execute runs the current command.
func (c *CreateListCommand) Execute(w io.Writer, s Storage) {

	list := make(map[string][]map[string]string, 1)

	list[c.List] = make([]map[string]string, 0)

	s.Lists = append(s.Lists, list)

	//Persist our data back.
	s.Save()

	//Notify the user.
	fmt.Fprintf(w, "List '%v' successfully Created!\n", c.List)

}
