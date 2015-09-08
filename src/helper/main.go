package helper

import (
	"time"

	"github.com/AlexanderThaller/lablog/src/data"
	"github.com/AlexanderThaller/lablog/src/store"

	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/now"
	"github.com/juju/errgo"
)

func ErrExit(err error) {
	cerr := err.(errgo.Wrapper)

	if cerr.Underlying() != nil {
		log.Debug(errgo.Details(err))
		log.Fatal(err)
	}
}

func Debug(args ...interface{}) {
	log.Debug(args)
}

func Fatal(args ...interface{}) {
	log.Fatal(args)
}

func DefaultOrRawTimestamp(timestamp time.Time, raw string) (time.Time, error) {
	if timestamp.String() == raw {
		return timestamp, nil
	}

	parsed, err := now.Parse(raw)
	if err != nil {
		return time.Time{}, errgo.Notef(err, "can not parse timestamp")
	}

	return parsed, nil
}

func RecordEntry(datadir string, project data.ProjectName, entry data.Entry) error {
	store, err := store.NewFolderStore(datadir)
	if err != nil {
		return errgo.Notef(err, "can not get data store")
	}

	err = store.AddEntry(project, entry)
	if err != nil {
		return errgo.Notef(err, "can not write note to data store")
	}

	return nil
}
