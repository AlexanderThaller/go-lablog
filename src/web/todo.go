package web

import (
	"net/http"

	"github.com/AlexanderThaller/logger"
	"github.com/juju/errgo"
)

func todoForm(w http.ResponseWriter, r *http.Request) {
	l := logger.New(Name, "todoForm")
	printerr(l, w, errgo.New("not implemented"))
}

func todoParser(w http.ResponseWriter, r *http.Request) {
	l := logger.New(Name, "todoParser")
	printerr(l, w, errgo.New("not implemented"))
}
