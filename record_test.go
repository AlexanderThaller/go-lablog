package main

import (
	"testing"
	"time"

	"github.com/AlexanderThaller/logger"
)

func Test_NoteCSV(t *testing.T) {
	l := logger.New(Name, "Test", "Note", "CSV")

	project := "TestProject"
	timestamp := time.Time{}
	value := "TestValue"

	note := Note{
		Project:   project,
		TimeStamp: timestamp,
		Value:     value,
	}

	out := note.CSV()

	{
		message := "Length of csv for note is wrong"
		got := len(out)
		expected := 3
		test(t, l, message, got, expected)
	}

	{
		message := "TimeStamp of note does not match timestamp of csv"
		got := out[0]
		expected := "0001-01-01T00:00:00Z"
		test(t, l, message, got, expected)
	}

	{
		message := "Action of output csv is not note"
		got := out[1]
		expected := "note"
		test(t, l, message, got, expected)
	}

	{
		message := "Value of note does not match note of csv"
		got := out[2]
		expected := "TestValue"
		test(t, l, message, got, expected)
	}
}

func test(t *testing.T, l logger.Logger, message string, got, expected interface{}) {
	if got != expected {
		l.Error("MESSAGE : ", message)
		l.Error("GOT     : ", got)
		l.Error("EXPECTED: ", expected)
		t.Fail()
	}
}
