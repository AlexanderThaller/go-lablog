package web

import (
	"bytes"
	"fmt"
	"net/http"
	"text/template"

	"github.com/AlexanderThaller/lablog/src/project"
	"github.com/AlexanderThaller/logger"
	"github.com/juju/errgo"
	"gopkg.in/pipe.v2"
)

type Page interface {
	Template() string
}

func WriteTemplateHTML(w http.ResponseWriter, r *http.Request, page Page) {
	template := template.New("templatehtmlpage")
	template.Parse(page.Template())

	buffer := bytes.NewBufferString("")
	err := template.Execute(buffer, page)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", errgo.Details(err))
		return
	}

	output, err := StringToAsciiDoctor(buffer.String())
	if err != nil {
		fmt.Fprintf(w, "Error: %s", errgo.Details(err))
		return
	}

	fmt.Fprintf(w, output)
}

func StringToAsciiDoctor(input string) (string, error) {
	l := logger.New("web", "StringToAsciiDoctor")

	p := pipe.Line(
		pipe.Print(input),
		pipe.Exec("asciidoctor", "-"),
	)

	l.Debug("Reading output")
	output, err := pipe.CombinedOutput(p)
	if err != nil {
		err = errgo.New("problem while getting asciidoc output: " + err.Error() +
			" - " + string(output))

		return "", err
	}

	l.Debug("Returning output")
	return string(output), nil
}

type RootPage struct {
	Projects []string
}

const TemplatePageRoot = `= Lablog -- Projects

[cols="4*", options="header"]
|===
|Project
|Notes
|Todos
|Tracks

{{ range .Projects }}
|{{ . }}
|link:/notes/{{ . }}[Notes]
|link:/todos/{{ . }}[Todos]
|link:/tracks/{{ . }}[Tracks]
{{ end }}
|===`

func (page RootPage) Template() string {
	return TemplatePageRoot
}

type PageNotes struct {
	Project string
	Notes   []project.Note
}

func (page PageNotes) Template() string {
	buffer := bytes.NewBufferString("")

	err := project.FormatNotes(buffer, page.Project, "Notes", page.Notes, 1)
	if err != nil {
		fmt.Fprintf(buffer, "Error: %s", errgo.Details(err))
		return buffer.String()
	}

	return buffer.String()
}
