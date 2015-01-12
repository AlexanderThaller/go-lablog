package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"sort"
	"time"

	"github.com/AlexanderThaller/logger"
	"github.com/gorilla/mux"
	"github.com/jinzhu/now"
	"github.com/juju/errgo"
)

type Command struct {
	Action        string
	Args          []string
	DataPath      string
	EndTime       time.Time
	Project       string
	SCM           string
	SCMAutoCommit bool
	SCMAutoPush   bool
	StartTime     time.Time
	TimeStamp     time.Time
	Value         string
	NoSubprojects bool
	writer        io.Writer
}

const (
	CommitMessageTimeStampFormat = RecordTimeStampFormat
	DateFormat                   = "2006-01-02"
)

const (
	ActionDates           = "dates"
	ActionDone            = "done"
	ActionList            = "list"
	ActionMerge           = "merge"
	ActionNote            = "note"
	ActionNotes           = "notes"
	ActionProjects        = "projects"
	ActionRename          = "rename"
	ActionTodos           = "todos"
	ActionTodo            = "todo"
	ActionTracksActive    = "tracksactive"
	ActionTracksDurations = "durations"
	ActionTrackStop       = "trackstop"
	ActionTracks          = "tracks"
	ActionTrack           = "track"
)

func NewCommand(writer io.Writer) *Command {
	command := new(Command)
	command.writer = writer

	return command
}

func (com *Command) Run() error {
	l := logger.New(Name, "Command", "Run")
	if com.DataPath == "" {
		return errgo.New("the datapath can not be empty")
	}

	l.Debug("Will now run the action ", com.Action)
	switch com.Action {
	case ActionDates:
		return com.runListDates()
	case ActionDone:
		return com.runDone()
	case ActionList:
		return com.List()
	case ActionMerge:
		return com.runMerge()
	case ActionNote:
		return com.runNote()
	case ActionNotes:
		return com.Notes()
	case ActionProjects:
		return com.Projects()
	case ActionRename:
		return com.runRename()
	case ActionTodo:
		return com.runTodo()
	case ActionTodos:
		return com.Todos()
	case ActionTrack:
		return com.runTrack()
	case ActionTrackStop:
		return com.runListCommand(com.runProjectTrackStop)
	case ActionTracks:
		return com.runListCommand(com.runListProjectTracks)
	case ActionTracksActive:
		return com.runListCommand(com.runListProjectTracksActive)
	case ActionTracksDurations:
		return com.runListCommand(com.runListProjectTracksDurations)
	default:
		return errgo.New("Do not recognize the action: " + com.Action)
	}
}

func (com *Command) Commit(record Record) error {
	if !com.SCMAutoCommit {
		return nil
	}

	if com.SCM == "" {
		return errgo.New("Can not use an empty scm for commiting")
	}

	filename := record.GetProject() + ".csv"
	err := scmAdd(com.SCM, com.DataPath, filename)
	if err != nil {
		return err
	}

	message := record.GetProject() + " - "
	message += record.GetAction() + " - "
	message += com.TimeStamp.Format(CommitMessageTimeStampFormat)
	err = scmCommit(com.SCM, com.DataPath, message)
	if err != nil {
		return err
	}

	return nil
}

func (com *Command) Write(record Record) error {
	l := logger.New(Name, "Command", "Write")
	if com.DataPath == "" {
		return errgo.New("datapath can not be empty")
	}

	if record.GetProject() == "" {
		l.Debug("Record: ", record)
		return errgo.New("project name can not be empty")
	}

	path := com.DataPath
	project := record.GetProject()

	err := os.MkdirAll(path, 0750)
	if err != nil {
		return err
	}

	filepath := filepath.Join(path, project+".csv")
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0640)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.Write(record.CSV())
	if err != nil {
		return err
	}
	writer.Flush()

	err = com.Commit(record)
	if err != nil {
		return err
	}

	return nil
}

type listCommand func(io.Writer, string, int) error

func (com *Command) runListCommand(command listCommand, projects ...string) error {
	l := logger.New(Name, "Command", "runListCommand")

	l.Trace("Length of projects: ", len(projects))
	if len(projects) == 0 {
		l.Debug("Trying to get projects")

		projcts, err := com.getProjects()
		if err != nil {
			return err
		}
		projects = projcts
	}

	l.Debug("Will now format the header")
	FormatHeader(com.writer, "Lablog", com.Action, 1)

	l.Debug("Will now run the command for projects")
	for _, project := range projects {
		l.Trace("Run the command for the project ", project)
		err := command(com.writer, project, 2)
		if err != nil {
			return err
		}
	}
	l.Debug("Finished running the command")

	return nil
}

func (com *Command) List() error {
	if com.Project == "" {
		return com.Projects()
	}

	if ProjectHasNotes(com.Project, com.DataPath, com.StartTime, com.EndTime) {
		return com.runListCommand(com.runListProjectNotesAndSubnotes, com.Project)
	}

	if ProjectHasTodos(com.Project, com.DataPath, com.StartTime, com.EndTime) {
		return com.runListCommand(com.runListProjectTodosAndSubtodos, com.Project)
	}

	return com.runListCommand(com.runListProjectTracksActiveAndSubtracks, com.Project)
}

func (com *Command) Notes() error {
	l := logger.New(Name, "Command", "Notes")
	if com.Project == "" {
		l.Debug("Will list notes for all projects")
		return com.runListCommand(com.runListProjectNotes)
	}

	l.Debug("Will list notes for project ", com.Project)
	return com.runListCommand(com.runListProjectNotesAndSubnotes, com.Project)
}

func (com *Command) Todos() error {
	l := logger.New(Name, "Command", "Todos")
	if com.Project == "" {
		l.Debug("Will list todos for all projects")
		return com.runListCommand(com.runListProjectTodos)
	}

	l.Debug("Will list todos for project ", com.Project)
	return com.runListCommand(com.runListProjectTodosAndSubtodos, com.Project)
}

func (com *Command) Tracks() error {
	l := logger.New(Name, "Command", "Tracks")
	if com.Project == "" {
		l.Debug("Will list tracks for all projects")
		return com.runListCommand(com.runListProjectTracks)
	}

	l.Debug("Will list tracks for project ", com.Project)
	return com.runListCommand(com.runListProjectTracksAndSubtracks, com.Project)
}

func (com *Command) Projects() error {
	projects, err := com.getProjects()
	if err != nil {
		return err
	}

	for _, project := range projects {
		com.writer.Write([]byte(project + "\n"))
	}

	return nil
}

func (com *Command) runListDates() error {
	l := logger.New(Name, "Command", "run", "ListDates")

	var dates []string
	var err error

	if com.Project == "" {
		dates, err = com.getDates()
	} else {
		dates, err = com.getProjectDates(com.Project)
	}

	if err != nil {
		return err
	}

	sort.Strings(dates)
	for _, date := range dates {
		timestamp, err := now.Parse(date)
		if err != nil {
			l.Warning("Can not parse timestamp: ", errgo.Details(err))
			continue
		}

		if timestamp.Before(com.StartTime) {
			continue
		}

		if timestamp.After(com.EndTime) {
			continue
		}

		com.writer.Write([]byte(date + "\n"))
	}

	return nil
}

func (com *Command) runListProjectNotes(writer io.Writer, project string, indent int) error {
	notes, err := com.getProjectNotes(project)
	if err != nil {
		return err
	}
	sort.Sort(NotesByDate(notes))

	err = FormatNotes(writer, project, com.Action, notes, indent)
	if err != nil {
		return err
	}

	return nil
}

func (com *Command) runListProjectNotesAndSubnotes(writer io.Writer, project string, indent int) error {
	err := com.runListProjectNotes(writer, project, indent)
	if err != nil {
		return err
	}

	if com.NoSubprojects {
		return nil
	}

	subprojects, err := com.getProjectSubprojects(project)
	if err != nil {
		return err
	}

	for _, subproject := range subprojects {
		err := com.runListProjectNotes(writer, subproject, indent)
		if err != nil {
			return err
		}
	}

	return nil
}

func (com *Command) runListProjectTodos(writer io.Writer, project string, indent int) error {
	todos, err := com.getProjectTodos(project)
	if err != nil {
		return err
	}
	todos = FilterInactiveTodos(todos)

	sort.Sort(TodoByValue(todos))
	err = FormatTodos(writer, project, com.Action, todos, indent)
	if err != nil {
		return err
	}

	return nil
}

func (com *Command) runListProjectTodosAndSubtodos(writer io.Writer, project string, indent int) error {
	err := com.runListProjectTodos(writer, project, indent)
	if err != nil {
		return err
	}

	if com.NoSubprojects {
		return nil
	}

	subprojects, err := com.getProjectSubprojects(project)
	if err != nil {
		return err
	}

	for _, subproject := range subprojects {
		err := com.runListProjectTodos(writer, subproject, indent)
		if err != nil {
			return err
		}
	}

	return nil
}

func (com *Command) runListProjectTracksAndSubtracks(writer io.Writer, project string, indent int) error {
	err := com.runListProjectTracks(writer, project, indent)
	if err != nil {
		return err
	}

	if com.NoSubprojects {
		return nil
	}

	subprojects, err := com.getProjectSubprojects(project)
	if err != nil {
		return err
	}

	for _, subproject := range subprojects {
		err := com.runListProjectTracks(writer, subproject, indent)
		if err != nil {
			return err
		}
	}

	return nil
}

func (com *Command) runListProjectTracks(writer io.Writer, project string, indent int) error {
	tracks, err := com.getProjectTracks(project)
	if err != nil {
		return err
	}

	sort.Sort(TracksByDate(tracks))
	err = FormatTracks(writer, project, com.Action, tracks, indent)
	if err != nil {
		return err
	}

	return nil
}

func (com *Command) runListProjectTracksActive(writer io.Writer, project string, indent int) error {
	active, err := com.getProjectActiveTracks(project)
	if err != nil {
		return err
	}

	sort.Sort(TracksByDate(active))
	err = FormatTracks(writer, project, com.Action, active, indent)

	return nil
}

func (com *Command) runListProjectTracksActiveAndSubtracks(writer io.Writer, project string, indent int) error {
	err := com.runListProjectTracksActive(writer, project, indent)
	if err != nil {
		return err
	}

	if com.NoSubprojects {
		return nil
	}

	subprojects, err := com.getProjectSubprojects(project)
	if err != nil {
		return err
	}

	for _, subproject := range subprojects {
		err := com.runListProjectTracksActive(writer, subproject, indent)
		if err != nil {
			return err
		}
	}

	return nil
}

func (com *Command) runProjectTrackStop(writer io.Writer, project string, indent int) error {
	active, err := com.getProjectActiveTracks(project)
	if err != nil {
		return err
	}

	if len(active) == 0 {
		return errgo.New("no active tracks we could stop")
	}

	FormatTracks(writer, project, com.Action, active, indent)
	for _, track := range active {
		track := Track{
			Project:   project,
			TimeStamp: com.TimeStamp,
			Value:     track.Value,
		}

		err := com.Write(track)
		if err != nil {
			return err
		}
	}

	return nil
}

func (com *Command) runListProjectTracksDurations(writer io.Writer, project string, indent int) error {
	durations, err := com.getProjectDurations(project)
	if err != nil {
		return err
	}

	sort.Sort(DurationsByValue(durations))
	FormatDurations(writer, project, durations, indent)

	return nil
}

func (com *Command) runTrack() error {
	if com.Project == "" {
		return errgo.New("track command needs an project")
	}

	track := new(Track)
	track.Project = com.Project
	track.TimeStamp = com.TimeStamp
	track.Value = com.Value

	return com.Write(track)
}

func (com *Command) runMerge() error {
	if com.Project == "" {
		return errgo.New("Project name can not be empty")
	}
	srcproject := com.Project
	dstproject := com.Value

	if !com.checkProjectExists(srcproject) {
		return errgo.New("no project with the name " + srcproject)
	}

	if !com.checkProjectExists(dstproject) {
		return errgo.New("the project " + dstproject + " already exists")
	}

	srcpath := path.Join(com.DataPath, srcproject+".csv")
	dstpath := path.Join(com.DataPath, dstproject+".csv")

	err := MergeFiles(srcpath, dstpath)
	if err != nil {
		return err
	}

	srcfile := srcproject + ".csv"
	err = scmRemove(com.SCM, srcfile, com.DataPath)
	if err != nil {
		return err
	}

	dstfile := dstproject + ".csv"
	err = scmAdd(com.SCM, com.DataPath, dstfile)
	if err != nil {
		return err
	}

	message := srcproject + " - merged - " + dstproject
	err = scmCommit(com.SCM, com.DataPath, message)
	if err != nil {
		return err
	}

	return nil
}

func (com *Command) runRename() error {
	if com.Project == "" {
		return errgo.New("Project name can not be empty")
	}
	oldproject := com.Project
	newproject := com.Value

	if !com.checkProjectExists(oldproject) {
		return errgo.New("no project with the name " + oldproject)
	}

	if com.checkProjectExists(newproject) {
		return errgo.New("the project " + newproject + " already exists")
	}

	oldpath := oldproject + ".csv"
	newpath := newproject + ".csv"

	err := scmRename(com.SCM, oldpath, newpath, com.DataPath)
	if err != nil {
		return err
	}

	message := oldproject + " - renamed - " + newproject
	err = scmCommit(com.SCM, com.DataPath, message)
	if err != nil {
		return err
	}

	return nil
}

func (com *Command) runDone() error {
	l := logger.New(Name, "Command", "run", "Done")

	l.Trace("Args length: ", len(com.Args))
	if com.Value == "" {
		return errgo.New("todo command needs a value")
	}
	l.Trace("Project: ", com.Project)
	if com.Project == "" {
		return errgo.New("todo command needs an project")
	}

	done := new(Todo)
	done.Project = com.Project
	done.TimeStamp = com.TimeStamp
	done.Value = com.Value
	done.Done = true
	l.Trace("Done: ", fmt.Sprintf("%+v", done))

	return com.Write(done)
}

func (com *Command) runNote() error {
	l := logger.New(Name, "Command", "run", "Note")

	l.Trace("Args length: ", len(com.Args))
	if com.Value == "" {
		return errgo.New("note command needs a value")
	}
	l.Trace("Project: ", com.Project)
	if com.Project == "" {
		return errgo.New("note command needs an project")
	}

	note := new(Note)
	note.Project = com.Project
	note.TimeStamp = com.TimeStamp
	note.Value = com.Value
	l.Trace("Note: ", fmt.Sprintf("%+v", note))

	return com.Write(note)
}

func (com *Command) runTodo() error {
	l := logger.New(Name, "Command", "run", "Todo")

	l.Trace("Args length: ", len(com.Args))
	if com.Value == "" {
		return errgo.New("todo command needs a value")
	}
	l.Trace("Project: ", com.Project)
	if com.Project == "" {
		return errgo.New("todo command needs an project")
	}

	todo := new(Todo)
	todo.Project = com.Project
	todo.TimeStamp = com.TimeStamp
	todo.Value = com.Value
	todo.Done = false
	l.Trace("Todo: ", fmt.Sprintf("%+v", todo))

	return com.Write(todo)
}

func (com *Command) RunWeb(binding string) error {
	r := mux.NewRouter()
	r.HandleFunc("/", com.webRootHandler)
	r.HandleFunc("/notes/{project}", com.webNotesHandler)

	http.Handle("/", r)
	err := http.ListenAndServe(binding, nil)
	if err != nil {
		return err
	}

	return nil
}

func (com *Command) getProjects() ([]string, error) {
	return Projects(com.DataPath, com.StartTime, com.EndTime)
}

func (com *Command) checkProjectExists(project string) bool {
	return ProjectExists(project, com.DataPath)
}

func (com *Command) getProjectNotes(project string) ([]Note, error) {
	return ProjectNotes(project, com.DataPath, com.StartTime, com.EndTime)
}

func (com *Command) getProjectTodos(project string) ([]Todo, error) {
	return ProjectTodos(project, com.DataPath, com.StartTime, com.EndTime)
}

func (com *Command) getProjectDates(project string) ([]string, error) {
	return ProjectDates(project, com.DataPath, com.StartTime, com.EndTime)
}

func (com *Command) getProjectActiveTracks(project string) ([]Track, error) {
	return ProjectActiveTracks(project, com.DataPath)
}

func (com *Command) getProjectTracks(project string) ([]Track, error) {
	return ProjectTracks(project, com.DataPath, com.StartTime, com.EndTime)
}

func (com *Command) getProjectDurations(project string) ([]Duration, error) {
	return ProjectDurations(project, com.DataPath, com.StartTime, com.EndTime)
}

func (com *Command) getProjectSubprojects(project string) ([]string, error) {
	return ProjectSubprojects(project, com.DataPath, com.StartTime, com.EndTime)
}

func (com *Command) getDates() ([]string, error) {
	return Dates(com.Project, com.DataPath, com.StartTime, com.EndTime)
}
