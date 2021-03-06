package web

import (
	"bytes"
	"html/template"
	"io"
	"net/http"
	"os/exec"

	"github.com/AlexanderThaller/httphelper"
	"github.com/AlexanderThaller/lablog/src/helper"
	"github.com/AlexanderThaller/lablog/src/store"
	log "github.com/Sirupsen/logrus"
	"github.com/juju/errgo"
	"github.com/julienschmidt/httprouter"
)

var (
	dataStore store.Store
)

func Listen(datadir, binding string, loglevel log.Level) error {
	var err error
	dataStore, err = helper.DefaultStore(datadir)
	if err != nil {
		return errgo.Notef(err, "can not get data store")
	}

	router := httprouter.New()

	// Router handler
	router.MethodNotAllowed = httphelper.HandlerLoggerHTTP(httphelper.PageRouterMethodNotAllowed)
	router.NotFound = httphelper.HandlerLoggerHTTP(httphelper.PageRouterNotFound)

	// Root and Favicon
	router.GET("/", httphelper.HandlerLoggerRouter(pageRoot))
	router.GET("/favicon.ico", httphelper.HandlerLoggerRouter(pageFavicon))

	// Show
	router.GET("/show/:type/", httphelper.HandlerLoggerRouter(pageShow))
	router.GET("/show/:type/:project", httphelper.HandlerLoggerRouter(pageShow))

	log.Info("Listening on ", binding)
	err = http.ListenAndServe(binding, router)
	if err != nil {
		return errgo.Notef(err, "can not listen to binding")
	}

	return nil
}

func getAssetTemplate(asset string) (*template.Template, error) {
	rawtmpl, err := Asset(asset)
	if err != nil {
		return nil, errgo.Notef(err, "can not get asset: "+asset)
	}

	tmpl, err := template.New(asset).Parse(string(rawtmpl))
	if err != nil {
		return nil, errgo.Notef(err, "can not parse template for asset: "+asset)
	}

	return tmpl, nil
}

func asciiDoctor(reader io.Reader, writer io.Writer) error {
	stderr := new(bytes.Buffer)

	command := exec.Command("asciidoctor", "-")
	command.Stdin = reader
	command.Stdout = writer
	command.Stderr = stderr

	err := command.Run()
	if err != nil {
		return errgo.Notef(errgo.Notef(err, "can not run asciidoctor"),
			stderr.String())
	}

	return nil
}
