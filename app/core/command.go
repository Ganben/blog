package core

import "github.com/codegangsta/cli"

type Command struct {
	*cli.App
	commands []cli.Command
}

func NewCommand() *Command {
	app := cli.NewApp()
	app.Name = APP_NAME
	app.Version = APP_VERSION
	return &Command{app, nil}
}

// add command simply
func (c *Command) Register(cmd cli.Command) {
	c.commands = append(c.commands, cmd)
}

// run command line
func (c *Command) Run() {
	c.App.Commands = c.commands
	c.App.RunAndExitOnError()
}
