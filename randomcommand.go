package main

import (
	"fmt"
	"io"
	"math/rand"
	"time"
)

//RandomCommand facilitates printing a random
//entry from either a list or all entries.
type RandomCommand struct {
	List        string
	description string
	definition  string
}

//NewRandomEntryCommand creates a command that will
//return a random entry from all entries.
func NewRandomEntryCommand() *RandomCommand {
	command := new(RandomCommand)

	command.definition = "goboom random"
	command.description = "Retrieves a random entry."

	return command
}

//NewRandomSpecificEntryCommand creates a command that will
//return a random entry from a specific list.
func NewRandomSpecificEntryCommand(list string) *RandomCommand {
	command := new(RandomCommand)

	command.definition = "goboom random <list>"
	command.description = "Retrieves a random entry from the specified list."
	command.List = list

	return command
}

//GetDescription retrieves a description of what the command does.
func (c *RandomCommand) GetDescription() string {
	return c.description
}

//GetDefinition retrieves definition to use for this command.
func (c *RandomCommand) GetDefinition() string {
	return c.definition
}

//Execute runs the current command.
func (c *RandomCommand) Execute(w io.Writer, s Storage) {

	rand.Seed(time.Now().UnixNano())

	//values := make([]string, 0)
	var values []string

	//Gather all of the values in our lists.
	for _, list := range s.Lists {
		for key, entries := range list {
			if c.List == "" || c.List == key {
				for _, entry := range entries {
					for _, value := range entry {
						values = append(values, value)
					}
				}
			}
		}
	}

	index := rand.Intn(len(values))
	fmt.Fprintf(w, "%v\n", values[index])

}
