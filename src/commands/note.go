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
	Run: func(cmd *cobra.Command, args []string) {
		l := logger.New("commands", "note")
		l.SetLevel(logger.Trace)

		l.Trace("Args: ", args)

		if len(args) < 1 {
			l.Alert("note command needs a project")
			os.Exit(1)
		}

		if len(args) < 2 {
			l.Alert("note command needs a text")
			os.Exit(1)
		}

		project := args[0]
		text := args[1:]

		l.Debug("Project: ", project)
		l.Debug("Text: ", text[0])
	},
}
