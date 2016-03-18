package data

import (
	"time"

	"github.com/juju/errgo"
)

type Note struct {
	Value     string
	TimeStamp time.Time
}

func (note Note) Type() EntryType {
	return EntryTypeNote
}

func (note Note) Values() []string {
	return []string{
		note.Type().String(),
		note.TimeStamp.Format(TimeStampFormat),
		note.Value,
	}
}

func (note Note) GetTimeStamp() time.Time {
	return note.TimeStamp
}

func ParseNote(values []string) (Note, error) {
	if len(values) != 3 {
		return Note{}, errgo.New("entry with the type note needs exactly three fields")
	}

	etype, err := ParseEntryType(values[0])
	if err != nil {
		return Note{}, errgo.Notef(err, "can not parse entry type")
	}
	if etype != EntryTypeNote {
		return Note{}, errgo.New("tried to parse a note but got the entry type " + etype.String())
	}

	timestamp, err := time.Parse(TimeStampFormat, values[1])
	if err != nil {
		return Note{}, errgo.Notef(err, "can not parse timestamp")
	}

	return Note{TimeStamp: timestamp, Value: values[2]}, nil
}
