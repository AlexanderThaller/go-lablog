package commands

import (
	"os"

	"github.com/AlexanderThaller/logger"
	"github.com/spf13/cobra"
)

var cmdDone = &cobra.Command{
	Use:   "done [project] [text]",
	Short: "Done track",
	Long:  `Done track`,
	Run:   runDone,
}

func runDone(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "done")
	l.Alert("not implemented")
	os.Exit(1)
}
