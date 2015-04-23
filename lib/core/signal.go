package core

import (
	"os"
	"os/signal"
	"reflect"
	"syscall"

	"github.com/gofxh/blog/lib/log"
)

type Service interface {
	Start()
	Stop()
}

//  start a service with ctrl+c signal notify
func Start(s Service) {
	signalChan := make(chan os.Signal)
	sName := reflect.TypeOf(s).String()

	// notify exit signal
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	log.Debug("Start|%s", sName)

	// start
	s.Start()
	<-signalChan
	s.Stop()
	WrapWait() // wait global wait group

	log.Debug("Stop|%s", sName)
}
