package main

import (
	"encoding/csv"
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/AlexanderThaller/logger"
	"github.com/jinzhu/now"
	"github.com/juju/errgo"
)

var (
	flagSource        = flag.String("src", "source.markdown", "")
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

	data, err := ioutil.ReadFile(*flagSource)
	if err != nil {
		l.Alert(errgo.New(err.Error()))
		os.Exit(1)
	}

	l.Trace("Data: ", string(data))

	regex := `# (?P<timestamp>.*)\n## (?P<project>.*)\n(?P<value>[^#]*)`
	matcher, err := regexp.Compile(regex)
	if err != nil {
		l.Alert("Can not compile regex: ", err)
		os.Exit(1)
	}

	var notes []Note
	for _, match := range matcher.FindAllStringSubmatch(string(data), -1) {
		note := Note{}

		timestamp, err := now.Parse(match[1])
		if err != nil {
			l.Warning("Can not parse timestamp: ", err)
			continue
		}
		note.TimeStamp = timestamp

		note.Project = match[2]
		note.Value = match[3]

		if note.Project == "" {
			l.Alert("Project can not be empty")
			os.Exit(1)
		}

		if note.Value == "" {
			l.Alert("Value can not be empty")
			os.Exit(1)
		}

		notes = append(notes, note)
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
