package commands

import (
	"github.com/AlexanderThaller/cobra"
	"github.com/AlexanderThaller/lablog/src/web"
	"github.com/AlexanderThaller/logger"
)

var cmdWeb = &cobra.Command{
	Use:    "web",
	Short:  "Serve the lablog data as a webpage.",
	Long:   `Will listen and serve all notes and todos formatted as html.`,
	Run:    runWeb,
	PreRun: setLogLevel,
}

var flagWebBind string

func init() {
	cmdWeb.Flags().StringVarP(&flagWebBind, "bind", "b", ":18080",
		"The address and port to bind to.")
}

func runWeb(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "web")

	l.Notice("Listening on ", flagWebBind)
	err := web.Listen(flagLablogDataDir, flagWebBind)
	errexit(l, err, "can not serve content")
}
