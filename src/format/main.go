package format

import (
	"bytes"
	"io"
	"os/exec"
	"sort"
	"time"

	"github.com/AlexanderThaller/lablog/src/data"
	"github.com/AlexanderThaller/lablog/src/helper"
	"github.com/AlexanderThaller/logger"
	"github.com/juju/errgo"
)

const (
	Name = "format"
)

const AsciiDocSettings = `:toc: right
:toclevels: 2
:sectanchors:
:sectlink:
:icons: font
:linkattrs:
:numbered:
:idprefix:
:idseparator: -
:doctype: book
:source-highlighter: pygments
:listing-caption: Listing`

func ProjectsEntries(writer io.Writer, projects []data.Project, start, end time.Time) error {
	l := logger.New(Name, "ProjectsEntries")

	io.WriteString(writer, AsciiDocSettings+"\n\n")
	io.WriteString(writer, "= Entries \n\n")

	for _, project := range projects {
		notes, err := helper.FilteredNotesByStartEnd(project, start, end)
		if err != nil {
			err := errgo.Notef(err, "can not get filtered notes")

			l.Debug(err)
			l.Trace(errgo.Details(err))
		}

		todos, err := helper.FilteredTodosByStartEnd(project, start, end)
		if err != nil {
			err := errgo.Notef(err, "can not get filtered todos")

			l.Debug(err)
			l.Trace(errgo.Details(err))
		}
		todos = data.FilterTodosLatest(todos)
		todos = data.FilterTodosAreNotDone(todos)

		project.Format(writer, 1)
		if len(todos) != 0 {
			Todos(writer, todos)
			io.WriteString(writer, "\n")
		}

		if len(notes) != 0 {
			Notes(writer, notes)
		}
	}

	return nil
}

func ProjectsNotes(writer io.Writer, projects []data.Project, start, end time.Time) error {
	l := logger.New(Name, "ProjectsNotes")

	io.WriteString(writer, AsciiDocSettings+"\n\n")
	io.WriteString(writer, "= Notes \n\n")

	for _, project := range projects {
		notes, err := helper.FilteredNotesByStartEnd(project, start, end)
		if err != nil {
			err := errgo.Notef(err, "can not get filtered notes")

			l.Debug(err)
			l.Trace(errgo.Details(err))
			continue
		}

		if len(notes) == 0 {
			continue
		}

		project.Format(writer, 1)
		Notes(writer, notes)
	}

	return nil
}

func ProjectsTodos(writer io.Writer, projects []data.Project, start, end time.Time) error {
	l := logger.New(Name, "ProjectTodos")

	io.WriteString(writer, AsciiDocSettings+"\n\n")
	io.WriteString(writer, "= Todos \n\n")

	for _, project := range projects {
		todos, err := helper.FilteredTodosByStartEnd(project, start, end)
		if err != nil {
			err := errgo.Notef(err, "can not get filtered todos")

			l.Debug(err)
			l.Trace(errgo.Details(err))
			continue
		}
		todos = data.FilterTodosLatest(todos)
		todos = data.FilterTodosAreNotDone(todos)

		if len(todos) == 0 {
			continue
		}

		project.Format(writer, 1)
		Todos(writer, todos)
		io.WriteString(writer, "\n")
	}

	return nil
}

func ProjectsTracks(writer io.Writer, projects []data.Project, start, end time.Time) error {
	l := logger.New(Name, "ProjectTodos")

	io.WriteString(writer, AsciiDocSettings+"\n\n")
	io.WriteString(writer, "= Tracks \n\n")

	for _, project := range projects {
		tracks, err := helper.FilteredTracksByStartEnd(project, start, end)
		if err != nil {
			err := errgo.Notef(err, "can not get filtered tracks")

			l.Debug(err)
			l.Trace(errgo.Details(err))
			continue
		}

		if len(tracks) == 0 {
			continue
		}

		project.Format(writer, 1)
		tracks = data.MergeTracks(tracks)
		Tracks(writer, tracks)
		io.WriteString(writer, "\n")
	}

	return nil
}

func ProjectsDurations(writer io.Writer, projects []data.Project, start, end time.Time) error {
	l := logger.New(Name, "ProjectsDurations")

	io.WriteString(writer, AsciiDocSettings+"\n\n")
	io.WriteString(writer, "= Durations \n\n")

	for _, project := range projects {
		tracks, err := helper.FilteredTracksByStartEnd(project, start, end)
		if err != nil {
			err := errgo.Notef(err, "can not get filtered tracks")

			l.Debug(err)
			l.Trace(errgo.Details(err))
			continue
		}

		if len(tracks) == 0 {
			continue
		}

		project.Format(writer, 1)
		Duration(writer, tracks)
		io.WriteString(writer, "\n")
	}

	return nil
}

func ProjectsDates(writer io.Writer, projects []data.Project, start, end time.Time) error {
	io.WriteString(writer, AsciiDocSettings+"\n\n")
	io.WriteString(writer, "= Dates \n\n")

	dates, err := helper.ProjectsDates(projects, start, end)
	if err != nil {
		return errgo.Notef(err, "can not get dates for projects")
	}

	sort.Strings(dates)

	for _, date := range dates {
		io.WriteString(writer, "* "+date+"\n")
	}

	return nil
}

func Todos(writer io.Writer, todos []data.Todo) {
	io.WriteString(writer, "=== Todos\n\n")

	sort.Sort(data.TodosByName(todos))
	for _, todo := range todos {
		todo.Format(writer, 2)
	}
}

func Notes(writer io.Writer, notes []data.Note) {
	io.WriteString(writer, "=== Notes\n\n")

	notes = data.FilterNotesNotEmpty(notes)

	sort.Sort(data.NotesByTimeStamp(notes))
	for _, note := range notes {
		note.Format(writer, 2)
	}
}

func Tracks(writer io.Writer, tracks []data.Track) {
	io.WriteString(writer, "=== Tracks\n\n")

	for i, track := range tracks {
		track.Format(writer, 2)
		if !track.Active {
			duration := data.TracksDuration(tracks[i-1 : i+1])
			io.WriteString(writer, " ["+duration.String()+"]")
		}

		if track.Active {
			if i == len(tracks)-1 {
				duration := time.Since(track.TimeStamp)
				io.WriteString(writer, " ["+duration.String()+"]")
			}
		}

		io.WriteString(writer, "\n")
	}
}

func Duration(writer io.Writer, tracks []data.Track) {
	io.WriteString(writer, "=== Duration\n\n")

	sort.Sort(data.TracksByTimeStamp(tracks))
	duration := data.TracksDuration(tracks)
	io.WriteString(writer, duration.String()+"\n")
}

func AsciiDoctor(reader io.Reader, writer io.Writer) error {
	stderr := new(bytes.Buffer)

	command := exec.Command("asciidoctor", "-")
	command.Stdin = reader
	command.Stdout = writer
	command.Stderr = stderr

	err := command.Run()
	if err != nil {
		return errgo.Notef(errgo.Notef(err, "can not run asciidoctor"),
			stderr.String())
	}

	return nil
}

func Timeline(writer io.Writer, projects []data.Project, start, end time.Time) error {
	var allnotes []data.Note

	for _, project := range projects {
		notes, err := helper.FilteredNotesByStartEnd(project, start, end)
		if err != nil {
			return errgo.Notef(err, "can not get filtered notes")
		}

		for _, note := range notes {
			allnotes = append(allnotes, note)
		}
	}

	allnotes = data.FilterNotesNotEmpty(allnotes)
	sort.Sort(data.NotesByTimeStamp(allnotes))

	io.WriteString(writer, AsciiDocSettings+"\n\n")
	io.WriteString(writer, "= Timeline \n\n")
	for _, note := range allnotes {
		note.FormatTimeStamp(writer, 2)
		note.Project.Format(writer, 2)
		note.FormatText(writer, 4)
	}

	return nil
}
