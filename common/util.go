package common

import 	"github.com/libgit2/git2go"

func (app *AppData) emptyTree() (*git.Oid, error) {
	tb, err := app.Repo.TreeBuilder()
	if err != nil {
		return nil, err
	}
	oid, err := tb.Write()
	if err != nil {
		return nil, err
	}
	return oid, nil
}
