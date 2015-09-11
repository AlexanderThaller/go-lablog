package formatting

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/AlexanderThaller/lablog/src/data"
	"github.com/sergi/go-diff/diffmatchpatch"
)

var testProject = data.Project{
	Name: data.ProjectName([]string{"Test", "Project", "A"}),
	Entries: data.Entries([]data.Entry{
		data.Todo{
			Active:    true,
			TimeStamp: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
			Value:     "todo todo todo",
		},
		data.Note{
			TimeStamp: time.Date(2010, time.November, 10, 23, 0, 0, 0, time.UTC),
			Value:     "note note note",
		},
	}),
}

func Test_Todos(t *testing.T) {
	expected := `* todo todo todo` + "\n"

	got := new(bytes.Buffer)
	Todos(got, testProject.Todos())

	compare(t, nil, got.String(), expected)
}

func Test_Notes(t *testing.T) {
	expected := `== 2010-11-10 23:00:00 +0000 UTC
note note note` + "\n"

	got := new(bytes.Buffer)
	Notes(got, testProject.Notes())

	compare(t, nil, got.String(), expected)
}

func Test_ProjectTodos(t *testing.T) {
	expected := `= Test.Project.A
* todo todo todo` + "\n"

	got := new(bytes.Buffer)
	ProjectTodos(got, 0, testProject)

	compare(t, nil, got.String(), expected)
}

func Test_ProjectNotes(t *testing.T) {
	expected := `= Test.Project.A
== 2010-11-10 23:00:00 +0000 UTC
note note note` + "\n"

	got := new(bytes.Buffer)
	ProjectNotes(got, 0, testProject)

	compare(t, nil, got.String(), expected)
}

func Test_HeaderIndent(t *testing.T) {
	tests := map[int]string{
		-1: "",
		0:  "",
		1:  "=",
		2:  "==",
		3:  "===",
		4:  "====",
		5:  "=====",
		6:  "======",
	}

	for input, expected := range tests {
		got := HeaderIndent(input)
		compare(t, nil, got, expected)
	}
}

func compare(t *testing.T, err error, got, expected interface{}) {
	if reflect.DeepEqual(got, expected) {
		return
	}

	t.Logf("got:\n%v", got)
	t.Logf("expected:\n%v", expected)

	differ := diffmatchpatch.New()
	diff := differ.DiffMain(fmt.Sprintf("%+v", expected),
		fmt.Sprintf("%+v", got), true)

	var diffout string
	for _, line := range diff {
		switch line.Type {
		case diffmatchpatch.DiffDelete:
			diffout += fmt.Sprintf("\033[32m%v\033[0m", line.Text)
		case diffmatchpatch.DiffInsert:
			diffout += fmt.Sprintf("\033[31m%v\033[0m", line.Text)
		default:
			diffout += line.Text
		}
	}

	t.Logf("diff:\n%v", diffout)
	t.Fatal("got is not like expected")
}
