package web

import (
	"net/http"

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
