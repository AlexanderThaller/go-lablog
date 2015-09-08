package commands

import (
	"fmt"
	"sort"

	"github.com/AlexanderThaller/lablog/src/data"
	"github.com/AlexanderThaller/lablog/src/helper"
	"github.com/AlexanderThaller/lablog/src/store"
	"github.com/juju/errgo"

	"github.com/AlexanderThaller/cobra"
)

var cmdShow = &cobra.Command{
	Use:   "show [command]",
	Short: "Show current projects and entries.",
	Long:  `Show a list of currently available projects or their entries like notes, todos, tracks, etc., see help for all options.`,
	Run:   runCmdShow,
}

var cmdShowProjects = &cobra.Command{
	Use:   "projects",
	Short: "Show current projects.",
	Long:  `Show all projects which currently have any type of entry.`,
	Run:   runCmdShowProjects,
}

func runCmdShow(cmd *cobra.Command, args []string) {
}

func runCmdShowProjects(cmd *cobra.Command, args []string) {
	store, err := store.NewFolderStore(flagDataDir)
	helper.ErrExit(errgo.Notef(err, "can not get data store"))

	projects, err := store.ListProjects()
	helper.ErrExit(errgo.Notef(err, "can not get list of projects"))

	sort.Strings(data.ProjectNamesToString(projects))

	for _, project := range projects {
		fmt.Println(project.String())
	}
}
