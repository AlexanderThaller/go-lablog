package formatting

import (
	"io"

	"github.com/AlexanderThaller/lablog/src/data"
)

func Todos(writer io.Writer, todos []data.Todo) {
	for _, todo := range todos {
		io.WriteString(writer, "* "+todo.Value+"\n")
	}
}

func Notes(writer io.Writer, notes []data.Note) {
	for _, note := range notes {
		io.WriteString(writer, "== "+note.TimeStamp.String()+"\n")
		io.WriteString(writer, note.Value+"\n")
	}
}

func ProjectTodos(writer io.Writer, indent int, project data.Project) {
	ProjectHeader(writer, indent+1, project.Name)
	Todos(writer, project.Todos())
}

func ProjectNotes(writer io.Writer, indent int, project data.Project) {
	ProjectHeader(writer, indent+1, project.Name)
	Notes(writer, project.Notes())
}

func ProjectHeader(writer io.Writer, indent int, name data.ProjectName) {
	io.WriteString(writer, HeaderIndent(indent)+" "+name.String()+"\n")
}

func HeaderIndent(indent int) string {
	var out string

	for i := 0; i < indent; i++ {
		out += "="
	}

	return out
}
