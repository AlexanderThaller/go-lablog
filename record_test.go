package main

import (
	"testing"
	"time"

	"github.com/AlexanderThaller/logger"
)

func Test_NoteCSV(t *testing.T) {
	l := logger.New(Name, "Test", "Note", "CSV")

	timestamp := time.Date(2009, time.November, 10, 23, 1, 2, 3, time.UTC)
	value := "NoteTestValue"

	note := Note{
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
		expected := timestamp.Format(time.RFC3339Nano)
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
		expected := value
		test(t, l, message, got, expected)
	}
}

func Test_TodoCSV(t *testing.T) {
	l := logger.New(Name, "Test", "Todo", "CSV")

	timestamp := time.Date(2009, time.November, 10, 23, 1, 2, 3, time.UTC)
	value := "TodoTestValue"
	done := true

	todo := Todo{
		TimeStamp: timestamp,
		Value:     value,
		Done:      done,
	}

	out := todo.CSV()

	{
		message := "Length of csv for todo is wrong"
		got := len(out)
		expected := 4
		test(t, l, message, got, expected)
	}

	{
		message := "TimeStamp of todo does not match timestamp of csv"
		got := out[0]
		expected := timestamp.Format(time.RFC3339Nano)
		test(t, l, message, got, expected)
	}

	{
		message := "Action of output csv is not todo"
		got := out[1]
		expected := "todo"
		test(t, l, message, got, expected)
	}

	{
		message := "Value of todo does not match todo of csv"
		got := out[2]
		expected := value
		test(t, l, message, got, expected)
	}

	{
		message := "Done of todo does not match todo of csv"
		got := out[3]
		expected := "true"
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
