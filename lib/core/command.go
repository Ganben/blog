package core

import "github.com/codegangsta/cli"

// command line struct
type Command struct {
	*cli.App
}

// new command line
func NewCommand() *Command {
	app := cli.NewApp()
	// fill app info
	app.Name = APP_NAME
	app.Version = APP_VERSION
	app.Usage = APP_USAGE
	return &Command{app}
}

// run command line
func (c *Command) Run() {
	c.RunAndExitOnError()
}
