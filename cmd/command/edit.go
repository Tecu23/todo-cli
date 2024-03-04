package command

import (
	"flag"
	"fmt"
	"os"
)

var editUsage = `Edit a todo name using the todo id

Usage: todo-cli edit ID NAME

Options:
`

var editFunc = func(cmd *Command, args []string) {
	if cmd.flags.NArg() == 0 {
		cmd.flags.Usage()
	}

	// Find Todo by id

	// Change name of todo

	// Save todos to file
	os.Exit(0)
}

// EditTodo edits a todo
func EditTodo() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("edit", flag.ExitOnError),
		Execute: editFunc,
	}

	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, editUsage)
	}

	return cmd
}
