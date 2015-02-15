package web

import (
	"net/http"

	"github.com/AlexanderThaller/lablog/src/data"
	"github.com/AlexanderThaller/logger"
	"github.com/gorilla/schema"
	"github.com/juju/errgo"
)

func NoteFormPresenterHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.New(Name, "NoteFormPresenterHandler")

	data, err := html_note_html_bytes()
	if err != nil {
		printerr(l, w, errgo.Notef(err, "can not read note html"))
		return
	}

	_, err = w.Write(data)
	if err != nil {
		printerr(l, w, errgo.Notef(err, "can not write file to response writer"))
		return
	}
}

func NoteFormParserHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.New(Name, "NoteFormParserHandler")

	err := r.ParseForm()
	if err != nil {
		printerr(l, w, errgo.Notef(err, "can not parse form"))
		return
	}

	decoder := schema.NewDecoder()
	rawnote := new(Note)

	err = decoder.Decode(rawnote, r.PostForm)
	if err != nil {
		printerr(l, w, errgo.Notef(err, "can not decode form to note"))
		return
	}

	l.Trace("Note: ", rawnote)

	note, err := rawnote.ToData()
	if err != nil {
		printerr(l, w, errgo.Notef(err, "can not convert note to data note"))
		return
	}

	err = data.Record(_datadir, note)
	if err != nil {
		printerr(l, w, errgo.Notef(err, "can not record note"))
		return
	}
}
