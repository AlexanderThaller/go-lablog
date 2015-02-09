package project

import (
	"io"
	"reflect"
	"regexp"
	"strings"
	"unicode"
)

const AsciiDocSettings = `:toc: right
:toclevels: 2
:sectanchors:
:sectlink:
:icons: font
:linkattrs:
:numbered:
:idprefix:
:idseparator: -
:doctype: book
:source-highlighter: pygments
:listing-caption: Listing`

func FormatAction(action string) string {
	switch action {
	case ActionTracksActive:
		return "TracksActive"

	default:
		capaction := []rune(action)
		capaction[0] = unicode.ToUpper(capaction[0])
		return string(capaction)
	}
}

func FormatHeader(writer io.Writer, project, action string, indent int) {
	if indent != 1 {
		return
	}

	formaction := FormatAction(action)
	writer.Write([]byte("= link:/[" + project + " -- " + string(formaction) + "]\n"))
	writer.Write([]byte(AsciiDocSettings + "\n\n"))
}

func FormatNotes(writer io.Writer, project, action string, notes []Note, indent int) error {
	records := make([]Record, len(notes))
	for i, v := range notes {
		records[i] = Record(v)
	}

	return FormatRecords(writer, project, action, records, indent)
}

func FormatTodos(writer io.Writer, project, action string, todos []Todo, indent int) error {
	records := make([]Record, len(todos))
	for i, v := range todos {
		records[i] = Record(v)
	}

	return FormatRecords(writer, project, action, records, indent)
}

func FormatTracks(writer io.Writer, project, action string, tracks []Track, indent int) error {
	records := make([]Record, len(tracks))
	for i, v := range tracks {
		records[i] = Record(v)
	}

	return FormatRecords(writer, project, action, records, indent)
}

func FormatDurations(writer io.Writer, project string, durations []Duration, indent int) error {
	if durations == nil {
		return nil
	}
	if len(durations) == 0 {
		return nil
	}

	section := strings.Repeat("=", indent)
	writer.Write([]byte(section + " " + project + "\n\n"))

	for _, duration := range durations {
		writer.Write([]byte(duration.GetFormattedValue() + "\n"))
	}

	writer.Write([]byte("\n"))

	return nil
}

func FormatRecords(writer io.Writer, project, action string, records []Record, indent int) error {
	if len(records) == 0 {
		return nil
	}

	reg, err := regexp.Compile("(?m)^=")
	if err != nil {
		return err
	}

	section := strings.Repeat("=", indent)
	if indent == 1 {
		FormatHeader(writer, project, action, indent)
	} else {
		writer.Write([]byte(section + " " + project + "\n\n"))
	}

	for index, record := range records {
		if reflect.TypeOf(record).Name() == "Note" {
			writer.Write([]byte(section + "= " + record.GetTimeStamp() + "\n"))
		}

		out := reg.ReplaceAllString(record.GetFormattedValue(), section+"==")
		writer.Write([]byte(out + "\n"))
		if reflect.TypeOf(record).Name() == "Note" {
			writer.Write([]byte("\n"))
		} else if len(records) == index+1 {
			writer.Write([]byte("\n"))
		}
	}

	return nil
}
