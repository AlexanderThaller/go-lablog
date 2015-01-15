package commands

import "github.com/spf13/cobra"

var cmdTrack = &cobra.Command{
	Use:   "track (start|stop) (project)",
	Short: "Track projects",
	Long:  `Track projects`,
	Run:   runTrack,
}

var cmdTrackStart = &cobra.Command{
	Use:   "start (project)",
	Short: "Start Track projects",
	Long:  `Start Track projects`,
	Run:   runTrackStart,
}

var cmdTrackStop = &cobra.Command{
	Use:   "stop (project)",
	Short: "Stop Track projects",
	Long:  `Stop Track projects`,
	Run:   runTrackStop,
}

func init() {
	cmdTrack.AddCommand(cmdTrackStart)
	cmdTrack.AddCommand(cmdTrackStop)
}

func runTrack(cmd *cobra.Command, args []string) {
}

func runTrackStart(cmd *cobra.Command, args []string) {
}

func runTrackStop(cmd *cobra.Command, args []string) {
}
