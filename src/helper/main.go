package helper

import (
	"time"

	"github.com/AlexanderThaller/lablog/src/data"
	"github.com/jinzhu/now"
	"github.com/juju/errgo"
)

func DefaultOrRawTimestamp(timestamp time.Time, raw string) (time.Time, error) {
	if timestamp.String() == raw {
		return timestamp, nil
	}

	parsed, err := now.Parse(raw)
	if err != nil {
		return time.Time{}, errgo.Notef(err, "can not parse timestamp")
	}

	return parsed, nil
}

func ProjectsDates(projects []data.Project, start, end time.Time) ([]string, error) {
	filter := make(map[string]struct{})

	for _, project := range projects {
		entries, err := FilteredEntriesByStartEnd(project, start, end)
		if err != nil {
			return nil, errgo.Notef(err, "can not get filtered entries for project")
		}

		if len(entries) == 0 {
			continue
		}

		for _, entry := range entries {
			date := entry.GetTimeStamp().Format("2006-01-02")
			filter[date] = struct{}{}
		}
	}

	var dates []string
	for date := range filter {
		dates = append(dates, date)
	}

	return dates, nil
}

func FilteredEntriesByStartEnd(project data.Project, start, end time.Time) ([]data.Entry, error) {
	var out []data.Entry

	notes, err := FilteredNotesByStartEnd(project, start, end)
	if err != nil {
		return nil, errgo.Notef(err, "can not get todos from project "+project.Name)
	}

	for _, note := range notes {
		out = append(out, note)
	}

	todos, err := FilteredTodosByStartEnd(project, start, end)
	if err != nil {
		return nil, errgo.Notef(err, "can not get todos from project "+project.Name)
	}
	for _, todo := range todos {
		out = append(out, todo)
	}

	tracks, err := FilteredTracksByStartEnd(project, start, end)
	if err != nil {
		return nil, errgo.Notef(err, "can not get tracks from project "+project.Name)
	}
	for _, track := range tracks {
		out = append(out, track)
	}

	return out, nil
}

func FilteredNotesByStartEnd(project data.Project, start, end time.Time) ([]data.Note, error) {
	notes, err := project.Notes()
	if err != nil {
		return nil, errgo.Notef(err, "can not get notes from project "+project.Name)
	}

	notes = data.FilterNotesBeforeTimeStamp(notes, start)
	notes = data.FilterNotesAfterTimeStamp(notes, end)

	return notes, nil
}

func FilteredTodosByStartEnd(project data.Project, start, end time.Time) ([]data.Todo, error) {
	todos, err := project.Todos()
	if err != nil {
		return nil, errgo.Notef(err, "can not get todos from project "+project.Name)
	}

	todos = data.FilterTodosBeforeTimeStamp(todos, start)
	todos = data.FilterTodosAfterTimeStamp(todos, end)

	return todos, nil
}

func FilteredTracksByStartEnd(project data.Project, start, end time.Time) ([]data.Track, error) {
	tracks, err := project.Tracks()
	if err != nil {
		return nil, errgo.Notef(err, "can not get tracks from project "+project.Name)
	}

	tracks = data.FilterTracksBeforeTimeStamp(tracks, start)
	tracks = data.FilterTracksAfterTimeStamp(tracks, end)

	return tracks, nil
}
