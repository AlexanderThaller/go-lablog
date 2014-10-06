package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/AlexanderThaller/logger"
	"github.com/juju/errgo"
)

var (
	flagSourceJSON    = flag.String("json", "source.json", "")
	flagOutFolderPath = flag.String("out", "out", "")
)

func init() {
	flag.Parse()
}

type Project struct {
	Name   string
	Notes  []Note
	Tracks []Track
	Tags   []Tag
	Todos  []Todo
}

type Note struct {
	TimeStamp time.Time
	Value     string
}

type Tag struct {
	Value  string
	Action TagAction
}

type TagAction uint8

const (
	TagActionAppend TagAction = iota
	TagActionRemove
	TagActionReplace
)

type Todo struct {
	Value  string
	Action TodoAction
}

type TodoAction bool

const (
	TodoActionStart TodoAction = true
	TodoActionStop  TodoAction = false
)

type Track struct {
	Action    TrackAction
	TimeStamp time.Time
}

type TrackAction uint8

const (
	TrackActionTrackingStop TrackAction = iota
	TrackActionTrackingStart
	TrackActionTrackingToggle
)

func main() {
	l := logger.New("main")
	l.SetLevel(logger.Trace)

	i, err := ioutil.ReadFile(*flagSourceJSON)
	if err != nil {
		l.Alert(errgo.New(err.Error()))
		os.Exit(1)
	}

	var b bytes.Buffer
	json.Compact(&b, i)

	projects := make(map[string]Project)
	err = json.Unmarshal(b.Bytes(), &projects)
	if err != nil {
		l.Alert(errgo.New(err.Error()))
		os.Exit(1)
	}

	l.Trace("Projects: ", projects)

	err = os.MkdirAll(*flagOutFolderPath, 0750)
	if err != nil {
		l.Alert(errgo.New(err.Error()))
		os.Exit(1)
	}

	for _, project := range projects {
		filepath := filepath.Join(*flagOutFolderPath, project.Name+".csv")
		file, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0640)
		if err != nil {
			l.Alert(errgo.New(err.Error()))
			os.Exit(1)
		}

		writer := csv.NewWriter(file)

		for _, note := range project.Notes {
			message := []string{
				note.TimeStamp.Format(time.RFC3339Nano),
				project.Name,
				note.Value,
			}

			err = writer.Write(message)
			if err != nil {
				l.Alert(errgo.New(err.Error()))
				os.Exit(1)
			}
		}

		writer.Flush()
		file.Close()
	}
}
