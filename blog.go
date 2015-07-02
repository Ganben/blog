package main

import (
	"github.com/gofxh/blog/app"
	"github.com/gofxh/blog/app/cmd"
)

func main() {
	app.Command.Register(cmd.Init)
	app.Command.Register(cmd.Serv)
	app.Command.Run()
}
