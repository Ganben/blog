package core

import (
	"github.com/gofxh/blog/app/log"
	"os"
	"os/signal"
	"reflect"
	"runtime"
	"sync"
	"syscall"
	"time"
)

var wg WaitGroup

// wrap a goroutine to global group
func Wrap(funcName string, fn func()) {
	wg.Wrap(funcName, fn)
}

// wait global group
func WrapWait() {
	wg.Wait()
}

// wait group
type WaitGroup struct {
	sync.WaitGroup
}

// wrap a function in global wait group
func (w *WaitGroup) Wrap(funName string, fn func()) {
	w.Add(1)
	go func() {
		t := time.Now()
		fn()
		w.Done()
		log.Debug("%s|%.1fms|%d", funName, time.Since(t).Seconds()*1000, runtime.NumGoroutine())
		// exit goroutine
		runtime.Goexit()
	}()
}

// service struct
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
