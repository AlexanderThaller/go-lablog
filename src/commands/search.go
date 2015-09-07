package commands

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/AlexanderThaller/lablog/src/data"

	"github.com/AlexanderThaller/cobra"
	"github.com/AlexanderThaller/logger"
)

var cmdSearch = &cobra.Command{
	Use:   "search (text)",
	Short: "Search in notes, todos, tracks, etc., for a given string.",
	Long:  "Search in notes, todos, tracks, etc., for a given string.",
	Run:   runSearch,
}

func runSearch(cmd *cobra.Command, args []string) {
	l := logger.New("commands", "list", "entries")

	projects, err := data.GetProjects(flagLablogDataDir)
	errexit(l, err, "can not get projects")

	sort.Sort(data.ProjectsByName(projects))

	buffer := new(bytes.Buffer)

	for _, project := range projects {
		var projfindings []string

		for _, searchstr := range args {
			findings, err := project.Search(searchstr)
			errexit(l, err, "can not search in project")

			for _, finding := range findings {
				projfindings = append(projfindings, finding)
			}
		}

		if len(projfindings) != 0 {
			buffer.WriteString("= " + project.Name + "\n")

			for _, finding := range projfindings {
				buffer.WriteString(finding + "\n")
			}

			buffer.WriteString("\n")
		}
	}

	fmt.Print(buffer.String())
}
