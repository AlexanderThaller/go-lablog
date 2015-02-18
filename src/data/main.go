package data

import (
	"encoding/csv"
	"os"
	"path/filepath"

	"github.com/juju/errgo"
)

func Record(datadir string, entry Entry) error {
	if datadir == "" {
		return errgo.New("datadir can not be emtpy")
	}

	project := entry.GetProject()
	if project.Name == "" {
		return errgo.New("the project name can not be empty")
	}

	err := os.MkdirAll(datadir, 0750)
	if err != nil {
		return errgo.Notef(err, "can not create datadir")
	}

	filepath := filepath.Join(datadir, entry.GetProject().Name+".csv")
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0640)
	if err != nil {
		return errgo.Notef(err, "can not open file for writing")
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.Write(entry.ValueArray())
	if err != nil {
		return errgo.Notef(err, "can not write entry")
	}
	writer.Flush()

	return nil
}
