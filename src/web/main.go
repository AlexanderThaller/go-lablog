package web

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/AlexanderThaller/lablog/src/data"
	"github.com/AlexanderThaller/lablog/src/scm"
	"github.com/AlexanderThaller/logger"
	"github.com/gorilla/mux"
	"github.com/juju/errgo"
)

const (
	Name = "web"
)

var _datadir string

func Listen(datadir, binding string) error {
	_datadir = datadir

	router := mux.NewRouter()
	router.HandleFunc("/", listProjects)
	router.HandleFunc("/note", noteForm)
	router.HandleFunc("/note/", noteForm)
	router.HandleFunc("/note/record", noteParser)
	router.HandleFunc("/note/record", noteParser)
	router.HandleFunc("/todo", todoForm)
	router.HandleFunc("/todo/", todoForm)
	router.HandleFunc("/todo/record", todoParser)
	router.HandleFunc("/todo/record", todoParser)
	router.HandleFunc("/list/notes", listNotes)
	router.HandleFunc("/list/notes/", listNotes)
	router.HandleFunc("/list/todos", listTodos)
	router.HandleFunc("/list/todos/", listTodos)

	http.Handle("/", router)

	err := http.ListenAndServe(binding, nil)
	if err != nil {
		return errgo.Notef(err, "can not listen on binding")
	}

	return nil
}

func printerr(l logger.Logger, w http.ResponseWriter, err error) {
	l.Error(errgo.Details(err))
	fmt.Fprintf(w, errgo.Details(err))

	return
}

func defquery(r *http.Request, key string, defvalue ...string) []string {
	queries, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		return defvalue
	}

	value, exists := queries[key]
	if !exists {
		return defvalue
	}

	return value
}

func recordAndCommit(datadir string, entry data.Entry) error {
	err := data.Record(datadir, entry)
	if err != nil {
		return errgo.Notef(err, "can not record note")
	}

	err = scm.Commit(datadir, entry)
	if err != nil {
		return errgo.Notef(err, "can not commit entry")
	}

	return nil
}
