package commands

import (
	"fmt"

	"github.com/AlexanderThaller/cobra"
)

var cmdVersion = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of lablog.",
	Long:  `All software has versions. This is the version of lablog.`,
	Run:   runVersion,
}

func runVersion(cmd *cobra.Command, args []string) {
	fmt.Println("lablog v0.0.1")
}
