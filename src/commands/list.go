package commands

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/AlexanderThaller/lablog/src/project"
	"github.com/AlexanderThaller/logger"
	"github.com/juju/errgo"
	"github.com/spf13/cobra"
)

var cmdList = &cobra.Command{
	Use:   "list (dates|notes|projects|todos|tracks)",
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
	Run:   runListTracksDurations,
}

var flagListStartTime time.Time
var flagListEndTime time.Time

func init() {
	cmdList.AddCommand(cmdListDates)
	cmdList.AddCommand(cmdListNotes)
	cmdList.AddCommand(cmdListProjects)
	cmdList.AddCommand(cmdListTodos)

	cmdListTracks.AddCommand(cmdListTracksActive)
	cmdListTracks.AddCommand(cmdListTracksDurations)
	cmdList.AddCommand(cmdListTracks)

	flagListStartTime = time.Time{}
	flagListEndTime = time.Now()
}

func runListDates(cmd *cobra.Command, args []string) {
}

func runListDurations(cmd *cobra.Command, args []string) {
}

func runListNotes(cmd *cobra.Command, args []string) {
}

func runListProjects(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "list", "projects")

	output := bytes.NewBufferString("")

	projects, err := project.Projects(flagLablogDataDir,
		flagListStartTime, flagListEndTime)
	if err != nil {
		l.Alert("can not get projects: ", errgo.Details(err))
		os.Exit(1)
	}

	for _, project := range projects {
		output.WriteString(project + "\n")
	}

	fmt.Print(output.String())
}

func runListTodos(cmd *cobra.Command, args []string) {
}

func runListTracks(cmd *cobra.Command, args []string) {
}

func runListTracksActive(cmd *cobra.Command, args []string) {
}

func runListTracksDurations(cmd *cobra.Command, args []string) {
}
