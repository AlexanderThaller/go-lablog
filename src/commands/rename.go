package commands

import (
	"os"

	"github.com/AlexanderThaller/cobra"
	"github.com/AlexanderThaller/logger"
)

var cmdRename = &cobra.Command{
	Use:    "rename [src] [dst]",
	Short:  "Rename project from src to dst.",
	Long:   "Rename project from src to dst.",
	Run:    runRename,
	PreRun: setLogLevel,
}

func runRename(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "rename")
	l.Alert("not implemented")
	os.Exit(1)
}
