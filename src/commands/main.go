package commands

import (
	"os"

	"github.com/AlexanderThaller/cobra"
	"github.com/AlexanderThaller/lablog/src/data"
	"github.com/AlexanderThaller/lablog/src/store"
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
	lablogCmd.AddCommand(cmdSearch)
	lablogCmd.AddCommand(cmdTodo)
	lablogCmd.AddCommand(cmdTrack)
	lablogCmd.AddCommand(cmdVersion)
	lablogCmd.AddCommand(cmdWeb)

	cmdList.AddCommand(cmdListDates)
	cmdList.AddCommand(cmdListEntries)
	cmdList.AddCommand(cmdListTimeline)
	cmdList.AddCommand(cmdListNotes)
	cmdList.AddCommand(cmdListProjects)
	cmdList.AddCommand(cmdListTodos)
	cmdList.AddCommand(cmdListTracks)
	cmdList.AddCommand(cmdListTracksActive)
	cmdList.AddCommand(cmdListTracksDurations)

	cmdTodo.AddCommand(cmdTodoStart)
	cmdTodo.AddCommand(cmdTodoDone)
	cmdTodo.AddCommand(cmdTodoToggle)

	cmdTrack.AddCommand(cmdTrackStart)
	cmdTrack.AddCommand(cmdTrackStop)

	lablogCmd.Execute()
}

func errexit(l logger.Logger, err error, message ...string) {
	if err != nil {
		l.Alert(message, ": ", err)
		l.Trace(message, ": ", errgo.Details(err))
		os.Exit(1)
	}
}

func setLogLevel(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "setLogLevel")
	prio, err := logger.ParsePriority(flagLablogLogLevel)
	errexit(l, err, "can not parse loglevel")

	logger.SetLevel(".", prio)
	l.Debug("New loglevel is: ", flagLablogLogLevel)
}

func finished(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "finished")
	l.Info("Finished")
}

func recordAndCommit(l logger.Logger, datadir string, entry data.Entry) {
	store := store.NewFlatFile(datadir)
	err := store.Write(entry)
	errexit(l, err, "can not record note")
}
