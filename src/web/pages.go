package web

import (
	"bytes"
	"net/http"

	"github.com/AlexanderThaller/lablog/src/formatting"
	"github.com/AlexanderThaller/lablog/src/helper"

	"github.com/AlexanderThaller/httphelper"
	"github.com/juju/errgo"
	"github.com/julienschmidt/httprouter"
)

func pageRoot(w http.ResponseWriter, r *http.Request, p httprouter.Params) *httphelper.HandlerError {
	projects, err := dataStore.ListProjects(false)
	if err != nil {
		return httphelper.NewHandlerErrorDef(errgo.Notef(err, "can not get list of projects"))
	}

	tmpl, err := getAssetTemplate("templates/html_pageRoot.html")
	if err != nil {
		return httphelper.NewHandlerErrorDef(errgo.Notef(err, "can not execute pageRoot template with project list"))
	}

	err = tmpl.Execute(w, projects.List())
	if err != nil {
		return httphelper.NewHandlerErrorDef(errgo.Notef(err, "can not execute template pageRoot"))
	}

	return nil
}

func pageShow(w http.ResponseWriter, r *http.Request, p httprouter.Params) *httphelper.HandlerError {
	l := httphelper.NewHandlerLogEntry(r)

	etype := p.ByName("type")
	project := p.ByName("project")

	l.Debug("Type: ", etype)
	l.Debug("Project: ", project)

	projects, err := helper.ProjectNamesFromArgs(dataStore, []string{project}, false)
	if err != nil {
		return httphelper.NewHandlerErrorDef(errgo.Notef(err, "can not get list of projects"))
	}

	err = dataStore.PopulateProjects(&projects)
	if err != nil {
		return httphelper.NewHandlerErrorDef(errgo.Notef(err, "can not populate projects with entries"))
	}

	buffer := new(bytes.Buffer)
	formatting.Projects(buffer, "Entries", 0, &projects)

	err = asciiDoctor(buffer, w)
	if err != nil {
		return httphelper.NewHandlerErrorDef(errgo.Notef(err, "can not format entries with asciidoctor"))
	}

	return nil
}
