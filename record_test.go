package main

import (
	"errors"
	"strconv"
	"testing"
	"time"

	"github.com/AlexanderThaller/logger"
)

func Test_RecordFromCSV(t *testing.T) {
	l := logger.New(Name, "Test", "RecordFromCSV")

	timestamp := time.Date(2009, time.November, 10, 23, 1, 2, 3, time.UTC)

	{
		value := "NoteTestValue"

		record := Note{
			TimeStamp: timestamp,
			Value:     value,
		}

		recordcsv := record.CSV()
		newrecord, err := RecordFromCSV(recordcsv)
		compareRecord(t, l, err, newrecord, record)
	}

	{
		value := "TodoTestValue"
		done := true

		record := Todo{
			TimeStamp: timestamp,
			Value:     value,
			Done:      done,
		}

		recordcsv := record.CSV()
		newrecord, err := RecordFromCSV(recordcsv)
		compareRecord(t, l, err, newrecord, record)
	}
}

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
		expected := timestamp.Format(RecordTimeStampFormat)
		test(t, l, message, got, expected)
	}

	{
		message := "Action of output csv is not note"
		got := out[1]
		expected := ActionNote
		test(t, l, message, got, expected)
	}

	{
		message := "Value of note does not match note of csv"
		got := out[2]
		expected := value
		test(t, l, message, got, expected)
	}
}

func Test_NoteGet(t *testing.T) {
	l := logger.New(Name, "Test", "Note", "Get")

	timestamp := time.Date(2009, time.November, 10, 23, 1, 2, 3, time.UTC)
	value := "NoteTestValue"
	project := "NoteTestProject"

	note := Note{
		TimeStamp: timestamp,
		Value:     value,
		Project:   project,
	}

	{
		message := "Note action is not the expected action"
		got := note.GetAction()
		expected := ActionNote
		test(t, l, message, got, expected)
	}

	{
		message := "Note project is not the expected project"
		got := note.GetProject()
		expected := project
		test(t, l, message, got, expected)
	}

	{
		message := "Note timestamp is not the expected timestamp"
		got := note.GetTimeStamp()
		expected := timestamp.Format(RecordTimeStampFormat)
		test(t, l, message, got, expected)
	}

	{
		message := "Note value is not the expected value"
		got := note.GetValue()
		expected := value
		test(t, l, message, got, expected)
	}
}

func Test_NoteSet(t *testing.T) {
	l := logger.New(Name, "Test", "Note", "Set")

	project := "NoteTestProject"

	note := Note{}
	note.SetProject(project)

	{
		message := "Note project is not equal to project that was set"
		got := note.GetProject()
		expected := project
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
		expected := strconv.FormatBool(done)
		test(t, l, message, got, expected)
	}
}

func Test_TodoGet(t *testing.T) {
	l := logger.New(Name, "Test", "Todo", "Get")

	timestamp := time.Date(2009, time.November, 10, 23, 1, 2, 3, time.UTC)
	value := "TodoTestValue"
	project := "TodoTestProject"

	note := Todo{
		TimeStamp: timestamp,
		Value:     value,
		Project:   project,
	}

	{
		message := "Todo action is not the expected action"
		got := note.GetAction()
		expected := ActionTodo
		test(t, l, message, got, expected)
	}

	{
		message := "Todo project is not the expected project"
		got := note.GetProject()
		expected := project
		test(t, l, message, got, expected)
	}

	{
		message := "Todo timestamp is not the expected timestamp"
		got := note.GetTimeStamp()
		expected := timestamp.Format(RecordTimeStampFormat)
		test(t, l, message, got, expected)
	}

	{
		message := "Todo value is not the expected value"
		got := note.GetValue()
		expected := value
		test(t, l, message, got, expected)
	}
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

func test(t *testing.T, l logger.Logger, message string, got, expected interface{}) {
	testerr(t, l, message, errors.New("no error"), got, expected)
}

func testerr(t *testing.T, l logger.Logger, message string, err error, got, expected interface{}) {
	if got != expected {
		l.Error("MESSAGE : ", message)
		l.Error("ERROR   : ", err)
		l.Error("GOT     : ", got)
		l.Error("EXPECTED: ", expected)
		t.Fail()
	}
}
