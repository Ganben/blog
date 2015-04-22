package core

import "github.com/codegangsta/cli"

type Command struct {
	*cli.App
}

func NewCommand() *Command {
	app := cli.NewApp()
	app.Name = APP_NAME
	app.Version = APP_VERSION
	app.Usage = APP_USAGE
	return &Command{app}
}

func (c *Command) Run() {
	c.RunAndExitOnError()
}
