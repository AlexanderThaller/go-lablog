package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/AlexanderThaller/logger"
	"github.com/juju/errgo"
)

type Command struct {
	Action        string
	Args          []string
	DataPath      string
	EndTime       string
	Project       string
	SCM           string
	SCMAutoCommit bool
	SCMAutoPush   bool
	StartTime     string
	TimeStamp     time.Time
	Value         string
}

const (
	CommitMessageTimeStampFormat = RecordTimeStampFormat
	DateFormat                   = "2006-01-02"
)

const (
	ActionListDates = "listdates"
	ActionList      = "list"
	ActionListTodos = "listtodos"
	ActionNote      = "note"
	ActionTodo      = "todo"
)

func NewCommand() *Command {
	return new(Command)
}

func (com *Command) Run() error {
	if com.DataPath == "" {
		return errgo.New("the datapath can not be empty")
	}

	switch com.Action {
	case ActionNote:
		return com.runNote()
	case ActionListDates:
		return com.runListDates()
	case ActionList:
		return com.runList()
	case ActionListTodos:
		return com.runListTodos()
	case ActionTodo:
		return com.runTodo()
	default:
		return errgo.New("Do not recognize the action: " + com.Action)
	}
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

func (com *Command) runList() error {
	if com.Project == "" {
		return com.runListProjects()
	} else {
		return com.runListProjectNotes(com.Project)
	}
}

func (com *Command) runListTodos() error {
	if com.Project != "" {
		return com.runListProjectTodos(com.Project)
	}

	projects, err := com.getProjects()
	if err != nil {
		return err
	}

	for _, project := range projects {
		err := com.runListProjectTodos(project)
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
		fmt.Println(date)
	}

	return nil
}

func (com *Command) getDates() ([]string, error) {
	projects, err := com.getProjects()
	if err != nil {
		return nil, err
	}

	datemap := make(map[string]struct{})
	for _, project := range projects {
		dates, err := com.getProjectDates(project)
		if err != nil {
			return nil, err
		}

		for _, date := range dates {
			datemap[date] = struct{}{}
		}
	}

	var out []string
	for date := range datemap {
		out = append(out, date)
	}

	return out, nil
}

func (com *Command) getProjectDates(project string) ([]string, error) {
	if com.DataPath == "" {
		return nil, errgo.New("datapath can not be empty")
	}
	if project == "" {
		return nil, errgo.New("project name can not be empty")
	}
	if !com.checkProjectExists(project) {
		return nil, errgo.New("project does not exist")
	}

	var out []string

	records, err := com.getProjectNotes(project)
	if err != nil {
		return nil, err
	}

	datemap := make(map[string]struct{})

	for _, record := range records {
		date, err := time.Parse(RecordTimeStampFormat, record.GetTimeStamp())
		if err != nil {
			return nil, err
		}

		datemap[date.Format(DateFormat)] = struct{}{}
	}

	for date := range datemap {
		out = append(out, date)
	}

	return out, nil
}

func (com *Command) getProjects() ([]string, error) {
	dir, err := ioutil.ReadDir(com.DataPath)
	if err != nil {
		return nil, err
	}

	var out []string
	for _, d := range dir {
		file := d.Name()

		// Skip dotfiles
		if strings.HasPrefix(file, ".") {
			continue
		}

		ext := filepath.Ext(file)
		name := file[0 : len(file)-len(ext)]

		out = append(out, name)
	}

	sort.Strings(out)
	return out, nil
}

func (com *Command) runListProjectTodos(project string) error {
	if project == "" {
		return errgo.New("project name can not be empty")
	}
	if !com.checkProjectExists(project) {
		return errgo.New("no notes for this project")
	}

	todos, err := com.getProjectTodos(project)
	if err != nil {
		return err
	}

	if len(todos) == 0 {
		return nil
	}

	fmt.Println("#", project, "- Todos")
	for _, todo := range todos {
		if todo.Done {
			continue
		}

		fmt.Println("  *", todo.GetValue())
	}
	fmt.Println("")

	return nil
}

func (com *Command) runListProjectNotes(project string) error {
	if project == "" {
		return errgo.New("project name can not be empty")
	}
	if !com.checkProjectExists(project) {
		return errgo.New("no notes for this project")
	}

	notes, err := com.getProjectNotes(project)
	if err != nil {
		return err
	}

	for _, note := range notes {
		fmt.Println("#", note.GetTimeStamp())
		fmt.Println(note.GetValue())
		fmt.Println("")
	}

	return nil
}

func (com *Command) checkProjectExists(project string) bool {
	projects, err := com.getProjects()
	if err != nil {
		return false
	}

	for _, d := range projects {
		if d == project {
			return true
		}
	}

	return false
}

func (com *Command) getProjectNotes(project string) ([]Note, error) {
	l := logger.New(Name, "Command", "get", "ProjectRecords")
	l.SetLevel(logger.Debug)

	if com.DataPath == "" {
		return nil, errgo.New("datapath can not be empty")
	}
	if project == "" {
		return nil, errgo.New("project name can not be empty")
	}
	if !com.checkProjectExists(project) {
		return nil, errgo.New("project does not exist")
	}

	filepath := filepath.Join(com.DataPath, project+".csv")
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0640)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 3

	var out []Note
	for {
		csv, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			continue
		}

		note, err := NoteFromCSV(csv)
		if err != nil {
			continue
		}
		note.SetProject(project)

		out = append(out, note)
	}

	return out, err
}

func (com *Command) getProjectTodos(project string) ([]Todo, error) {
	l := logger.New(Name, "Command", "get", "ProjectRecords")
	l.SetLevel(logger.Debug)

	if com.DataPath == "" {
		return nil, errgo.New("datapath can not be empty")
	}
	if project == "" {
		return nil, errgo.New("project name can not be empty")
	}
	if !com.checkProjectExists(project) {
		return nil, errgo.New("project does not exist")
	}

	filepath := filepath.Join(com.DataPath, project+".csv")
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0640)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 4

	var out []Todo
	for {
		csv, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			continue
		}

		todo, err := TodoFromCSV(csv)
		if err != nil {
			continue
		}

		out = append(out, todo)
	}

	return out, err
}

func (com *Command) Write(record Record) error {
	if com.DataPath == "" {
		return errgo.New("datapath can not be empty")
	}

	if com.Project == "" {
		return errgo.New("project name can not be empty")
	}

	path := com.DataPath
	project := com.Project

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

	message := com.Project + " - "
	message += record.GetAction() + " - "
	message += com.TimeStamp.Format(CommitMessageTimeStampFormat)
	err = scmCommit(com.SCM, com.DataPath, message)
	if err != nil {
		return err
	}

	return nil
}
