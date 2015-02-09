package commands

import (
	"os"

	"github.com/AlexanderThaller/lablog/src/web"
	"github.com/AlexanderThaller/logger"
	"github.com/spf13/cobra"
)

var cmdWeb = &cobra.Command{
	Use:   "web",
	Short: "Serve the lablog data as a webpage",
	Long:  `Will listen and serve all notes and todos formatted as html`,
	Run:   runWeb,
}

var flagWebBind string

func init() {
	cmdWeb.Flags().StringVarP(&flagWebBind, "bind", "b", ":8080",
		"the address and port to bind to")
}

func runWeb(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "web")
	l.Alert("not implemented")
	os.Exit(1)

	err := web.Listen(flagLablogDataDir, flagWebBind)
	errexit(l, err, "can not serve content")
}
