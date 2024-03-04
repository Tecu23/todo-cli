package command

import (
	"flag"
	"fmt"
	"os"
)

var completeUsage = `Complete a todo by specifying the id

Usage: todo-cli complete TODO_ID

Options:
`

var completeFunc = func(cmd *Command, args []string) {
	if cmd.flags.NArg() == 0 {
		cmd.flags.Usage()
	}

	// Find todo by id

	// Mark todo as completed

	// Store new todos

	os.Exit(0)
}

// CompleteTodo comples a todo
func CompleteTodo() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("complete", flag.ExitOnError),
		Execute: completeFunc,
	}

	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, completeUsage)
	}

	return cmd
}
