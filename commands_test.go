package main

import (
	"testing"

	"github.com/AlexanderThaller/logger"
)

func Test_runListProjects(t *testing.T) {
	l := logger.New(Name, "Test", "Command", "runListProjects")

	command, buffer := testCommand()

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

	err := command.runListProjects()
	got := buffer.String()

	testerr_output(t, l, err, got, expected)
}
