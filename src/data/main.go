package data

import (
	"fmt"

	"github.com/juju/errgo"
)

func Record(datadir string, entry Entry) error {
	fmt.Println(entry.CSV())
	return errgo.New("not implemented")
}
