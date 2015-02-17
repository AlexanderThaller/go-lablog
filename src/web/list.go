package web

import (
	"bytes"
	"html/template"
	"net/http"
	"sort"

	"github.com/AlexanderThaller/lablog/src/data"
	"github.com/AlexanderThaller/lablog/src/format"
	"github.com/AlexanderThaller/logger"
	"github.com/juju/errgo"
)

func listProjects(w http.ResponseWriter, r *http.Request) {
	l := logger.New(Name, "listProjects")

	projects, err := data.Projects(_datadir)
	if err != nil {
		printerr(l, w, errgo.Notef(err, "can not get projects"))
		return
	}

	rawtmpl, err := html_listprojects_html_bytes()
	if err != nil {
		printerr(l, w, errgo.Notef(err, "can not get projects html template"))
		return
	}

	sort.Sort(data.ProjectsByName(projects))
	tmpl, err := template.New("name").Parse(string(rawtmpl))
	if err != nil {
		printerr(l, w, errgo.Notef(err, "can not parse projects html template"))
		return
	}

	l.Info("Serving project list")
	err = tmpl.Execute(w, projects)
	if err != nil {
		printerr(l, w, errgo.Notef(err, "can not execute projects html template"))
		return
	}
}

func listNotes(w http.ResponseWriter, r *http.Request) {
	l := logger.New(Name, "listProjects")

	args := defquery(r, "project")
	projects, err := data.ProjectsOrArgs(args, _datadir)
	if err != nil {
		printerr(l, w, errgo.Notef(err, "can not get projects or args"))
		return
	}

	sort.Sort(data.ProjectsByName(projects))

	var buffer = new(bytes.Buffer)
	err = format.ProjectsNotes(buffer, projects)
	if err != nil {
		printerr(l, w, errgo.Notef(err, "can not format projects"))
		return
	}

	err = format.AsciiDoctor(buffer, w)
	if err != nil {
		printerr(l, w, errgo.Notef(err, "can not format with asciidoctor"))
		return
	}
}

func listTodos(w http.ResponseWriter, r *http.Request) {
	l := logger.New(Name, "listProjects")

	args := defquery(r, "project")
	projects, err := data.ProjectsOrArgs(args, _datadir)
	if err != nil {
		printerr(l, w, errgo.Notef(err, "can not get projects or args"))
		return
	}

	sort.Sort(data.ProjectsByName(projects))
	err = format.ProjectsNotes(w, projects)
	if err != nil {
		printerr(l, w, errgo.Notef(err, "can not format projects"))
		return
	}
}
