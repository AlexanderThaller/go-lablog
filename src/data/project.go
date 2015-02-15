package data

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/jinzhu/now"
	"github.com/juju/errgo"
)

// Projects returns all projects which can be read from the given datadir.
func Projects(datadir string) ([]Project, error) {
	names, err := projectNames(datadir)
	if err != nil {
		return nil, errgo.Notef(err, "can not get project files")
	}

	var projects []Project
	for _, name := range names {
		project := Project{
			Name:    name,
			Datadir: datadir,
		}

		projects = append(projects, project)
	}

	return projects, nil
}

// Project represents a project.
type Project struct {
	Name    string
	Datadir string
}

// ProjectsByName allows sorting project slices by name.
type ProjectsByName []Project

func (by ProjectsByName) Len() int {
	return len(by)
}

func (by ProjectsByName) Swap(i, j int) {
	by[i], by[j] = by[j], by[i]
}

func (by ProjectsByName) Less(i, j int) bool {
	return by[i].Name < by[j].Name
}

const (
	projectFileExtention = ".csv"
)

// projectNames read the files in the datadir and returns every filename
// (without the extention) that ends with projectFileExtention and does not
// start with a dot.
func projectNames(datadir string) ([]string, error) {
	dir, err := ioutil.ReadDir(datadir)
	if err != nil {
		return nil, errgo.Notef(err, "can not read files from directory")
	}

	var names []string
	for _, file := range dir {
		filename := file.Name()

		// Skip dotfiles
		if strings.HasPrefix(filename, ".") {
			continue
		}

		// Skip file not ending with right extention
		if !strings.HasSuffix(filename, projectFileExtention) {
			continue
		}

		ext := filepath.Ext(filename)

		// Strip extention from filename to get only the name
		name := filename[0 : len(filename)-len(ext)]

		names = append(names, name)
	}

	return names, nil
}

func (project Project) Notes() ([]Note, error) {
	file, err := project.File()
	if err != nil {
		return nil, errgo.Notef(err, "can not open project file")
	}
	defer file.Close()

	parser := csv.NewReader(file)
	parser.FieldsPerRecord = 3

	var out []Note

	for {
		csv, err := parser.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			// an error would mean that the csv line is not a note so we can skip it
			continue
		}

		timestamp, err := now.Parse(csv[0])
		if err != nil {
			return nil, errgo.Notef(err, "can not parse timestamp from csv")
		}

		note := Note{
			TimeStamp: timestamp,
			Text:      csv[2],
			Project:   project,
		}

		out = append(out, note)
	}

	return out, nil
}

func (project Project) File() (*os.File, error) {
	if project.Datadir == "" {
		return nil, errgo.New("path to datadir can not be emtpy")
	}

	filepath := filepath.Join(project.Datadir, project.Name+".csv")
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0640)
	if err != nil {
		return nil, errgo.Notef(err, "can not open file")
	}

	return file, nil
}
