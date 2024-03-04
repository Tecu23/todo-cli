package command

import (
	"flag"
	"fmt"
	"os"
)

var listUsage = `List all currently available templates.

Usage: todo-cli list

Options:
`

var listFunc = func(cmd *Command, args []string) {
	if cmd.flags.NArg() == 0 {
		cmd.flags.Usage()
	}

	// Print all todos in a nice format

	os.Exit(1)
}

// ListTodos prints all todos in order
func ListTodos() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("list", flag.ExitOnError),
		Execute: listFunc,
	}

	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, listUsage)
	}

	return cmd
}
