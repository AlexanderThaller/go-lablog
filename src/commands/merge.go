package commands

import (
	"os"

	"github.com/AlexanderThaller/logger"
	"github.com/spf13/cobra"
)

var cmdMerge = &cobra.Command{
	Use:   "merge [src] [dst]",
	Short: "Merge projects",
	Long:  `Merge projects`,
	Run:   runMerge,
}

func runMerge(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "done")
	l.Alert("not implemented")
	os.Exit(1)
}
