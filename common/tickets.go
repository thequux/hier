package common

import (
	"github.com/libgit2/git2go"
)

type Ticket struct {
	Hash *git.Oid
	Title string
	Type string
	Status string
	artifacts []ticketArtifact
}

type ticketArtifact struct {
	
}

func (app *AppData) Tickets() []Ticket {
	if app.TicketBranch == nil {
		// No tickets
		return nil
	}
	return []Ticket{}
}
