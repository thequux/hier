package webui

// This is a bit of an abuse of "go generate", but having the entire
// asset pipeline excecutable with one command outweighs the
// (practically non-existant) disadvantages.

//go:generate compass compile

//go:generate -command bindata go-bindata -pkg webui -ignore .*~ -ignore .*/\..*
//go:generate bindata -tags debug -debug -o assets-debug.gen.go static/... templates/...
//go:generate bindata -tags !debug -o assets-release.gen.go static/... templates/...
