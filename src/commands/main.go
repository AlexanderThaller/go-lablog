package commands

import "github.com/AlexanderThaller/cobra"

var mainCmd = &cobra.Command{
	Use:   "lablog",
	Short: "lablog makes taking notes and todos easy.",
	Long:  `lablog orders notes and todos into projects and subprojects without dictating a specific format.`,
}

func Run() {
	mainCmd.AddCommand(cmdVersion)

	mainCmd.Execute()
}
