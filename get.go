package main

import (
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"
)

func GetProjectSubprojects(project string, projects []string) []string {
	if len(projects) == 0 {
		return []string{}
	}

	var out []string
	for _, subproject := range projects {
		if subproject == project {
			continue
		}

		if !strings.HasPrefix(subproject, project) {
			continue
		}

		out = append(out, subproject)
	}

	sort.Strings(out)
	return out
}

func GetProjects(datapath string) ([]string, error) {
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
