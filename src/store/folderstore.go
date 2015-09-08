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

func (store FolderStore) AddEntry(project data.ProjectName, entry data.Entry) error {
	db := store.db()
	err := db.Put(entry.Values(), project.Values()...)
	if err != nil {
		return errgo.Notef(err, "can not put entry into the database")
	}

	return nil
}

func (store FolderStore) GetProjects() (data.Projects, error) {
	return nil, errgo.New("not implemented")
}

func (store FolderStore) GetProject(name data.ProjectName) (data.Project, error) {
	return data.Project{}, errgo.New("not implemented")
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
