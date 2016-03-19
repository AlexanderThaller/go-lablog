package store

import (
	"io/ioutil"
	"testing"

	"github.com/AlexanderThaller/lablog/src/data"
	testhelper "github.com/AlexanderThaller/lablog/src/testing"
	"github.com/juju/errgo"
)

func tmp_folderstore() (Store, error) {
	tmpdir, err := ioutil.TempDir("/tmp/", "folderstore_test")
	if err != nil {
		return nil, errgo.Notef(err, "can not open tmpdir")
	}

	store, err := NewFolderStore(tmpdir)
	if err != nil {
		return nil, errgo.Notef(err, "can not create new store")
	}

	return store, nil
}

func Test_PutProject(t *testing.T) {
	store, err := tmp_folderstore()
	if err != nil {
		t.Fatal("can not get tmp folderstore: ", err)
	}

	project := testhelper.GetTestProject("A", 1, 1)
	err = store.PutProject(project)
	if err != nil {
		t.Fatal("can not put test project into store", err)
	}

	got, err := store.GetProject(project.Name)
	testhelper.CompareGotExpected(t, err, got, project)
}

func Test_AddEntry(t *testing.T) {
	store, err := tmp_folderstore()
	if err != nil {
		t.Fatal("can not get tmp folderstore: ", err)
	}

	project := testhelper.GetTestProject("A", 1, 1)
	for _, entry := range project.Entries {
		err = store.AddEntry(project.Name, entry)
		if err != nil {
			t.Fatal("can not put test project into store", err)
		}
	}

	got, err := store.GetProject(project.Name)
	testhelper.CompareGotExpected(t, err, got, project)
}

func Test_ListProjects(t *testing.T) {
	store, err := tmp_folderstore()
	if err != nil {
		t.Fatal("can not get tmp folderstore: ", err)
	}

	project := testhelper.GetTestProject("A", 1, 1)
	err = store.PutProject(project)
	if err != nil {
		t.Fatal("can not put test project into store", err)
	}

	projects := data.NewProjects()
	projects.Add(project)

	expected := []data.Project{data.Project{Name: project.Name}}
	got, err := store.ListProjects(false)
	testhelper.CompareGotExpected(t, err, got.List(), expected)
}
