package main

import (
	"os/exec"
	"time"

	"github.com/juju/errgo"
)

func scmCommit(scm, datapath, message string) error {
	switch scm {
	case "git":
		return gitCommit(datapath, message)
	default:
		return errgo.New("do not know the scm " + scm)
	}
}

func gitCommit(datapath, message string) error {
	command := exec.Command("git", "commit", "-m", message)
	command.Dir = datapath

	output, err := command.CombinedOutput()
	if err != nil {
		err = errgo.New("problem when commiting to git: " + err.Error() + " - " +
			string(output))

		return err
	}

	// Give git time to commit everything and remove the lockfile.
	time.Sleep(100 * time.Millisecond)
	return nil
}

func scmAdd(scm, datapath string) error {
	switch scm {
	case "git":
		return gitAdd(datapath)
	default:
		return errgo.New("do not know the scm " + scm)
	}
}

func gitAdd(datapath string) error {
	command := exec.Command("git", "add", "--all", ".")
	command.Dir = datapath

	output, err := command.CombinedOutput()
	if err != nil {
		err = errgo.New("problem when adding to git: " + err.Error() + " - " +
			string(output))

		return err
	}

	// Give git time to commit everything and remove the lockfile.
	time.Sleep(100 * time.Millisecond)
	return nil
}
