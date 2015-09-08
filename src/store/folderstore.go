package store

import (
	"github.com/AlexanderThaller/dbfiles"
	"github.com/AlexanderThaller/lablog/src/data"
	"github.com/juju/errgo"
)

func NewFolderStore(datadir string) (Store, error) {
	return FolderStore{datadir}, nil
}

type FolderStore struct {
	datadir string
}

func (store FolderStore) AddEntry(name data.ProjectName, entry data.Entry) error {
	db := store.db()
	err := db.Put(entry.Values(), name.Values()...)
	if err != nil {
		return errgo.Notef(err, "can not put entry into the database")
	}

	return nil
}

func (store FolderStore) GetProjects() (data.Projects, error) {
	return nil, errgo.New("not implemented")
}

func (store FolderStore) GetProject(name data.ProjectName) (data.Project, error) {
	db := store.db()
	values, err := db.Get(name.Values()...)
	if err != nil {
		return data.Project{}, errgo.Notef(err, "can not put entry into the database")
	}

	var entries []data.Entry
	for _, value := range values {
		entry, err := data.ParseEntry(value)
		if err != nil {
			return data.Project{}, errgo.Notef(err, "can not parse entry from value")
		}

		entries = append(entries, entry)
	}

	return data.Project{Name: name, Entries: entries}, nil
}

func (store FolderStore) ListProjects() ([]data.ProjectName, error) {
	db := store.db()
	keys, err := db.Keys()
	if err != nil {
		return nil, errgo.Notef(err, "can not get keys from database")
	}

	var out []data.ProjectName
	for _, key := range keys {
		name := data.ProjectName(key)
		out = append(out, name)
	}

	return out, nil
}

func (store FolderStore) db() *dbfiles.DBFiles {
	db := dbfiles.New()
	db.BaseDir = store.datadir

	return db
}
