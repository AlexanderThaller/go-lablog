package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"time"
	"unicode"

	"github.com/AlexanderThaller/logger"
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
}

const (
	CommitMessageTimeStampFormat = RecordTimeStampFormat
	DateFormat                   = "2006-01-02"
)

const (
	ActionDone                = "done"
	ActionList                = "list"
	ActionListActiveTracks    = "activetracks"
	ActionListDates           = "dates"
	ActionListNotes           = "notes"
	ActionListProjects        = "projects"
	ActionListTodos           = "todos"
	ActionListTracks          = "tracks"
	ActionListTracksDurations = "durations"
	ActionMerge               = "merge"
	ActionNote                = "note"
	ActionRename              = "rename"
	ActionStopTracking        = "stoptrack"
	ActionTodo                = "todo"
	ActionTrack               = "track"
)

func NewCommand() *Command {
	return new(Command)
}

func (com *Command) Run() error {
	if com.DataPath == "" {
		return errgo.New("the datapath can not be empty")
	}

	switch com.Action {
	case ActionDone:
		return com.runDone()
	case ActionNote:
		return com.runNote()
	case ActionListDates:
		return com.runListDates()
	case ActionList:
		return com.runList()
	case ActionListNotes:
		return com.runNotes()
	case ActionListProjects:
		return com.runListProjects()
	case ActionListTodos:
		return com.runListCommand(com.runListProjectTodosAndSubtodos)
	case ActionListTracks:
		return com.runListCommand(com.runListProjectTracks)
	case ActionTodo:
		return com.runTodo()
	case ActionRename:
		return com.runRename()
	case ActionMerge:
		return com.runMerge()
	case ActionTrack:
		return com.runTrack()
	case ActionStopTracking:
		return com.runListCommand(com.runProjectStopTracking)
	case ActionListActiveTracks:
		return com.runListCommand(com.runListProjectActiveTracks)
	case ActionListTracksDurations:
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

type listCommand func(string, int) error

func (com *Command) runList() error {
	if com.Project == "" {
		return com.runListProjects()
	}

	if !com.checkProjectExists(com.Project) {
		return errgo.New("project " + com.Project + " does not exist")
	}

	notes, err := com.getProjectNotes(com.Project)
	if err != nil {
		return err
	}
	if len(notes) != 0 {
		return com.runListProjectNotesAndSubnotes(com.Project, 1)
	}

	return com.runListProjectTodosAndSubtodos(com.Project, 1)
}

func (com *Command) runNotes() error {
	if com.Project == "" {
		return com.runListCommand(com.runListProjectNotes)
	}

	fmt.Println("=", com.Project)
	fmt.Println(AsciiDocSettings)
	fmt.Println("")

	return com.runListProjectNotesAndSubnotes(com.Project, 2)
}

func (com *Command) runListCommand(command listCommand) error {
	if com.Project != "" {
		return command(com.Project, 1)
	}

	projects, err := com.getProjects()
	if err != nil {
		return err
	}

	FormatHeader(os.Stdout, com.Action)

	for _, project := range projects {
		err := command(project, 2)
		if err != nil {
			return err
		}
	}

	return nil
}

func (com *Command) runListProjects() error {
	projects, err := com.getProjects()
	if err != nil {
		return err
	}

	for _, project := range projects {
		fmt.Println(project)
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

		fmt.Println(date)
	}

	return nil
}

func (com *Command) runListProjectNotes(project string, indent int) error {
	notes, err := com.getProjectNotes(project)
	if err != nil {
		return err
	}
	sort.Sort(NotesByDate(notes))

	buffer := bytes.NewBufferString("")
	err = FormatNotes(buffer, project, notes, indent)
	if err != nil {
		return err
	}
	fmt.Print(buffer.String())

	return nil
}

func (com *Command) runListProjectNotesAndSubnotes(project string, indent int) error {
	err := com.runListProjectNotes(project, indent)
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
		err := com.runListProjectNotes(subproject, indent+1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (com *Command) runListProjectTodos(project string, indent int) error {
	if project == "" {
		return errgo.New("project name can not be empty")
	}
	if !com.checkProjectExists(project) {
		return errgo.New("the project does not exist")
	}

	todos, err := com.getProjectTodos(project)
	if err != nil {
		return err
	}
	todos = FilterInactiveTodos(todos)

	if len(todos) == 0 {
		return nil
	}

	section := strings.Repeat("=", indent)

	var action string
	if indent == 1 {
		capcommand := []rune(com.Action)
		capcommand[0] = unicode.ToUpper(capcommand[0])
		action = "-- " + string(capcommand) + "\n"
	}

	fmt.Println(section, project, action)

	var out []string
	for _, todo := range todos {
		out = append(out, "  * "+todo.GetValue())
	}
	sort.Strings(out)

	for _, todo := range out {
		fmt.Println(todo)
	}
	fmt.Println("")

	return nil
}

func (com *Command) runListProjectTodosAndSubtodos(project string, indent int) error {
	err := com.runListProjectTodos(project, indent)
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
		err := com.runListProjectTodos(subproject, indent+1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (com *Command) runListProjectTracks(project string, indent int) error {
	if project == "" {
		return errgo.New("project name can not be empty")
	}
	if !com.checkProjectExists(project) {
		return errgo.New("the project does not exist")
	}

	tracks, err := com.getProjectTracks(project)
	if err != nil {
		return err
	}

	if len(tracks) == 0 {
		return nil
	}

	section := strings.Repeat("=", indent)
	fmt.Println(section, project, "- Tracks")

	for _, track := range tracks {
		out := "  * " + track.GetTimeStamp()

		if track.Value != "" {
			out += " - " + track.Value
		}

		fmt.Println(out)
	}
	fmt.Println("")

	return nil
}

func (com *Command) runListProjectActiveTracks(project string, indent int) error {
	if project == "" {
		return errgo.New("project name can not be empty")
	}
	if !com.checkProjectExists(project) {
		return errgo.New("the project does not exist")
	}

	tracks, err := com.getProjectActiveTracks(project)
	if err != nil {
		return err
	}

	if len(tracks) == 0 {
		return nil
	}

	section := strings.Repeat("=", indent)
	fmt.Println(section, project, "- ActiveTracks")

	for _, track := range tracks {
		out := "* " + track.GetTimeStamp()

		if track.Value != "" {
			out += " -- " + track.Value
		}

		fmt.Println(out)
	}

	return nil
}

func (com *Command) runProjectStopTracking(project string, indent int) error {
	active, err := com.getProjectActiveTracks(project)
	if err != nil {
		return err
	}

	if len(active) == 0 {
		return nil
	}

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

func (com *Command) runListProjectTracksDurations(project string, indent int) error {
	tracks, err := com.getProjectTracks(project)
	if err != nil {
		return err
	}

	starttracks := make(map[string]Track)
	durations := make(map[string]time.Duration)
	active := make(map[string]bool)
	for _, track := range tracks {
		if !active[track.Value] {
			starttracks[track.Value] = track
			active[track.Value] = true
			continue
		}

		startrack := starttracks[track.Value]
		duration := track.TimeStamp.Sub(startrack.TimeStamp)
		durations[track.Value] += duration
		active[track.Value] = false
	}

	if len(durations) == 0 {
		return nil
	}

	section := strings.Repeat("=", indent)
	fmt.Println(section, project, "- Durations")
	for value, duration := range durations {
		fmt.Println(value, "-", duration)
	}
	fmt.Println()

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

func (com *Command) getProjectSubprojects(project string) ([]string, error) {
	return ProjectSubprojects(project, com.DataPath, com.StartTime, com.EndTime)
}

func (com *Command) getDates() ([]string, error) {
	return Dates(com.Project, com.DataPath, com.StartTime, com.EndTime)
}
