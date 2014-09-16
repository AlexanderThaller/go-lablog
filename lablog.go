package main

import (
	"os"

	"github.com/AlexanderThaller/logger"
	"github.com/juju/errgo"
)

const (
	Name = "lablog"
)

func init() {
	logger.SetLevel(".", logger.Trace)
}

func main() {
	l := logger.New(Name, "main")
	args := os.Args

	var command Command
	if len(args) < 1 {
		command = NewCommand(CommandList, []string{})
	} else {
		comm, err := parseCommand(args)
		if err != nil {
			l.Alert("Problem while parsing command: ", errgo.Details(err))
			os.Exit(1)
		}
		command = comm
	}

	switch command.Type {
	case CommandList:
		l.Debug("will run CommandList")
	default:
		l.Alert("do not recognize the command: ", command)
	}
}
