package commands

import (
	"os"
	"strings"
	"time"

	"github.com/AlexanderThaller/lablog/src/project"
	"github.com/AlexanderThaller/logger"
	"github.com/jinzhu/now"
	"github.com/juju/errgo"
	"github.com/spf13/cobra"
)

var cmdNote = &cobra.Command{
	Use:   "note [project] [text]",
	Short: "Create a new note for the project",
	Long: `Create a note which will record the current timestamp and the given
  text for the given project`,
	Run: runNote,
}

var flagNoteSCM string
var flagNoteSCMAutoCommit bool
var flagNoteSCMAutoPush bool
var flagNoteTimeStamp string
var flagNoteTimeStampParsed time.Time

func init() {
	cmdNote.Flags().StringVarP(&flagNoteSCM, "scm", "s",
		"git", "Which scm to use for the repository")
	cmdNote.Flags().BoolVarP(&flagNoteSCMAutoCommit, "autocommit", "a",
		true, "Auto commit the note to the repository")
	cmdNote.Flags().BoolVarP(&flagNoteSCMAutoPush, "autopush", "p",
		false, "Auto push the note to the origin remote of the repository")

	flagNoteTimeStampParsed = time.Now()
	cmdNote.Flags().StringVarP(&flagNoteTimeStamp, "timestamp", "t",
		flagNoteTimeStampParsed.Format(project.RecordTimeStampFormat), "The timestamp for which the note will be recorded")
}

func runNote(cmd *cobra.Command, args []string) {
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

	var timestamp time.Time
	if flagNoteTimeStampParsed.Format(project.RecordTimeStampFormat) == flagNoteTimeStamp {
		timestamp = flagNoteTimeStampParsed
	} else {
		var err error
		timestamp, err = now.Parse(flagNoteTimeStamp)
		if err != nil {
			l.Alert("can not parse timestamp: ", err)
			os.Exit(1)
		}
	}

	note := project.Note{
		Project:   args[0],
		Value:     strings.Join(args[1:], " "),
		TimeStamp: timestamp,
	}

	l.Debug("Note: ", note)

	err := project.WriteRecord(note, flagLablogDataDir, flagNoteSCM,
		flagNoteSCMAutoCommit, flagNoteSCMAutoPush)
	if err != nil {
		l.Alert("can not write note: ", errgo.Details(err))
		os.Exit(1)
	}
}
