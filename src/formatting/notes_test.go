package formatting

import (
	"bytes"
	"testing"

	testhelper "github.com/AlexanderThaller/lablog/src/testing"
)

func Test_Notes(t *testing.T) {
	expected := `== 2010-11-10 23:00:00 +0000 UTC
note note note` + "\n\n"

	got := new(bytes.Buffer)
	Notes(got, 2, testhelper.GetTestProject("A", 1, 1).Notes())

	testhelper.CompareGotExpected(t, nil, got.String(), expected)
}

func Test_NotesValue(t *testing.T) {
	expected := `== 2010-11-10 23:00:00 +0000 UTC
=== Header In Note Value
note note note

==== SubHeader In Note Value
note note note` + "\n\n"

	value := `= Header In Note Value
note note note

== SubHeader In Note Value
note note note`

	got := new(bytes.Buffer)

	project := testhelper.GetTestProject("A", 0, 0)
	project.AddNote(testhelper.GetTestNote(0, value))

	Notes(got, 2, project.Notes())

	testhelper.CompareGotExpected(t, nil, got.String(), expected)
}

func Test_NotesMultiple5(t *testing.T) {
	expected := `== 2010-11-10 23:00:00 +0000 UTC
note note note` + "\n\n"

	expected += `== 2011-11-10 23:00:00 +0000 UTC
note note note` + "\n\n"

	expected += `== 2012-11-10 23:00:00 +0000 UTC
note note note` + "\n\n"

	expected += `== 2013-11-10 23:00:00 +0000 UTC
note note note` + "\n\n"

	expected += `== 2014-11-10 23:00:00 +0000 UTC
note note note` + "\n\n"

	got := new(bytes.Buffer)
	Notes(got, 2, testhelper.GetTestProject("A", 5, 1).Notes())

	testhelper.CompareGotExpected(t, nil, got.String(), expected)
}

func Test_ProjectNotes(t *testing.T) {
	expected := `= Test.Project.A
== 2010-11-10 23:00:00 +0000 UTC
note note note` + "\n\n"

	got := new(bytes.Buffer)
	project := testhelper.GetTestProject("A", 1, 1)
	ProjectNotes(got, 0, &project)

	testhelper.CompareGotExpected(t, nil, got.String(), expected)
}

func Test_ProjectNotesMultiple5(t *testing.T) {
	expected := `= Test.Project.A
== 2010-11-10 23:00:00 +0000 UTC
note note note` + "\n\n"

	expected += `== 2011-11-10 23:00:00 +0000 UTC
note note note` + "\n\n"

	expected += `== 2012-11-10 23:00:00 +0000 UTC
note note note` + "\n\n"

	expected += `== 2013-11-10 23:00:00 +0000 UTC
note note note` + "\n\n"

	expected += `== 2014-11-10 23:00:00 +0000 UTC
note note note` + "\n\n"

	got := new(bytes.Buffer)
	project := testhelper.GetTestProject("A", 5, 1)
	ProjectNotes(got, 0, &project)

	testhelper.CompareGotExpected(t, nil, got.String(), expected)
}

func Test_ProjectsNotes(t *testing.T) {
	expected := `= Test.Project.A
== 2010-11-10 23:00:00 +0000 UTC
note note note` + "\n\n"

	expected += `= Test.Project.B
== 2010-11-10 23:00:00 +0000 UTC
note note note` + "\n\n"

	expected += `= Test.Project.C
== 2010-11-10 23:00:00 +0000 UTC
note note note` + "\n\n"

	expected += `= Test.Project.D
== 2010-11-10 23:00:00 +0000 UTC
note note note` + "\n\n"

	got := new(bytes.Buffer)

	projects := testhelper.GetTestProjects(1, 1, "A", "B", "C", "D")
	for _, project := range projects.List() {
		ProjectNotes(got, 0, &project)
	}

	testhelper.CompareGotExpected(t, nil, got.String(), expected)
}
