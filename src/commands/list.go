package commands

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
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
var flagListProject string
var flagListNoSubprojects bool

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

	cmdList.PersistentFlags().StringVarP(&flagListProject, "projects", "p",
		"", "The project to list for")
	cmdList.PersistentFlags().BoolVarP(&flagListNoSubprojects, "nosubprojects", "s",
		false, "Don't print subprojects if true")
}

func runListDates(cmd *cobra.Command, args []string) {
}

func runListDurations(cmd *cobra.Command, args []string) {
}

func runListNotes(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "list", "notes")

	output := bytes.NewBufferString("")

	if flagListProject == "" {
		err := runListCommand(output, runListProjectNotes)
		errexit(l, err, "can not list notes for projects")
	} else {
		err := runListCommand(output, runListProjectNotesAndSubnotes, flagListProject)
		errexit(l, err, "can not list notes for project")
	}

	fmt.Print(output.String())
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

type listCommand func(io.Writer, string, int) error

func runListCommand(writer io.Writer, command listCommand, projects ...string) error {
	l := logger.New("commands", "list", "runListCommand")

	l.Trace("Length of projects: ", len(projects))
	if len(projects) == 0 {
		l.Debug("Trying to get projects")

		projcts, err := project.Projects(flagLablogDataDir, flagListStartTime,
			flagListEndTime)
		if err != nil {
			return errgo.Notef(err, "can not get projects")
		}
		projects = projcts
	}

	l.Debug("Will now format the header")
	project.FormatHeader(writer, "Lablog", "list", 1)

	l.Debug("Will now run the command for projects")
	for _, project := range projects {
		l.Trace("Run the command for the project ", project)
		err := command(writer, project, 2)
		if err != nil {
			return errgo.Notef(err, "problem while running list command")
		}
	}
	l.Debug("Finished running the command")

	return nil
}

func runListProjectNotes(writer io.Writer, proj string, indent int) error {
	notes, err := project.ProjectNotes(proj, flagLablogDataDir,
		flagListStartTime, flagListEndTime)
	if err != nil {
		return errgo.Notef(err, "can not get notes")
	}

	sort.Sort(project.NotesByDate(notes))

	err = project.FormatNotes(writer, proj, "list", notes, indent)
	if err != nil {
		return errgo.Notef(err, "can not format notes")
	}

	return nil
}

func runListProjectNotesAndSubnotes(writer io.Writer, proj string, indent int) error {
	err := runListProjectNotes(writer, proj, indent)
	if err != nil {
		return errgo.Notef(err, "can not write project notes")
	}

	if flagListNoSubprojects {
		return nil
	}

	subprojects, err := project.ProjectSubprojects(proj, flagLablogDataDir,
		flagListStartTime, flagListEndTime)
	if err != nil {
		return errgo.Notef(err, "can not get subprojects")
	}

	for _, subproj := range subprojects {
		err := runListProjectNotes(writer, subproj, indent)
		if err != nil {
			return errgo.Notef(err, "can not list notes for subproject")
		}
	}

	return nil
}
