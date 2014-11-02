package main

import (
	"io"
	"regexp"
	"strings"
	"unicode"
)

const AsciiDocSettings = `:toc: right
:toclevels: 1
:sectanchors:
:sectlink:
:icons: font
:linkattrs:
:numbered:
:idprefix:
:idseparator: -
:doctype: article
:source-highlighter: coderay
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
	writer.Write([]byte("= " + project + " -- " + string(formaction) + "\n"))
	writer.Write([]byte(AsciiDocSettings + "\n\n"))
}

func FormatNotes(writer io.Writer, project string, notes []Note, indent int) error {
	records := make([]Record, len(notes))
	for i, v := range notes {
		records[i] = Record(v)
	}

	return FormatRecords(writer, project, records, indent)
}

func FormatTodos(writer io.Writer, project string, todos []Todo, indent int) error {
	if len(todos) == 0 {
		return nil
	}

	section := strings.Repeat("=", indent)
	writer.Write([]byte(section + " " + project + "\n\n"))

	for _, record := range todos {
		out := record.GetFormattedValue()
		writer.Write([]byte(out + "\n"))
	}
	writer.Write([]byte("\n"))

	return nil
}

func FormatTracks(writer io.Writer, project string, tracks []Track, indent int) error {
	if len(tracks) == 0 {
		return nil
	}

	section := strings.Repeat("=", indent)
	writer.Write([]byte(section + " " + project + "\n\n"))

	for _, record := range tracks {
		out := record.GetFormattedValue()
		writer.Write([]byte(out + "\n"))
	}
	writer.Write([]byte("\n"))

	return nil
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

func FormatRecords(writer io.Writer, project string, records []Record, indent int) error {
	if len(records) == 0 {
		return nil
	}

	reg, err := regexp.Compile("(?m)^=")
	if err != nil {
		return err
	}

	section := strings.Repeat("=", indent)
	writer.Write([]byte(section + " " + project + "\n\n"))

	for _, record := range records {
		writer.Write([]byte(section + "= " + record.GetTimeStamp() + "\n"))

		out := reg.ReplaceAllString(record.GetFormattedValue(), section+"==")
		writer.Write([]byte(out + "\n\n"))
	}

	return nil
}
