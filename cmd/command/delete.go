package command

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	todo "github.com/Tecu23/todo-cli/internal"
)

var deleteUsage = `Delete a todo by id

Usage: todo-cli delete TODO_ID

Options:
`

var deleteFunc = func(cmd *Command, args []string) {
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

	err = todos.Delete(id)
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
