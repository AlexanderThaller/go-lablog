package commands

import "github.com/spf13/cobra"

var cmdTodo = &cobra.Command{
	Use:   "todo [project] [text]",
	Short: "Todo projects",
	Long:  `Todo projects`,
	Run:   runTodo,
}

func runTodo(cmd *cobra.Command, args []string) {
}
