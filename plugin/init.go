package plugin

import (
	"github.com/gofxh/blog/lib/base"
	"github.com/gofxh/blog/lib/core"
	"github.com/gofxh/blog/lib/log"
)

func Init(_ interface{}) *core.ActionResult {
	for k, p := range *(base.Plugin) {
		p.Init()
		log.Debug("Plugin|Init|%s", k)
	}
	return core.NewOKActionResult(nil)
}
