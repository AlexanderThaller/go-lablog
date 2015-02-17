package data

import (
	"io"
	"regexp"
	"strings"
	"time"

	"github.com/jinzhu/now"
	"github.com/juju/errgo"
)

type Note struct {
	Project
	Text      string
	TimeStamp time.Time
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

func (note Note) Format(writer io.Writer, indent uint) {
	indentchar := strings.Repeat("=", int(indent))
	reg, _ := regexp.Compile("(?m)^=")

	io.WriteString(writer, indentchar+"== "+note.TimeStamp.Format(EntryCSVTimeStampFormat)+"\n")
	io.WriteString(writer, reg.ReplaceAllString(note.Text, indentchar+"==="))
	io.WriteString(writer, "\n\n")
}

// NotesByTimeStamp allows sorting project slices by timestamp.
type NotesByTimeStamp []Note

func (by NotesByTimeStamp) Len() int {
	return len(by)
}

func (by NotesByTimeStamp) Swap(i, j int) {
	by[i], by[j] = by[j], by[i]
}

func (by NotesByTimeStamp) Less(i, j int) bool {
	return by[i].TimeStamp.Before(by[j].TimeStamp)
}

func ParseNote(project Project, values []string) (Note, error) {
	timestamp, err := now.Parse(values[0])
	if err != nil {
		return Note{}, errgo.Notef(err, "can not parse timestamp from csv")
	}

	note := Note{
		Project:   project,
		Text:      values[2],
		TimeStamp: timestamp,
	}

	return note, nil
}
