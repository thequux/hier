// +build debug

package webui

import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
	"path"
	"path/filepath"
)

// bindata_read reads the given file from disk. It returns an error on failure.
func bindata_read(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

// static_css_ie_css reads file data from disk. It returns an error on failure.
func static_css_ie_css() (*asset, error) {
	path := "/home/thequux/Projects/go/src/github.com/thequux/hier/webui/static/css/ie.css"
	name := "static/css/ie.css"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_css_print_css reads file data from disk. It returns an error on failure.
func static_css_print_css() (*asset, error) {
	path := "/home/thequux/Projects/go/src/github.com/thequux/hier/webui/static/css/print.css"
	name := "static/css/print.css"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_css_screen_css reads file data from disk. It returns an error on failure.
func static_css_screen_css() (*asset, error) {
	path := "/home/thequux/Projects/go/src/github.com/thequux/hier/webui/static/css/screen.css"
	name := "static/css/screen.css"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_hier_css reads file data from disk. It returns an error on failure.
func static_hier_css() (*asset, error) {
	path := "/home/thequux/Projects/go/src/github.com/thequux/hier/webui/static/hier.css"
	name := "static/hier.css"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_images_favicon_16_png reads file data from disk. It returns an error on failure.
func static_images_favicon_16_png() (*asset, error) {
	path := "/home/thequux/Projects/go/src/github.com/thequux/hier/webui/static/images/favicon-16.png"
	name := "static/images/favicon-16.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_images_favicon_32_png reads file data from disk. It returns an error on failure.
func static_images_favicon_32_png() (*asset, error) {
	path := "/home/thequux/Projects/go/src/github.com/thequux/hier/webui/static/images/favicon-32.png"
	name := "static/images/favicon-32.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_images_logo_large_png reads file data from disk. It returns an error on failure.
func static_images_logo_large_png() (*asset, error) {
	path := "/home/thequux/Projects/go/src/github.com/thequux/hier/webui/static/images/logo-large.png"
	name := "static/images/logo-large.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_images_logo_small_png reads file data from disk. It returns an error on failure.
func static_images_logo_small_png() (*asset, error) {
	path := "/home/thequux/Projects/go/src/github.com/thequux/hier/webui/static/images/logo-small.png"
	name := "static/images/logo-small.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_js_hier_js reads file data from disk. It returns an error on failure.
func static_js_hier_js() (*asset, error) {
	path := "/home/thequux/Projects/go/src/github.com/thequux/hier/webui/static/js/hier.js"
	name := "static/js/hier.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templates_footer_html reads file data from disk. It returns an error on failure.
func templates_footer_html() (*asset, error) {
	path := "/home/thequux/Projects/go/src/github.com/thequux/hier/webui/templates/footer.html"
	name := "templates/footer.html"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templates_header_html reads file data from disk. It returns an error on failure.
func templates_header_html() (*asset, error) {
	path := "/home/thequux/Projects/go/src/github.com/thequux/hier/webui/templates/header.html"
	name := "templates/header.html"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templates_new_ticket_html reads file data from disk. It returns an error on failure.
func templates_new_ticket_html() (*asset, error) {
	path := "/home/thequux/Projects/go/src/github.com/thequux/hier/webui/templates/new_ticket.html"
	name := "templates/new_ticket.html"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templates_ticket_list_html reads file data from disk. It returns an error on failure.
func templates_ticket_list_html() (*asset, error) {
	path := "/home/thequux/Projects/go/src/github.com/thequux/hier/webui/templates/ticket_list.html"
	name := "templates/ticket_list.html"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templates_view_ticket_html reads file data from disk. It returns an error on failure.
func templates_view_ticket_html() (*asset, error) {
	path := "/home/thequux/Projects/go/src/github.com/thequux/hier/webui/templates/view_ticket.html"
	name := "templates/view_ticket.html"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"static/css/ie.css": static_css_ie_css,
	"static/css/print.css": static_css_print_css,
	"static/css/screen.css": static_css_screen_css,
	"static/hier.css": static_hier_css,
	"static/images/favicon-16.png": static_images_favicon_16_png,
	"static/images/favicon-32.png": static_images_favicon_32_png,
	"static/images/logo-large.png": static_images_logo_large_png,
	"static/images/logo-small.png": static_images_logo_small_png,
	"static/js/hier.js": static_js_hier_js,
	"templates/footer.html": templates_footer_html,
	"templates/header.html": templates_header_html,
	"templates/new_ticket.html": templates_new_ticket_html,
	"templates/ticket_list.html": templates_ticket_list_html,
	"templates/view_ticket.html": templates_view_ticket_html,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"static": &_bintree_t{nil, map[string]*_bintree_t{
		"css": &_bintree_t{nil, map[string]*_bintree_t{
			"ie.css": &_bintree_t{static_css_ie_css, map[string]*_bintree_t{
			}},
			"print.css": &_bintree_t{static_css_print_css, map[string]*_bintree_t{
			}},
			"screen.css": &_bintree_t{static_css_screen_css, map[string]*_bintree_t{
			}},
		}},
		"hier.css": &_bintree_t{static_hier_css, map[string]*_bintree_t{
		}},
		"images": &_bintree_t{nil, map[string]*_bintree_t{
			"favicon-16.png": &_bintree_t{static_images_favicon_16_png, map[string]*_bintree_t{
			}},
			"favicon-32.png": &_bintree_t{static_images_favicon_32_png, map[string]*_bintree_t{
			}},
			"logo-large.png": &_bintree_t{static_images_logo_large_png, map[string]*_bintree_t{
			}},
			"logo-small.png": &_bintree_t{static_images_logo_small_png, map[string]*_bintree_t{
			}},
		}},
		"js": &_bintree_t{nil, map[string]*_bintree_t{
			"hier.js": &_bintree_t{static_js_hier_js, map[string]*_bintree_t{
			}},
		}},
	}},
	"templates": &_bintree_t{nil, map[string]*_bintree_t{
		"footer.html": &_bintree_t{templates_footer_html, map[string]*_bintree_t{
		}},
		"header.html": &_bintree_t{templates_header_html, map[string]*_bintree_t{
		}},
		"new_ticket.html": &_bintree_t{templates_new_ticket_html, map[string]*_bintree_t{
		}},
		"ticket_list.html": &_bintree_t{templates_ticket_list_html, map[string]*_bintree_t{
		}},
		"view_ticket.html": &_bintree_t{templates_view_ticket_html, map[string]*_bintree_t{
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
        if err != nil {
                return err
        }
        err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
        if err != nil {
                return err
        }
        err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
        if err != nil {
                return err
        }
        return nil
}

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

