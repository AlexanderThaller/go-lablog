package data

import "strings"

type Projects []Project

type Project struct {
	Entries
	Name ProjectName
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
