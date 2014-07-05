package main

import (
	"github.com/codegangsta/cli"
	"github.com/mephux/todo/lib"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Name = "todo"
	app.Usage = "simple todo management"
	app.Version = todo.VERSION
	app.Commands = todo.Commands()

	app.Run(os.Args)
}
