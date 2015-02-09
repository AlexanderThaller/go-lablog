package commands

import (
	"os"

	"github.com/AlexanderThaller/logger"
	"github.com/spf13/cobra"
)

var cmdRename = &cobra.Command{
	Use:   "rename [src] [dst]",
	Short: "Rename project from src to dst",
	Long:  "Rename project from src to dst",
	Run:   runRename,
}

func runRename(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "rename")
	l.Alert("not implemented")
	os.Exit(1)
}
