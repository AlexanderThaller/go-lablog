package formatting

import (
	"io"

	"github.com/AlexanderThaller/lablog/src/data"
)

func Todos(writer io.Writer, todos []data.Todo) {
	var printedtodos bool

	for _, todo := range todos {
		if todo.Active {
			io.WriteString(writer, "* "+todo.Value+"\n")
			printedtodos = true
		}
	}

	if printedtodos {
		io.WriteString(writer, "\n")
	}
}

func ProjectTodos(writer io.Writer, indent int, project *data.Project) {
	if len(project.Todos()) == 0 {
		return
	}

	HeaderProject(writer, indent+1, project)
	Todos(writer, project.Todos())
}
