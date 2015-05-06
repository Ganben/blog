package helper

import (
	"github.com/gofxh/blog/lib/base"
	"github.com/gofxh/blog/lib/core"
)

func Init(_ interface{}) *core.ActionResult {
	// init theme manager
	initTheme(base.Config, "")

	// add user auth handler
	base.Server.Use(UseAuth())

	return core.NewOKActionResult(nil)
}
