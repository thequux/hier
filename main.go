package main

import (
	"github.com/codegangsta/cli"
	"github.com/thequux/hier/webui"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name: "ui",
			Usage: "Run a web UI",
			Action: webui.Run,
		},
	}
	app.Run(os.Args)
}
