package core

import (
	"github.com/gofxh/blog/lib/log"
	"reflect"
	"runtime"
)

// action manager struct
type Action struct {
	before  map[string][]ActionBeforeFunc // before caller
	after   map[string][]ActionAfterFunc  // after caller
	actions map[string]bool               // actions' list
}

// new action manager
func NewAction() *Action {
	return &Action{
		actions: make(map[string]bool),
		before:  make(map[string][]ActionBeforeFunc),
		after:   make(map[string][]ActionAfterFunc),
	}
}

// call action with params,
// usage:
//    a.Call(method,params)
//
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

// add before caller
func (a *Action) Before(fn ActionFunc, beforeFunc ActionBeforeFunc) {
	name := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	a.actions[name] = true
	a.before[name] = append(a.before[name], beforeFunc)
}

// add after caller
func (a *Action) After(fn ActionFunc, afterFunc ActionAfterFunc) {
	name := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	a.actions[name] = true
	a.after[name] = append(a.after[name], afterFunc)
}

// action function type
type ActionFunc func(v interface{}) *ActionResult

// action before caller type
type ActionBeforeFunc func(interface{})

// action after caller type
type ActionAfterFunc func(*ActionResult)

// action result meta,
// with status and error message
type ActionResultMeta struct {
	Status       bool   `json:"status"`
	ErrorCode    int    `json:"error_code,emitempty"`
	ErrorMessage string `json:"error_message,emitempty"`
}

// action result data
type ActionResult struct {
	Meta ActionResultMeta `json:"meta"`
	Data AData            `json:"data"`
}

// action result map data type
type AData map[string]interface{}

// new success result with data
func NewOKActionResult(data AData) *ActionResult {
	return &ActionResult{
		Meta: ActionResultMeta{
			Status: true,
		},
		Data: data,
	}
}

// new error result with err data
func NewSystemErrorResult(err error) *ActionResult {
	return &ActionResult{
		Meta: ActionResultMeta{
			Status:       false,
			ErrorMessage: err.Error(),
		},
		Data: nil,
	}
}
