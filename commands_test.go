package main

import (
	"testing"

	"github.com/AlexanderThaller/logger"
)

const (
	ExpectedDates = `2014-10-30
2014-10-31
2014-11-01
2014-11-05
`

	ExpectedProjects = `TestNotes
TestNotes.Subproject
TestTodos
TestTodos.Subproject
TestTracks
TestTracks.Subproject
`

	ExpectedNotes = `== TestNotes

=== 2014-10-30T21:36:31.49146148+01:00
Test1

=== 2014-10-30T21:36:33.49871531+01:00
Test2

=== 2014-10-30T21:36:35.138412374+01:00
Test3

=== 2014-10-30T21:36:36.810478305+01:00
Test4

=== 2014-10-30T21:36:38.450479686+01:00
Test5

== TestNotes.Subproject

=== 2014-11-05T22:39:54.162785484+01:00
Test

`

	ExpectedTodos = `== TestTodos

* Test1
* Test2
* Test3
* Test4
* Test5
* Test7

== TestTodos.Subproject

* Test

`

	ExpectedTracks = `== TestTracks

* 2014-11-01T00:46:27.250010094+01:00 -- Test1
* 2014-11-01T00:46:31.186052306+01:00 -- Test2
* 2014-11-01T00:46:32.794131714+01:00 -- Test3
* 2014-11-01T00:46:34.322047221+01:00 -- Test4
* 2014-11-01T00:46:35.658221386+01:00 -- Test5
* 2014-11-01T00:46:38.385921461+01:00
* 2014-11-01T00:46:51.833861322+01:00 -- Test6
* 2014-11-01T00:46:53.713763133+01:00 -- Test6
* 2014-11-01T00:46:55.609688812+01:00 -- Test7
* 2014-11-01T00:46:56.713668447+01:00 -- Test7
* 2014-11-01T00:46:57.953493565+01:00 -- Test7
* 2014-11-01T00:54:54.778921093+01:00

== TestTracks.Subproject

* 2014-11-05T23:27:37.354857614+01:00 -- Test

`

	ExpectedTracksActive = `== TestTracks

* 2014-11-01T00:46:27.250010094+01:00 -- Test1
* 2014-11-01T00:46:31.186052306+01:00 -- Test2
* 2014-11-01T00:46:32.794131714+01:00 -- Test3
* 2014-11-01T00:46:34.322047221+01:00 -- Test4
* 2014-11-01T00:46:35.658221386+01:00 -- Test5
* 2014-11-01T00:46:57.953493565+01:00 -- Test7

== TestTracks.Subproject

* 2014-11-05T23:27:37.354857614+01:00 -- Test

`

	ExpectedTracksDurations = `== TestTracks

* 8m16.392999632s
* Test6 -- 1.879901811s
* Test7 -- 1.103979635s

`
)

func Test_RunNoDataPath(t *testing.T) {
	l := logger.New(Name, "Test", "Command", "Run", "NoDataPath")

	command := new(Command)
	err := command.Run()

	got := err.Error()
	expected := "the datapath can not be empty"

	testerr_output(t, l, err, got, expected)
}

func Test_RunDates(t *testing.T) {
	l := logger.New(Name, "Test", "Command", "Run", "Dates")

	action := ActionDates

	expected := ExpectedDates

	testCommandRunOutput(t, l, action, expected)
}

func Test_RunList(t *testing.T) {
	l := logger.New(Name, "Test", "Command", "Run", "List")

	action := ActionList

	expected := ExpectedProjects

	testCommandRunOutput(t, l, action, expected)
}

func Test_RunListProject(t *testing.T) {
	l := logger.New(Name, "Test", "Command", "Run", "List", "Project")

	action := ActionList

	command, buffer := testCommand(action)
	command.Project = "TestNotes"

	expected := "= Lablog -- List\n"
	expected += AsciiDocSettings + "\n\n"
	expected += ExpectedNotes

	err := command.Run()
	got := buffer.String()

	testerr_output(t, l, err, got, expected)
}

func Test_RunListProjectNoDataDir(t *testing.T) {
	l := logger.New(Name, "Test", "Command", "Run", "List", "Project", "NoDataDir")

	action := ActionList

	command, buffer := testCommand(action)
	command.DataPath = TestDataPathFail

	expected := ""

	err := command.Run()
	got := buffer.String()

	testerr_output(t, l, err, got, expected)
}

func Test_RunListProjectNoNotes(t *testing.T) {
	l := logger.New(Name, "Test", "Command", "Run", "List", "Project", "NoNotes")

	action := ActionList

	command, buffer := testCommand(action)
	command.Project = "TestTodos"

	expected := "= Lablog -- List\n"
	expected += AsciiDocSettings + "\n\n"
	expected += ExpectedTodos

	err := command.Run()
	got := buffer.String()

	testerr_output(t, l, err, got, expected)
}

func Test_RunListProjectNoNotesNoTodos(t *testing.T) {
	l := logger.New(Name, "Test", "Command", "Run", "List", "Project", "NoNotesNoTodos")

	action := ActionList

	command, buffer := testCommand(action)
	command.Project = "TestTracks"

	expected := "= Lablog -- List\n"
	expected += AsciiDocSettings + "\n\n"
	expected += ExpectedTracksActive

	err := command.Run()
	got := buffer.String()

	testerr_output(t, l, err, got, expected)
}

func Test_RunListProjectNoExists(t *testing.T) {
	l := logger.New(Name, "Test", "Command", "Run", "List", "Project", "NoExists")

	action := ActionList

	command, buffer := testCommand(action)
	command.Project = "DOES NOT EXIST"

	expected := "= Lablog -- List\n"
	expected += AsciiDocSettings + "\n\n"

	err := command.Run()
	got := buffer.String()

	testerr_output(t, l, err, got, expected)
}

func Test_RunNotes(t *testing.T) {
	l := logger.New(Name, "Test", "Command", "Run", "Notes")
	l.SetLevel(logger.Info)

	action := ActionNotes

	expected := "= Lablog -- Notes\n"
	expected += AsciiDocSettings + "\n\n"
	expected += ExpectedNotes

	testCommandRunOutput(t, l, action, expected)
}

func Benchmark_RunNotes(b *testing.B) {
	action := ActionNotes

	command, _ := testCommand(action)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		command.Run()
	}
}

func Test_RunNotesProject(t *testing.T) {
	l := logger.New(Name, "Test", "Command", "Run", "Notes", "Project")
	l.SetLevel(logger.Info)

	action := ActionNotes

	command, buffer := testCommand(action)
	command.Project = "TestNotes"

	expected := "= Lablog -- Notes\n"
	expected += AsciiDocSettings + "\n\n"
	expected += ExpectedNotes

	err := command.Run()
	got := buffer.String()

	testerr_output(t, l, err, got, expected)
}

func Test_RunProjects(t *testing.T) {
	l := logger.New(Name, "Test", "Command", "Run", "Projects")

	action := ActionProjects

	expected := ExpectedProjects

	testCommandRunOutput(t, l, action, expected)
}

func Test_RunTodos(t *testing.T) {
	l := logger.New(Name, "Test", "Command", "Run", "Todos")
	l.SetLevel(logger.Info)

	action := ActionTodos

	expected := "= Lablog -- Todos\n"
	expected += AsciiDocSettings + "\n\n"
	expected += ExpectedTodos

	testCommandRunOutput(t, l, action, expected)
}

func Benchmark_RunTodos(b *testing.B) {
	action := ActionTodos
	command, _ := testCommand(action)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		command.Run()
	}
}

func Test_RunTracks(t *testing.T) {
	l := logger.New(Name, "Test", "Command", "Run", "Tracks")
	l.SetLevel(logger.Info)

	action := ActionTracks

	expected := "= Lablog -- Tracks\n"
	expected += AsciiDocSettings + "\n\n"
	expected += ExpectedTracks

	testCommandRunOutput(t, l, action, expected)
}

func Test_RunTracksActive(t *testing.T) {
	l := logger.New(Name, "Test", "Command", "Run", "TracksActive")
	l.SetLevel(logger.Info)

	action := ActionTracksActive

	expected := "= Lablog -- TracksActive\n"
	expected += AsciiDocSettings + "\n\n"
	expected += ExpectedTracksActive

	testCommandRunOutput(t, l, action, expected)
}

func Test_RunTracksDurations(t *testing.T) {
	l := logger.New(Name, "Test", "Command", "Run", "TracksDurations")
	l.SetLevel(logger.Info)

	action := ActionTracksDurations

	expected := "= Lablog -- Durations\n"
	expected += AsciiDocSettings + "\n\n"
	expected += ExpectedTracksDurations

	testCommandRunOutput(t, l, action, expected)
}
