package webui

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-gonic/gin"
	"github.com/thequux/hier/common"
	"github.com/thequux/hier/data"
	"time"
)

type WebApp struct {
	App *common.AppData
}

func (app *WebApp) ListTickets(c *gin.Context) {
	for ticket := range app.App.Tickets() {
		fmt.Println(ticket)
	}
}

type TicketParams struct {
	Title string `json:"title" form:"title" binding:"required"`
	Type string `json:"type" form:"type" binding:"required"`
	Status string `json:"status" form:"status" binding:"required"`
	Substatus string `json:"substatus" form:"substatus"`
	Message string `json:"message" form:"message"`
}

func (app *WebApp) NewTicket(c *gin.Context) {
	var params TicketParams
	if !c.Bind(&params) {
		// TODO: redisplay the New Ticket page
		c.String(500, "Error!")
	}
	now := time.Now().Format(time.RFC3339)
	artifact := &data.TicketArtifact{
		Title: &params.Title,
		Type: &params.Type,
		Status: &params.Status,
		Message: &params.Message,
		Date: &now,
	}
	ticket, err := app.App.NewTicket(artifact)
	if err != nil {
		panic(err)
	}
	c.Redirect(303, fmt.Sprintf("/ticket/%s", ticket.Hash))
}

func (app *WebApp) GetTicket(c *gin.Context) {
	id := c.Params.ByName("id")
	if id == "new" {
		c.HTML(200, "templates/new_ticket.html", SkeletonParams{
			Title: "New Ticket",
			Content: app.App.TicketConfig(),
		})
	} else {
		ticketId, err := common.ParseTicketId(id)
		if err != nil {
			c.String(404, "Invalid ticket ID")
			return
		}
		ticket, err := app.App.GetTicket(ticketId)
		if err != nil {
			c.String(500, err.Error())
		}
		
		c.HTML(200, "templates/view_ticket.html", SkeletonParams{
			Title: "View Ticket",
			Content: ticket,
		})
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
	r.HTMLRender = &TemplateRenderer{Template:LoadTemplates()}
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "Pong")
	})
	r.GET("/js/data.js", app.StaticData)
	r.GET("/ticket/:id", app.GetTicket)
	r.GET("/ticket", app.ListTickets)
	r.POST("/ticket/new", app.NewTicket)
	r.ServeFiles("/static/*filepath",
		&assetfs.AssetFS{Asset:Asset,AssetDir:AssetDir,Prefix:"static"})
	r.Run(":8082")
}
