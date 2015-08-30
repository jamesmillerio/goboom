package main

import (
	"io"
	"strings"
)

//CatchAllCommand is a struct that facilitates
//several functions including the creation of
//lists, showing list items, creating new list
//items, and copying entry's to the clipboard.
type CatchAllCommand struct {
	definition  string
	description string
	commands    []Command
	storage     Storage
}

//NewCatchAllCommand creates a command that
//can be used to delete the specified list.
func NewCatchAllCommand(args []string, translator Translator) *CatchAllCommand {
	command := new(CatchAllCommand)
	count := len(args)

	command.storage = NewDefaultListStorage()

	switch count {
	case 1:

		value := command.storage.FindValueByEntry(args[0])
		list := command.storage.FindList(args[0])

		command.commands = make([]Command, 1)

		if value != "" {
			//Copy entry commnd
			command.commands[0] = NewCopyEntryCommand(args[0])
		} else if list != nil {
			//Print entries in list command.
			command.commands[0] = NewAllListCommand(args[0], translator)
		} else if list == nil {
			command.commands[0] = NewCreateListCommand(args[0])
		}

		break
	case 2:
		command.commands = make([]Command, 1)

		//Copy entry to clipboard.
		command.commands[0] = NewCopySpecificEntryCommand(args[0], args[1])
		break
	default:

		list := command.storage.FindList(args[0])
		value := strings.Join(args[2:], " ")

		if list == nil {

			command.commands = make([]Command, 2)

			command.commands[0] = NewCreateListCommand(args[0])
			command.commands[1] = NewCreateEntryCommand(args[0], args[1], value)

		} else {

			command.commands = make([]Command, 1)

			command.commands[0] = NewCreateEntryCommand(args[0], args[1], value)

		}

		//Create new list entry
		break
	}

	return command
}

//GetDescription retrieves a description of what the command does.
func (c *CatchAllCommand) GetDescription() string {
	return c.description
}

//GetDefinition retrieves definition to use for this command.
func (c *CatchAllCommand) GetDefinition() string {
	return c.definition
}

//Execute runs the current command.
func (c *CatchAllCommand) Execute(w io.Writer, s Storage) {

	if c.commands == nil {
		return
	}

	for _, command := range c.commands {
		if command != nil {
			command.Execute(w, s)
		}
	}

}
