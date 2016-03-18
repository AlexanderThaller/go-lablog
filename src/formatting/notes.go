package formatting

import (
	"io"
	"regexp"
	"strings"

	"github.com/AlexanderThaller/lablog/src/data"
)

func Notes(writer io.Writer, indent int, notes []data.Note) {
	for _, note := range notes {
		io.WriteString(writer, HeaderIndent(indent)+" ")
		io.WriteString(writer, note.TimeStamp.String()+"\n")

		NotesValue(writer, note.Value, indent+1)
		io.WriteString(writer, "\n")
	}
}

func NotesValue(writer io.Writer, value string, indent int) {
	indentchar := strings.Repeat("=", int(indent))
	indentreg, _ := regexp.Compile("(?m)^=")
	value = indentreg.ReplaceAllString(value, indentchar)

	leadinghashchar := `\#`
	leadinghashreg, _ := regexp.Compile("(?m)^#")
	value = leadinghashreg.ReplaceAllString(value, leadinghashchar)

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
