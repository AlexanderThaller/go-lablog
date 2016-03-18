package vcs

import (
	"bytes"
	"os/exec"
	"time"

	"github.com/AlexanderThaller/lablog/src/data"
	"github.com/juju/errgo"
)

//Commit will add and commit the given entry into the repository that lays unter
//the given datadir.
func Commit(datadir string, project data.ProjectName, entry data.Entry) error {
	err := gitAdd(datadir, ".")
	if err != nil {
		return errgo.Notef(err, "can not add file to repository")
	}

	message := project.String() + " - " + entry.Type().String() + " - " +
		entry.GetTimeStamp().Format(data.TimeStampFormat)

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
