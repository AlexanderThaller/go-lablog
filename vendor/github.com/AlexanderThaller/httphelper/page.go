package httphelper

import (
	"net/http"

	"github.com/juju/errgo"
	"github.com/julienschmidt/httprouter"
)

func PageRouterNotFound(w http.ResponseWriter, r *http.Request, p httprouter.Params) *HandlerError {
	return NewHandlerError(errgo.New("page not found"), http.StatusNotFound)
}

func PageRouterMethodNotAllowed(w http.ResponseWriter, r *http.Request, p httprouter.Params) *HandlerError {
	return NewHandlerError(errgo.New("method not allowed"), http.StatusMethodNotAllowed)
}

func PageRouterStopping(w http.ResponseWriter, r *http.Request, p httprouter.Params) *HandlerError {
	return NewHandlerErrorDef(errgo.New("service is stopping"))
}

func PageMinimalFavicon(w http.ResponseWriter, r *http.Request, p httprouter.Params) *HandlerError {
	raw, err := Asset("data/favicon.ico")
	if err != nil {
		return NewHandlerErrorDef(errgo.Notef(err, "can not read raw page"))
	}

	_, err = w.Write(raw)
	if err != nil {
		return NewHandlerErrorDef(errgo.Notef(err, "can not write raw data to responsewriter"))
	}

	return nil
}
