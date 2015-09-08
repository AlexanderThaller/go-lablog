package data

import "time"

type Note struct {
	Value     string
	TimeStamp time.Time
}

func (note Note) Type() EntryType {
	return EntryTypeNote
}

func (note Note) Values() []string {
	return []string{
		note.TimeStamp.Format(TimeStampFormat),
		note.Type().String(),
		note.Value,
	}
}
