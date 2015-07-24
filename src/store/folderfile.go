package store

import (
	"github.com/AlexanderThaller/lablog/src/data"
	"github.com/juju/errgo"
)

type FolderFile struct {
	DataDir    string
	AutoCommit bool
	AutoPush   bool
}

func (store *FolderFile) Write(entry data.Entry) error {
	return errgo.New("not implemented")
}
