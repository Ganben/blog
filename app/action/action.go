// all actions in app.
// actions are real functionality which called by code to execute.
package action

import (
	"errors"
	"fmt"
	"github.com/gofxh/blog/app/log"
	"reflect"
	"runtime"
	"strings"
)

// action result
type Result struct {
	Data   map[string]interface{} `json:"data,omitempty"`
	Status bool                   `json:"status"`          // result status, true to success, false to fail
	Error  string                 `json:"error,omitempty"` // error message
}

// new ok result
func OkResult(data map[string]interface{}) *Result {
	return &Result{
		Data:   data,
		Status: true,
	}
}

// error result
func ErrorResult(err error) *Result {
	return &Result{
		Data:   nil,
		Status: false,
		Error:  err.Error(),
	}
}

// action func, use params to solve result
type Func func(interface{}) *Result

// action before func, check and change params to determine going on or not (true or false)
type BeforeFunc func(*interface{}) bool

// action after func, change final result
type AfterFunc func(*Result, *interface{})

// get function name
func funcName(fn Func) string {
	name := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	nameData := strings.Split(name, "/")
	if len(nameData) > 3 {
		nameData = nameData[len(nameData)-3:]
	}
	return strings.Join(nameData, "/")
}

// global action manager
var a *actionManager = new(actionManager)

// action manager, maintain before and after functions
type actionManager struct {
	before map[string][]BeforeFunc
	after  map[string][]AfterFunc
}

// add before caller
func Before(fn Func, beforeFunc BeforeFunc) {
	name := funcName(fn)
	a.before[name] = append(a.before[name], beforeFunc)
}

// add after caller
func After(fn Func, afterFunc AfterFunc) {
	name := funcName(fn)
	a.after[name] = append(a.after[name], afterFunc)
}

// call an action
func Call(fn Func, param interface{}) *Result {
	name := funcName(fn)
	log.Debug("Action|Call|%s", name)
	if len(a.before[name]) > 0 {
		for _, b := range a.before[name] {
			b(&param)
		}
	}
	result := fn(param)
	if len(a.after[name]) > 0 {
		for _, af := range a.after[name] {
			af(result, &param)
		}
	}
	if !result.Status {
		log.Warn("Action|Call|%s|%s", name, result.Error)
	}
	return result
}

// action param type error
func paramTypeError(v interface{}) error {
	return errors.New(fmt.Sprintf("need %s", reflect.TypeOf(v)))
}
