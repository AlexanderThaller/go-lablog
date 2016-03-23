package formatting

import (
	"io"
	"regexp"

	"github.com/AlexanderThaller/lablog/src/data"
)

const (
	HeaderTimeFormat = "2006-01-02 15:04:05"
)

func Notes(writer io.Writer, indent int, notes []data.Note) {
	for _, note := range notes {
		if note.Value == "" {
			continue
		}

		io.WriteString(writer, HeaderIndent(indent)+" ")
		io.WriteString(writer, note.TimeStamp.Format(HeaderTimeFormat)+"\n")

		NotesValue(writer, note.Value, indent+1)
		io.WriteString(writer, "\n")
	}
}

func NotesValue(writer io.Writer, value string, indent int) {
	indentchar := HeaderIndent(indent)
	indentreg, _ := regexp.Compile("(?m)^=")
	value = indentreg.ReplaceAllString(value, indentchar)

	io.WriteString(writer, value)
	io.WriteString(writer, "\n")
}

func ProjectNotes(writer io.Writer, indent int, project *data.Project) {
	if len(project.Notes()) == 0 {
		return
	}

	HeaderProject(writer, indent+1, project)
	Notes(writer, indent+2, project.Notes())
}
