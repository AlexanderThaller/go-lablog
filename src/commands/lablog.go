package commands

import "github.com/spf13/cobra"

var lablogCmd = &cobra.Command{
	Use:   BuildName,
	Short: BuildName + " makes taking notes and todos simple",
	Long: BuildName + ` orders notes and todos into projects and subprojects
  without dictating a specific format`,
	Run: runListProjects,
}
