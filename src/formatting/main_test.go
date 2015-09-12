package formatting

import (
	"bytes"
	"testing"

	testhelper "github.com/AlexanderThaller/lablog/src/testing"
)

func Test_Todos(t *testing.T) {
	expected := `* todo todo todo` + "\n"

	got := new(bytes.Buffer)
	Todos(got, testhelper.TestProject.Todos())

	testhelper.CompareGotExpected(t, nil, got.String(), expected)
}

func Test_Notes(t *testing.T) {
	expected := `== 2010-11-10 23:00:00 +0000 UTC
note note note` + "\n"

	got := new(bytes.Buffer)
	Notes(got, testhelper.TestProject.Notes())

	testhelper.CompareGotExpected(t, nil, got.String(), expected)
}

func Test_ProjectTodos(t *testing.T) {
	expected := `= Test.Project.A
* todo todo todo` + "\n"

	got := new(bytes.Buffer)
	ProjectTodos(got, 0, testhelper.TestProject)

	testhelper.CompareGotExpected(t, nil, got.String(), expected)
}

func Test_ProjectNotes(t *testing.T) {
	expected := `= Test.Project.A
== 2010-11-10 23:00:00 +0000 UTC
note note note` + "\n"

	got := new(bytes.Buffer)
	ProjectNotes(got, 0, testhelper.TestProject)

	testhelper.CompareGotExpected(t, nil, got.String(), expected)
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
		testhelper.CompareGotExpected(t, nil, got, expected)
	}
}
