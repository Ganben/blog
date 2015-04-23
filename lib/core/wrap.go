package core

import (
	"github.com/gofxh/blog/lib/log"
	"runtime"
	"sync"
	"time"
)

var wg WaitGroup

// wrap a goroutine to global group
func Wrap(funcName string, fn func()) {
	wg.Wrap(funcName, fn)
}

func WrapWait() {
	wg.Wait()
}

type WaitGroup struct {
	sync.WaitGroup
}

func (w *WaitGroup) Wrap(funName string, fn func()) {
	w.Add(1)
	go func() {
		t := time.Now()
		fn()
		w.Done()
		log.Debug("%s|%.1f|%d", funName, time.Since(t).Seconds()*1000, runtime.NumGoroutine())
		// exit goroutine
		runtime.Goexit()
	}()
}
