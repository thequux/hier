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
	// The Artifacts array is guaranteed to be sorted in some
	// topological order.
	Artifacts []*data.TicketArtifact
	no_copy bool
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
	if err := tb.Insert(hash.String(), oid, 0100644); err != nil {
		panic(err)
	}
	oid, err = tb.Write()
	if err != nil {panic(err)}

	// Add the new tree to the branch
	tb, err = app.Repo.TreeBuilderFromTree(oldTree)
	if err != nil {panic(err)}
	err = tb.Insert(hash.String(), oid, 0040000)
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
	return &Ticket{
		Hash: hash,
		Title: *artifact.Title,
		Type: *artifact.Type,
		Status: *artifact.Status,
		Artifacts: []*data.TicketArtifact{artifact},
	}, nil
}

func (app *AppData) parseTicket(id TicketId, tree *git.Tree) (*Ticket, error) {
	// For now, we just look at the root artifact.  Later, we'll
	// parse all the artifacts on each ticket to get a overall
	// status.
	rootArtifact := tree.EntryByName(id.String())
	if rootArtifact == nil {return nil, fmt.Errorf("Ticket is missing root artifact")}
	if rootArtifact.Type != git.ObjectBlob {
		return nil, fmt.Errorf("Malformed ticket: root artifact is not a blob")
	}
	blob, err := app.Repo.LookupBlob(rootArtifact.Id)
	if err != nil {
		return nil, err
	}
	var artifact data.TicketArtifact
	if err := proto.Unmarshal(blob.Contents(), &artifact); err != nil {
		return nil, err
	}
	return &Ticket{
		Hash: id,
		Title: *artifact.Title,
		Type: *artifact.Type,
		Status: *artifact.Status,
		Artifacts: []*data.TicketArtifact{&artifact},
	}, nil
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
