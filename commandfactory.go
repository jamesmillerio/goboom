package main

import "strings"

//CreateCommand is a factory method that returns the command that
//is infered from the passed command line arguments.
func CreateCommand(args []string, translator Translator) Command {

	if args == nil || len(args) == 0 {
		return new(OverviewCommand)
	}

	count := len(args)
	command := strings.ToLower(args[0])

	switch command {
	case "all":
		return NewAllCommand(translator)
	case "delete":
		if count == 2 { //Delete list
			return NewDeleteListCommand(args[1])
		} else if count == 3 { //Delete list item
			return NewDeleteListEntryCommand(args[1], args[2])
		} else { //Do nothing.
			return nil
		}
	case "open":
		if count == 2 { //Open all in a list.
			return NewOpenListCommand(args[1])
		} else if count == 3 { //Open a specific entry.
			return NewOpenListEntryCommand(args[1], args[2])
		}
	case "random":
		if count == 1 {
			return NewRandomEntryCommand()
		} else if count == 2 {
			return NewRandomSpecificEntryCommand(args[1])
		}
	case "echo":
		if count == 2 {
			return NewEchoEntryCommand(args[1])
		}
		return NewEchoSpecificEntryCommand(args[1], args[2])
	case "copy":
		if count == 2 {
			return NewCopyEntryCommand(args[1])
		}
		return NewCopySpecificEntryCommand(args[1], args[2])
	default:
		return NewCatchAllCommand(args, translator)
	}

	return nil

}
