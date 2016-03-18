package data

import (
	"time"

	"github.com/juju/errgo"
)

type Entries []Entry

type Entry interface {
	Type() EntryType
	Values() []string
	GetTimeStamp() time.Time
}

const TimeStampFormat = time.RFC3339Nano

type EntryType int

const (
	EntryTypeNote EntryType = iota
	EntryTypeTodo
	EntryTypeUnkown
)

func (etype EntryType) String() string {
	switch etype {
	case EntryTypeNote:
		return "note"
	case EntryTypeTodo:
		return "todo"
	default:
		return "unkown"
	}
}

func ParseEntryType(value string) (EntryType, error) {
	switch value {
	case "note":
		return EntryTypeNote, nil
	case "todo":
		return EntryTypeTodo, nil
	default:
		return EntryTypeUnkown, errgo.New("the entry type " + value + " is not known")
	}
}

func ParseEntry(values []string) (Entry, error) {
	if len(values) < 1 {
		return nil, errgo.New("entry values need at least one field")
	}

	etype, err := ParseEntryType(values[0])
	if err != nil {
		return nil, errgo.Notef(err, "can not parse entry type")
	}

	switch etype {
	case EntryTypeNote:
		return ParseNote(values)
	case EntryTypeTodo:
		return ParseTodo(values)
	default:
		return nil, errgo.New("do not know how to parse this entry type")
	}
}
