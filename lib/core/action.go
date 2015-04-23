package core

import (
	"github.com/gofxh/blog/lib/log"
	"reflect"
	"runtime"
)

type Action struct {
	before  map[string][]ActionBeforeFunc
	after   map[string][]ActionAfterFunc
	actions map[string]bool
}

func NewAction() *Action {
	return &Action{
		actions: make(map[string]bool),
		before:  make(map[string][]ActionBeforeFunc),
		after:   make(map[string][]ActionAfterFunc),
	}
}

func (a *Action) Call(fn ActionFunc, param interface{}) *ActionResult {
	name := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	log.Debug("Action|Call|%s", name)
	a.actions[name] = true
	if len(a.before[name]) > 0 {
		for _, b := range a.before[name] {
			b(&param)
		}
	}
	result := fn(param)
	if len(a.after[name]) > 0 {
		for _, af := range a.after[name] {
			af(result)
		}
	}
	return result
}

func (a *Action) Before(fn ActionFunc, beforeFunc ActionBeforeFunc) {
	name := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	a.actions[name] = true
	a.before[name] = append(a.before[name], beforeFunc)
}

func (a *Action) After(fn ActionFunc, afterFunc ActionAfterFunc) {
	name := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	a.actions[name] = true
	a.after[name] = append(a.after[name], afterFunc)
}

type ActionFunc func(v interface{}) *ActionResult

type ActionBeforeFunc func(interface{})

type ActionAfterFunc func(*ActionResult)

type ActionResultMeta struct {
	Status       bool   `json:"status"`
	ErrorCode    int    `json:"error_code,emitempty"`
	ErrorMessage string `json:"error_message,emitempty"`
}

type ActionResult struct {
	Meta ActionResultMeta       `json:"meta"`
	Data map[string]interface{} `json:"data"`
}

func NewOKActionResult(data map[string]interface{}) *ActionResult {
	return &ActionResult{
		Meta: ActionResultMeta{
			Status: true,
		},
		Data: data,
	}
}
