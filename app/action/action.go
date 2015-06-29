// all actions in app.
// actions are real functionality which called by code to execute.
package action

// action params
type Param interface{}

// action result
type Result struct {
	Data   map[string]interface{} `json:"data,omitempty"`
	Status bool                   `json:"status"`          // result status, true to success, false to fail
	Error  string                 `json:"error,omitempty"` // error message
}

// action func, use params to solve result
type Func func(Param) *Result

// action before func, check and change params to determine going on or not (true or false)
type BeforeFunc func(*Param) bool

// action after func, change final result
type AfterFunc func(*Result, *Param)
