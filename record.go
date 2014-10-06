package main

import (
	"strconv"
	"time"

	"github.com/juju/errgo"
)

const (
	RecordTimeStampFormat = time.RFC3339Nano
)

type Record interface {
	CSV() []string
	GetAction() string
	GetProject() string
	GetTimeStamp() string
	GetValue() string
}

func RecordFromCSV(values []string) (Record, error) {
	if len(values) < 2 {
		return nil, errgo.New("we need at least two fields for parsing")
	}

	recordtype := values[1]
	switch recordtype {
	case ActionNote:
		return NoteFromCSV(values)
	case ActionTodo:
		return TodoFromCSV(values)
	default:
		return nil, errgo.New("can not parse record type " + recordtype)
	}
}

func NoteFromCSV(values []string) (Note, error) {
	if len(values) != 3 {
		return Note{}, errgo.New("we need three fields for parsing a note")
	}

	if values[1] != "note" {
		return Note{}, errgo.New("second field has to have the string 'note' in it")
	}

	timestamp, err := time.Parse(RecordTimeStampFormat, values[0])
	if err != nil {
		return Note{}, err
	}

	note := Note{
		TimeStamp: timestamp,
		Value:     values[2],
	}

	return note, nil
}

func TodoFromCSV(values []string) (Todo, error) {
	if len(values) != 4 {
		return Todo{}, errgo.New("we need three fields for parsing a todo")
	}

	if values[1] != "todo" {
		return Todo{}, errgo.New("second field has to have the string 'todo' in it")
	}

	timestamp, err := time.Parse(RecordTimeStampFormat, values[0])
	if err != nil {
		return Todo{}, err
	}
	done, err := strconv.ParseBool(values[3])
	if err != nil {
		return Todo{}, err
	}

	todo := Todo{
		TimeStamp: timestamp,
		Value:     values[2],
		Done:      done,
	}

	return todo, nil
}

type Note struct {
	Project   string
	TimeStamp time.Time
	Value     string
}

func (note Note) CSV() []string {
	return []string{
		note.GetTimeStamp(),
		note.GetAction(),
		note.GetValue(),
	}
}

func (note Note) GetAction() string {
	return ActionNote
}

func (note Note) GetProject() string {
	return note.Project
}

func (note Note) GetTimeStamp() string {
	return note.TimeStamp.Format(RecordTimeStampFormat)
}

func (note Note) GetValue() string {
	return note.Value
}

func (note *Note) SetProject(project string) {
	note.Project = project
}

type Todo struct {
	Project   string
	TimeStamp time.Time
	Value     string
	Done      bool
}

func (todo Todo) CSV() []string {
	return []string{
		todo.GetTimeStamp(),
		todo.GetAction(),
		todo.GetValue(),
		strconv.FormatBool(todo.Done),
	}
}

func (todo Todo) GetAction() string {
	return ActionTodo
}

func (todo Todo) GetProject() string {
	return todo.Project
}

func (todo Todo) GetTimeStamp() string {
	return todo.TimeStamp.Format(RecordTimeStampFormat)
}

func (todo Todo) GetValue() string {
	return todo.Value
}

func (todo Todo) SetProject(project string) {
	todo.Project = project
}

type TodoByDate []Todo

func (todo TodoByDate) Len() int {
	return len(todo)
}

func (todo TodoByDate) Swap(i, j int) {
	todo[i], todo[j] = todo[j], todo[i]
}

func (todo TodoByDate) Less(i, j int) bool {
	return todo[j].TimeStamp.After(todo[i].TimeStamp)
}

type NotesByDate []Note

func (note NotesByDate) Len() int {
	return len(note)
}

func (note NotesByDate) Swap(i, j int) {
	note[i], note[j] = note[j], note[i]
}

func (note NotesByDate) Less(i, j int) bool {
	return note[j].TimeStamp.After(note[i].TimeStamp)
}