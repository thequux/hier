package webui

// The go-bindata utility is available via
// $ go get -u github.com/jteeuwen/go-bindata/...

//go:generate -command bindata go-bindata -pkg webui -ignore .*~ -ignore .*/\..*
//go:generate bindata -tags debug -debug -o assets-debug.gen.go static templates
//go:generate bindata -tags -debug -o assets-release.gen.go static templates
