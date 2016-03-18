package data

import (
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/armon/go-radix"
)

type Projects struct {
	tree *radix.Tree
}

func NewProjects() Projects {
	projects := Projects{
		tree: radix.New(),
	}

	return projects
}

func (proj Projects) Add(project Project) {
	proj.tree.Insert(project.Name.String(), project)
}

func (proj Projects) Set(project Project) {
	proj.tree.Insert(project.Name.String(), project)
}

func (proj Projects) List(projects ...Project) []Project {
	var out []Project
	walkFn := func(prefix string, item interface{}) bool {
		log.Debug("Appending ", prefix, " to out")
		out = append(out, item.(Project))
		return false
	}

	if len(projects) == 0 {
		proj.tree.Walk(walkFn)
	} else {
		for _, project := range projects {
			proj.tree.WalkPrefix(project.Name.String(), walkFn)
		}
	}

	return out
}

type Project struct {
	Entries Entries
	Name    ProjectName
}

func (project Project) Notes() []Note {
	var out []Note
	for _, entry := range project.Entries {
		if entry.Type() == EntryTypeNote {
			out = append(out, entry.(Note))
		}
	}

	return out
}

func (project *Project) AddNote(note Note) {
	project.Entries = append(project.Entries, note)
}

func (project Project) Todos() []Todo {
	var out []Todo
	for _, entry := range project.Entries {
		if entry.Type() == EntryTypeTodo {
			out = append(out, entry.(Todo))
		}
	}

	return out
}

func (project *Project) AddTodo(todo Todo) {
	project.Entries = append(project.Entries, todo)
}

type ProjectName []string

const ProjectNameSepperator = "."

func ParseProjectName(name string) (ProjectName, error) {
	splitted := strings.Split(name, ProjectNameSepperator)

	return ProjectName(splitted), nil
}

func (name ProjectName) Values() []string {
	return []string(name)
}

func (name ProjectName) String() string {
	return strings.Join(name, ProjectNameSepperator)
}

func ProjectNamesToString(names []ProjectName) []string {
	var out []string

	for _, name := range names {
		out = append(out, name.String())
	}

	return out
}
