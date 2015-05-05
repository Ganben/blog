package helper

import (
	"github.com/gofxh/blog/lib/base"
	"github.com/gofxh/blog/lib/core"
)

func Init(_ interface{}) *core.ActionResult {
	initTheme(base.Config, "")
	return core.NewOKActionResult(nil)
}
