package commands

import "github.com/spf13/cobra"

var cmdList = &cobra.Command{
	Use:   "list (dates|notes|projects|todos|tracks)",
	Short: "List x",
	Long:  `List x`,
	Run:   runListProjects,
}

var cmdListDates = &cobra.Command{
	Use:   "dates",
	Short: "List dates",
	Long:  `List dates`,
	Run:   runListDates,
}

var cmdListNotes = &cobra.Command{
	Use:   "notes",
	Short: "List notes",
	Long:  `List notes`,
	Run:   runListNotes,
}

var cmdListProjects = &cobra.Command{
	Use:   "projects",
	Short: "List projects",
	Long:  `List projects`,
	Run:   runListProjects,
}

var cmdListTodos = &cobra.Command{
	Use:   "todos",
	Short: "List todos",
	Long:  `List todos`,
	Run:   runListTodos,
}

var cmdListTracks = &cobra.Command{
	Use:   "tracks",
	Short: "List tracks",
	Long:  `List tracks`,
	Run:   runListTracks,
}

var cmdListTracksActive = &cobra.Command{
	Use:   "active",
	Short: "List tracks",
	Long:  `List tracks`,
	Run:   runListTracksActive,
}

var cmdListTracksDurations = &cobra.Command{
	Use:   "durations",
	Short: "List durations",
	Long:  `List durations`,
	Run:   runListTracksDurations,
}

func init() {
	cmdList.AddCommand(cmdListDates)
	cmdList.AddCommand(cmdListNotes)
	cmdList.AddCommand(cmdListProjects)
	cmdList.AddCommand(cmdListTodos)

	cmdListTracks.AddCommand(cmdListTracksActive)
	cmdListTracks.AddCommand(cmdListTracksDurations)
	cmdList.AddCommand(cmdListTracks)
}

func runListDates(cmd *cobra.Command, args []string) {
}

func runListDurations(cmd *cobra.Command, args []string) {
}

func runListNotes(cmd *cobra.Command, args []string) {
}

func runListProjects(cmd *cobra.Command, args []string) {
}

func runListTodos(cmd *cobra.Command, args []string) {
}

func runListTracks(cmd *cobra.Command, args []string) {
}

func runListTracksActive(cmd *cobra.Command, args []string) {
}

func runListTracksDurations(cmd *cobra.Command, args []string) {
}
