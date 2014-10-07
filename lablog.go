package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/AlexanderThaller/logger"
	"github.com/jinzhu/now"
	"github.com/juju/errgo"
)

var (
	buildVersion string
	buildTime    string

	flagAction        = flag.String("c", "list", "")
	flagDataPath      = flag.String("datapath", "/home/thalleralexander/.lablog", "")
	flagEndTime       = flag.String("endtime", time.Now().String(), "EndTime")
	flagProject       = flag.String("p", "", "")
	flagSCM           = flag.String("scm", "git", "")
	flagSCMAutoCommit = flag.Bool("autocommit", true, "")
	flagSCMAutoPush   = flag.Bool("autopush", false, "")
	flagStartTime     = flag.String("starttime", time.Time{}.String(), "StartTime")
	flagValue         = flag.String("v", "", "")
	flagLogLevel      = flag.String("loglevel", "Notice", "")
)

const (
	Name = "lablog"
)

func init() {
	l := logger.New(Name, "init")
	flag.Parse()

	priority, err := logger.ParsePriority(*flagLogLevel)
	if err != nil {
		l.Alert("Can not parse loglevel: ", errgo.Details(err))
		os.Exit(1)
	}
	logger.SetLevel(".", priority)

	now.TimeFormats = append(now.TimeFormats, "2006-01-02 15:04:05 -0700 MST")
}

func main() {
	l := logger.New(Name, "main")
	l.Info("Version: ", buildVersion)
	l.Info("Buildtime: ", buildTime)

	command := NewCommand()
	command.Action = *flagAction
	command.Args = flag.Args()
	command.DataPath = *flagDataPath
	command.Project = *flagProject
	command.SCM = *flagSCM
	command.SCMAutoCommit = *flagSCMAutoCommit
	command.SCMAutoPush = *flagSCMAutoPush
	command.TimeStamp = time.Now()
	command.Value = *flagValue

	starttime, err := now.Parse(*flagStartTime)
	if err != nil {
		l.Alert("Can not parse starttime: ", err)
		os.Exit(1)
	}
	command.StartTime = starttime

	endtime, err := now.Parse(*flagEndTime)
	if err != nil {
		l.Alert("Can not parse endtime: ", err)
		os.Exit(1)
	}
	command.EndTime = endtime

	l.Trace("Command: ", fmt.Sprintf("%+v", command))

	err = command.Run()
	if err != nil {
		l.Alert("Problem while running command: ", errgo.Details(err))
		os.Exit(1)
	}

	os.Exit(0)
}
