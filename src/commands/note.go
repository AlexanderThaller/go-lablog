package commands

import (
	"os"

	"github.com/AlexanderThaller/logger"
	"github.com/spf13/cobra"
)

var cmdNote = &cobra.Command{
	Use:   "note [project] [text]",
	Short: "Create a new note for the project",
	Long: `Create a note which will record the current timestamp and the given
  text for the given project`,
	Run: runNote,
}

func runNote(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "note")
	l.Alert("not implemented")
	os.Exit(1)
}
