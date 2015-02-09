package commands

import (
	"os"

	"github.com/AlexanderThaller/cobra"
	"github.com/AlexanderThaller/logger"
)

var cmdTrack = &cobra.Command{
	Use:    "track (command) [project]",
	Short:  "Track projects.",
	Long:   `Track projects.`,
	Run:    runTrackToggle,
	PreRun: setLogLevel,
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

func init() {
	cmdTrack.AddCommand(cmdTrackStart)
	cmdTrack.AddCommand(cmdTrackStop)
}

func runTrackToggle(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "track", "toggle")
	l.Alert("not implemented")
	os.Exit(1)
}

func runTrackStart(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "track", "start")
	l.Alert("not implemented")
	os.Exit(1)
}

func runTrackStop(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "track", "stop")
	l.Alert("not implemented")
	os.Exit(1)
}
