package main

import "github.com/juju/errgo"

type Command struct {
	DataPath      string
	SCM           string
	SCMAutoCommit bool
	SCMAutoPush   bool
	Project       string
	Action        string
	StartTime     string
	EndTime       string
}

const (
	ActionNote = "note"
	ActionList = "list"
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
	return errgo.New("Not implemented")
}

func (com *Command) runList() error {
	return errgo.New("Not implemented")
}
