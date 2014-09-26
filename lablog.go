package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/AlexanderThaller/logger"
	"github.com/juju/errgo"
)

var (
	flagAction        = flag.String("c", "list", "")
	flagDataPath      = flag.String("datapath", "/home/thalleralexander/.lablog", "")
	flagEndTime       = flag.String("endtime", "", "EndTime")
	flagProject       = flag.String("p", "", "")
	flagSCM           = flag.String("scm", "git", "")
	flagSCMAutoCommit = flag.Bool("autocommit", true, "")
	flagSCMAutoPush   = flag.Bool("autopush", false, "")
	flagStartTime     = flag.String("starttime", "", "StartTime")
	flagValue         = flag.String("v", "", "")
)

const (
	Name = "lablog"
)

func init() {
	flag.Parse()
}

func main() {
	l := logger.New(Name, "main")

	command := NewCommand()
	command.Action = *flagAction
	command.Args = flag.Args()
	command.DataPath = *flagDataPath
	command.EndTime = *flagEndTime
	command.Project = *flagProject
	command.SCM = *flagSCM
	command.SCMAutoCommit = *flagSCMAutoCommit
	command.SCMAutoPush = *flagSCMAutoPush
	command.StartTime = *flagStartTime
	command.TimeStamp = time.Now()
	command.Value = *flagValue

	l.Trace("Command: ", fmt.Sprintf("%+v", command))

	err := command.Run()
	if err != nil {
		l.Alert("Problem while running command: ", errgo.Details(err))
		os.Exit(1)
	}

	os.Exit(0)
}
