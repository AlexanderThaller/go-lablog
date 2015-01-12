package commands

//BuildName represents the name of the software
var BuildName string

//BuildVersion represents the version of the software
var BuildVersion string

//BuildHash represents the hash from git with which the software was build
var BuildHash string

//BuildTime represents the time when the software was build
var BuildTime string

func Execute() {
	AddCommands()
	lablogCmd.Execute()
}

func AddCommands() {
	lablogCmd.AddCommand(cmdVersion)
	lablogCmd.AddCommand(cmdWeb)
	lablogCmd.AddCommand(cmdNote)
}
