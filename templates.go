package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"
	"text/template"

	"github.com/AlexanderThaller/logger"
	"github.com/juju/errgo"
)

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

type RootPage struct {
	Projects []string
}

func WriteTemplateHTMLRoot(w http.ResponseWriter, r *http.Request, page RootPage) {
	template := template.New("templatehtmlroot")
	template.Parse(TemplatePageRoot)

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
	l := logger.New(Name, "StringToAsciiDoctor")
	l.SetLevel(logger.Debug)

	l.Debug("Starting command")
	command := exec.Command("asciidoctor", "-")

	l.Debug("Opening pipe to command")
	pipe, err := command.StdinPipe()
	if err != nil {
		return "", err
	}

	l.Debug("Writing to pipe")
	_, err = pipe.Write([]byte(input))
	if err != nil {
		return "", err
	}
	pipe.Close()

	l.Debug("Reading output")
	output, err := command.CombinedOutput()
	if err != nil {
		err = errgo.New("problem while getting raw snapshots: " + err.Error() +
			" - " + string(output))

		return "", err
	}

	l.Debug("Returning output")
	return string(output), nil
}
