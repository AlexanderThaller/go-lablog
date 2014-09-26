package main

import (
	"os"

	"github.com/AlexanderThaller/logger"
	"github.com/juju/errgo"
)

const (
	Name = "lablog"
)

func main() {
	l := logger.New(Name, "main")
	args := os.Args
	l.Debug("Args: ", args)

	conf := NewConfig()
	conf.DataPath = "/home/thalleralexander/docs/lablog"
	conf.SCM = "git"
	conf.SCMAutoCommit = true
	conf.SCMAutoPush = false

	command, err := parseCommand(args)
	if err != nil {
		l.Alert("Problem while parsing command: ", errgo.Details(err))
		os.Exit(1)
	}
	command.Config = conf

	err = command.Run()
	if err != nil {
		l.Alert("Problem while running command: ", errgo.Details(err))
		os.Exit(1)
	}

	os.Exit(0)
}
