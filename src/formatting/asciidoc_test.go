package formatting

import (
	"bytes"
	"testing"

	testhelper "github.com/AlexanderThaller/lablog/src/testing"
)

func Test_AsciidocHeaderIndent(t *testing.T) {
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
		got := AsciidocHeader(input)
		testhelper.CompareGotExpected(t, nil, got, expected)
	}
}

func Test_AsciidocNotes(t *testing.T) {
	expected := `== 2010-11-10 23:00:00 +0000 UTC
note note note` + "\n\n"

	got := new(bytes.Buffer)
	AsciidocNotes(got, 2, testhelper.GetTestProject("A", 1, 1).Notes())

	testhelper.CompareGotExpected(t, nil, got.String(), expected)
}

func Test_AsciidocNotesValue(t *testing.T) {
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

	AsciidocNotes(got, 2, project.Notes())

	testhelper.CompareGotExpected(t, nil, got.String(), expected)
}

func Test_AsciidocNotesMultiple5(t *testing.T) {
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
	AsciidocNotes(got, 2, testhelper.GetTestProject("A", 5, 1).Notes())

	testhelper.CompareGotExpected(t, nil, got.String(), expected)
}

func Test_AsciidocProjectNotes(t *testing.T) {
	expected := `= Test.Project.A
== 2010-11-10 23:00:00 +0000 UTC
note note note` + "\n\n"

	got := new(bytes.Buffer)
	project := testhelper.GetTestProject("A", 1, 1)
	AsciidocProjectNotes(got, 0, &project)

	testhelper.CompareGotExpected(t, nil, got.String(), expected)
}

func Test_AsciidocProjectNotesMultiple5(t *testing.T) {
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
	AsciidocProjectNotes(got, 0, &project)

	testhelper.CompareGotExpected(t, nil, got.String(), expected)
}

func Test_AsciidocProjectsNotes(t *testing.T) {
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
		AsciidocProjectNotes(got, 0, &project)
	}

	testhelper.CompareGotExpected(t, nil, got.String(), expected)
}

func Test_AsciidocProject(t *testing.T) {
	expected := `= Test.Project.A
== Todos
* todo todo todo

== Notes
=== 2010-11-10 23:00:00 +0000 UTC
note note note` + "\n\n"

	got := new(bytes.Buffer)

	project := testhelper.GetTestProject("A", 1, 1)
	AsciidocProject(got, 0, &project)

	testhelper.CompareGotExpected(t, nil, got.String(), expected)
}

func Test_AsciidocProjectNoNotesNoTodos(t *testing.T) {
	expected := ``

	got := new(bytes.Buffer)

	project := testhelper.GetTestProject("A", 0, 0)
	AsciidocProject(got, 0, &project)

	testhelper.CompareGotExpected(t, nil, got.String(), expected)
}

func Test_AsciidocProjectNoNotes(t *testing.T) {
	expected := `= Test.Project.A
== Todos
* todo todo todo` + "\n\n"

	got := new(bytes.Buffer)

	project := testhelper.GetTestProject("A", 0, 1)
	AsciidocProject(got, 0, &project)

	testhelper.CompareGotExpected(t, nil, got.String(), expected)
}

func Test_AsciidocProjectNoTodos(t *testing.T) {
	expected := `= Test.Project.A
== Notes
=== 2010-11-10 23:00:00 +0000 UTC
note note note` + "\n\n"

	got := new(bytes.Buffer)

	project := testhelper.GetTestProject("A", 1, 0)
	AsciidocProject(got, 0, &project)

	testhelper.CompareGotExpected(t, nil, got.String(), expected)
}

func Test_AsciidocProjectMultiple(t *testing.T) {
	expected := `:toc: right
:toclevels: 2
:sectanchors:
:sectlink:
:icons: font
:linkattrs:
:numbered:
:idprefix:
:idseparator: -
:doctype: book
:source-highlighter: pygments
:listing-caption: Listing` + "\n\n"

	expected += "= Entries\n\n"

	expected += `== Test.Project.A
=== Todos
* todo todo todo

=== Notes
==== 2010-11-10 23:00:00 +0000 UTC
note note note` + "\n\n"

	expected += `== Test.Project.B
=== Todos
* todo todo todo

=== Notes
==== 2010-11-10 23:00:00 +0000 UTC
note note note` + "\n\n"

	expected += `== Test.Project.C
=== Todos
* todo todo todo

=== Notes
==== 2010-11-10 23:00:00 +0000 UTC
note note note` + "\n\n"

	expected += `== Test.Project.D
=== Todos
* todo todo todo

=== Notes
==== 2010-11-10 23:00:00 +0000 UTC
note note note` + "\n\n"

	got := new(bytes.Buffer)

	projects := testhelper.GetTestProjects(1, 1, "A", "B", "C", "D")
	AsciidocProjects(got, "Entries", 0, &projects)

	testhelper.CompareGotExpected(t, nil, got.String(), expected)
}

func Test_AsciidocTodos(t *testing.T) {
	expected := `* todo todo todo` + "\n\n"

	got := new(bytes.Buffer)
	AsciidocTodos(got, testhelper.GetTestProject("A", 1, 1).Todos())

	testhelper.CompareGotExpected(t, nil, got.String(), expected)
}

func Test_AsciidocTodosDone(t *testing.T) {
	expected := ``

	got := new(bytes.Buffer)
	todos := testhelper.GetTestProject("A", 1, 1).Todos()
	todos[0].Active = false

	AsciidocTodos(got, todos)

	testhelper.CompareGotExpected(t, nil, got.String(), expected)
}

func Test_AsciidocTodosMultiple(t *testing.T) {
	expected := `* todo todo todo` + "\n"
	expected += `* todo todo todo` + "\n"
	expected += `* todo todo todo` + "\n"
	expected += `* todo todo todo` + "\n"
	expected += `* todo todo todo` + "\n\n"

	got := new(bytes.Buffer)
	AsciidocTodos(got, testhelper.GetTestProject("A", 1, 5).Todos())

	testhelper.CompareGotExpected(t, nil, got.String(), expected)
}

func Test_AsciidocProjectTodos(t *testing.T) {
	expected := `= Test.Project.A
* todo todo todo` + "\n\n"

	got := new(bytes.Buffer)
	project := testhelper.GetTestProject("A", 1, 1)
	AsciidocProjectTodos(got, 0, &project)

	testhelper.CompareGotExpected(t, nil, got.String(), expected)
}
