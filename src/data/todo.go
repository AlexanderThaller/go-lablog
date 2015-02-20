package data

import (
	"io"
	"sort"
	"strconv"
	"time"

	"github.com/AlexanderThaller/logger"
	"github.com/jinzhu/now"
	"github.com/juju/errgo"
)

type Todo struct {
	Done bool
	Project
	Text      string
	TimeStamp time.Time
}

func (todo Todo) ValueArray() []string {
	return []string{
		todo.TimeStamp.Format(EntryCSVTimeStampFormat),
		"todo",
		todo.Text,
		strconv.FormatBool(todo.Done),
	}
}

func (todo Todo) GetProject() Project {
	return todo.Project
}

func (todo Todo) Type() string {
	return "todo"
}

func (todo Todo) GetTimeStamp() time.Time {
	return todo.TimeStamp
}

func (todo Todo) Format(writer io.Writer, indent uint) {
	io.WriteString(writer, "* "+todo.Text)
	io.WriteString(writer, "\n")
}

// TodosByName allows sorting todo slices by name.
type TodosByName []Todo

func (by TodosByName) Len() int {
	return len(by)
}

func (by TodosByName) Swap(i, j int) {
	by[i], by[j] = by[j], by[i]
}

func (by TodosByName) Less(i, j int) bool {
	return by[i].Text < by[j].Text
}

// TodosByTimeStamp allows sorting todo slices by timestamp.
type TodosByTimeStamp []Todo

func (by TodosByTimeStamp) Len() int {
	return len(by)
}

func (by TodosByTimeStamp) Swap(i, j int) {
	by[i], by[j] = by[j], by[i]
}

func (by TodosByTimeStamp) Less(i, j int) bool {
	return by[i].TimeStamp.Before(by[j].TimeStamp)
}

func ParseTodo(project Project, values []string) (Todo, error) {
	timestamp, err := now.Parse(values[0])
	if err != nil {
		return Todo{}, errgo.Notef(err, "can not parse timestamp")
	}

	inactive, err := strconv.ParseBool(values[3])
	if err != nil {
		return Todo{}, errgo.Notef(err, "can not parse active status")
	}

	todo := Todo{
		Done:      inactive,
		Project:   project,
		Text:      values[2],
		TimeStamp: timestamp,
	}

	return todo, nil
}

func FilterTodosLatest(todos []Todo) []Todo {
	l := logger.New(Name, "todo", "FilterTodosLatest")

	sort.Sort(TodosByTimeStamp(todos))

	filter := make(map[string]Todo)
	for _, todo := range todos {
		l.Trace("Todo: ", todo)
		filter[todo.Text] = todo
	}

	l.Trace("Filter: ", filter)

	var out []Todo
	for _, todo := range filter {
		out = append(out, todo)
	}

	return out
}

func FilterTodosAreNotDone(todos []Todo) []Todo {
	var out []Todo

	for _, todo := range todos {
		if todo.Done {
			continue
		}

		out = append(out, todo)
	}

	return out
}

func FilterTodosBeforeTimeStamp(todos []Todo, start time.Time) []Todo {
	var out []Todo

	for _, todo := range todos {
		if todo.TimeStamp.Before(start) {
			continue
		}

		out = append(out, todo)
	}

	return out
}

func FilterTodosAfterTimeStamp(todos []Todo, end time.Time) []Todo {
	var out []Todo

	for _, todo := range todos {
		if todo.TimeStamp.After(end) {
			continue
		}

		out = append(out, todo)
	}

	return out
}
