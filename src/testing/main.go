package testing

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/AlexanderThaller/lablog/src/data"
	"github.com/sergi/go-diff/diffmatchpatch"
)

var TestProject = data.Project{
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

func CompareGotExpected(t *testing.T, err error, got, expected interface{}) {
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
