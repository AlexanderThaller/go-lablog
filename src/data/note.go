package data

import "time"

type Note struct {
	Project
	TimeStamp time.Time
	Text      string
}

func (note Note) ValueArray() []string {
	return []string{
		note.TimeStamp.Format(EntryCSVTimeStampFormat),
		"note",
		note.Text,
	}
}

func (note Note) GetProject() Project {
	return note.Project
}
