package command

import (
	"flag"
	"fmt"
	"os"

	todo "github.com/Tecu23/todo-cli/internal"
)

var addUsage = `Add one or more todos to the list

Usage: todo-cli add TODO_NAME1 TODO_NAME2 ...

Options:
`

var addFunc = func(cmd *Command, args []string) {
	if cmd.flags.NArg() == 0 {
		cmd.flags.Usage()
	}

	todos := todo.Todos{}

	err := todos.Load(todoFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, a := range args {
		todos.Add(a)
	}

	err = todos.Store(todoFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Exit(0)
}

// AddTodos adds todos to the list and saves them in the file
func AddTodos() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("add", flag.ExitOnError),
		Execute: addFunc,
	}

	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, addUsage)
	}

	return cmd
}
