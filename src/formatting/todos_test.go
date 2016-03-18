package formatting

import (
	"bytes"
	"testing"

	testhelper "github.com/AlexanderThaller/lablog/src/testing"
)

func Test_Todos(t *testing.T) {
	expected := `* todo todo todo` + "\n\n"

	got := new(bytes.Buffer)
	Todos(got, testhelper.GetTestProject("A", 1, 1).Todos())

	testhelper.CompareGotExpected(t, nil, got.String(), expected)
}

func Test_TodosDone(t *testing.T) {
	expected := ``

	got := new(bytes.Buffer)
	todos := testhelper.GetTestProject("A", 1, 1).Todos()
	todos[0].Active = false

	Todos(got, todos)

	testhelper.CompareGotExpected(t, nil, got.String(), expected)
}

func Test_TodosMultiple(t *testing.T) {
	expected := `* todo todo todo` + "\n"
	expected += `* todo todo todo` + "\n"
	expected += `* todo todo todo` + "\n"
	expected += `* todo todo todo` + "\n"
	expected += `* todo todo todo` + "\n\n"

	got := new(bytes.Buffer)
	Todos(got, testhelper.GetTestProject("A", 1, 5).Todos())

	testhelper.CompareGotExpected(t, nil, got.String(), expected)
}

func Test_ProjectTodos(t *testing.T) {
	expected := `= Test.Project.A
* todo todo todo` + "\n\n"

	got := new(bytes.Buffer)
	project := testhelper.GetTestProject("A", 1, 1)
	ProjectTodos(got, 0, &project)

	testhelper.CompareGotExpected(t, nil, got.String(), expected)
}
