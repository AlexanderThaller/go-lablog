package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/AlexanderThaller/logger"
	"github.com/juju/errgo"
)

var (
	flagDataPath      = flag.String("datapath", "/home/thalleralexander/.lablog", "")
	flagSCM           = flag.String("scm", "git", "")
	flagSCMAutoCommit = flag.Bool("autocommit", true, "")
	flagSCMAutoPush   = flag.Bool("autopush", false, "")
	flagAction        = flag.String("c", "list", "")
	flagProject       = flag.String("p", "", "")
	flagStartTime     = flag.String("starttime", "", "StartTime")
	flagEndTime       = flag.String("endtime", "", "EndTime")
)

const (
	Name = "lablog"
)

func init() {
	flag.Parse()
	logger.SetLevel(logger.New(Name, "main"), logger.Trace)
}

func main() {
	l := logger.New(Name, "main")

	command := NewCommand()
	command.DataPath = *flagDataPath
	command.SCM = *flagSCM
	command.SCMAutoCommit = *flagSCMAutoCommit
	command.SCMAutoPush = *flagSCMAutoPush
	command.Project = *flagProject
	command.Action = *flagAction
	command.StartTime = *flagStartTime
	command.EndTime = *flagEndTime

	l.Trace("Command: ", fmt.Sprintf("%#v", command))

	err := command.Run()
	if err != nil {
		l.Alert("Problem while running command: ", errgo.Details(err))
		os.Exit(1)
	}

	os.Exit(0)
}
