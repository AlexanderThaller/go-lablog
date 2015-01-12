package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func getCmdVersion() *cobra.Command {
	cmdVersion.Flags().BoolVarP(&flagVersionBuildTime, "build", "b", false,
		"also print when the software was build")
	cmdVersion.Flags().BoolVarP(&flagVersionBuildHash, "git", "g", false,
		"also print with which git version the software was build")

	return cmdVersion
}

var cmdVersion = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of aptly-manager",
	Long:  `All software has versions. This is aptly-manager's`,
	Run: func(cmd *cobra.Command, args []string) {
		out := "v" + BuildVersion

		if flagVersionBuildHash {
			out += "-" + BuildHash
		}

		if flagVersionBuildTime {
			out += " b" + BuildTime
		}

		fmt.Println(out)
	},
}

var flagVersionBuildTime bool
var flagVersionBuildHash bool
