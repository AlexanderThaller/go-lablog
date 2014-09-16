package main

import (
	"io/ioutil"
	"sort"
)

func GetProjects(datapath string) ([]string, error) {
	dir, err := ioutil.ReadDir(datapath)
	if err != nil {
		return []string{}, err
	}

	var out []string
	for _, d := range dir {
		out = append(out, d.Name())
	}

	sort.Strings(out)
	return out, nil
}
