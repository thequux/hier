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

type  NewTicketParams struct{
}

var _ = fmt.Errorf
func Run(ctx *cli.Context) {
	var app WebApp
	var err error
	app.App, err = common.OpenRepo(ctx.Args().First())
	if err != nil {
		panic(err)
	}
	_ = app

	
	println("Running!")
	r := gin.Default()
	r.SetHTMLTemplate(LoadTemplates())
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "Pong")
	})
	r.GET("/new_ticket", func(c *gin.Context) {
		c.HTML(200, "templates/new_ticket.html", app.App.TicketConfig())
	})
	r.Run(":8082")
}
