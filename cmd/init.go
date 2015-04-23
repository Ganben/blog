package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/gofxh/blog/lib/base"
	"github.com/gofxh/blog/lib/core"
)

func Init() {
	base.Action = core.NewAction()

	if base.Command == nil {
		base.Command = core.NewCommand()
	}

	base.Command.Commands = []cli.Command{
		installCommand,
	}
}

func Run() {
	if base.Command == nil {
		Init()
	}
	base.Command.Run()
}
