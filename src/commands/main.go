package commands

import (
	"os"

	"github.com/AlexanderThaller/logger"
	"github.com/juju/errgo"
)

func init() {
	logger.SetLevel(".", logger.Trace)
}

//BuildName represents the name of the software
var BuildName string

//BuildVersion represents the version of the software
var BuildVersion string

//BuildHash represents the hash from git with which the software was build
var BuildHash string

//BuildTime represents the time when the software was build
var BuildTime string

func Execute() {
	AddCommands()
	lablogCmd.Execute()
}

func AddCommands() {
	lablogCmd.AddCommand(cmdDone)
	lablogCmd.AddCommand(cmdList)
	lablogCmd.AddCommand(cmdMerge)
	lablogCmd.AddCommand(cmdNote)
	lablogCmd.AddCommand(cmdRename)
	lablogCmd.AddCommand(cmdTodo)
	lablogCmd.AddCommand(cmdTrack)
	lablogCmd.AddCommand(cmdVersion)
	lablogCmd.AddCommand(cmdWeb)
}

func errexit(l logger.Logger, err error, message string) {
	if err != nil {
		l.Alert(message, ": ", err)
		l.Debug(message, ": ", errgo.Details(err))
		os.Exit(1)
	}
}
