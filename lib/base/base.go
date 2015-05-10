package base

import (
	"github.com/gofxh/blog/lib/core"
	"github.com/gofxh/blog/lib/entity"
)

var (
	Action *core.Action // action manager

	Command *core.Command // command line
	Config  *core.Config  // config data
	Storage *core.Storage // storage engine
	Server  *core.Server  // http server
	Cron    *core.Cron    // cron task
	Plugin  *core.Plugins // plugins

	Max *entity.Maxer // max-id generator
)

func init() {
	Action = core.NewAction()
	Plugin = core.NewPlugins()
}
