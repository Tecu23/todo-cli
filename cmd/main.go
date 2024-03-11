// main package
package main

import (
	"flag"
	"fmt"
	"os"

	command "github.com/Tecu23/todo-cli/cmd/command"
)

var usage = `Usage: todo-cli command [options]

A simple tool to manage custom todos from the terminal

Options:

Commands:
  add		Adds one or more todos to the todo list
  edit		Edits the name or the category of the todo
  list		Lists all todos stored
  complete	Marks a todo as completed
  delete	Removes a stored todo
  version	Prints version info to console
`

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprint(usage))
	}

	flag.Parse()
	if flag.NArg() < 1 {
		usageAndExit("")
	}

	var cmd *command.Command

	switch os.Args[1] {
	case "version":
		cmd = command.PrintVersion()
	case "add":
		cmd = command.AddTodos()
	case "edit":
		cmd = command.EditTodo()
	case "list":
		cmd = command.ListTodos()
	case "complete":
		cmd = command.CompleteTodo()
	case "delete":
		cmd = command.DeleteTodo()
	default:
		usageAndExit(fmt.Sprintf("gupi: %s is not a gupi command.\n", os.Args[1]))
	}

	cmd.Init(os.Args[2:])
	cmd.Run()
}

func usageAndExit(msg string) {
	if msg != "" {
		fmt.Fprint(os.Stderr, msg)
		fmt.Fprintf(os.Stderr, "\n")
	}

	flag.Usage()
	os.Exit(0)
}
