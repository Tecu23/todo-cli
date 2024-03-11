// Package todo maintains the Todo Item and List data structures
package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
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
func (t *Todos) Complete(idx int) error {
	tmpTodo := *t

	if idx <= 0 || idx > len(tmpTodo) {
		return errors.New("invalid idx")
	}

	fmt.Printf("%v\n", tmpTodo[idx-1])

	tmpTodo[idx-1].CompletedAt = time.Now()
	tmpTodo[idx-1].Done = true

	return nil
}

// Edit -> Edits the name of the todo
func (t *Todos) Edit(task string, idx int) error {
	tmpTodo := *t

	if idx <= 0 || idx > len(tmpTodo) {
		return errors.New("invalid index")
	}

	tmpTodo[idx-1].Task = task

	return nil
}

// Delete -> deletes a todo
func (t *Todos) Delete(idx int) error {
	tmpTodo := *t

	if idx <= 0 || idx > len(tmpTodo) {
		return errors.New("invalid index")
	}

	*t = append(tmpTodo[:idx-1], tmpTodo[idx:]...)

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
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignRight, Text: "CreatedAt"},
			{Align: simpletable.AlignRight, Text: "CompletedAt"},
		},
	}

	var cells [][]*simpletable.Cell

	for idx, item := range *t {
		idx++
		task := blue(item.Task)
		done := blue("no")
		if item.Done {
			task = green(fmt.Sprintf("\u2705 %s", item.Task))
			done = green("yes")
		}
		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: task},
			{Text: done},
			{Text: item.CreatedAt.Format(time.RFC822)},
			{Text: item.CompletedAt.Format(time.RFC822)},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{
			Align: simpletable.AlignCenter,
			Span:  5,
			Text:  red(fmt.Sprintf("You have %d pending todos", t.CountPending())),
		},
	}}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}

// CountPending count the numbers of uncompleted todos
func (t *Todos) CountPending() int {
	total := 0
	for _, item := range *t {
		if !item.Done {
			total++
		}
	}

	return total
}
