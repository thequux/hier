// +build -debug

package webui

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _templates_new_ticket_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x53\x4d\x4b\x03\x31\x10\x3d\xdb\x5f\x31\x86\x1e\xa5\xc1\x6b\xc9\xe6\x52\x0b\x7a\xd0\x16\xdb\x8b\x20\x48\xb6\x3b\x76\x17\xf7\xa3\x6c\xb2\x6a\x59\xf6\xbf\x3b\x93\xa4\xa5\x42\x85\x1e\xbc\x6c\x26\x99\xf7\xe6\xcd\xcb\x4e\xd4\xf5\xdd\x62\xb6\x7e\x59\xce\x21\x77\x55\xa9\x47\x2a\x2c\x00\x2a\x47\x93\x71\x40\xa1\x2b\x5c\x89\xfa\xfe\x61\xfe\x0c\xd3\x29\x3c\xe1\x17\xb8\x62\xf3\x81\x4e\xc9\x90\x61\xb8\x3c\xe0\x55\xda\x64\xfb\x48\xcc\x6f\xf5\xac\x45\xe3\x10\x0c\xd4\x27\x34\x3a\x0f\x80\xf7\xa6\xad\xc0\x6c\x5c\xd1\xd4\x89\x20\xc4\x5b\x40\x08\xa8\xd0\xe5\x4d\x96\x88\xe5\x62\xb5\x16\x01\xcc\x8d\x98\x94\xe5\xae\x94\x6b\xe9\xcb\x07\xb9\x56\xa5\x49\xb1\x04\xaa\x94\x08\xdf\x8e\xd0\x6b\x5e\x94\xf4\x09\x4d\x4d\xe6\x11\x9c\x69\x55\xd4\xbb\xce\x81\xdb\xef\x90\xd0\xf8\x4d\x4a\xb5\xa9\xf0\xc8\x24\x70\xc6\xf5\xa5\x17\xf8\x53\x86\xe8\xa4\x42\xdf\xf3\x22\x16\x4b\xdc\xb8\x43\x65\x0f\x1e\xf5\x7d\x6b\xea\x2d\xc2\x84\x69\x76\x18\x5e\xa3\xa7\xe8\xac\xd9\xf1\x1d\xc0\xa7\x29\x3b\xe2\xf4\xfd\x64\x18\x84\xf6\x8b\x92\x21\xc7\x25\xb0\xce\x98\xc9\x3a\x32\x88\x5c\xd8\xb1\x75\xc6\x75\x56\xe8\x95\x5f\x2f\xe8\xfa\x40\x38\xf6\x3d\xe6\xf3\x1b\x18\xdb\x2e\x0d\x39\x98\x26\x30\x09\xf5\x2e\xf0\xe3\xe9\xc1\x53\x0c\xff\xc5\x57\x85\xd6\x9a\x2d\xdd\xef\x63\x08\xce\x3b\xe3\x1f\x6d\x68\x0e\xa3\xb7\x23\x89\x50\x31\xf3\x5b\x2e\x7a\x90\x71\xdc\xe2\xf6\x74\x74\xe8\x16\xaa\x82\x86\x27\xfa\x5b\x85\x6d\x1c\x6a\xc9\x53\x1d\x1e\x45\x78\x0b\x34\xf0\xfe\x55\xfd\x04\x00\x00\xff\xff\x70\x04\x48\x3e\x6d\x03\x00\x00")

func templates_new_ticket_html_bytes() ([]byte, error) {
	return bindata_read(
		_templates_new_ticket_html,
		"templates/new_ticket.html",
	)
}

func templates_new_ticket_html() (*asset, error) {
	bytes, err := templates_new_ticket_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "templates/new_ticket.html", size: 877, mode: os.FileMode(420), modTime: time.Unix(1420829909, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
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
	"templates/new_ticket.html": templates_new_ticket_html,
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
	"templates": &_bintree_t{nil, map[string]*_bintree_t{
		"new_ticket.html": &_bintree_t{templates_new_ticket_html, map[string]*_bintree_t{
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

