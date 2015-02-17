package format

import (
	"io"
	"sort"

	"github.com/AlexanderThaller/lablog/src/data"
	"github.com/juju/errgo"
)

const AsciiDocSettings = `:toc: right
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

func ProjectsNotes(writer io.Writer, projects []data.Project) error {
	io.WriteString(writer, AsciiDocSettings+"\n\n")

	for _, project := range projects {
		notes, err := project.Notes()
		if err != nil {
			return errgo.Notef(err, "can not get notes from project "+project.Name)
		}

		if len(notes) == 0 {
			continue
		}

		project.Format(writer, 1)
		Notes(writer, notes)
	}

	return nil
}

func ProjectsTodos(writer io.Writer, projects []data.Project) error {
	io.WriteString(writer, AsciiDocSettings+"\n\n")

	for _, project := range projects {
		todos, err := project.Todos()
		if err != nil {
			return errgo.Notef(err, "can not get todos from project "+project.Name)
		}

		if len(todos) == 0 {
			continue
		}

		project.Format(writer, 1)
		Todos(writer, todos)
	}

	return nil
}

func Todos(writer io.Writer, todos []data.Todo) {
	io.WriteString(writer, "=== Todos\n\n")

	sort.Sort(data.TodosByName(todos))
	for _, todo := range todos {
		todo.Format(writer, 2)
	}
}

func Notes(writer io.Writer, notes []data.Note) {
	io.WriteString(writer, "=== Notes\n\n")

	sort.Sort(data.NotesByTimeStamp(notes))
	for _, note := range notes {
		note.Format(writer, 2)
	}
}
