package httphelper

import (
	"os"
	"os/signal"
	"syscall"
)

var (
	Stopping bool
)

func WaitForStopSignal() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	<-signals
}
