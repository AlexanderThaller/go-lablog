package formatting

import (
	"bytes"
	"testing"

	testhelper "github.com/AlexanderThaller/lablog/src/testing"
)

func Test_Project(t *testing.T) {
	expected := `= Test.Project.A
== Todos
* todo todo todo

== Notes
=== 2010-11-10 23:00:00
note note note` + "\n\n"

	got := new(bytes.Buffer)

	project := testhelper.GetTestProject("A", 1, 1)
	Project(got, 0, &project)

	testhelper.CompareGotExpected(t, nil, got.String(), expected)
}

func Test_ProjectNoNotesNoTodos(t *testing.T) {
	expected := ``

	got := new(bytes.Buffer)

	project := testhelper.GetTestProject("A", 0, 0)
	Project(got, 0, &project)

	testhelper.CompareGotExpected(t, nil, got.String(), expected)
}

func Test_ProjectNoNotes(t *testing.T) {
	expected := `= Test.Project.A
== Todos
* todo todo todo` + "\n\n"

	got := new(bytes.Buffer)

	project := testhelper.GetTestProject("A", 0, 1)
	Project(got, 0, &project)

	testhelper.CompareGotExpected(t, nil, got.String(), expected)
}

func Test_ProjectNoTodos(t *testing.T) {
	expected := `= Test.Project.A
== Notes
=== 2010-11-10 23:00:00
note note note` + "\n\n"

	got := new(bytes.Buffer)

	project := testhelper.GetTestProject("A", 1, 0)
	Project(got, 0, &project)

	testhelper.CompareGotExpected(t, nil, got.String(), expected)
}

func Test_ProjectMultiple(t *testing.T) {
	expected := `:toc: right
:toclevels: 4
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
==== 2010-11-10 23:00:00
note note note` + "\n\n"

	expected += `== Test.Project.B
=== Todos
* todo todo todo

=== Notes
==== 2010-11-10 23:00:00
note note note` + "\n\n"

	expected += `== Test.Project.C
=== Todos
* todo todo todo

=== Notes
==== 2010-11-10 23:00:00
note note note` + "\n\n"

	expected += `== Test.Project.D
=== Todos
* todo todo todo

=== Notes
==== 2010-11-10 23:00:00
note note note` + "\n\n"

	got := new(bytes.Buffer)

	projects := testhelper.GetTestProjects(1, 1, "A", "B", "C", "D")
	Projects(got, "Entries", 0, &projects)

	testhelper.CompareGotExpected(t, nil, got.String(), expected)
}
