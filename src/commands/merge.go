package commands

import (
	"os"

	"github.com/AlexanderThaller/cobra"
	"github.com/AlexanderThaller/logger"
)

var cmdMerge = &cobra.Command{
	Use:   "merge [src] [dst]",
	Short: "Merge all entries of src project into the dst project.",
	Long:  "Merge all entries of src project into the dst project.",
	Run:   runMerge,
}

func runMerge(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "done")
	l.Alert("not implemented")
	os.Exit(1)
}
