package main

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/AlexanderThaller/logger"
	"github.com/juju/errgo"
)

func Dates(project, datapath string, start, end time.Time) ([]string, error) {
	projects, err := Projects(datapath, start, end)
	if err != nil {
		return nil, err
	}

	datemap := make(map[string]struct{})
	for _, project := range projects {
		dates, err := ProjectDates(project, datapath, start, end)
		if err != nil {
			return nil, err
		}

		for _, date := range dates {
			datemap[date] = struct{}{}
		}
	}

	var out []string
	for date := range datemap {
		out = append(out, date)
	}

	return out, nil
}

func MergeFiles(srcpath, dstpath string) error {
	srcdata, err := ioutil.ReadFile(srcpath)
	if err != nil {
		return errgo.New(err.Error())
	}

	dstfile, err := os.OpenFile(dstpath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return errgo.New(err.Error())
	}
	defer dstfile.Close()

	_, err = dstfile.Write(srcdata)
	if err != nil {
		return errgo.New(err.Error())
	}

	return nil
}

func ProjectsFiles(datapath string) ([]string, error) {
	dir, err := ioutil.ReadDir(datapath)
	if err != nil {
		return nil, err
	}

	var out []string
	for _, d := range dir {
		file := d.Name()

		// Skip dotfiles
		if strings.HasPrefix(file, ".") {
			continue
		}

		// Skip files not ending with .csv
		if !strings.HasSuffix(file, ".csv") {
			continue
		}

		ext := filepath.Ext(file)
		name := file[0 : len(file)-len(ext)]

		out = append(out, name)
	}

	sort.Strings(out)
	return out, nil
}

func Projects(datapath string, start, end time.Time) ([]string, error) {
	projects, err := ProjectsFiles(datapath)
	if err != nil {
		return nil, err
	}

	checker := make(chan string, 5000)
	var wg sync.WaitGroup

	for _, project := range projects {
		wg.Add(1)
		go func(project string) {
			if ProjectWasActive(project, datapath, start, end) {
				checker <- project
			} else {
				wg.Done()
			}
		}(project)
	}

	var out []string
	go func() {
		for {
			project := <-checker
			out = append(out, project)
			wg.Done()
		}
	}()

	wg.Wait()
	sort.Strings(out)

	return out, nil
}

func ProjectWasActive(project, datapath string, start, end time.Time) bool {
	if ProjectHasNotes(project, datapath, start, end) {
		return true
	}

	if ProjectHasTodos(project, datapath, start, end) {
		return true
	}

	if ProjectHasTracks(project, datapath, start, end) {
		return true
	}

	return false
}

func ProjectHasNotes(project, datapath string, start, end time.Time) bool {
	notes, _ := ProjectNotes(project, datapath, start, end)
	if len(notes) != 0 {
		return true
	}

	return false
}

func ProjectHasTodos(project, datapath string, start, end time.Time) bool {
	todos, _ := ProjectTodos(project, datapath, start, end)
	if len(todos) != 0 {
		return true
	}

	return false
}

func ProjectHasTracks(project, datapath string, start, end time.Time) bool {
	tracks, _ := ProjectTracks(project, datapath, start, end)
	if len(tracks) != 0 {
		return true
	}

	return false
}

func ProjectSubprojects(project, datapath string, start, end time.Time) ([]string, error) {
	projects, err := Projects(datapath, start, end)
	if err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return []string{}, nil
	}

	var out []string
	for _, subproject := range projects {
		if subproject == project {
			continue
		}

		if !strings.HasPrefix(subproject, project+".") {
			continue
		}

		out = append(out, subproject)
	}

	sort.Strings(out)
	return out, nil
}

func ProjectNotes(project, datapath string, start, end time.Time) ([]Note, error) {
	if datapath == "" {
		return nil, errgo.New("datapath can not be empty")
	}
	if project == "" {
		return nil, errgo.New("project name can not be empty")
	}
	if !ProjectExists(project, datapath) {
		return nil, errgo.New("project does not exist")
	}

	filepath := filepath.Join(datapath, project+".csv")
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0640)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return ProjectNotesFromReader(project, start, end, file), nil
}

func ProjectTodos(project, datapath string, start, end time.Time) ([]Todo, error) {
	if datapath == "" {
		return nil, errgo.New("datapath can not be empty")
	}
	if project == "" {
		return nil, errgo.New("project name can not be empty")
	}
	if !ProjectExists(project, datapath) {
		return nil, errgo.New("project does not exist")
	}

	filepath := filepath.Join(datapath, project+".csv")
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0640)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 4

	var out []Todo
	for {
		csv, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			continue
		}

		todo, err := TodoFromCSV(csv)
		if err != nil {
			continue
		}

		if todo.TimeStamp.Before(start) {
			continue
		}

		if todo.TimeStamp.After(end) {
			continue
		}

		out = append(out, todo)
	}

	return out, err
}

func ProjectTracks(project, datapath string, start, end time.Time) ([]Track, error) {
	if datapath == "" {
		return nil, errgo.New("datapath can not be empty")
	}
	if project == "" {
		return nil, errgo.New("project name can not be empty")
	}
	if !ProjectExists(project, datapath) {
		return nil, errgo.New("project does not exist")
	}

	filepath := filepath.Join(datapath, project+".csv")
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0640)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 3

	var out []Track
	for {
		csv, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			continue
		}

		track, err := TrackFromCSV(csv)
		if err != nil {
			continue
		}

		if track.TimeStamp.Before(start) {
			continue
		}

		if track.TimeStamp.After(end) {
			continue
		}

		out = append(out, track)
	}
	sort.Sort(TracksByDate(out))

	return out, err
}

func ProjectNotesFromReader(project string, start, end time.Time, reader io.Reader) []Note {
	parser := csv.NewReader(reader)
	parser.FieldsPerRecord = 3

	var out []Note
	for {
		csv, err := parser.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			continue
		}

		note, err := NoteFromCSV(csv)
		if err != nil {
			continue
		}
		note.SetProject(project)

		if note.TimeStamp.Before(start) {
			continue
		}

		if note.TimeStamp.After(end) {
			continue
		}

		out = append(out, note)
	}

	return out
}

func ProjectExists(project, datapath string) bool {
	projects, err := ProjectsFiles(datapath)
	if err != nil {
		return false
	}

	for _, d := range projects {
		if d == project {
			return true
		}
	}

	return false
}

func ProjectDates(project, datapath string, start, end time.Time) ([]string, error) {
	if datapath == "" {
		return nil, errgo.New("datapath can not be empty")
	}
	if project == "" {
		return nil, errgo.New("project name can not be empty")
	}
	if ProjectExists(project, datapath) {
		return nil, errgo.New("project does not exist")
	}

	var out []string

	notes, err := ProjectNotes(project, datapath, start, end)
	if err != nil {
		return nil, err
	}

	todos, err := ProjectTodos(project, datapath, start, end)
	if err != nil {
		return nil, err
	}
	todos = FilterInactiveTodos(todos)

	datemap := make(map[string]struct{})

	for _, note := range notes {
		date, err := time.Parse(RecordTimeStampFormat, note.GetTimeStamp())
		if err != nil {
			return nil, err
		}

		datemap[date.Format(DateFormat)] = struct{}{}
	}

	for _, todo := range todos {
		datemap[todo.TimeStamp.Format(DateFormat)] = struct{}{}
	}

	for date := range datemap {
		out = append(out, date)
	}

	return out, nil
}

func ProjectActiveTracks(project, datapath string) ([]Track, error) {
	l := logger.New(Name, "Command", "Get", "Project", "ActiveTracks")

	tracks, err := ProjectTracks(project, datapath, time.Time{}, time.Now())
	if err != nil {
		return nil, err
	}

	filter := make(map[string]bool)
	for _, track := range tracks {
		value, _ := filter[track.Value]

		if value {
			filter[track.Value] = false
		} else {
			filter[track.Value] = true
		}
		l.Debug(track.Value, " is ", filter[track.Value])
	}

	sort.Sort(TracksByDate(tracks))
	latesttracks := make(map[string]Track)
	for _, track := range tracks {
		latesttracks[track.Value] = track
	}

	var out []Track
	for value, track := range latesttracks {
		if !filter[value] {
			continue
		}

		out = append(out, track)
	}

	return out, nil
}
