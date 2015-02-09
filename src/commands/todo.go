package commands

import (
	"os"

	"github.com/AlexanderThaller/cobra"
	"github.com/AlexanderThaller/logger"
)

var cmdTodo = &cobra.Command{
	Use:    "todo [project] [text]",
	Short:  "Todo projects.",
	Long:   `Todo projects.`,
	Run:    runTodo,
	PreRun: setLogLevel,
}

func runTodo(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "todo")
	l.Alert("not implemented")
	os.Exit(1)
}
