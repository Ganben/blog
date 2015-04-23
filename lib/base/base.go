package base

import (
	"github.com/gofxh/blog/lib/core"
	"github.com/gofxh/blog/lib/entity"
)

var (
	Action *core.Action

	Command *core.Command
	Config  *core.Config
	Storage *core.Storage

	Max *entity.Maxer
)
