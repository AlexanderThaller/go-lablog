package main

import (
	"runtime"

	"github.com/AlexanderThaller/lablog/src/commands"
)

const (
	BuildName    = "lablog"
	BuildVersion = "0.0.1"
)

var (
	BuildHash string
	BuildTime string
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// See commands/main.go
	commands.BuildName = BuildName
	commands.BuildVersion = BuildVersion
	commands.BuildHash = BuildHash
	commands.BuildTime = BuildTime
	commands.Execute()
}
