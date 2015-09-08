package commands

import (
	"strings"
	"time"

	"github.com/AlexanderThaller/lablog/src/data"
	"github.com/AlexanderThaller/lablog/src/helper"
	"github.com/juju/errgo"

	"github.com/AlexanderThaller/cobra"
)

var flagAddTimeStamp time.Time
var flagAddTimeStampRaw string

func init() {
	flagAddTimeStamp = time.Now()

	cmdAdd.PersistentFlags().StringVarP(&flagAddTimeStampRaw, "timestamp", "t",
		flagAddTimeStamp.String(), "The timestamp for which to record the note.")
}

var cmdAdd = &cobra.Command{
	Use:   "add [command]",
	Short: "Add a new entry to the log",
	Long:  `Add a new entry like a note or a todo to the log. You have to specify a project for which we want to record the log for.`,
	Run:   runCmdAdd,
}

func runCmdAdd(cmd *cobra.Command, args []string) {
}

var cmdAddNote = &cobra.Command{
	Use:   "note",
	Short: "Add current projects.",
	Long:  `Add all projects which currently have any type of entry.`,
	Run:   runCmdAddNote,
}

func runCmdAddNote(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		helper.Fatal(errgo.New("need at least two arguments to run"))
	}

	project, err := data.ParseProjectName(args[0])
	helper.ErrExit(errgo.Notef(err, "can not parse project name"))

	value := strings.Join(args[1:], " ")

	timestamp, err := helper.DefaultOrRawTimestamp(flagAddTimeStamp, flagAddTimeStampRaw)
	helper.ErrExit(errgo.Notef(err, "can not get timestamp"))

	note := data.Note{
		Value:     value,
		TimeStamp: timestamp,
	}

	err = helper.RecordEntry(flagDataDir, project, note)
	helper.ErrExit(errgo.Notef(err, "can not record note to store"))
}
