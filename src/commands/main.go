package commands

import (
	"os"
	"path"

	"github.com/AlexanderThaller/logger"
	"github.com/juju/errgo"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
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
	AddCommands()
	getMain().Execute()
}

func AddCommands() {
	MainCmd.AddCommand(getCmdVersion())
	MainCmd.AddCommand(getCmdWeb())
}

func getMain() *cobra.Command {
	l := logger.New("commands", "getMain")
	homepath, err := homedir.Dir()
	if err != nil {
		l.Alert("can not get homepath: ", errgo.Details(err))
		os.Exit(1)
	}

	datadir := path.Join(homepath, ".lablog")
	MainCmd.PersistentFlags().StringVarP(&flagMainDataDir, "datadir", "d",
		datadir, "The path to the datadir we will use with lablog")

	return MainCmd
}

var MainCmd = &cobra.Command{
	Use:   BuildName,
	Short: BuildName + " makes taking notes and todos simple",
	Long: BuildName + ` orders notes and todos into projects and subprojects
  without dictating a specific format`,
}

var flagMainDataDir string
