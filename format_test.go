package main

import (
	"bytes"
	"testing"
	"time"

	"github.com/AlexanderThaller/logger"
)

func Test_FormatHeader(t *testing.T) {
	l := logger.New(Name, "Test", "FormatHeader")

	expected := "= Lablog -- Notes\n"
	expected += AsciiDocSettings + "\n\n"

	buffer := bytes.NewBufferString("")

	FormatHeader(buffer, "Lablog", ActionNotes, 1)
	got := buffer.String()

	testerr_output(t, l, nil, got, expected)
}

func Test_FormatNotes(t *testing.T) {
	l := logger.New(Name, "Test", "FormatNotes")

	note := Note{
		Project:   "TestProjectNote",
		TimeStamp: time.Time{},
		Value:     "TestValue",
	}
	expected := `= TestProject

== 0001-01-01T00:00:00Z
TestValue

`

	records := []Note{note}
	buffer := bytes.NewBufferString("")
	err := FormatNotes(buffer, "TestProject", records, 1)
	got := buffer.String()
	testerr_output(t, l, err, got, expected)
}

func Test_FormatRecordsNote(t *testing.T) {
	l := logger.New(Name, "Test", "FormatRecords", "Note")

	note := Note{
		Project:   "TestProjectNote",
		TimeStamp: time.Time{},
		Value:     "TestValue",
	}
	expected := `= TestProject

== 0001-01-01T00:00:00Z
TestValue

`

	records := []Record{note}
	buffer := bytes.NewBufferString("")
	err := FormatRecords(buffer, "TestProject", records, 1)
	got := buffer.String()
	testerr_output(t, l, err, got, expected)
}
