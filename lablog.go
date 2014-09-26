package main

import (
	"flag"
	"os"
	"time"

	"github.com/AlexanderThaller/logger"
	"github.com/juju/errgo"
)

var (
	argStartTime = flag.String("starttime", "", "StartTime")
	argEndTime   = flag.String("endtime", "", "EndTime")
)

const (
	Name = "lablog"
)

func init() {
	flag.Parse()
}

func main() {
	l := logger.New(Name, "main")
	args := os.Args
	l.Debug("Args: ", args)
	l.Debug("Args without flags: ", flag.Args())

	conf := NewConfig()
	conf.DataPath = "/home/thalleralexander/docs/lablog"
	conf.SCM = "git"
	conf.SCMAutoCommit = true
	conf.SCMAutoPush = false

	command, err := parseCommand(flag.Args())
	if err != nil {
		l.Alert("Problem while parsing command: ", errgo.Details(err))
		os.Exit(1)
	}
	command.Config = conf

	l.Debug("StartDate: ", *argStartTime)
	if *argStartTime != "" {
		l.Debug("StartDate: ", *argStartTime)
		command.StartTime, err = time.Parse(DateFormatDay, *argStartTime)
		if err != nil {
			l.Alert("Can not parse startDate: ", err)
			os.Exit(1)
		}
		l.Debug("StartDate: ", command.StartTime)
	}

	if *argEndTime != "" {
		command.EndTime, err = time.Parse(DateFormatDay, *argEndTime)
		if err != nil {
			l.Alert("can not parse endDate: ", err)
			os.Exit(1)
		}
	}

	err = command.Run()
	if err != nil {
		l.Alert("Problem while running command: ", errgo.Details(err))
		os.Exit(1)
	}

	os.Exit(0)
}
