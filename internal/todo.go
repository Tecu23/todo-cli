// Package todo maintains the Todo Item and List data structures
package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

// Todos is the list of todos
type Todos []item

// Add -> adds a new todo to the list
func (t *Todos) Add(task string) {
	todo := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}

// Complete -> marks a todo as completed
func (t *Todos) Complete(index int) error {
	tmpTodo := *t

	if index <= 0 || index > len(tmpTodo) {
		return errors.New("invalid index")
	}

	tmpTodo[index-1].CompletedAt = time.Now()
	tmpTodo[index-1].Done = true

	return nil
}

// Delete -> deletes a todo
func (t *Todos) Delete(index int) error {
	tmpTodo := *t

	if index <= 0 || index > len(tmpTodo) {
		return errors.New("invalid index")
	}

	*t = append(tmpTodo[:index-1], tmpTodo[index:]...)

	return nil
}

// Load -> load the stored todo data from file
func (t *Todos) Load(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}

		return err
	}

	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}

	return nil
}

// Store -> save the todo data in a file
func (t *Todos) Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

// Print -> displays the todos in order
func (t *Todos) Print() {
	for i, item := range *t {
		fmt.Printf("%d - %s\n", i+1, item.Task)
	}
}
