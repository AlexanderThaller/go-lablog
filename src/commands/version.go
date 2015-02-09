package commands

import (
	"fmt"

	"github.com/AlexanderThaller/cobra"
)

var cmdVersion = &cobra.Command{
	Use:    "version",
	Short:  "Print the version number of aptly-manager.",
	Long:   `All software has versions. This is aptly-manager's.`,
	Run:    runVersion,
	PreRun: setLogLevel,
}

var flagVersionBuildTime bool
var flagVersionBuildHash bool

func init() {
	cmdVersion.Flags().BoolVarP(&flagVersionBuildTime, "build", "b", false,
		"also print when the software was build.")
	cmdVersion.Flags().BoolVarP(&flagVersionBuildHash, "git", "g", false,
		"also print with which git version the software was build.")
}

func runVersion(cmd *cobra.Command, args []string) {
	out := "v" + BuildVersion

	if flagVersionBuildHash {
		out += "-" + BuildHash
	}

	if flagVersionBuildTime {
		out += " b" + BuildTime
	}

	fmt.Println(out)
}
