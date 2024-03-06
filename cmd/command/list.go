package command

import (
	"flag"
	"fmt"
	"os"

	todo "github.com/Tecu23/todo-cli/internal"
)

var listUsage = `List all currently available templates.

Usage: todo-cli list

Options:
`

var listFunc = func(cmd *Command, args []string) {
	todos := todo.Todos{}

	err := todos.Load(todoFile)
	if err != nil {
		cmd.flags.Usage()
	}

	todos.Print()

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
