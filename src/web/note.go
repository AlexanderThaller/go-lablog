package web

import (
	"net/http"
	"text/template"
	"time"

	"github.com/AlexanderThaller/lablog/src/data"
	"github.com/AlexanderThaller/logger"
	"github.com/gorilla/schema"
	"github.com/jinzhu/now"
	"github.com/juju/errgo"
)

type Note struct {
	Project   string
	TimeStamp string
	Text      string
}

func (note Note) ToData() (data.Note, error) {
	timestamp := time.Now()

	if note.TimeStamp != "" {
		var err error
		timestamp, err = now.Parse(note.TimeStamp)
		if err != nil {
			return data.Note{}, errgo.Notef(err, "can not parse timestamp")
		}
	}

	data := data.Note{
		Project:   data.Project{Name: note.Project},
		TimeStamp: timestamp,
		Text:      note.Text,
	}

	return data, nil
}

func noteForm(w http.ResponseWriter, r *http.Request) {
	l := logger.New(Name, "noteForm")

	rawtmpl, err := html_note_html_bytes()
	if err != nil {
		printerr(l, w, errgo.Notef(err, "can not read note html"))
		return
	}

	tmpl, err := template.New("name").Parse(string(rawtmpl))
	if err != nil {
		printerr(l, w, errgo.Notef(err, "can not parse noteform html template"))
		return
	}

	project := defquery(r, "project", "")
	note := Note{Project: project[0]}

	l.Info("Serving noteform")
	err = tmpl.Execute(w, note)
	if err != nil {
		printerr(l, w, errgo.Notef(err, "can not execute projects html template"))
		return
	}
}

func noteParser(w http.ResponseWriter, r *http.Request) {
	l := logger.New(Name, "noteParser")

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

	l.Info("Recording new note: ", note)

	err = data.Record(_datadir, note)
	if err != nil {
		printerr(l, w, errgo.Notef(err, "can not record note"))
		return
	}

	http.Redirect(w, r, "/note?project="+note.Project.Name, 302)
}
