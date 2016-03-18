package testing

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/AlexanderThaller/lablog/src/data"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func GetTestProjects(notes, todos int, suffixes ...string) data.Projects {
	out := data.NewProjects()

	for _, suffix := range suffixes {
		out.Add(GetTestProject(suffix, notes, todos))
	}

	return out
}

func GetTestProject(suffix string, notes, todos int) data.Project {
	project := data.Project{
		Name: data.ProjectName([]string{"Test", "Project", suffix}),
	}

	for i := 0; i != notes; i++ {
		project.AddNote(GetTestNote(i, "note note note"))
	}

	for i := 0; i != todos; i++ {
		project.AddTodo(data.Todo{
			TimeStamp: time.Date(2010+i, time.November, 10, 23, 0, 0, 0, time.UTC),
			Value:     "todo todo todo",
			Active:    true,
		})
	}

	return project
}

func GetTestNote(increment int, value string) data.Note {
	return data.Note{
		TimeStamp: time.Date(2010+increment, time.November, 10, 23, 0, 0, 0, time.UTC),
		Value:     value,
	}
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
