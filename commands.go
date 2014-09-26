package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
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
}

const (
	ActionList = "list"
	ActionNote = "note"
)

func NewCommand() *Command {
	return new(Command)
}

func (com *Command) Run() error {
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
	if len(com.Args) != 1 {
		return errgo.New("note command needs one argument")
	}
	l.Trace("Project: ", com.Project)
	if com.Project == "" {
		return errgo.New("note command needs an project")
	}

	note := new(Note)
	note.Project = com.Project
	note.TimeStamp = com.TimeStamp
	note.Value = com.Args[0]
	l.Trace("Note: ", fmt.Sprintf("%+v", note))

	return com.Write(note)
}

func (com *Command) runList() error {
	return errgo.New("Not implemented")
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
