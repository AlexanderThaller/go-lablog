package scm

import (
	"bytes"
	"os/exec"
	"time"

	"github.com/AlexanderThaller/lablog/src/data"
	"github.com/juju/errgo"
)

const (
	//Name is the name of the current package. Used in logging.
	Name = "scm"
)

//Commit will add and commit the given entry into the repository that lays unter
//the given datadir.
func Commit(datadir string, entry data.Entry) error {
	filename := entry.GetProject().Name + ".csv"

	err := gitAdd(datadir, filename)
	if err != nil {
		return errgo.Notef(err, "can not add file to repository")
	}

	message := entry.GetProject().Name + " - " + entry.Type() + " - " +
		entry.GetTimeStamp().Format(data.EntryCSVTimeStampFormat)

	err = gitCommit(datadir, message)
	if err != nil {
		return errgo.Notef(err, "can not commit file to repository")
	}

	return nil
}

func gitAdd(datadir, filename string) error {
	command := exec.Command("git", "add", filename)
	command.Dir = datadir

	stderr := new(bytes.Buffer)
	command.Stderr = stderr

	err := command.Run()
	if err != nil {
		return errgo.Notef(errgo.Notef(err, "can not add file with git"),
			stderr.String())
	}

	// Give git time to add everything and remove the lockfile.
	time.Sleep(5 * time.Millisecond)
	return nil
}

func gitCommit(datadir, message string) error {
	command := exec.Command("git", "commit", "-m", message)
	command.Dir = datadir

	stderr := new(bytes.Buffer)
	command.Stderr = stderr

	err := command.Run()
	if err != nil {
		return errgo.Notef(errgo.Notef(err, "can not add file with git"),
			stderr.String())
	}

	// Give git time to commit everything and remove the lockfile.
	time.Sleep(5 * time.Millisecond)
	return nil
}
