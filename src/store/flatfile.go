package store

import (
	"encoding/csv"
	"os"
	"path/filepath"

	"github.com/AlexanderThaller/lablog/src/data"
	"github.com/AlexanderThaller/lablog/src/scm"
	"github.com/juju/errgo"
)

type FlatFile struct {
	dataDir    string
	AutoCommit bool
	AutoPush   bool
}

func NewFlatFile(datadir string) *FlatFile {
	store := new(FlatFile)
	store.dataDir = datadir

	return store
}

func (store *FlatFile) Write(entry data.Entry) error {
	if store.dataDir == "" {
		return errgo.New("datadir can not be emtpy")
	}

	project := entry.GetProject()
	if project.Name == "" {
		return errgo.New("the project name can not be empty")
	}

	datadir := store.dataDir

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

	err = scm.Commit(store.dataDir, entry)
	if err != nil {
		return errgo.Notef(err, "can not commit entry")
	}

	return nil
}
