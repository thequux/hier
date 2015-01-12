package common

import (
	"container/heap"
	"crypto/sha1"
	"fmt"
	"github.com/libgit2/git2go"
	"github.com/thequux/hier/data"
	"github.com/gogo/protobuf/proto"
	"encoding/hex"
	"time"

)

// Todo: include a filesystem-like abstraction for modifying git trees.

type TicketId [20]byte
func ParseTicketId(id string) (TicketId, error) {
	var ticketId TicketId
	if l, err := hex.Decode(ticketId[:], []byte(id)); err != nil {
		return ticketId, err
	} else if l != 20 {
		return ticketId, fmt.Errorf("TicketID is the wrong length")
	}
	return ticketId, nil
}
func TicketIdFromRaw(src []byte) TicketId {
	var ret TicketId
	copy(ret[:], src[0:20])
	return ret
}

func (id TicketId) String() string {
	return hex.EncodeToString(id[:])
}



type Ticket struct {
	Hash TicketId
	Title string
	Type string
	Status string
	Resolution string
	// The Artifacts array is guaranteed to be sorted in some
	// topological order.
	Artifacts []*TicketArtifact
	heads []*TicketArtifact
	app *AppData
	no_copy bool
}

type TicketArtifact struct {
	Title *string
	Type *string
	Status string
	Resolution string
	AuthorName string
	AuthorEmail string
	Message string
	Date time.Time
	After []*TicketArtifact
	artifact *data.TicketArtifact
	hash *TicketId
}

type TicketConfig struct {
	Types []string
	Statuses map[string][]string
}

const ticketBranchName string = HierBranchPrefix + "/tickets"
const ticketRefName string = "refs/heads/" + ticketBranchName

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

func (app *AppData) TicketBranch() *git.Branch {
	app.Repo.EnsureLog(ticketRefName)
	ticketBranch, err := app.Repo.LookupBranch(ticketBranchName, git.BranchLocal)
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
		commit_oid, err := app.Repo.CreateCommit(ticketRefName, author, author, "Common base commit for Hier ticketing", tree)
		targetOid, _ := git.NewOid("1c93808bd6f8f577e3c78cdf82e8335e1d64976b")
		if !commit_oid.Equal(targetOid) {
			panic("Generated inconsistent ticket base. This is a bug in Hier")
		}
		if err != nil {
			panic(err)
		}
		ticketBranch, err = app.Repo.LookupBranch(ticketBranchName, git.BranchLocal)
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

	config := app.TicketConfig()
	// As this is the first artifact, we require title to be set.
	if artifact.Title == nil {panic("Missing title!")}
	if artifact.Type == nil {panic("Missing type!")}
	if artifact.Status == nil {panic("Missing status!")}
	if resolutions, ok := config.Statuses[*artifact.Status]; ok {
		if len(resolutions) != 0 {
			if artifact.Resolution == nil {panic("Missing resolution")}
			res_found := false
			for _, res := range resolutions {
				if res == *artifact.Resolution {
					res_found = true
					break
				}
			}
			if !res_found {
				panic("Invalid resolution")
			}
		} else {
			if artifact.Resolution != nil {
				panic("Unexpected resolution")
			}
		}
	} else {
		panic("Unknown status")
	}
	if artifact.Author == nil {
		config, _ := app.Repo.Config()
		name, _ := config.LookupString("user.name")
		email, _ := config.LookupString("user.email")
		artifact.Author = &data.Author{
			Name: &name,
			Email: &email,
		}
	}

	marshaled, err := proto.Marshal(artifact)
	if err != nil {panic(err)}

	odb, err := app.Repo.Odb()
	if err != nil {panic(err)}
	oid, err := odb.Write(marshaled, git.ObjectBlob)
	hash := TicketId(sha1.Sum(marshaled))
	if err != nil {
		return nil, err
	}
	oldCommitObj, _ := app.TicketBranch().Peel(git.ObjectCommit)
	oldCommit := (oldCommitObj).(*git.Commit)
	oldTree, _ := oldCommit.Tree()

	// create the new ticket tree
	tb, err := app.Repo.TreeBuilder()
	if err != nil {panic(err)}
	if err := tb.Insert(hash.String(), oid, FilemodeNormal); err != nil {
		panic(err)
	}
	oid, err = tb.Write()
	if err != nil {panic(err)}

	// Add the new tree to the branch
	tb, err = app.Repo.TreeBuilderFromTree(oldTree)
	if err != nil {panic(err)}
	err = tb.Insert(hash.String(), oid, FilemodeDirectory)
	if err != nil {
		panic(err)
	}
	tree_oid, err := tb.Write()
	if err != nil {panic(err)}
	tree, err := app.Repo.LookupTree(tree_oid)

	// Commit
	author := &git.Signature{
		Name: *artifact.Author.Name,
		Email: *artifact.Author.Email,
		When: time.Now(),
	}
	if err != nil {panic(err)}
	// TODO: Come up with a useful commit message
	
	_, err = app.Repo.CreateCommit(ticketRefName, author, author, "", tree, oldCommit)
	if err != nil {
		panic(err)
	}

	ticketTree, _ := app.Repo.LookupTree(oid)
	return app.parseTicket(hash, ticketTree)
}

// This doesn't completely parse the artifact; it still needs to have
// its predecessors stitched in.
func (app *AppData) parseTicketArtifact(id TicketId, blob []byte) (*TicketArtifact, error) {
	var artifact data.TicketArtifact

	if err := proto.Unmarshal(blob, &artifact); err != nil {
		return nil, err
	}
	when, err := time.Parse(time.RFC3339, *artifact.Date)
	if err != nil {
		when = time.Unix(0,0)
	}
	
	ret := TicketArtifact{
		Title: artifact.Title,
		Type: artifact.Type,
		AuthorName: *artifact.Author.Name,
		AuthorEmail: *artifact.Author.Email,
		Date: when,
		Message: *artifact.Message,
		artifact: &artifact,
		hash: &id,
	}

	if artifact.Status != nil {
		ret.Status = *artifact.Status
	}
	if artifact.Resolution != nil {
		ret.Resolution = *artifact.Resolution
	}
	return &ret, nil
}

func (app *AppData) parseTicket(id TicketId, tree *git.Tree) (*Ticket, error) {
	// Parse all the artifacts.

	// TODO: this structure isn't particularly efficient; optimize
	// it.
	artifacts := map[TicketId]*TicketArtifact{}
	ticket := &Ticket{
		Hash: id,
		app: app,
	}
	for i := uint64(0); i < tree.EntryCount(); i++ {
		
		entry := tree.EntryByIndex(i)
		blob, err := app.Repo.LookupBlob(entry.Id)
		if err != nil {
			panic("Malformed repo")
		}
		aid, err := ParseTicketId(entry.Name)
		if err != nil {
			panic(err)
		}
		artifact, err := app.parseTicketArtifact(aid, blob.Contents())
		if err != nil {
			panic(err)
		}
		artifacts[aid] = artifact
		ticket.Artifacts = append(ticket.Artifacts, artifact)
	}
	for _, artifact := range artifacts {
		for _, after := range artifact.artifact.After {
			artifact.After = append(artifact.After, artifacts[TicketIdFromRaw(after)])
		}
	}
	ticket.sortArtifacts()
	rootArtifact := artifacts[id]
	
	if rootArtifact == nil {return nil, fmt.Errorf("Ticket is missing root artifact")}
	ticket.Hash = id
	ticket.Title = *rootArtifact.Title
	ticket.Type = *rootArtifact.Type
	
	for _, artifact := range ticket.Artifacts[1:] {
		// [1:] to skip the root artifact
		status := artifact.After[0].Status
		resolution := artifact.After[0].Resolution
		for _, pred := range artifact.After[1:] {
			status, resolution = app.MergeStatus(status, resolution, pred.Status, pred.Resolution)
		}
		if artifact.Status == "" {
			artifact.Status = status
			artifact.Resolution = resolution
		}
	}
	status := ticket.heads[0].Status
	resolution := ticket.heads[0].Resolution
	for _, pred := range ticket.heads[1:] {
		status, resolution = app.MergeStatus(status, resolution, pred.Status, pred.Resolution)
	}
	ticket.Status = status
	ticket.Resolution = resolution
	return ticket, nil
}


func (app *AppData) Tickets() []*Ticket {
	ticketRef := app.TicketBranch()
	ticketTreeObj, _ := ticketRef.Peel(git.ObjectTree)
	ticketTree := ticketTreeObj.(*git.Tree)
	result := []*Ticket{}
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
		ticketId, err := ParseTicketId(entry.Name)
		if err != nil {continue}

		ticket, err := app.parseTicket(ticketId, ticketTree)
		if err != nil {continue}

		result = append(result, ticket)
	}
	return result
}

func (app *AppData) GetTicket(id TicketId) (*Ticket, error) {
	ticketRootObj, _ := app.TicketBranch().Peel(git.ObjectTree)
	ticketRoot := ticketRootObj.(*git.Tree)
	entry := ticketRoot.EntryByName(id.String())
	if entry == nil {
		return nil, fmt.Errorf("No such ticket")
	} else if entry.Type != git.ObjectTree {
		return nil, fmt.Errorf("Ticket has wrong type")
	}
	tree, err := app.Repo.LookupTree(entry.Id)
	if err != nil {
		return nil, err
	}
	return app.parseTicket(id, tree)
}


// Sort artifacts in semantic (causal) order.
func (ticket *Ticket) sortArtifacts() {
	heads := []*TicketArtifact{}
	available := &artifactHeap{}
	inserted := map[*TicketArtifact]bool{}
	tails := map[*TicketArtifact][]*TicketArtifact{}
	sortedArtifacts := []*TicketArtifact{}
 
	for _, artifact := range ticket.Artifacts {
		if len(artifact.After) == 0 {
			heap.Push(available, artifact)
		}
		for _, next := range artifact.After {
			tails[next] = append(tails[next], artifact)
		}
	}

	// TODO: verify that there is only one root artifact
	
	for len(*available) != 0 {
		next := heap.Pop(available).(*TicketArtifact)
		sortedArtifacts = append(sortedArtifacts, next)
		inserted[next] = true
		if len(tails[next]) == 0 {
			heads = append(heads, next)
		}
		for _, potential := range tails[next] {
			insert := true
			for _, prereq := range potential.After {
				if !inserted[prereq] {
					insert = false
					break
				}
			}
			if insert {
				heap.Push(available, potential)
			}
		}
	}

	ticket.Artifacts = sortedArtifacts
	ticket.heads = heads
}
	

func (ticket *Ticket) NewComment(artifact *TicketArtifact) error {
	repo := ticket.app.Repo
	// Fill in missing fields...
	artifact.After = ticket.heads
	config, _ := repo.Config()
	artifact.AuthorName, _ = config.LookupString("user.name")
	artifact.AuthorEmail, _ = config.LookupString("user.email")
	artifact.Date = time.Now()
	
	// Create an appropriate raw artifact...
	when := artifact.Date.Format(time.RFC3339)
	artifact.artifact = &data.TicketArtifact{
		Author: &data.Author{
			Name: &artifact.AuthorName,
			Email: &artifact.AuthorEmail,
		},
		Message: &artifact.Message,
		Date: &when,
	}
	if artifact.Status != "" {
		artifact.artifact.Status = &artifact.Status
		if artifact.Resolution != "" {
			artifact.artifact.Resolution = &artifact.Resolution
		}
	}

	for _, head := range artifact.After {
		artifact.artifact.After = append(artifact.artifact.After, head.hash[:])
	}
	ticket.Artifacts = append(ticket.Artifacts, artifact)
	ticket.heads = nil

	// find branch to be modifying
	ticketBranch := ticket.app.TicketBranch()
	var oldCommit *git.Commit
	if oldCommitObj, err := ticketBranch.Peel(git.ObjectCommit); err != nil {
		return err
	} else {
		oldCommit = oldCommitObj.(*git.Commit)
	}
	oldTree, _ := oldCommit.Tree()
	oldTreeEntry := oldTree.EntryByName(ticket.Hash.String())
	if oldTreeEntry == nil {
		panic("Ticket missing")
	}
	var ticketTB *git.TreeBuilder
	if ticketTree, err := repo.LookupTree(oldTreeEntry.Id); err != nil {
		return err // TODO: Handle this better
	} else {
		// TODO: handle error
		ticketTB, _ = repo.TreeBuilderFromTree(ticketTree)
	}

	// Write new artifact...
	blobContent, err := proto.Marshal(artifact.artifact)
	hash := TicketId(sha1.Sum(blobContent))
	artifact.hash = &hash
	if err != nil {return err}
	blobOid, err := repo.CreateBlobFromBuffer(blobContent)
	if err != nil {return err}
	
	// Put it in the ticket tree
	ticketTB.Insert(artifact.hash.String(), blobOid, FilemodeNormal)
	ticketTreeOid, err := ticketTB.Write()
	if err != nil {return err}

	// update the root tree
	rootTB, err := repo.TreeBuilderFromTree(oldTree)
	rootTB.Insert(ticket.Hash.String(), ticketTreeOid, FilemodeDirectory)
	newTreeId, err := rootTB.Write()
	if err != nil {return err}

	if !newTreeId.Equal(oldTree.Id()) {
		// the trees might be equal in case of a double post
		author := &git.Signature{
			Name: artifact.AuthorName,
			Email: artifact.AuthorEmail,
			When: time.Now(),
		}
		newTree, err := repo.LookupTree(newTreeId)
		if err != nil {return err}
		_, err = repo.CreateCommit(ticketRefName, author, author, fmt.Sprintf("Added comment to %s", ticket.Hash), newTree, oldCommit)
		if err != nil {return err}
	}
	return nil
}
