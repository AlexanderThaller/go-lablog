package commands

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/AlexanderThaller/cobra"
	"github.com/AlexanderThaller/lablog/src/data"
	"github.com/AlexanderThaller/lablog/src/format"
	"github.com/AlexanderThaller/lablog/src/helper"
	"github.com/AlexanderThaller/logger"
)

var cmdList = &cobra.Command{
	Use:    "list (command)",
	Short:  "List projects, notes, todos, dates, tracks, etc., see help for all options.",
	Long:   "List projects, notes, todos, dates, tracks, etc., see help for all options.",
	Run:    runListProjects,
	PreRun: runListParseTimeStamps,
}

var cmdListDates = &cobra.Command{
	Use:    "dates",
	Short:  "List dates.",
	Long:   `List dates.`,
	Run:    runListDates,
	PreRun: runListParseTimeStamps,
}

var cmdListNotes = &cobra.Command{
	Use:    "notes",
	Short:  "List notes.",
	Long:   `List notes.`,
	Run:    runListNotes,
	PreRun: runListParseTimeStamps,
}

var cmdListProjects = &cobra.Command{
	Use:    "projects",
	Short:  "List projects.",
	Long:   `List projects.`,
	Run:    runListProjects,
	PreRun: runListParseTimeStamps,
}

var cmdListSubProjects = &cobra.Command{
	Use:    "subprojects",
	Short:  "List subprojects of projects",
	Long:   `List subprojects of projects`,
	Run:    runListSubProjects,
	PreRun: runListParseTimeStamps,
}

var cmdListTodos = &cobra.Command{
	Use:    "todos",
	Short:  "List todos.",
	Long:   `List todos.`,
	Run:    runListTodos,
	PreRun: runListParseTimeStamps,
}

var cmdListTracks = &cobra.Command{
	Use:    "tracks",
	Short:  "List tracks.",
	Long:   `List tracks.`,
	Run:    runListTracks,
	PreRun: runListParseTimeStamps,
}

var cmdListTracksActive = &cobra.Command{
	Use:    "active",
	Short:  "List tracks.",
	Long:   `List tracks.`,
	Run:    runListTracksActive,
	PreRun: runListParseTimeStamps,
}

var cmdListTracksDurations = &cobra.Command{
	Use:    "durations",
	Short:  "List durations.",
	Long:   `List durations.`,
	Run:    runListDurations,
	PreRun: runListParseTimeStamps,
}

var cmdListEntries = &cobra.Command{
	Use:    "entries",
	Short:  "List all entries (notes, todos, tracks).",
	Long:   `List all entries (notes, todos, trackss).`,
	Run:    runListEntries,
	PreRun: runListParseTimeStamps,
}

var cmdListTimeline = &cobra.Command{
	Use:    "timeline",
	Short:  "List all entries (notes, todos, tracks) ordered by date.",
	Long:   `List all entries (notes, todos, trackss) ordered by date.`,
	Run:    runListTimeline,
	PreRun: runListParseTimeStamps,
}

var flagListStart time.Time
var flagListStartRaw string
var flagListEnd time.Time
var flagListEndRaw string

func init() {
	flagListStart = time.Time{}
	flagListEnd = time.Now()

	cmdList.PersistentFlags().StringVarP(&flagListStartRaw, "start", "s",
		flagListStart.String(),
		"The timestamp after which entries have to be taken to be displayes.")

	cmdList.PersistentFlags().StringVarP(&flagListEndRaw, "end", "e",
		flagListEnd.String(),
		"The timestamp before which entries have to be taken to be displayes.")
}

func runListParseTimeStamps(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "list", "ParseTimeStamps")

	start, err := helper.DefaultOrRawTimestamp(flagListStart, flagListStartRaw)
	errexit(l, err, "can not get start timestamp")
	flagListStart = start

	end, err := helper.DefaultOrRawTimestamp(flagListEnd, flagListEndRaw)
	errexit(l, err, "can not get end timestamp")
	flagListEnd = end
}

func runListDates(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "list", "dates")

	projects, err := data.ProjectsOrArgs(args, flagLablogDataDir)
	errexit(l, err, "can not get projects")

	sort.Sort(data.ProjectsByName(projects))

	buffer := new(bytes.Buffer)
	err = format.ProjectsDates(buffer, projects, flagListStart, flagListEnd)
	errexit(l, err, "can not format projects")

	fmt.Print(buffer.String())
}

func runListDurations(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "list", "durations")
	l.Alert("not implemented")
	os.Exit(1)
}

func runListNotes(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "list", "notes")

	projects, err := data.ProjectsOrArgs(args, flagLablogDataDir)
	errexit(l, err, "can not get projects")

	sort.Sort(data.ProjectsByName(projects))

	buffer := new(bytes.Buffer)
	err = format.ProjectsNotes(buffer, projects, flagListStart, flagListEnd)
	errexit(l, err, "can not format projects")

	fmt.Print(buffer.String())
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

func runListSubProjects(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "list", "subprojects")

	projects, err := data.ProjectsOrArgs(args, flagLablogDataDir)
	errexit(l, err, "can not get projects")

	sort.Sort(data.ProjectsByName(projects))

	for _, project := range projects {
		fmt.Println(project.Name)
	}
}

func runListTodos(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "list", "todos")

	projects, err := data.ProjectsOrArgs(args, flagLablogDataDir)
	errexit(l, err, "can not get projects")

	sort.Sort(data.ProjectsByName(projects))

	buffer := new(bytes.Buffer)
	err = format.ProjectsTodos(buffer, projects, flagListStart, flagListEnd)
	errexit(l, err, "can not format projects")

	fmt.Print(buffer.String())
}

func runListTracks(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "list", "tracks")

	projects, err := data.ProjectsOrArgs(args, flagLablogDataDir)
	errexit(l, err, "can not get projects")

	sort.Sort(data.ProjectsByName(projects))

	buffer := new(bytes.Buffer)
	err = format.ProjectsTracks(buffer, projects, flagListStart, flagListEnd)
	errexit(l, err, "can not format projects")

	fmt.Print(buffer.String())
}

func runListTracksActive(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "list", "active")
	l.Alert("not implemented")
	os.Exit(1)
}

func runListEntries(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "list", "entries")

	projects, err := data.ProjectsOrArgs(args, flagLablogDataDir)
	errexit(l, err, "can not get projects")

	sort.Sort(data.ProjectsByName(projects))

	buffer := new(bytes.Buffer)
	err = format.ProjectsEntries(buffer, projects, flagListStart, flagListEnd)
	errexit(l, err, "can not format projects")

	fmt.Print(buffer.String())
}

func runListTimeline(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "list", "log")

	projects, err := data.ProjectsOrArgs(args, flagLablogDataDir)
	errexit(l, err, "can not get projects")

	buffer := new(bytes.Buffer)
	err = format.Timeline(buffer, projects, flagListStart, flagListEnd)
	errexit(l, err, "can not format log")

	fmt.Print(buffer.String())
}
