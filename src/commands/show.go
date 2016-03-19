package commands

import (
	"fmt"
	"os"

	"github.com/AlexanderThaller/lablog/src/formatting"
	"github.com/AlexanderThaller/lablog/src/helper"
	"github.com/juju/errgo"

	"github.com/AlexanderThaller/cobra"
)

var flagShowArchive bool

func init() {
	cmdShow.PersistentFlags().BoolVarP(&flagShowArchive, "archive", "a",
		false, "Determines if entries from the archive will be shown. (default is false)")
}

var cmdShow = &cobra.Command{
	Use:   "show [command]",
	Short: "Show current projects and entries.",
	Long:  `Show a list of currently available projects or their entries like notes, todos, tracks, etc., see help for all options.`,
	Run:   runCmdShow,
}

func runCmdShow(cmd *cobra.Command, args []string) {
	cmd.Help()
}

var cmdShowProjects = &cobra.Command{
	Use:   "projects",
	Short: "Show current projects.",
	Long:  `Show all projects which currently have any type of entry.`,
	Run:   runCmdShowProjects,
}

func runCmdShowProjects(cmd *cobra.Command, args []string) {
	store, err := helper.DefaultStore(flagDataDir)
	helper.ErrExit(errgo.Notef(err, "can not get data store"))

	projects, err := helper.ProjectNamesFromArgs(store, args, flagShowArchive)
	helper.ErrExit(errgo.Notef(err, "can not get list of projects"))

	for _, project := range projects.List() {
		fmt.Println(project.Name)
	}
}

var cmdShowEntries = &cobra.Command{
	Use:   "entries",
	Short: "Show entries",
	Long:  `Show all entries`,
	Run:   runCmdShowEntries,
}

func runCmdShowEntries(cmd *cobra.Command, args []string) {
	store, err := helper.DefaultStore(flagDataDir)
	helper.ErrExit(errgo.Notef(err, "can not get data store"))

	projects, err := helper.ProjectNamesFromArgs(store, args, flagShowArchive)
	helper.ErrExit(errgo.Notef(err, "can not get list of projects"))

	err = store.PopulateProjects(&projects)
	helper.ErrExit(errgo.Notef(err, "can not populate projects with entries"))

	formatting.Projects(os.Stdout, "Entries", 0, &projects)
}

var cmdShowNotes = &cobra.Command{
	Use:   "notes",
	Short: "Show notes",
	Long:  `Show all notes`,
	Run:   runCmdShowNotes,
}

func runCmdShowNotes(cmd *cobra.Command, args []string) {
	store, err := helper.DefaultStore(flagDataDir)
	helper.ErrExit(errgo.Notef(err, "can not get data store"))

	projects, err := helper.ProjectNamesFromArgs(store, args, flagShowArchive)
	helper.ErrExit(errgo.Notef(err, "can not get list of projects"))

	err = store.PopulateProjects(&projects)
	helper.ErrExit(errgo.Notef(err, "can not populate projects with entries"))

	formatting.ProjectsNotes(os.Stdout, "Notes", 0, &projects)
}

var cmdShowTodos = &cobra.Command{
	Use:   "todos",
	Short: "Show todos",
	Long:  `Show all todos`,
	Run:   runCmdShowTodos,
}

func runCmdShowTodos(cmd *cobra.Command, args []string) {
	store, err := helper.DefaultStore(flagDataDir)
	helper.ErrExit(errgo.Notef(err, "can not get data store"))

	projects, err := helper.ProjectNamesFromArgs(store, args, flagShowArchive)
	helper.ErrExit(errgo.Notef(err, "can not get list of projects"))

	err = store.PopulateProjects(&projects)
	helper.ErrExit(errgo.Notef(err, "can not populate projects with entries"))

	formatting.ProjectsTodos(os.Stdout, "Todos", 0, &projects)
}