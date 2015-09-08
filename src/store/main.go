package store

import "github.com/AlexanderThaller/lablog/src/data"

type Store interface {
	AddEntry(data.ProjectName, data.Entry) error
	GetProject(data.ProjectName) (data.Project, error)
	GetProjects() (data.Projects, error)
	ListProjects() ([]data.ProjectName, error)
}
