package data

import (
	"io/ioutil"
	"path/filepath"
	"strings"

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
			Name: name,
		}

		projects = append(projects, project)
	}

	return projects, nil
}

// Project represents a project.
type Project struct {
	Name string
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
