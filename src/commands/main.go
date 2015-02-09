package commands

import (
	"os"

	"github.com/AlexanderThaller/cobra"
	"github.com/AlexanderThaller/logger"
	"github.com/juju/errgo"
)

//BuildName represents the name of the software
var BuildName string

//BuildVersion represents the version of the software
var BuildVersion string

//BuildHash represents the hash from git with which the software was build
var BuildHash string

//BuildTime represents the time when the software was build
var BuildTime string

func Execute() {
	lablogCmd.AddCommand(cmdList)
	lablogCmd.AddCommand(cmdMerge)
	lablogCmd.AddCommand(cmdNote)
	lablogCmd.AddCommand(cmdRename)
	lablogCmd.AddCommand(cmdTodo)
	lablogCmd.AddCommand(cmdTrack)
	lablogCmd.AddCommand(cmdVersion)
	lablogCmd.AddCommand(cmdWeb)

	lablogCmd.Execute()
}

func errexit(l logger.Logger, err error, message string) {
	if err != nil {
		l.Alert(message, ": ", err)
		l.Debug(message, ": ", errgo.Details(err))
		os.Exit(1)
	}
}

func setLogLevel(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "setLogLevel")
	prio, err := logger.ParsePriority(flagLablogLogLevel)
	errexit(l, err, "can not parse loglevel")

	logger.SetLevel(".", prio)
	l.Debug("New loglevel is: ", flagLablogLogLevel)

	l.Debug("Args: ", args)
}

func finished(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "finished")
	l.Debug("Args: ", args)
	l.Info("Finished")
}
