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

func FormatHeader(writer io.Writer, action string) {
	capaction := []rune(action)
	capaction[0] = unicode.ToUpper(capaction[0])

	writer.Write([]byte("= Lablog -- " + string(capaction) + "\n"))
	writer.Write([]byte(AsciiDocSettings + "\n\n"))
}

func FormatNotes(writer io.Writer, project string, notes []Note, indent int) error {
	records := make([]Record, len(notes))
	for i, v := range notes {
		records[i] = Record(v)
	}

	return FormatRecords(writer, project, records, indent)
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

		out := reg.ReplaceAllString(record.GetValue(), section+"==")
		writer.Write([]byte(out + "\n\n"))
	}

	return nil
}
