package app

import "github.com/gofxh/blog/app/core"

// global variables
var (
	// something can init
	Command *core.Command = core.NewCommand()
	Config  *core.Config  = core.NewConfig()

	Db     *core.Database
	Server *core.Server
)
