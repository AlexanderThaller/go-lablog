package main

import (
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
	if len(values) < 4 {
		return nil, errgo.New("we need at least 4 fields for parsing")
	}

	recordtype := values[1]
	switch recordtype {
	case ActionNote:
		return NoteFromCSV(values)
	default:
		return nil, errgo.New("can not parse record type " + recordtype)
	}
}

func NoteFromCSV(values []string) (Note, error) {
	note := new(Note)
	timestamp, err := time.Parse(RecordTimeStampFormat, values[0])
	if err != nil {
		return Note{}, err
	}

	note.TimeStamp = timestamp
	note.Project = values[2]
	note.Value = values[3]

	return *note, nil
}

type Note struct {
	Project   string
	TimeStamp time.Time
	Value     string
}

func (note Note) CSV() []string {
	return []string{
		note.TimeStamp.Format(RecordTimeStampFormat),
		ActionNote,
		note.Project,
		note.Value,
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
