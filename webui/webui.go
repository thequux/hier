package webui

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-gonic/gin"
	"github.com/thequux/hier/common"
	"github.com/thequux/hier/data"
	"sort"
	"time"
)

type WebApp struct {
	App *common.AppData
}

type TicketSortOrder int
const (
	TicketSortByID TicketSortOrder = iota
	TicketSortByName
	TicketSortByCTime
	TicketSortByMTime
)

type TicketSort struct{
	list []*common.Ticket
	order TicketSortOrder
}

func (s TicketSort) Sort() { sort.Sort(s) }
func (s TicketSort) Len() int { return len(s.list) }
func (s TicketSort) Swap(i,j int) { s.list[i],s.list[j] = s.list[j],s.list[i] }
func (s TicketSort) Less(i,j int) bool {
	a,b := s.list[i], s.list[j]
	switch s.order {
	default: fallthrough
	case TicketSortByID:
		for i := 0; i < 20; i++ {
			if a.Hash[i] == b.Hash[i] {
				continue
			}
			return a.Hash[i] < b.Hash[i]
		}
		return false
	case TicketSortByName:
		return a.Title < b.Title
	}
}

func (app *WebApp) ListTickets(c *gin.Context) {
	tickets := app.App.Tickets()
	TicketSort{tickets, TicketSortByName}.Sort()
	c.HTML(200, "templates/ticket_list.html", SkeletonParams{
		Title: "Tickets",
		Content: tickets,
	})
}

type TicketParams struct {
	Title string `json:"title" form:"title" binding:"required"`
	Type string `json:"type" form:"type" binding:"required"`
	Status string `json:"status" form:"status" binding:"required"`
	Resolution string `json:"resolution" form:"resolution"`
	Message string `json:"message" form:"message"`
}

type TicketCommentParams struct {
	Status string `json:"status" form:"status" binding:"required"`
	Resolution string `json:"resolution" form:"resolution"`
	Message string `json:"message" form:"message"`
}

func (app *WebApp) NewTicketComment(c *gin.Context) {
	if c.Params.ByName("id") == "new" {
		var params TicketParams
		if !c.Bind(&params) {
			// TODO: redisplay the New Ticket page
			c.String(500, "Error!")
			return
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
			panic(err) // TODO: Handle better
		}
		c.Redirect(303, fmt.Sprintf("/ticket/%s", ticket.Hash))
	} else {
		var params TicketCommentParams
		if !c.Bind(&params) {
			// TODO: redisplay the New Ticket page
			c.String(500, "Error!")
			return
		}
		ticketId, err := common.ParseTicketId(c.Params.ByName("id"))
		if err != nil {
			panic(err) // TODO: Handle better
		}
		ticket, err := app.App.GetTicket(ticketId)
		if err != nil {
			panic(err) // TODO: Handle better
		}
		artifact := &common.TicketArtifact{
			Message: params.Message,
		}
		if ticket.Status != params.Status || ticket.Resolution != params.Resolution {
			artifact.Status = params.Status
			artifact.Resolution = params.Resolution
		}			
		ticket.NewComment(artifact)
			
		// Redirect to the same page; we are always requested
		// via POST, so this won't introduce a redirect loop.
		c.Redirect(303, c.Request.URL.String())
	}
}

type TicketDisplayParams struct {
	Ticket *common.Ticket
	Config common.TicketConfig
}

func (app *WebApp) GetTicket(c *gin.Context) {
	id := c.Params.ByName("id")
	if id == "new" {
		c.HTML(200, "templates/new_ticket.html", SkeletonParams{
			Title: "New Ticket",
			Content: TicketDisplayParams{
				Config: app.App.TicketConfig(),
			},
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
			return
		}
		
		c.HTML(200, "templates/view_ticket.html", SkeletonParams{
			Title: "View Ticket",
			Content: TicketDisplayParams{
				Config: app.App.TicketConfig(),
				Ticket: ticket,
			},
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
	// Reloads templates if built with debug
	r.Use(TemplateReloader)
	r.HTMLRender = &TemplateRenderer{Template:LoadTemplates()}
	r.GET("/js/data.js", app.StaticData)
	r.GET("/ticket/:id", app.GetTicket)
	r.GET("/ticket", app.ListTickets)
	r.POST("/ticket/:id", app.NewTicketComment)
	r.ServeFiles("/static/*filepath",
		&assetfs.AssetFS{Asset:Asset,AssetDir:AssetDir,Prefix:"static"})
	r.Run(":8082")
}
