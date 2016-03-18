package store

import (
	"github.com/AlexanderThaller/dbfiles"
	"github.com/AlexanderThaller/lablog/src/data"
	log "github.com/Sirupsen/logrus"
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
	return data.Projects{}, errgo.New("not implemented")
}

func (store FolderStore) PutProject(project data.Project) error {
	for _, entry := range project.Entries {
		err := store.AddEntry(project.Name, entry)
		if err != nil {
			return errgo.Notef(err, "can not add entry for project "+project.Name.String())
		}
	}

	return nil
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

func (store FolderStore) ListProjects(showarchive bool) (data.Projects, error) {
	db := store.db()
	keys, err := db.Keys()
	if err != nil {
		return data.Projects{}, errgo.Notef(err, "can not get keys from database")
	}

	out := data.NewProjects()
	for _, key := range keys {
		if !showarchive {
			// Skipping archived projects
			if len(key) > 0 {
				log.Debug("Key: ", key[0])

				if key[0] == ".archive" {
					continue
				}
			}
		}

		name := data.ProjectName(key)
		out.Add(data.Project{Name: name})
	}

	return out, nil
}

func (store FolderStore) PopulateProjects(projects *data.Projects) error {
	for _, project := range projects.List() {
		filled, err := store.GetProject(project.Name)
		if err != nil {
			return errgo.Notef(err, "can not get project: "+project.Name.String())
		}

		projects.Set(filled)
	}

	return nil
}

func (store FolderStore) db() *dbfiles.DBFiles {
	db := dbfiles.New()
	db.BaseDir = store.datadir

	return db
}
