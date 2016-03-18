package data

import (
	"strconv"
	"time"

	"github.com/juju/errgo"
)

type Todo struct {
	Active    bool
	TimeStamp time.Time
	Value     string
}

func (todo Todo) Type() EntryType {
	return EntryTypeTodo
}

func (todo Todo) Values() []string {
	return []string{
		todo.Type().String(),
		todo.TimeStamp.Format(TimeStampFormat),
		strconv.FormatBool(todo.Active),
		todo.Value,
	}
}

func (todo Todo) GetTimeStamp() time.Time {
	return todo.TimeStamp
}

func ParseTodo(values []string) (Todo, error) {
	if len(values) != 4 {
		return Todo{}, errgo.New("entry with the type todo needs exactly four fields")
	}

	etype, err := ParseEntryType(values[0])
	if err != nil {
		return Todo{}, errgo.Notef(err, "can not parse entry type")
	}
	if etype != EntryTypeTodo {
		return Todo{}, errgo.New("tried to parse a todo but got the entry type " + etype.String())
	}

	timestamp, err := time.Parse(TimeStampFormat, values[1])
	if err != nil {
		return Todo{}, errgo.Notef(err, "can not parse timestamp")
	}

	active, err := strconv.ParseBool(values[2])
	if err != nil {
		return Todo{}, errgo.Notef(err, "can not parse active state")
	}

	return Todo{Active: active, TimeStamp: timestamp, Value: values[3]}, nil
}
