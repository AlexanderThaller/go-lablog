package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"time"

	"bitbucket.org/kardianos/osext"

	"github.com/AlexanderThaller/go-lablog/Godeps/_workspace/src/github.com/mitchellh/go-homedir"
	"github.com/AlexanderThaller/logger"
	"github.com/davecheney/profile"
	"github.com/jinzhu/now"
	"github.com/juju/errgo"
)

var (
	buildVersion string
	buildTime    string

	flagAction        = flag.String("c", "list", "The action to run")
	flagDataPath      *string
	flagEndTime       = flag.String("endtime", time.Now().String(), "The endtime for the timerange filter")
	flagProject       = flag.String("p", "", "The project to use")
	flagSCM           = flag.String("scm", "git", "The sourcecode management to use")
	flagSCMAutoCommit = flag.Bool("autocommit", true, "Auto commit new records to the scm")
	flagSCMAutoPush   = flag.Bool("autopush", false, "Auto push new records to the remote of the scm")
	flagStartTime     = flag.String("starttime", time.Time{}.String(), "The starttime for the timerange filter")
	flagValue         = flag.String("v", "", "The value which is used by certain commands")
	flagLogLevel      = flag.String("loglevel", "Notice", "The loglevel")
	flagNoSubprojects = flag.Bool("nosubprojects", false, "If true we will not print records for subprojects")
	flagProfile       = flag.String("profile", "", "Path to folder where profile data will be saved.")
)

const (
	Name = "lablog"
)

func init() {
	l := logger.New(Name, "init")
	runtime.GOMAXPROCS(runtime.NumCPU())

	home, err := homedir.Dir()
	if err != nil {
		home = ""
		l.Warning("Can not get homedir: ", err)
	}
	flagDataPath = flag.String("datapath", home+"/.lablog", "The path to the datafolder")

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

	if *flagProfile != "" {
		prof, err := configProfile(*flagProfile)
		if err != nil {
			l.Alert("Can not start profiling: ", errgo.Details(err))
			os.Exit(1)
		}

		defer prof.Stop()
	}

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
	command.NoSubprojects = *flagNoSubprojects

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
}

// configProfile will start profiling based on the default profile
// settings.
func configProfile(folder string) (interface {
	Stop()
}, error) {

	timestamp := time.Now().Format(time.RFC3339Nano)
	folderpath := path.Join(folder, timestamp)

	prof := profile.Config{
		CPUProfile:     true,
		MemProfile:     true,
		NoShutdownHook: true,       // do not hook SIGINT
		ProfilePath:    folderpath, // store profiles in current directory
		Quiet:          true,
	}

	err := os.MkdirAll(folderpath, 0755)
	if err != nil {
		return nil, err
	}

	binarypath, err := osext.Executable()
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(binarypath)
	if err != nil {
		return nil, err
	}

	filepath := path.Join(folderpath, "binary")
	err = ioutil.WriteFile(filepath, data, 0755)
	if err != nil {
		return nil, err
	}

	return profile.Start(&prof), nil
}
