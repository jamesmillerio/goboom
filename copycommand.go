package main

import (
	"io"
	"os/exec"
)

//CopyCommand is a structure to hold
//arugments related to copying values.
type CopyCommand struct {
	definition  string
	description string
	List        string
	Entry       string
}

//NewCopyEntryCommand creates a CopyCommand
//object that copies the first entry found
//regardless of the list it is in.
func NewCopyEntryCommand(entry string) *CopyCommand {
	command := new(CopyCommand)

	command.definition = "goboom copy <entry>"
	command.description = "Copies the first found entry to the clipboard."
	command.Entry = entry

	return command
}

//NewCopySpecificEntryCommand creates a CopyCommand
//object that copies the entry found in the
//specified list.
func NewCopySpecificEntryCommand(list string, entry string) *CopyCommand {
	command := new(CopyCommand)

	command.definition = "goboom copy <list> <entry>"
	command.description = "Copies the specified list entry to the clipboard."
	command.List = list
	command.Entry = entry

	return command
}

//GetDescription retrieves a description of what the command does.
func (c *CopyCommand) GetDescription() string {
	return c.description
}

//GetDefinition retrieves definition to use for this command.
func (c *CopyCommand) GetDefinition() string {
	return c.definition
}

//Execute runs the current command.
func (c *CopyCommand) Execute(w io.Writer, s Storage) {

	copy := exec.Command("pbcopy")
	in, err := copy.StdinPipe()

	if err != nil {
		panic(err)
	}

	for _, list := range s.Lists {
		for key, entries := range list {
			if key == c.List || c.List == "" {
				for _, entry := range entries {
					for key, value := range entry {
						if key == c.Entry {
							//fmt.Fprint(w, value+"\n")
							if err := copy.Start(); err != nil {
								panic(err)
							}
							if _, err := in.Write([]byte(value)); err != nil {
								panic(err)
							}
							if err := in.Close(); err != nil {
								panic(err)
							}
						}
					}
				}
			}
		}
	}
}
