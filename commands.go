package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/AlexanderThaller/logger"
	"github.com/juju/errgo"
)

type Command struct {
	Action        string
	Args          []string
	DataPath      string
	EndTime       string
	Project       string
	SCM           string
	SCMAutoCommit bool
	SCMAutoPush   bool
	StartTime     string
	TimeStamp     time.Time
	Value         string
}

const (
	ActionList = "list"
	ActionNote = "note"
)

func NewCommand() *Command {
	return new(Command)
}

func (com *Command) Run() error {
	if com.DataPath == "" {
		return errgo.New("the datapath can not be empty")
	}

	switch com.Action {
	case ActionNote:
		return com.runNote()
	case ActionList:
		return com.runList()
	default:
		return errgo.New("Do not recognize the action: " + com.Action)
	}
}

func (com *Command) runNote() error {
	l := logger.New(Name, "Command", "run", "Note")

	l.Trace("Args length: ", len(com.Args))
	if com.Value == "" {
		return errgo.New("note command needs a value")
	}
	l.Trace("Project: ", com.Project)
	if com.Project == "" {
		return errgo.New("note command needs an project")
	}

	note := new(Note)
	note.Project = com.Project
	note.TimeStamp = com.TimeStamp
	note.Value = com.Value
	l.Trace("Note: ", fmt.Sprintf("%+v", note))

	return com.Write(note)
}

func (com *Command) runList() error {
	if com.Project == "" {
		return com.runListProjects()
	} else {
		return com.runListProjectNotes()
	}
}

func (com *Command) runListProjects() error {
	projects, err := com.getProjects()
	if err != nil {
		return err
	}

	for _, project := range projects {
		fmt.Println(project)
	}

	return nil
}

func (com *Command) getProjects() ([]string, error) {
	dir, err := ioutil.ReadDir(com.DataPath)
	if err != nil {
		return nil, err
	}

	var out []string
	for _, d := range dir {
		file := d.Name()

		// Skip dotfiles
		if strings.HasPrefix(file, ".") {
			continue
		}

		ext := filepath.Ext(file)
		name := file[0 : len(file)-len(ext)]

		out = append(out, name)
	}

	sort.Strings(out)
	return out, nil
}

func (com *Command) runListProjectNotes() error {
	if com.Project == "" {
		return errgo.New("project name can not be empty")
	}

	if !com.checkProjectExists() {
		return errgo.New("no notes for this project")
	}

	records, err := com.getProjectRecords(com.Project)
	if err != nil {
		return err
	}

	for _, record := range records {
		if record.GetAction() != ActionNote {
			continue
		}

		fmt.Println("#", record.GetTimeStamp())
		fmt.Println(record.GetValue())
		fmt.Println("")
	}

	return nil
}

func (com *Command) checkProjectExists() bool {
	projects, err := com.getProjects()
	if err != nil {
		return false
	}

	for _, project := range projects {
		if project == com.Project {
			return true
		}
	}

	return false
}

func (com *Command) getProjectRecords(project string) ([]Record, error) {
	if com.DataPath == "" {
		return nil, errgo.New("datapath can not be empty")
	}
	if project == "" {
		return nil, errgo.New("project name can not be empty")
	}

	filepath := filepath.Join(com.DataPath, project+".csv")
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0640)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, errgo.New(err.Error())
	}

	var out []Record
	for _, d := range records {
		record, err := RecordFromCSV(d)
		if err != nil {
			return nil, err
		}
		record.SetProject(project)

		out = append(out, record)
	}

	return out, err
}

func (com *Command) Write(record Record) error {
	if com.DataPath == "" {
		return errgo.New("datapath can not be empty")
	}

	if com.Project == "" {
		return errgo.New("project name can not be empty")
	}

	path := com.DataPath
	project := com.Project

	err := os.MkdirAll(path, 0750)
	if err != nil {
		return err
	}

	filepath := filepath.Join(path, project+".csv")
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0640)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.Write(record.CSV())
	if err != nil {
		return err
	}
	writer.Flush()

	return nil
}
