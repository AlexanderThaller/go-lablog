package formatting

import (
	"strings"

	log "github.com/Sirupsen/logrus"
)

func HeaderIndent(char string, indent int) string {
	log.Debug("Indent: ", indent)
	if indent < 1 {
		return ""
	}

	out := strings.Repeat(char, indent)

	return out
}
