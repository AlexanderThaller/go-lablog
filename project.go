package main

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/juju/errgo"
)

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

	out := make(map[string]struct{})
	for _, project := range projects {
		notes, err := ProjectNotes(project, datapath, start, end)
		if err != nil {
			return nil, err
		}

		if len(notes) == 0 {
			continue
		}

		out[project] = struct{}{}
	}

	for _, project := range projects {
		todos, err := ProjectTodos(project, datapath)
		if err != nil {
			return nil, err
		}

		if len(todos) == 0 {
			continue
		}

		out[project] = struct{}{}
	}

	var outsort []string
	for project := range out {
		outsort = append(outsort, project)
	}
	sort.Strings(outsort)

	return outsort, nil
}

func ProjectSubprojects(project string, projects []string) []string {
	if len(projects) == 0 {
		return []string{}
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
	return out
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

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 3

	var out []Note
	for {
		csv, err := reader.Read()
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

	return out, err
}

func ProjectTodos(project, datapath string) ([]Todo, error) {
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

		out = append(out, todo)
	}

	return out, err
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
