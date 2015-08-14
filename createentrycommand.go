package main

import (
	"fmt"
	"io"
)

//CreateEntryCommand is a struct that
//facilitates the deletion of Entrys.
type CreateEntryCommand struct {
	definition  string
	description string
	Value       string
	Entry       string
	List        string
}

//NewCreateEntryCommand creates a command that
//can be used to Create the specified Entry.
func NewCreateEntryCommand(list string, entry string, value string) *CreateEntryCommand {
	command := new(CreateEntryCommand)

	command.definition = "goboom <list> <entry> <value>"
	command.description = "Creates an entry in the specified list."
	command.Entry = entry
	command.List = list
	command.Value = value

	return command
}

//GetDescription retrieves a description of what the command does.
func (c *CreateEntryCommand) GetDescription() string {
	return c.description
}

//GetDefinition retrieves definition to use for this command.
func (c *CreateEntryCommand) GetDefinition() string {
	return c.definition
}

//Execute runs the current command.
func (c *CreateEntryCommand) Execute(w io.Writer, s Storage) {

	for _, list := range s.Lists {
		for name, entries := range list {
			if name == c.List {

				found := false

				for _, entry := range entries {
					for key := range entry {
						if key == c.Entry {
							entry[key] = c.Value
							found = true
						}
					}
				}

				if !found {
					entry := make(map[string]string)
					entry[c.Entry] = c.Value

					list[c.List] = append(list[c.List], entry)
				}

			}
		}
	}

	//Persist our data back.
	s.Save()

	//Notify the user.
	fmt.Fprintf(w, "Entry '%v' successfully Created!\n", c.Entry)

}
