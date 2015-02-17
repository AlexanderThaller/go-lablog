package data

import (
	"io"
	"regexp"
	"strings"
	"time"
)

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

func (note Note) Format(writer io.Writer, indent uint) {
	indentchar := strings.Repeat("=", int(indent))
	reg, _ := regexp.Compile("(?m)^=")

	io.WriteString(writer, indentchar+"== "+note.TimeStamp.Format(EntryCSVTimeStampFormat)+"\n")
	io.WriteString(writer, reg.ReplaceAllString(note.Text, indentchar+"==="))
	io.WriteString(writer, "\n\n")
}

// NotesByTimeStamp allows sorting project slices by name.
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
