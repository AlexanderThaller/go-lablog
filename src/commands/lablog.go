package commands

import (
	"os"
	"path"

	"github.com/AlexanderThaller/logger"
	"github.com/juju/errgo"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var lablogCmd = &cobra.Command{
	Use:   BuildName,
	Short: BuildName + " makes taking notes and todos simple",
	Long: BuildName + ` orders notes and todos into projects and subprojects
  without dictating a specific format`,
}

var flagLablogDataDir string
var flagLablogLogLevel string

func init() {
	l := logger.New("commands")
	homepath, err := homedir.Dir()
	if err != nil {
		l.Alert("can not get homepath: ", errgo.Details(err))
		os.Exit(1)
	}

	datadir := path.Join(homepath, ".lablog")
	lablogCmd.PersistentFlags().StringVarP(&flagLablogDataDir, "datadir", "d",
		datadir, "The path to the datadir we will use with lablog")

	lablogCmd.PersistentFlags().StringVarP(&flagLablogLogLevel, "loglevel", "l",
		"Notice", "The loglevel for this run")
}
