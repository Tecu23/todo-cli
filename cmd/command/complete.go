package command

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	todo "github.com/Tecu23/todo-cli/internal"
)

var completeUsage = `Complete a todo by specifying the id

Usage: todo-cli complete TODO_ID

Options:
`

var completeFunc = func(cmd *Command, args []string) {
	if cmd.flags.NArg() == 0 {
		cmd.flags.Usage()
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	todos := todo.Todos{}

	err = todos.Load(todoFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = todos.Complete(id)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = todos.Store(todoFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

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
