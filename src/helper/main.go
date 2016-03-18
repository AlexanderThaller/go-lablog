package helper

import (
	"strings"
	"time"

	"github.com/AlexanderThaller/lablog/src/data"
	"github.com/AlexanderThaller/lablog/src/store"

	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/now"
	"github.com/juju/errgo"
)

// DefaultStore will return the default store used in the software. This is only
// to make it easy to change the store type.
func DefaultStore(datadir string) (store.Store, error) {
	return store.NewFolderStore(datadir)
}

// ErrExit will check if the underlying error is nil and if its not it will
// print a debug and fatal message and exit the program.
func ErrExit(err error) {
	cerr := err.(errgo.Wrapper)

	if cerr.Underlying() != nil {
		log.Debug(errgo.Details(err))
		log.Fatal(err)
	}
}

// Debug is a shortcut to log debug messages.
func Debug(args ...interface{}) {
	log.Debug(args)
}

// Fatal is a shortcut to log fatal messages. This will exit the program after
// printing the message.
func Fatal(args ...interface{}) {
	log.Fatal(args)
}

// DefaultOrRawTimestamp will compare the timestamp and the raw timestamp to
// each other and return the given timestamp if the raw and the timestamp are
// the same. Else it will parse the raw timestamp and return that. This is
// mostly used for determining if a timestamp flag was given or not (see
// commands package).
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

// RecordEntry will record the given entry for the given project using the
// specified datadir. This is mostly a helper function which will inizialize the
// store and then record the entry to the store.
func RecordEntry(datadir string, project data.ProjectName, entry data.Entry) error {
	store, err := DefaultStore(datadir)
	if err != nil {
		return errgo.Notef(err, "can not get data store")
	}

	err = store.AddEntry(project, entry)
	if err != nil {
		return errgo.Notef(err, "can not write note to data store")
	}

	return nil
}

// ProjectNamesFromArgs will return the list of all project names if the args
// are empty or return the parses args as a project name list.
func ProjectNamesFromArgs(store store.Store, args []string, showarchive bool) ([]data.ProjectName, error) {
	// If there where no projects specified in args we will just use the whole
	// list of projects.
	if len(args) == 0 {
		names, err := store.ListProjects(showarchive)
		if err != nil {
			return nil, errgo.Notef(err, "can not get list of projects")
		}

		return names, nil
	}

	var names []data.ProjectName
	for _, arg := range args {
		project, err := data.ParseProjectName(arg)
		if err != nil {
			return nil, errgo.Notef(err, "can not parse project name")
		}

		names = append(names, project)
	}

	return names, nil
}

//ArgsToEntryValues will take the given args and try to parse the parameters and
//flags to the values a normaly entry (note, todo, etc.) would need.
func ArgsToEntryValues(args []string, addTimeStamp time.Time, rawTimeStamp string) (
	data.ProjectName,
	time.Time,
	string,
	error) {

	if len(args) < 2 {
		return data.ProjectName{}, time.Time{}, "", errgo.New("need at least two arguments to run")
	}

	project, err := data.ParseProjectName(args[0])
	if err != nil {
		return data.ProjectName{}, time.Time{}, "", errgo.Notef(err, "can not parse project name")
	}

	value := strings.Join(args[1:], " ")

	timestamp, err := DefaultOrRawTimestamp(addTimeStamp, rawTimeStamp)
	if err != nil {
		return data.ProjectName{}, time.Time{}, "", errgo.Notef(err, "can not get timestamp")
	}

	return project, timestamp, value, nil
}

//ArgsToTodo will take the given args and parameters and try to convert them to
//a todo.
func ArgsToTodo(args []string, addTimeStamp time.Time, rawTimeStamp string) (data.ProjectName, data.Todo, error) {
	project, timestamp, value, err := ArgsToEntryValues(args, addTimeStamp, rawTimeStamp)
	if err != nil {
		return data.ProjectName{}, data.Todo{}, errgo.Notef(err, "can not convert args to entry usable values")
	}

	todo := data.Todo{
		TimeStamp: timestamp,
		Value:     value,
	}

	return project, todo, nil
}
