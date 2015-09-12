package commands

import (
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
	Short: "Add a new entry to the log.",
	Long:  `Add a new entry like a note or a todo to the log. You have to specify a project for which we want to record the log for.`,
	Run:   runCmdAdd,
}

func runCmdAdd(cmd *cobra.Command, args []string) {
	cmd.Help()
}

var cmdAddNote = &cobra.Command{
	Use:   "note",
	Short: "Add a new note to the log.",
	Long:  `Add a new note to the log which can have a timestamp and an free form value for text.`,
	Run:   runCmdAddNote,
}

func runCmdAddNote(cmd *cobra.Command, args []string) {
	project, timestamp, value, err := helper.ArgsToEntryValues(args, flagAddTimeStamp, flagAddTimeStampRaw)
	helper.ErrExit(errgo.Notef(err, "can not convert args to entry usable values"))

	note := data.Note{
		Value:     value,
		TimeStamp: timestamp,
	}

	err = helper.RecordEntry(flagDataDir, project, note)
	helper.ErrExit(errgo.Notef(err, "can not record note to store"))
}

var cmdAddTodo = &cobra.Command{
	Use:   "todo [command]",
	Short: "Add a new todo to the log.",
	Long:  `Add a new todo to the log which can have a timestamp, a toggle state (if its active or not) and an free form value for text.`,
	Run:   runCmdAddTodo,
}

func runCmdAddTodo(cmd *cobra.Command, args []string) {
	cmd.Help()
}

var cmdAddTodoActive = &cobra.Command{
	Use:   "active",
	Short: "Add a new todo to the log and mark it as active.",
	Long:  `Add a new todo to the log which can have a timestamp, is marked as active and an free form value for text.`,
	Run:   runCmdAddTodoActive,
}

func runCmdAddTodoActive(cmd *cobra.Command, args []string) {
	project, todo, err := helper.ArgsToTodo(args, flagAddTimeStamp, flagAddTimeStampRaw)
	helper.ErrExit(errgo.Notef(err, "can not convert args to todo"))

	todo.Active = true

	err = helper.RecordEntry(flagDataDir, project, todo)
	helper.ErrExit(errgo.Notef(err, "can not record todo to store"))
}

var cmdAddTodoInActive = &cobra.Command{
	Use:   "inactive",
	Short: "Add a new todo to the log and mark it as inactive.",
	Long:  `Add a new todo to the log which can have a timestamp, is marked as inactive and an free form value for text.`,
	Run:   runCmdAddTodoInActive,
}

func runCmdAddTodoInActive(cmd *cobra.Command, args []string) {
	project, todo, err := helper.ArgsToTodo(args, flagAddTimeStamp, flagAddTimeStampRaw)
	helper.ErrExit(errgo.Notef(err, "can not convert args to todo"))

	todo.Active = false

	err = helper.RecordEntry(flagDataDir, project, todo)
	helper.ErrExit(errgo.Notef(err, "can not record todo to store"))
}
