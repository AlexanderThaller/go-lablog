package data

import "time"

type Entries []Entry

type Entry interface {
	Type() EntryType
	Values() []string
}

const TimeStampFormat = time.RFC3339Nano

type EntryType int

const (
	EntryTypeNote EntryType = iota
)

func (etype EntryType) String() string {
	switch etype {
	case EntryTypeNote:
		return "note"
	default:
		return "unkown"
	}
}
