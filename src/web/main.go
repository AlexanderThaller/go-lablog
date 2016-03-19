package web

import (
	"time"

	"github.com/AlexanderThaller/httphelper"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
	"github.com/phyber/negroni-gzip/gzip"
	"gopkg.in/tylerb/graceful.v1"
)

func Listen(datadir, binding string, loglevel log.Level) error {
	router := httprouter.New()

	// Router handler
	router.MethodNotAllowed = httphelper.HandlerLoggerHTTP(httphelper.PageRouterMethodNotAllowed)
	router.NotFound = httphelper.HandlerLoggerHTTP(httphelper.PageRouterNotFound)

	// Root and Favicon
	router.GET("/", httphelper.HandlerLoggerRouter(pageRoot))
	router.GET("/favicon.ico", httphelper.HandlerLoggerRouter(httphelper.PageMinimalFavicon))

	server := negroni.New()
	server.UseHandler(router)

	// Recovery
	server.Use(negroni.NewRecovery())

	// GZIP
	server.Use(gzip.Gzip(gzip.DefaultCompression))

	log.Info("Listening on ", binding)
	graceful.Run(binding, 10*time.Second, server)

	return nil
}
