package commands

import "github.com/AlexanderThaller/cobra"

var cmdMain = &cobra.Command{
	Use:   "lablog",
	Short: "lablog makes taking notes and todos easy.",
	Long:  `lablog orders notes and todos into projects and subprojects without dictating a specific format.`,
}

func Run() {
	cmdMain.AddCommand(cmdVersion)

	cmdMain.AddCommand(cmdShow)
	cmdShow.AddCommand(cmdShowProjects)

	cmdMain.Execute()
}
