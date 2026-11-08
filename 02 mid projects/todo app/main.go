package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title string
	Completed bool
	CreatedAt time.Time
	CompletedAt time.Time
}

type Todos []Todo


func (todos *Todos) add(Title string) {
	todo := Todo{
		Title: Title,
		Completed: false,
		CompletedAt: time.Time{},
		CreatedAt: time.Now(),
	}
	*todos = append(*todos, todo)
}



func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("invilaid index")
		fmt.Println(err.Error())
		return err
	}
	return nil
}




func (todos *Todos) delete(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}
	*todos = append(t[:index], t[index+1:]...)
	return nil
}




func (todos *Todos) toggle(index int) error {
	t := todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	item := &(*todos)[index]
	if item.Completed {
		item.Completed = false
		item.CompletedAt = time.Time{} 
	} else {
		item.Completed = true
		item.CompletedAt = time.Now()
	}
	return nil
}

func (todos *Todos) edit(index int, title string) error {
t := *todos

if err := t.validateIndex(index); err != nil {
    return err
}

t[index].Title = title

return nil
}



func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created at", "Completed at")
	
	for index, t := range *todos {
		complated := "❌"
		complatedat := ""
		if t.Completed {
			complated = "✅"
			if !t.CompletedAt.IsZero() {
				complatedat = t.CompletedAt.Format(time.RFC1123)
			}
		}
		table.AddRow(strconv.Itoa(index), t.Title, complated, t.CreatedAt.Format(time.RFC1123), complatedat)

	}
	table.Render()
}
func main() {
	fmt.Println("welcome todo app \n by hackpatato")
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	CmdFlags := NewCmdFlags()
	CmdFlags.Execute(&todos)
	storage.Save(todos)
}