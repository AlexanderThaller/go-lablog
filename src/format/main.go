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

func Projects(writer io.Writer, projects []data.Project) error {
	io.WriteString(writer, AsciiDocSettings+"\n\n")

	io.WriteString(writer, "= Notes\n\n")

	for _, project := range projects {
		notes, err := project.Notes()
		if err != nil {
			return errgo.Notef(err, "can not get notes from project "+project.Name)
		}

		io.WriteString(writer, project.Format(1))

		sort.Sort(data.NotesByTimeStamp(notes))
		for _, note := range notes {
			io.WriteString(writer, note.Format(1))
		}
	}

	return nil
}
