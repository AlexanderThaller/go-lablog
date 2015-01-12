package commands

import (
	"github.com/AlexanderThaller/lablog/src/web"
	"github.com/AlexanderThaller/logger"
	"github.com/juju/errgo"
	"github.com/spf13/cobra"
)

func getCmdWeb() *cobra.Command {
	cmdWeb.Flags().StringVarP(&flagWebBind, "bind", "b", ":8080",
		"the address and port to bind to")

	return cmdWeb
}

var cmdWeb = &cobra.Command{
	Use:   "web",
	Short: "Serve the lablog data as a webpage",
	Long:  `Will listen and serve all notes and todos formatted as html`,
	Run: func(cmd *cobra.Command, args []string) {
		l := logger.New("commands", "web")
		l.Notice("listening on " + flagWebBind)

		web.DataDir = flagMainDataDir
		err := web.Listen(flagWebBind)
		if err != nil {
			l.Alert("can not serve content: ", errgo.Details(err))
		}
	},
}

var flagWebBind string
