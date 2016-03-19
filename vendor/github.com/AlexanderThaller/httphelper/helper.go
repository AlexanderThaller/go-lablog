package httphelper

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/juju/errgo"
)

func MarshalCompactJsonToWriter(writer io.Writer, data interface{}) *HandlerError {
	marshal, err := json.Marshal(&data)
	if err != nil {
		return NewHandlerErrorDef(errgo.Notef(err, "can not encode data"))
	}

	buffer := new(bytes.Buffer)
	err = json.Compact(buffer, marshal)
	if err != nil {
		return NewHandlerErrorDef(errgo.Notef(err, "can not compact json data"))
	}

	_, err = io.Copy(writer, buffer)
	if err != nil {
		return NewHandlerErrorDef(errgo.Notef(err, "can not copy buffer to writer"))
	}

	return nil
}
