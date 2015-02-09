package commands

import (
	"os"

	"github.com/AlexanderThaller/logger"
	"github.com/spf13/cobra"
)

var cmdList = &cobra.Command{
	Use:   "list (command)",
	Short: "List x",
	Long:  `List x`,
	Run:   runListProjects,
}

var cmdListDates = &cobra.Command{
	Use:   "dates",
	Short: "List dates",
	Long:  `List dates`,
	Run:   runListDates,
}

var cmdListNotes = &cobra.Command{
	Use:   "notes",
	Short: "List notes",
	Long:  `List notes`,
	Run:   runListNotes,
}

var cmdListProjects = &cobra.Command{
	Use:   "projects",
	Short: "List projects",
	Long:  `List projects`,
	Run:   runListProjects,
}

var cmdListTodos = &cobra.Command{
	Use:   "todos",
	Short: "List todos",
	Long:  `List todos`,
	Run:   runListTodos,
}

var cmdListTracks = &cobra.Command{
	Use:   "tracks",
	Short: "List tracks",
	Long:  `List tracks`,
	Run:   runListTracks,
}

var cmdListTracksActive = &cobra.Command{
	Use:   "active",
	Short: "List tracks",
	Long:  `List tracks`,
	Run:   runListTracksActive,
}

var cmdListTracksDurations = &cobra.Command{
	Use:   "durations",
	Short: "List durations",
	Long:  `List durations`,
	Run:   runListDurations,
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
	l.Alert("not implemented")
	os.Exit(1)
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
