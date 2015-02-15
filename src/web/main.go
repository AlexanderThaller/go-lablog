package web

import (
	"fmt"
	"net/http"

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
	router.HandleFunc("/new/note", NoteFormPresenterHandler)
	router.HandleFunc("/put/note", NoteFormParserHandler)

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
