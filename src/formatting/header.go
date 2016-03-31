package formatting

import (
	"io"
	"strings"

	"github.com/AlexanderThaller/lablog/src/data"
	log "github.com/Sirupsen/logrus"
)

func HeaderIndent(indent int) string {
	log.Debug("Indent: ", indent)
	if indent < 1 {
		return ""
	}

	out := strings.Repeat("=", indent)

	return out
}

func HeaderSettings(writer io.Writer) {
	settings := `:toc: right
:toclevels: 4
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

func HeaderProjects(writer io.Writer, command string, indent int, project *data.Projects) {
	io.WriteString(writer, HeaderIndent(indent)+" "+command+"\n\n")
}

func HeaderProject(writer io.Writer, indent int, project *data.Project) {
	io.WriteString(writer, HeaderIndent(indent)+" "+project.Name.String()+"\n")
}

func HeaderTodos(writer io.Writer, indent int) {
	io.WriteString(writer, HeaderIndent(indent)+" Todos\n")
}

func HeaderNotes(writer io.Writer, indent int) {
	io.WriteString(writer, HeaderIndent(indent)+" Notes\n")
}
