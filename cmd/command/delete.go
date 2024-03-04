package command

import (
	"flag"
	"fmt"
	"os"
)

var deleteUsage = `Delete a todo by id

Usage: todo-cli delete TODO_ID

Options:
`

var deleteFunc = func(cmd *Command, args []string) {
	if cmd.flags.NArg() == 0 {
		cmd.flags.Usage()
	}

	// Find todo by id

	// Delete todo

	// Save todo list

	os.Exit(0)
}

// DeleteTodo deletes todo
func DeleteTodo() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("delete", flag.ExitOnError),
		Execute: deleteFunc,
	}

	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, deleteUsage)
	}

	return cmd
}
