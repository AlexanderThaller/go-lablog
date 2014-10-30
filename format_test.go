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
	FormatHeader(buffer, "Lablog", ActionListNotes, 1)
	got := buffer.String()
	if got != expected {
		l.Error("Did not get the expected output")
		l.Notice("GOT: ", got)
		l.Notice("EXPECTED: ", expected)

		t.Fail()
	}
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
	if got != expected {
		l.Error("Did not get the expected output")
		l.Notice("ERROR: ", err)
		l.Notice("GOT: ", got)
		l.Notice("EXPECTED: ", expected)

		t.Fail()
	}
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
	if got != expected {
		l.Error("Did not get the expected output")
		l.Notice("ERROR: ", err)
		l.Notice("GOT: ", got)
		l.Notice("EXPECTED: ", expected)

		t.Fail()
	}
}
