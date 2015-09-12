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

	err = store.PutProject(testhelper.TestProject)
	if err != nil {
		t.Fatal("can not put test project into store", err)
	}

	got, err := store.GetProject(testhelper.TestProject.Name)
	testhelper.CompareGotExpected(t, err, got, testhelper.TestProject)
}

func Test_AddEntry(t *testing.T) {
	store, err := tmp_folderstore()
	if err != nil {
		t.Fatal("can not get tmp folderstore: ", err)
	}

	for _, entry := range testhelper.TestProject.Entries {
		err = store.AddEntry(testhelper.TestProject.Name, entry)
		if err != nil {
			t.Fatal("can not put test project into store", err)
		}
	}

	got, err := store.GetProject(testhelper.TestProject.Name)
	testhelper.CompareGotExpected(t, err, got, testhelper.TestProject)
}

func Test_ListProjects(t *testing.T) {
	store, err := tmp_folderstore()
	if err != nil {
		t.Fatal("can not get tmp folderstore: ", err)
	}

	err = store.PutProject(testhelper.TestProject)
	if err != nil {
		t.Fatal("can not put test project into store", err)
	}

	expected := []data.ProjectName{testhelper.TestProject.Name}
	got, err := store.ListProjects()
	testhelper.CompareGotExpected(t, err, got, expected)
}
