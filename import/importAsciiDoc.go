package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/AlexanderThaller/logger"
	"github.com/jinzhu/now"
	"github.com/juju/errgo"
)

var (
	flagSource        = flag.String("src", "source.asciidoc", "")
	flagOutFolderPath = flag.String("out", "out", "")
)

func init() {
	flag.Parse()

	now.TimeFormats = append(now.TimeFormats, time.ANSIC)
	now.TimeFormats = append(now.TimeFormats, time.UnixDate)
	now.TimeFormats = append(now.TimeFormats, time.RubyDate)
	now.TimeFormats = append(now.TimeFormats, time.RFC822)
	now.TimeFormats = append(now.TimeFormats, time.RFC822Z)
	now.TimeFormats = append(now.TimeFormats, time.RFC850)
	now.TimeFormats = append(now.TimeFormats, time.RFC1123)
	now.TimeFormats = append(now.TimeFormats, time.RFC1123Z)
	now.TimeFormats = append(now.TimeFormats, time.RFC3339)
	now.TimeFormats = append(now.TimeFormats, time.RFC3339Nano)
	now.TimeFormats = append(now.TimeFormats, "Mon, _2 Jan 2006 15:04:05 MST")
	now.TimeFormats = append(now.TimeFormats, "Mon, _2 Jan 2006 15:04:05 -0700")
	now.TimeFormats = append(now.TimeFormats, "Mon, _2 January 2006 15:04:05 MST")
	now.TimeFormats = append(now.TimeFormats, "Mon, _2 January 2006 15:04:05 -0700")
	now.TimeFormats = append(now.TimeFormats, "Mon, _2 Jan 06 15:04:05 -0700")
	now.TimeFormats = append(now.TimeFormats, "Mon _2. Jan 15:04:05 MST 2006")
	now.TimeFormats = append(now.TimeFormats, "_2 Jan 2006 15:04:05 -0700")
	now.TimeFormats = append(now.TimeFormats, "2006.01.02-15.04.05")
	now.TimeFormats = append(now.TimeFormats, "2006.01.02")
	now.TimeFormats = append(now.TimeFormats, "2006.01.02---- ")
	now.TimeFormats = append(now.TimeFormats, "2006.01.02 - 15:04:05")
}

func main() {
	l := logger.New("main")
	l.SetLevel(logger.Debug)

	file, err := os.OpenFile(*flagSource, os.O_RDONLY, 0640)
	if err != nil {
		l.Alert("Can not open file: ", file)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var notes []Note
	var project string
	var timestamp time.Time
	var value string
	for scanner.Scan() {
		entry := scanner.Text()

		if strings.HasPrefix(entry, "==== ") ||
			strings.HasPrefix(entry, "===== ") ||
			strings.HasPrefix(entry, "====== ") {

			value += strings.TrimPrefix(entry, "===")
			value += "\n"
			continue
		}

		if strings.HasPrefix(entry, "=== ") {
			oldtimestamp := timestamp
			entrytrim := strings.TrimPrefix(entry, "=== ")
			stamp, err := now.Parse(entrytrim)
			if err != nil {
				l.Alert("Can not parse timestamp: ", err)
				return
			}
			timestamp = stamp

			if oldtimestamp.Equal(time.Time{}) {
				continue
			}

			note := Note{
				Project:   project,
				TimeStamp: timestamp,
				Value:     value,
			}
			fmt.Printf("note: %+v\n", note)
			value = ""

			notes = append(notes, note)
			continue
		}

		if strings.HasPrefix(entry, "== ") {
			oldproject := project
			entrytrim := strings.TrimPrefix(entry, "== ")
			project = entrytrim

			if oldproject == "" {
				continue
			}

			note := Note{
				Project:   project,
				TimeStamp: timestamp,
				Value:     value,
			}
			fmt.Printf("note: %+v\n", note)
			value = ""
			timestamp = time.Time{}

			notes = append(notes, note)
			continue
		}

		value += entry
		value += "\n"
	}

	err = os.MkdirAll(*flagOutFolderPath, 0750)
	if err != nil {
		l.Alert(errgo.New(err.Error()))
		os.Exit(1)
	}

	for _, note := range notes {
		filepath := filepath.Join(*flagOutFolderPath, note.Project+".csv")
		file, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0640)
		if err != nil {
			l.Alert(errgo.New(err.Error()))
			os.Exit(1)
		}

		writer := csv.NewWriter(file)

		message := []string{
			note.TimeStamp.Format(time.RFC3339Nano),
			"note",
			strings.TrimSpace(note.Value),
		}

		err = writer.Write(message)
		if err != nil {
			l.Alert(errgo.New(err.Error()))
			os.Exit(1)
		}

		writer.Flush()
		file.Close()
	}
}

type Note struct {
	Project   string
	TimeStamp time.Time
	Value     string
}
