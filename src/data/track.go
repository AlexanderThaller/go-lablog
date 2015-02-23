package data

import (
	"io"
	"strconv"
	"time"
)

type Track struct {
	Active bool
	Project
	TimeStamp time.Time
}

func (track Track) ValueArray() []string {
	return []string{
		track.TimeStamp.Format(EntryCSVTimeStampFormat),
		"track",
		strconv.FormatBool(track.Active),
	}
}

func (track Track) GetProject() Project {
	return track.Project
}

func (track Track) Type() string {
	return "track"
}

func (track Track) GetTimeStamp() time.Time {
	return track.TimeStamp
}

func (track Track) Format(writer io.Writer, indent uint) {
	io.WriteString(writer, "* ")
	io.WriteString(writer, track.TimeStamp.Format(EntryCSVTimeStampFormat))

	if track.Active {
		io.WriteString(writer, " (+)")
	} else {
		io.WriteString(writer, " (-)")
	}

	io.WriteString(writer, "\n")
}

type TracksByTimeStamp []Track

func (by TracksByTimeStamp) Len() int {
	return len(by)
}

func (by TracksByTimeStamp) Swap(i, j int) {
	by[i], by[j] = by[j], by[i]
}

func (by TracksByTimeStamp) Less(i, j int) bool {
	return by[i].TimeStamp.Before(by[j].TimeStamp)
}
