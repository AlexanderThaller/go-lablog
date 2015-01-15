package commands

import "github.com/spf13/cobra"

var cmdDone = &cobra.Command{
	Use:   "done [project] [text]",
	Short: "Done track",
	Long:  `Done track`,
	Run:   runDone,
}

func runDone(cmd *cobra.Command, args []string) {
}
