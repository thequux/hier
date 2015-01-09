package common

import (
	"crypto/sha1"
	"fmt"
	"github.com/libgit2/git2go"
	"github.com/thequux/hier/data"
	"github.com/gogo/protobuf/proto"
	"encoding/hex"
	"time"
)

type TicketId [20]byte

type Ticket struct {
	Hash TicketId
	Title string
	Type string
	Status string
	artifacts []*data.TicketArtifact
}

type TicketConfig struct {
	Types []string
	Statuses map[string][]string
}

const ticketRefName string = HierBranchPrefix + "/tickets"

func (app *AppData) TicketConfig() TicketConfig {
	return TicketConfig{
		Types: []string{
			"Code Defect",
			"Build Problem",
			"Documentation",
			"Feature Request",
			"Technical Debt",
		},
		Statuses: map[string][]string{
			"New": nil,
			"Verified": nil,
			"In Progress": nil,
			"Closed": []string{
				"Fixed",
				"Wontfix",
				"Duplicate",
				"Invalid",
			},
		},
	}
}

func (app *AppData) TicketBranch() *git.Reference {
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
	return ticketBranch
}

func (app *AppData) NewTicket(artifact *data.TicketArtifact) (*Ticket, error) {
	if (artifact == nil) {
		panic("Can't file a ticket with no details.")
	}
	if len(artifact.After) != 0 {
		panic("Can't file a new ticket with history")
	}

	// As this is the first artifact, we require title to be set.
	if artifact.Title == nil {panic("Missing title!")}
	if artifact.Type == nil {panic("Missing type!")}
	if artifact.Status == nil {panic("Missing status!")}

	marshaled, err := proto.Marshal(artifact)
	if err != nil {panic(err)}

	odb, err := app.Repo.Odb()
	if err != nil {panic(err)}
	oid, err := odb.Write(marshaled, git.ObjectBlob)
	hash := sha1.Sum(marshaled)
	if err != nil {
		return nil, err
	}
	oldCommitObj, _ := app.TicketBranch().Peel(git.ObjectCommit)
	oldCommit := (oldCommitObj).(*git.Commit)
	oldTree, _ := oldCommit.Tree()
	tb, err := app.Repo.TreeBuilderFromTree(oldTree)
	if err != nil {panic(err)}

	author := &git.Signature{
		Name: *artifact.Author.Name,
		Email: *artifact.Author.Email,
		When: time.Now(),
	}
	tb.Insert(fmt.Sprintf("%x/%x", hash, hash), oid, 0444)
	tree_oid, err := tb.Write()
	if err != nil {panic(err)}
	tree, err := app.Repo.LookupTree(tree_oid)
	if err != nil {panic(err)}
	// TODO: Come up with a useful commit message
	
	_, err = app.Repo.CreateCommit(ticketRefName, author, author, "", tree, oldCommit)
	if err != nil {
		panic(err)
	}
	return &Ticket{
		Hash: hash,
		Title: *artifact.Title,
		Type: *artifact.Type,
		Status: *artifact.Status,
		artifacts: []*data.TicketArtifact{artifact},
	}, nil
}

func (app *AppData) parseTicket(id TicketId, tree *git.Tree) Ticket {
	// For now, we just look at the root artifact.  Later, we'll
	// parse all the artifacts on each ticket to get a overall
	// status.
	rootArtifact := tree.EntryByName(fmt.Sprintf("%x", id))
	if rootArtifact == nil {panic("Malformed ticket; missing root artifact")}
	if rootArtifact.Type != git.ObjectBlob {
		panic("Malformed ticket: contains non-blob root artifact")
	}
	blob, _ := app.Repo.LookupBlob(rootArtifact.Id)
	var artifact data.TicketArtifact
	if err := proto.Unmarshal(blob.Contents(), &artifact); err != nil {
		panic("Malformed artifact: " + err.Error())
	}
	return Ticket{
		Hash: id,
		Title: *artifact.Title,
		Type: *artifact.Type,
		Status: *artifact.Status,
		artifacts: []*data.TicketArtifact{&artifact},
	}
}

func (app *AppData) Tickets() []Ticket {
	ticketRef := app.TicketBranch()
	ticketTreeObj, _ := ticketRef.Peel(git.ObjectTree)
	ticketTree := ticketTreeObj.(*git.Tree)
	result := []Ticket{}
	for i := uint64(0); i < ticketTree.EntryCount(); i++ {
		entry := ticketTree.EntryByIndex(i)
		if entry == nil {panic("Should not happen")}
		if len(entry.Name) != 40 {
			// Not a hex repr of a SHA1 digest; must not
			// be relevant.
			continue
		}
		if entry.Type != git.ObjectTree {
			// Anything that's not a tree is irrelevant
			continue
		}
		ticketTree, _ := app.Repo.LookupTree(entry.Id)
		var ticketId TicketId
		if l, err := hex.Decode(ticketId[:], []byte(entry.Name)); err != nil {
			continue
		} else if l != 20 {
			continue
		}

		result = append(result, app.parseTicket(ticketId, ticketTree))
	}
	return []Ticket{}
}
