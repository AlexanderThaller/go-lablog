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

var (
	buildTime    string
	buildVersion string
)

func runVersion(cmd *cobra.Command, args []string) {
	fmt.Printf("lablog v%v-b%v\n", buildVersion, buildTime)
}
