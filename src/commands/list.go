package commands

import (
	"fmt"
	"os"
	"sort"

	"github.com/AlexanderThaller/cobra"
	"github.com/AlexanderThaller/lablog/src/data"
	"github.com/AlexanderThaller/logger"
)

var cmdList = &cobra.Command{
	Use:    "list (command)",
	Short:  "List projects, notes, todos, dates, tracks, etc., see help for all options.",
	Long:   "List projects, notes, todos, dates, tracks, etc., see help for all options.",
	Run:    runListProjects,
	PreRun: setLogLevel,
}

var cmdListDates = &cobra.Command{
	Use:    "dates",
	Short:  "List dates.",
	Long:   `List dates.`,
	Run:    runListDates,
	PreRun: setLogLevel,
}

var cmdListNotes = &cobra.Command{
	Use:    "notes",
	Short:  "List notes.",
	Long:   `List notes.`,
	Run:    runListNotes,
	PreRun: setLogLevel,
}

var cmdListProjects = &cobra.Command{
	Use:    "projects",
	Short:  "List projects.",
	Long:   `List projects.`,
	Run:    runListProjects,
	PreRun: setLogLevel,
}

var cmdListTodos = &cobra.Command{
	Use:    "todos",
	Short:  "List todos.",
	Long:   `List todos.`,
	Run:    runListTodos,
	PreRun: setLogLevel,
}

var cmdListTracks = &cobra.Command{
	Use:    "tracks",
	Short:  "List tracks.",
	Long:   `List tracks.`,
	Run:    runListTracks,
	PreRun: setLogLevel,
}

var cmdListTracksActive = &cobra.Command{
	Use:    "active",
	Short:  "List tracks.",
	Long:   `List tracks.`,
	Run:    runListTracksActive,
	PreRun: setLogLevel,
}

var cmdListTracksDurations = &cobra.Command{
	Use:    "durations",
	Short:  "List durations.",
	Long:   `List durations.`,
	Run:    runListDurations,
	PreRun: setLogLevel,
}

func runListDates(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "list", "dates")
	l.Alert("not implemented")
	os.Exit(1)
}

func runListDurations(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "list", "durations")
	l.Alert("not implemented")
	os.Exit(1)
}

func runListNotes(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "list", "notes")
	l.Alert("not implemented")
	os.Exit(1)
}

func runListProjects(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "list", "projects")
	projects, err := data.Projects(flagLablogDataDir)
	errexit(l, err, "can not get projects")

	sort.Sort(data.ProjectsByName(projects))

	for _, project := range projects {
		fmt.Println(project.Name)
	}
}

func runListTodos(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "list", "todos")
	l.Alert("not implemented")
	os.Exit(1)
}

func runListTracks(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "list", "tracks")
	l.Alert("not implemented")
	os.Exit(1)
}

func runListTracksActive(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "list", "active")
	l.Alert("not implemented")
	os.Exit(1)
}
