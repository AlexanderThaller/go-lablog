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

func init() {
	log.SetLevel(log.DebugLevel)

	homepath, err := homedir.Dir()
	helper.ErrExit(errgo.Notef(err, "can not get homepath"))

	datadir := path.Join(homepath, ".lablog")
	cmdMain.PersistentFlags().StringVarP(&flagDataDir, "datadir", "d",
		datadir, "The path to the datadir for retreiving and storing the data.")
}

var cmdMain = &cobra.Command{
	Use:   "lablog [command]",
	Short: "lablog makes taking notes and todos easy.",
	Long:  `lablog orders notes and todos into projects and subprojects without dictating a specific format.`,
}

func Run() {
	cmdMain.AddCommand(cmdVersion)

	// show
	cmdMain.AddCommand(cmdShow)
	cmdShow.AddCommand(cmdShowProjects)
	cmdShow.AddCommand(cmdShowNotes)

	// add
	cmdMain.AddCommand(cmdAdd)
	cmdAdd.AddCommand(cmdAddNote)

	cmdMain.Execute()
}
