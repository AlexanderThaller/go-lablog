package formatting

import (
	"io"

	"github.com/AlexanderThaller/lablog/src/data"
)

func Project(writer io.Writer, indent int, project *data.Project) {
	HeaderProject(writer, indent+1, project)
	HeaderTodos(writer, indent+2)
	Todos(writer, project.Todos())

	HeaderNotes(writer, indent+2)
	Notes(writer, indent+3, project.Notes())
}

func Projects(writer io.Writer, command string, indent int, projects *data.Projects) {
	HeaderSettings(writer)
	HeaderProjects(writer, command, indent+1, projects)
	for _, project := range projects.List() {
		Project(writer, indent+1, &project)
	}
}

func ProjectsNotes(writer io.Writer, command string, indent int, projects *data.Projects) {
	HeaderSettings(writer)
	HeaderProjects(writer, command, indent+1, projects)
	for _, project := range projects.List() {
		ProjectNotes(writer, indent+1, &project)
	}
}

func ProjectsTodos(writer io.Writer, command string, indent int, projects *data.Projects) {
	HeaderSettings(writer)
	HeaderProjects(writer, command, indent+1, projects)
	for _, project := range projects.List() {
		ProjectTodos(writer, indent+1, &project)
	}
}
