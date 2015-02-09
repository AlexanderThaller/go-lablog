package main

import (
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/gorilla/mux"
	"github.com/juju/errgo"
)

func (com *Command) runWeb() error {
	r := mux.NewRouter()
	r.HandleFunc("/", com.webRootHandler)
	r.HandleFunc("/notes/{project}", com.webNotesHandler)

	http.Handle("/", r)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return err
	}

	return nil
}

func (com *Command) webRootHandler(w http.ResponseWriter, r *http.Request) {
	projects, err := Projects(com.DataPath, time.Time{}, time.Now())
	if err != nil {
		fmt.Fprintf(w, "Error: %s", errgo.Details(err))
		return
	}

	page := RootPage{Projects: projects}
	WriteTemplateHTML(w, r, page)
}

func (com *Command) webNotesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	project := vars["project"]

	notes, err := ProjectNotes(project, com.DataPath, time.Time{}, time.Now())
	if err != nil {
		fmt.Fprintf(w, "Error: %s", errgo.Details(err))
		return
	}
	sort.Sort(NotesByDate(notes))

	page := PageNotes{Project: project, Notes: notes}
	WriteTemplateHTML(w, r, page)
}
