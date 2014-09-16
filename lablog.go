package main

import (
	"os"

	"github.com/AlexanderThaller/logger"
	"github.com/juju/errgo"
)

const (
	Name = "lablog"
)

func init() {
	logger.SetLevel(".", logger.Trace)
}

func main() {
	l := logger.New(Name, "main")
	args := os.Args
	conf := NewConfig()
	conf.DataPath = "/tmp/lablog"

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
