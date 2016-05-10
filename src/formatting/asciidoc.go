package formatting

import (
	"io"
	"regexp"

	"github.com/AlexanderThaller/lablog/src/data"
)

const AsciidocHeaderChar = "="

func AsciidocHeaderSettings(writer io.Writer) {
	settings := `:toc: right
:toclevels: 2
:sectanchors:
:sectlink:
:icons: font
:linkattrs:
:numbered:
:idprefix:
:idseparator: -
:doctype: book
:source-highlighter: pygments
:listing-caption: Listing`

	io.WriteString(writer, settings+"\n\n")
}

func AsciidocHeader(indent int) string {
	return HeaderIndent(AsciidocHeaderChar, indent)
}

func AsciidocHeaderProjects(writer io.Writer, command string, indent int, project *data.Projects) {
	io.WriteString(writer, AsciidocHeader(indent)+" "+command+"\n\n")
}

func AsciidocHeaderProject(writer io.Writer, indent int, project *data.Project) {
	io.WriteString(writer, AsciidocHeader(indent)+" "+project.Name.String()+"\n")
}

func AsciidocHeaderTodos(writer io.Writer, indent int) {
	io.WriteString(writer, AsciidocHeader(indent)+" Todos\n")
}

func AsciidocHeaderNotes(writer io.Writer, indent int) {
	io.WriteString(writer, AsciidocHeader(indent)+" Notes\n")
}

func AsciidocNotes(writer io.Writer, indent int, notes []data.Note) {
	for _, note := range notes {
		if note.Value == "" {
			continue
		}

		io.WriteString(writer, AsciidocHeader(indent)+" ")
		io.WriteString(writer, note.TimeStamp.String()+"\n")

		AsciidocNotesValue(writer, note.Value, indent+1)
		io.WriteString(writer, "\n")
	}
}

func AsciidocNotesValue(writer io.Writer, value string, indent int) {
	indentchar := AsciidocHeader(indent)
	indentreg, _ := regexp.Compile("(?m)^=")
	value = indentreg.ReplaceAllString(value, indentchar)

	io.WriteString(writer, value)
	io.WriteString(writer, "\n")
}

func AsciidocProjectNotes(writer io.Writer, indent int, project *data.Project) {
	if len(project.Notes()) == 0 {
		return
	}

	AsciidocHeaderProject(writer, indent+1, project)
	AsciidocNotes(writer, indent+2, project.Notes())
}

func AsciidocProject(writer io.Writer, indent int, project *data.Project) {
	todos := project.Todos()
	notes := project.Notes()

	if len(todos) == 0 && len(notes) == 0 {
		return
	}

	AsciidocHeaderProject(writer, indent+1, project)

	if len(todos) != 0 {
		AsciidocHeaderTodos(writer, indent+2)
		AsciidocTodos(writer, todos)
	}

	if len(notes) != 0 {
		AsciidocHeaderNotes(writer, indent+2)
		AsciidocNotes(writer, indent+3, notes)
	}
}

func AsciidocProjects(writer io.Writer, command string, indent int, projects *data.Projects) {
	AsciidocHeaderSettings(writer)
	AsciidocHeaderProjects(writer, command, indent+1, projects)
	for _, project := range projects.List() {
		AsciidocProject(writer, indent+1, &project)
	}
}

func AsciidocProjectsNotes(writer io.Writer, command string, indent int, projects *data.Projects) {
	AsciidocHeaderSettings(writer)
	AsciidocHeaderProjects(writer, command, indent+1, projects)
	for _, project := range projects.List() {
		AsciidocProjectNotes(writer, indent+1, &project)
	}
}

func AsciidocProjectsTodos(writer io.Writer, command string, indent int, projects *data.Projects) {
	AsciidocHeaderSettings(writer)
	AsciidocHeaderProjects(writer, command, indent+1, projects)
	for _, project := range projects.List() {
		AsciidocProjectTodos(writer, indent+1, &project)
	}
}

func AsciidocTodos(writer io.Writer, todos []data.Todo) {
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

func AsciidocProjectTodos(writer io.Writer, indent int, project *data.Project) {
	if len(project.Todos()) == 0 {
		return
	}

	AsciidocHeaderProject(writer, indent+1, project)
	AsciidocTodos(writer, project.Todos())
}
