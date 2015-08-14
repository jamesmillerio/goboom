package main

import (
	"flag"
	"os"
	"os/user"
	"text/tabwriter"
)

//Indent defines the indentation for printing across classes.
const defaultIndention = "  "

func main() {

	user, error := user.Current()

	if error != nil {
		panic(error)
	}

	var command Command

	//Define some common flags.
	data := flag.String("data", user.HomeDir+"/.boom", "Defines the path to the desired .boom file. Defaults to ~/.boom if not specified.")
	help := flag.Bool("help", false, "Displys usage information for goboom.")

	//Grab any defined flags/args.
	flag.Parse()

	//Create our tab writer so everything lines up.
	writer := new(tabwriter.Writer)

	//Get a reference to our storage instance.
	storage := NewCustomListStorage(*data)

	//Initialize our tab writer values.
	writer.Init(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)

	//If they want help information, process it and return.
	if *help {

		//Get all available commands.
		commands := getAllCommands()

		//Get our help command.
		command = NewHelpCommand(commands)

	} else {

		//Get the specified command.
		command = CreateCommand(flag.Args())

	}

	//Execute our command.
	command.Execute(writer, storage)

	//Flush any output.
	writer.Flush()

}

//getAllCommands retrieves the commands
//for use in displaying help information.
func getAllCommands() []Command {

	var commands []Command

	commands = append(commands, new(OverviewCommand))
	commands = append(commands, NewAllCommand())
	commands = append(commands, nil)
	commands = append(commands, NewCreateListCommand(""))
	commands = append(commands, NewAllListCommand(""))
	commands = append(commands, NewDeleteListCommand(""))
	commands = append(commands, nil)
	commands = append(commands, NewCreateEntryCommand("", "", ""))
	commands = append(commands, NewCopyEntryCommand(""))
	commands = append(commands, NewCopySpecificEntryCommand("", ""))
	commands = append(commands, NewOpenListCommand(""))
	commands = append(commands, NewOpenListEntryCommand("", ""))
	commands = append(commands, NewRandomEntryCommand())
	commands = append(commands, NewRandomSpecificEntryCommand(""))
	commands = append(commands, NewEchoEntryCommand(""))
	commands = append(commands, NewEchoSpecificEntryCommand("", ""))
	commands = append(commands, NewDeleteListEntryCommand("", ""))

	return commands
}
