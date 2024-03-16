package command

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	todo "github.com/Tecu23/todo-cli/internal"
)

var editUsage = `Edit a todo name using the todo id

Usage: todo-cli edit ID NAME

Options:
`

var editFunc = func(cmd *Command, args []string) {
	if cmd.flags.NArg() == 0 {
		cmd.flags.Usage()
		os.Exit(0)
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

	err = todos.Edit(args[1], id)
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
