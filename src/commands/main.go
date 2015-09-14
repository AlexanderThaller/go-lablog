package commands

import (
	"path"

	log "github.com/Sirupsen/logrus"

	"github.com/AlexanderThaller/lablog/src/helper"

	"github.com/AlexanderThaller/cobra"
	"github.com/juju/errgo"
	"github.com/mitchellh/go-homedir"
)

var flagDataDir string
var flagLogLevel string

func init() {
	homepath, err := homedir.Dir()
	helper.ErrExit(errgo.Notef(err, "can not get homepath"))

	datadir := path.Join(homepath, ".lablog")
	cmdMain.PersistentFlags().StringVarP(&flagDataDir, "datadir", "d",
		datadir, "The path to the datadir for retreiving and storing the data.")
	cmdMain.PersistentFlags().StringVarP(&flagLogLevel, "loglevel", "l",
		"warn", "The loglevel for which to run in. Default is warn. There are panic, fatal, error, warn info and debug as levels.")
}

var cmdMain = &cobra.Command{
	Use:              "lablog [command]",
	Short:            "lablog makes taking notes and todos easy.",
	Long:             `lablog orders notes and todos into projects and subprojects without dictating a specific format.`,
	PersistentPreRun: setLogLevel,
}

func Run() {
	cmdMain.AddCommand(cmdVersion)

	// show
	cmdMain.AddCommand(cmdShow)
	cmdShow.AddCommand(cmdShowProjects)
	cmdShow.AddCommand(cmdShowNotes)
	cmdShow.AddCommand(cmdShowTodos)

	// add
	cmdMain.AddCommand(cmdAdd)
	cmdAdd.AddCommand(cmdAddNote)
	cmdAdd.AddCommand(cmdAddTodo)

	// todo
	cmdAddTodo.AddCommand(cmdAddTodoActive)
	cmdAddTodo.AddCommand(cmdAddTodoInActive)

	cmdMain.Execute()
}

func setLogLevel(cmd *cobra.Command, args []string) {
	level, err := log.ParseLevel(flagLogLevel)
	if err != nil {
		log.Fatal(errgo.Notef(err, "can not parse loglevel from flag"))
	}

	log.SetLevel(level)
}
