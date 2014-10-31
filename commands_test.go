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

	expected := `2014-10-25
2014-10-31
`

	testCommandRunOutput(t, l, action, expected)
}

func Test_RunList(t *testing.T) {
	l := logger.New(Name, "Test", "Command", "Run", "List")

	action := ActionList
	expected := `Test1
Test10
Test2
Test3
Test4
Test5
Test6
Test7
Test8
Test9
`

	testCommandRunOutput(t, l, action, expected)
}
