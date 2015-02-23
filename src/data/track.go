package data

import (
	"io"
	"strconv"
	"time"

	"github.com/jinzhu/now"
	"github.com/juju/errgo"
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

func FilterTracksBeforeTimeStamp(tracks []Track, start time.Time) []Track {
	var entries []Entry
	for _, track := range tracks {
		entries = append(entries, track)
	}

	entries = FilterEntriesBeforeTimeStamp(entries, start)

	var out []Track
	for _, entry := range entries {
		out = append(out, entry.(Track))
	}

	return out
}

func FilterTracksAfterTimeStamp(tracks []Track, end time.Time) []Track {
	var entries []Entry
	for _, track := range tracks {
		entries = append(entries, track)
	}

	entries = FilterEntriesAfterTimeStamp(entries, end)

	var out []Track
	for _, entry := range entries {
		out = append(out, entry.(Track))
	}

	return out
}

func ParseTrack(project Project, values []string) (Track, error) {
	timestamp, err := now.Parse(values[0])
	if err != nil {
		return Track{}, errgo.Notef(err, "can not parse timestamp")
	}

	active, err := strconv.ParseBool(values[2])
	if err != nil {
		return Track{}, errgo.Notef(err, "can not parse active status")
	}

	track := Track{
		Project:   project,
		Active:    active,
		TimeStamp: timestamp,
	}

	return track, nil
}
