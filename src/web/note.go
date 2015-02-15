package web

import (
	"time"

	"github.com/AlexanderThaller/lablog/src/data"
	"github.com/jinzhu/now"
	"github.com/juju/errgo"
)

type Note struct {
	Project   string
	TimeStamp string
	Text      string
}

func (note Note) ToData() (data.Note, error) {
	timestamp := time.Now()

	if note.TimeStamp != "" {
		var err error
		timestamp, err = now.Parse(note.TimeStamp)
		if err != nil {
			return data.Note{}, errgo.Notef(err, "can not parse timestamp")
		}
	}

	data := data.Note{
		Project:   data.Project{Name: note.Project},
		TimeStamp: timestamp,
		Text:      note.Text,
	}

	return data, nil
}
