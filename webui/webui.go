package webui

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/elazarl/go-bindata-assetfs"
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

func (app *WebApp) NewTicket(c *gin.Context) {
	type TicketParams struct {
		Title string
		Type string
		Status string
		Substatus string
		
	}
}

func (app *WebApp) StaticData(c *gin.Context) {
	config := app.App.TicketConfig()
	// We generate this without a template to make things easier.
	
	statusList, err := json.Marshal(config.Statuses)
	typeList, err := json.Marshal(config.Types)
	_ = err
	w := c.Writer
	w.Header().Set("content-type", "text/javascript")
	w.WriteHeader(200)
	w.Write([]byte("if (typeof Hier === 'undefined') Hier = {};\n"))
	w.Write([]byte("Hier.statuses = "))
	w.Write(statusList)
	w.Write([]byte(";\nHier.types = "))
	w.Write(typeList)
	w.Write([]byte(";\n"))
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
	r.GET("/js/data.js", app.StaticData)
	r.GET("/ticket/new", func(c *gin.Context) {
		c.HTML(200, "templates/new_ticket.html", app.App.TicketConfig())
	})
	r.POST("/ticket/new", app.NewTicket)
	r.ServeFiles("/static/*filepath",
		&assetfs.AssetFS{Asset:Asset,AssetDir:AssetDir,Prefix:"static"})
	r.Run(":8082")
}
