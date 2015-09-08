package commands

import (
	"fmt"

	"github.com/AlexanderThaller/cobra"
)

var cmdShow = &cobra.Command{
	Use:   "show [command]",
	Short: "Show current projects and entries.",
	Long:  `Show a list of currently available projects or their entries like notes, todos, tracks, etc., see help for all options.`,
}

var cmdShowProjects = &cobra.Command{
	Use:   "projects",
	Short: "Show current projects.",
	Long:  `Show all projects which currently have any type of entry.`,
	Run:   runCmdShowProjects,
}

func runCmdShowProjects(cmd *cobra.Command, args []string) {
	fmt.Println("no projects yet")
}
