package main

import (
	"io/ioutil"
	"path/filepath"
	"sort"
)

func GetProjects(datapath string) ([]string, error) {
	dir, err := ioutil.ReadDir(datapath)
	if err != nil {
		return []string{}, err
	}

	var out []string
	for _, d := range dir {
		file := d.Name()
		l.Debug("DirName: ", file)

		if strings.HasPrefix(file, ".") {
			l.Trace(file, " has prefix '.'")
			continue
		}

		ext := filepath.Ext(file)
		name := file[0 : len(file)-len(ext)]

		out = append(out, name)
	}

	sort.Strings(out)
	return out, nil
}
