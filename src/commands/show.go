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

func runCmdShow(cmd *cobra.Command, args []string) {
}

var cmdShowProjects = &cobra.Command{
	Use:   "projects",
	Short: "Show current projects.",
	Long:  `Show all projects which currently have any type of entry.`,
	Run:   runCmdShowProjects,
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

var cmdShowNotes = &cobra.Command{
	Use:   "notes",
	Short: "Show notes",
	Long:  `Show all notes`,
	Run:   runCmdShowNotes,
}

func runCmdShowNotes(cmd *cobra.Command, args []string) {
	store, err := store.NewFolderStore(flagDataDir)
	helper.ErrExit(errgo.Notef(err, "can not get data store"))

	var names []data.ProjectName

	if len(args) == 0 {
		names, err = store.ListProjects()
		helper.ErrExit(errgo.Notef(err, "can not get list of projects"))
	} else {
		for _, arg := range args {
			project, err := data.ParseProjectName(arg)
			helper.ErrExit(errgo.Notef(err, "can not parse project name"))

			names = append(names, project)
		}
	}

	sort.Strings(data.ProjectNamesToString(names))

	for _, name := range names {
		project, err := store.GetProject(name)
		helper.ErrExit(errgo.Notef(err, "can not get list of projects"))

		fmt.Println("=", name.String())

		for _, note := range project.Notes() {
			fmt.Println("==", note.TimeStamp.String())
			fmt.Println(note.Value)
		}
	}
}
