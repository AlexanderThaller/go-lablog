package main

import (
	"io"
	"regexp"
	"strings"
	"time"
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

func FormatHeader(writer io.Writer, project, action string, indent int) {
	if indent != 1 {
		return
	}

	capaction := []rune(action)
	capaction[0] = unicode.ToUpper(capaction[0])

	writer.Write([]byte("= " + project + " -- " + string(capaction) + "\n"))
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
	records := make([]Record, len(tracks))
	for i, v := range tracks {
		records[i] = Record(v)
	}

	return FormatRecords(writer, project, records, indent)
}

func FormatDurations(writer io.Writer, project string, durations map[string]time.Duration, indent int) error {
	if durations == nil {
		return nil
	}
	if len(durations) == 0 {
		return nil
	}

	section := strings.Repeat("=", indent)
	writer.Write([]byte(section + " " + project + "\n\n"))

	for value, duration := range durations {
		if value != "" {
			value += " -- "
		}
		writer.Write([]byte("* " + value + duration.String() + "\n"))
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
