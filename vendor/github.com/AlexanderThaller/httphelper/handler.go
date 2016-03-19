package httphelper

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/juju/errgo"
	"github.com/julienschmidt/httprouter"
)

type Handler func(http.ResponseWriter, *http.Request, httprouter.Params) *HandlerError
type HandlerError struct {
	Error error
	Code  int
}

func NewHandlerErrorDef(err error) *HandlerError {
	return NewHandlerError(err, http.StatusInternalServerError)
}

func NewHandlerError(err error, code int) *HandlerError {
	return &HandlerError{Error: err, Code: code}
}

func HandlerLoggerHTTP(fn Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if Stopping {
			handlerLogger(PageRouterStopping, w, r, nil)
		} else {
			handlerLogger(fn, w, r, nil)
		}
	},
	)
}

func HandlerLoggerRouter(fn Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if Stopping {
			handlerLogger(PageRouterStopping, w, r, p)
		} else {
			handlerLogger(fn, w, r, p)
		}
	}
}

func NewHandlerLogEntry(r *http.Request) *log.Entry {
	remoteAddr := r.RemoteAddr
	if realIP := r.Header.Get("X-Real-IP"); realIP != "" {
		remoteAddr = realIP
	}

	entry := log.WithFields(log.Fields{
		"context": "http",
		"url":     r.URL,
		"remote":  remoteAddr,
	})

	var err error
	entry.Data["request_id"], err = generateID(16)
	if err != nil {
		panic("can not generate id for request")
	}

	return entry
}

func handlerLogger(fn Handler, w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	l := NewHandlerLogEntry(r)
	l.Info("started handling request")

	starttime := time.Now()
	err := fn(w, r, p)
	duration := time.Since(starttime)

	l.Data["duration"] = duration

	code := http.StatusOK

	if err != nil {
		code = err.Code
	}

	l.Data["status"] = code
	l.Data["text_status"] = http.StatusText(code)

	if err != nil {
		if err.Error != nil {
			l.Data["error_message"] = err.Error

			l.Debug(errgo.Details(err.Error))
			l.Error("completed handling request")

			scode := strconv.Itoa(err.Code)
			http.Error(w, scode+" ("+http.StatusText(err.Code)+") - "+err.Error.Error(), err.Code)
		}
	} else {
		l.Info("completed handling request")
	}
}

func generateID(n int) (string, error) {
	r := make([]byte, n)
	_, err := rand.Read(r)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(r), nil
}
