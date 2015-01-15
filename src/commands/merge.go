package commands

import "github.com/spf13/cobra"

var cmdMerge = &cobra.Command{
	Use:   "merge [src] [dst]",
	Short: "Merge projects",
	Long:  `Merge projects`,
	Run:   runMerge,
}

func runMerge(cmd *cobra.Command, args []string) {
}
