package data

import "time"

const EntryCSVTimeStampFormat = time.RFC3339Nano

type Entry interface {
	ValueArray() []string
	GetProject() Project
}
