package webui

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"
	"github.com/thequux/hier/common"
)

type WebApp struct {
	App *common.AppData
}

func (app *WebApp) ListTickets(c *gin.Context) {
	for ticket := range app.App.Tickets() {
		var _ = ticket
	}
}

func Run(ctx *cli.Context) {
	var app WebApp
	var err error
	app.App, err = common.OpenRepo(ctx.Args().First())
	fmt.Print(app.App.TicketBranch())
	if err != nil {
		panic(err)
	}
	_ = app

	
	println("Running!")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "Pong")
	})
	r.Run(":8082")
}
