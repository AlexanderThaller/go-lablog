package data

import (
	"fmt"
	"time"
)

type Note struct {
	Project
	TimeStamp time.Time
	Text      string
}

func (note Note) CSV() string {
	return fmt.Sprintf("%+v", note)
}
