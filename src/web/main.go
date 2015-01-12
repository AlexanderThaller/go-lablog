package web

import (
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/AlexanderThaller/lablog/src/project"
	"github.com/gorilla/mux"
	"github.com/juju/errgo"
)

// DataDir is the path to the datadir for lablog from which we will get the
// project data.
var DataDir string

func Listen(binding string) error {
	router := mux.NewRouter()
	router.HandleFunc("/", webRootHandler)
	router.HandleFunc("/notes/{project}", webNotesHandler)

	http.Handle("/", router)
	err := http.ListenAndServe(binding, nil)
	if err != nil {
		return err
	}

	return nil
}

func webRootHandler(w http.ResponseWriter, r *http.Request) {
	projects, err := project.Projects(DataDir, time.Time{}, time.Now())
	if err != nil {
		fmt.Fprintf(w, "can not get projects: %s", errgo.Details(err))
		return
	}

	page := RootPage{Projects: projects}
	WriteTemplateHTML(w, r, page)
}

func webNotesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	proj := vars["project"]

	notes, err := project.ProjectNotes(proj, DataDir, time.Time{}, time.Now())
	if err != nil {
		fmt.Fprintf(w, "can not get project notes: %s", errgo.Details(err))
		return
	}
	sort.Sort(project.NotesByDate(notes))

	page := PageNotes{Project: proj, Notes: notes}
	WriteTemplateHTML(w, r, page)
}
