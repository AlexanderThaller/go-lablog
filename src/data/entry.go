package data

import (
	"time"

	"github.com/jinzhu/now"
)

const EntryCSVTimeStampFormat = time.RFC3339Nano

func init() {
	now.TimeFormats = append(now.TimeFormats, EntryCSVTimeStampFormat)
}

type Entry interface {
	ValueArray() []string
	GetProject() Project
	Type() string
	GetTimeStamp() time.Time
}
