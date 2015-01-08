package common

import (
	"fmt"
	"github.com/libgit2/git2go"
	"github.com/thequux/hier/data"
	"time"
)

type Ticket struct {
	Hash *git.Oid
	Title string
	Type string
	Status string
	artifacts []data.TicketArtifact
}

type ticketArtifact struct {
	
}

const ticketRefName string = HierBranchPrefix + "/tickets"

func (app *AppData) TicketBranch() (*git.Reference, error) {
	app.Repo.EnsureLog(ticketRefName)
	ticketBranch, err := app.Repo.LookupReference(ticketRefName)
	if err != nil {
		// Every ticket repo must *always* start with the same commit.
		tree_oid, err := app.emptyTree()
		tree, err := app.Repo.LookupTree(tree_oid)
		// Create a commit...
		author := &git.Signature{
			Name: "Hier internals",
			Email: "hier-internals@hier.io",
			When: time.Unix(0,0).In(time.UTC),
		}
		targetOid, _ := git.NewOid("1c93808bd6f8f577e3c78cdf82e8335e1d64976b")
		commit_oid, err := app.Repo.CreateCommit(ticketRefName, author, author, "Common base commit for Hier ticketing", tree)
		if !commit_oid.Equal(targetOid) {
			panic("Generated inconsistent ticket base. This is a bug in Hier")
		}
		fmt.Print(commit_oid)
		if err != nil {
			panic(err)
		}
		ticketBranch, err = app.Repo.LookupReference(ticketRefName)
		if err != nil {
			panic(err)
		}
	}
	return ticketBranch, nil
}

func (app *AppData) Tickets() []Ticket {
	_, err := app.TicketBranch()
	if err != nil {
		// No tickets
		return nil
	}
	
	return []Ticket{}
}
