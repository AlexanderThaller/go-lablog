package data

import "time"

type Note struct {
	Project
	TimeStamp time.Time
	Text      string
}

func (note Note) CSV() string {
	return ""
}
