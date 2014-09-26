package main

import "time"

const (
	RecordTimeStampFormat = time.RFC3339Nano
)

type Record interface {
	CSV() []string
}

type Note struct {
	Project   string
	TimeStamp time.Time
	Value     string
}

func (note *Note) CSV() []string {
	return []string{
		note.TimeStamp.Format(RecordTimeStampFormat),
		ActionNote,
		note.Project,
		note.Value,
	}
}
