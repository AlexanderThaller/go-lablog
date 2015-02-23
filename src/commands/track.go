package commands

import (
	"sort"
	"time"

	"github.com/AlexanderThaller/cobra"
	"github.com/AlexanderThaller/lablog/src/data"
	"github.com/AlexanderThaller/lablog/src/helper"
	"github.com/AlexanderThaller/logger"
	"github.com/juju/errgo"
)

var cmdTrack = &cobra.Command{
	Use:   "track (command) [project]",
	Short: "Track projects.",
	Long:  `Track projects.`,
	Run:   runTrackToggle,
}

var cmdTrackStart = &cobra.Command{
	Use:   "start (project)",
	Short: "Start Track projects.",
	Long:  `Start Track projects.`,
	Run:   runTrackStart,
}

var cmdTrackStop = &cobra.Command{
	Use:   "stop (project)",
	Short: "Stop Track projects.",
	Long:  `Stop Track projects.`,
	Run:   runTrackStop,
}

var cmdTrackToggle = &cobra.Command{
	Use:   "toggle (project)",
	Short: "Toggle Track projects.",
	Long:  `Toggle Track projects.`,
	Run:   runTrackToggle,
}

var flagTrackTimeStamp time.Time
var flagTrackTimeStampRaw string

func init() {
	flagTrackTimeStamp = time.Now()

	cmdTrack.PersistentFlags().StringVarP(&flagTrackTimeStampRaw, "timestamp", "t",
		flagTrackTimeStamp.String(), "The timestamp for which to record the todo.")
}

func runTrackStart(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "track", "start")

	if len(args) < 1 {
		errexit(l, errgo.New("need at least one argument to run"))
	}

	project := args[0]

	timestamp, err := helper.DefaultOrRawTimestamp(flagTrackTimeStamp, flagTrackTimeStampRaw)
	errexit(l, err, "can not get timestamp")

	track := data.Track{
		Active:    true,
		Project:   data.Project{Name: project},
		TimeStamp: timestamp,
	}

	l.Trace("Track: ", track)
	recordAndCommit(l, flagLablogDataDir, track)
}

func runTrackStop(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "track", "stop")

	if len(args) < 1 {
		errexit(l, errgo.New("need at least one argument to run"))
	}

	project := args[0]

	timestamp, err := helper.DefaultOrRawTimestamp(flagTrackTimeStamp, flagTrackTimeStampRaw)
	errexit(l, err, "can not get timestamp")

	track := data.Track{
		Active:    false,
		Project:   data.Project{Name: project},
		TimeStamp: timestamp,
	}

	l.Trace("Track: ", track)
	recordAndCommit(l, flagLablogDataDir, track)
}

func runTrackToggle(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "track", "toggle")

	if len(args) < 1 {
		errexit(l, errgo.New("need at least one arguments to run"))
	}

	project := data.Project{Name: args[0], Datadir: flagLablogDataDir}

	timestamp, err := helper.DefaultOrRawTimestamp(flagTodoTimeStamp, flagTodoTimeStampRaw)
	errexit(l, err, "can not get timestamp")

	tracks, _ := project.Tracks()

	var active bool
	if len(tracks) >= 1 {
		sort.Sort(data.TracksByTimeStamp(tracks))
		active = tracks[len(tracks)-1].Active
	}

	track := data.Track{
		Active:    !active,
		Project:   project,
		TimeStamp: timestamp,
	}

	l.Trace("Track: ", track)
	recordAndCommit(l, flagLablogDataDir, track)
}
