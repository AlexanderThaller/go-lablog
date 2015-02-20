package helper

import (
	"time"

	"github.com/jinzhu/now"
	"github.com/juju/errgo"
)

func DefaultOrRawTimestamp(timestamp time.Time, raw string) (time.Time, error) {
	if timestamp.String() == raw {
		return timestamp, nil
	}

	parsed, err := now.Parse(raw)
	if err != nil {
		return time.Time{}, errgo.Notef(err, "can not parse timestamp")
	}

	return parsed, nil
}
