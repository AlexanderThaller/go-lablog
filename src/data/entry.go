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

func FilterEntriesBeforeTimeStamp(entries []Entry, start time.Time) []Entry {
	var out []Entry

	for _, entry := range entries {
		if entry.GetTimeStamp().Before(start) {
			continue
		}

		out = append(out, entry)
	}

	return out
}

func FilterEntriesAfterTimeStamp(entries []Entry, end time.Time) []Entry {
	var out []Entry

	for _, entry := range entries {
		if entry.GetTimeStamp().After(end) {
			continue
		}

		out = append(out, entry)
	}

	return out
}
