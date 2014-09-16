package main

type Command struct {
	Type CommandType
	Args []string
}

type CommandType uint

const (
	CommandList CommandType = iota
)

const (
	CommandListString = "list"
)

func NewCommand(typ CommandType, args []string) Command {
	command := new(Command)
	command.Type = typ
	command.Args = args

	return *command
}

func parseCommand(args []string) (Command, error) {
	command := NewCommand(CommandList, args)

	return command, nil
}
