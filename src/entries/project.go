package entries

import "github.com/juju/errgo"

func Projects(datadir string) ([]Project, error) {
	return nil, errgo.New("not implemented")
}

type Project struct{}

func (project Project) Name() string {
	return ""
}

type ProjectsByName []Project

func (by ProjectsByName) Len() int {
	return len(by)
}

func (by ProjectsByName) Swap(i, j int) {
	by[i], by[j] = by[j], by[i]
}

func (by ProjectsByName) Less(i, j int) bool {
	return by[i].Name() < by[j].Name()
}
