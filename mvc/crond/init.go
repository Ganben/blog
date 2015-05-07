package crond

import (
	"github.com/gofxh/blog/lib/base"
	"github.com/gofxh/blog/lib/core"
)

func Init(_ interface{}) *core.ActionResult {
	base.Cron.AddFunc("0 */30 * * * *", UserCron)
	base.Cron.Start()
	return nil
}
