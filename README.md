# TODO-CLI

A simple command-line todo application written in Go.

## Table of Contents

- [Installation](#installation)
- [Uninstallation](#uninstallation)
- [Usage](#usage)
  - [Add](#add)
  - [List](#list)
  - [Complete](#complete)
  - [Edit](#edit)
  - [Delete](#delete)

## Installation

To install the Todo CLI, you need to have [Go](https://go.dev/) installed on your
system. To install it, follow the [Go installation instructions](https://go.dev/doc/install).
Then, you can run the following command:

```bash
    git clone git@github.com:Tecu23/todo-cli.git
```

Then run the following commands:

```bash
    make build
    make install
```

## Uninstallation

To uninstall the Todo CLI, you can simply delete the binary file from your Go
bin directory. Or simply run the following commands:

```bash
    make uninstall
```

## Usage

### Add

To add a new todo, use the following command:

```bash
todo add "Task description"
```

Replace `"Task description"` with the description of the task you want to add.

### List

To list all todos, use the following command:

```bash
todo list
```

This will display all the todos along with their status (completed or pending).

### Complete

To mark a todo as completed, use the following command:

```bash
todo complete <todo_id>
```

Replace `<todo_id>` with the ID of the todo you want to mark as completed. You can find the ID by listing all todos.

### Edit

To edit the name of a todo, use the following command:

```bash
todo edit <todo_id> "New task description"
```

Replace `<todo_id>` with the ID of the todo you want to edit and `"New task description"` with the updated description.

### Delete

To delete a todo, use the following command:

```bash
todo delete <todo_id>
```

Replace `<todo_id>` with the ID of the todo you want to delete. Be careful, this action is irreversible.
