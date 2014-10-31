package main

import (
	"testing"

	"github.com/AlexanderThaller/logger"
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

	expected := `2014-10-31
`

	testCommandRunOutput(t, l, action, expected)
}

func Test_RunList(t *testing.T) {
	l := logger.New(Name, "Test", "Command", "Run", "List")

	action := ActionList
	expected := `TestNotes
`

	testCommandRunOutput(t, l, action, expected)
}

func Test_RunNotes(t *testing.T) {
	l := logger.New(Name, "Test", "Command", "Run", "Notes")
	l.SetLevel(logger.Info)

	action := ActionNotes
	expected := "= Lablog -- Notes\n"
	expected += AsciiDocSettings + "\n\n"
	expected += `== TestNotes

=== 2014-10-31T21:36:31.49146148+01:00
Test1

=== 2014-10-31T21:36:33.49871531+01:00
Test2

=== 2014-10-31T21:36:35.138412374+01:00
Test3

=== 2014-10-31T21:36:36.810478305+01:00
Test4

=== 2014-10-31T21:36:38.450479686+01:00
Test5

`

	testCommandRunOutput(t, l, action, expected)
}
