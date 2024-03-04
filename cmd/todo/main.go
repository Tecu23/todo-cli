// main package
package main

import (
	"flag"
	"fmt"
	"os"

	command "github.com/Tecu23/todo-cli/cmd/commands"
)

const (
	todoFile = ".todos.json"
)

var usage = `Usage: gupi command [options]

A simple tool to generate and manage custom templates

Options:

Commands:
  add		Adds a template to the collection from a local file
  edit		Uses the default text editor to modify a stored template
  list		Lists all stored templates
  create	Generates an instance of a template in the current directory
  delete	Removes a stored template
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
		cmd = command.NewVersionCommand()
	case "add":
		cmd = command.NewAddCommand()
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
