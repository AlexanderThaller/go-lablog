package main

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/AlexanderThaller/logger"
	"github.com/juju/errgo"
	"github.com/sergi/go-diff/diffmatchpatch"
)

const (
	Project          = "Test1"
	TestDataPath     = "testdata/tests"
	TestDataPathFail = "/FAILFAILFAIL_12345"
	BenchDataPath    = "testdata/benchmark"
)

var (
	StartTime = time.Time{}
	EndTime   = time.Time{}
)

func testCommand(action string) (*Command, *bytes.Buffer) {
	buffer := bytes.NewBufferString("")
	command := NewCommand(buffer)

	command.DataPath = TestDataPath
	command.EndTime = time.Now()
	command.StartTime = time.Time{}
	command.Action = action

	return command, buffer
}

func testCommandRunOutput(t *testing.T, l logger.Logger, action, expected string) {
	command, buffer := testCommand(action)

	err := command.Run()
	got := buffer.String()

	testerrOutput(t, l, err, got, expected)
}

func compareRecord(t *testing.T, l logger.Logger, err error, newrecord, record Record) {
	if record == nil {
		l.Alert("record is nil")
		t.Fail()
		return
	}
	if newrecord == nil {
		l.Alert("newrecord is nil")
		t.Fail()
		return
	}

	{
		message := "Newrecord action is not the same with record action"
		got := newrecord.GetAction()
		expected := record.GetAction()
		testerr(t, l, message, err, got, expected)
	}

	{
		message := "Newrecord project is not the same with record project"
		got := newrecord.GetProject()
		expected := record.GetProject()
		testerr(t, l, message, err, got, expected)
	}

	{
		message := "Newrecord timestamp is not the same with record timestamp"
		got := newrecord.GetTimeStamp()
		expected := record.GetTimeStamp()
		testerr(t, l, message, err, got, expected)
	}

	{
		message := "Newrecord value is not the same with record value"
		got := newrecord.GetValue()
		expected := record.GetValue()
		testerr(t, l, message, err, got, expected)
	}
}

func test_output(t *testing.T, l logger.Logger, got, expected interface{}) {
	message := "Did not get the expected output"
	test(t, l, message, got, expected)
}

func test(t *testing.T, l logger.Logger, message string, got, expected interface{}) {
	testerr(t, l, message, errors.New("no error"), got, expected)
}

func testerrOutput(t *testing.T, l logger.Logger, err error, got, expected interface{}) {
	message := "Did not get the expected output"
	testerr(t, l, message, err, got, expected)
}

func testerr(t *testing.T, l logger.Logger, message string, err error, got, expected interface{}) {
	if got == expected {
		return
	}

	l.Error("MESSAGE : ", message)
	if err != nil {
		l.Notice("ERROR: ", errgo.Details(err))
	}
	l.Notice("GOT:\n", got)
	l.Notice("EXPECTED:\n", expected)

	if reflect.TypeOf(got).Name() == "string" {
		differ := diffmatchpatch.New()
		diff := differ.DiffMain(expected.(string), got.(string), true)

		l.Notice("DIFF:")
		for _, line := range diff {
			switch line.Type {
			case diffmatchpatch.DiffDelete:
				fmt.Print("\033[32m" + line.Text + "\033[0m")
			case diffmatchpatch.DiffInsert:
				fmt.Print("\033[31m" + line.Text + "\033[0m")
			default:
				fmt.Print(line.Text)
			}
		}
	}

	t.Fail()
}
