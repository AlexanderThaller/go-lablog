package commands

import (
	"os"
	"strings"
	"time"

	"github.com/AlexanderThaller/cobra"
	"github.com/AlexanderThaller/lablog/src/data"
	"github.com/AlexanderThaller/lablog/src/helper"
	"github.com/AlexanderThaller/logger"
	"github.com/juju/errgo"
)

var cmdTodo = &cobra.Command{
	Use:   "todo [project] [text]",
	Short: "Todo projects.",
	Long:  `Todo projects.`,
	Run:   runTodoToggle,
}

var cmdTodoStart = &cobra.Command{
	Use:   "start [project] [text]",
	Short: "Start todo for project",
	Long:  `Start todo for project`,
	Run:   runTodoStart,
}

var cmdTodoStop = &cobra.Command{
	Use:   "stop [project] [text]",
	Short: "Stop todo for project",
	Long:  `Stop todo for project`,
	Run:   runTodoStop,
}

var cmdTodoToggle = &cobra.Command{
	Use:   "toggle [project] [text]",
	Short: "Toggle todo for project",
	Long:  `Toggle todo for project`,
	Run:   runTodoToggle,
}

var flagTodoTimeStamp time.Time
var flagTodoTimeStampRaw string

func init() {
	flagTodoTimeStamp = time.Now()

	cmdTodo.PersistentFlags().StringVarP(&flagTodoTimeStampRaw, "timestamp", "t",
		flagTodoTimeStamp.String(), "The timestamp for which to record the todo.")
}

func runTodoToggle(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "todo", "toggle")
	l.Alert("not implemented")
	os.Exit(1)
}

func runTodoStart(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "todo", "start")

	if len(args) < 2 {
		errexit(l, errgo.New("need at least two arguments to run"))
	}

	project := args[0]
	text := strings.Join(args[1:], " ")

	timestamp, err := helper.DefaultOrRawTimestamp(flagTodoTimeStamp, flagTodoTimeStampRaw)
	errexit(l, err, "can not get timestamp")

	todo := data.Todo{
		Done:      false,
		Project:   data.Project{Name: project},
		Text:      text,
		TimeStamp: timestamp,
	}

	l.Trace("Todo: ", todo)
	recordAndCommit(l, flagLablogDataDir, todo)
}

func runTodoStop(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "todo", "stop")
	l.Alert("not implemented")
	os.Exit(1)
}
