package commands

import "github.com/spf13/cobra"

var cmdRename = &cobra.Command{
	Use:   "rename [src] [dst]",
	Short: "Rename projects",
	Long:  `Rename projects`,
	Run:   runRename,
}

func runRename(cmd *cobra.Command, args []string) {
}
