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

var _static_hier_css_ = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func static_hier_css__bytes() ([]byte, error) {
	return bindata_read(
		_static_hier_css_,
		"static/#hier.css#",
	)
}

func static_hier_css_() (*asset, error) {
	bytes, err := static_hier_css__bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "static/#hier.css#", size: 0, mode: os.FileMode(420), modTime: time.Unix(1420836579, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _static_css_ie_css = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x8d\x31\x4f\xf3\x30\x18\x06\xf7\xfe\x8a\xe7\xf3\x18\x7d\x89\x77\xd4\x64\x41\x48\x54\x82\x0d\xc4\x50\x75\x88\x9c\xc7\xe4\x05\xc7\x0e\x7e\xdd\x96\xfe\x7b\x92\x56\x42\x8c\xa7\xd3\xe9\x6c\x85\x37\x06\x97\x26\xa2\x24\xdc\xa7\x69\xee\x55\x1b\xbc\xea\xc2\xa3\x28\xbc\x84\xab\x39\x67\x29\xc4\xee\x01\x3a\xd3\x89\x17\x87\x74\x62\xce\x32\x10\x5a\x2e\x81\xda\x6c\x50\x61\x37\xcd\x29\x97\x3f\xe5\x51\x25\xbe\x2f\x4c\xf8\x14\x42\x3a\xaf\xf4\xf8\xf2\xfc\x84\x94\xc1\xaf\xa3\x9c\xfa\xc0\x58\xee\xd6\x76\xfb\xaf\xae\xf7\xe2\x97\xc7\xa1\x5b\x19\xd8\x06\x89\x9f\x18\x33\x7d\x6b\xec\xed\x32\x92\x45\xad\xb0\x71\xaa\x06\x13\x07\xe9\x5b\xa3\x2e\x93\xf1\x3f\xe6\x9c\x3e\xe8\x8a\xa4\x68\x90\x19\x16\xf1\xdb\x18\x94\xcb\xcc\xd6\x14\x7e\x17\x7b\x6d\x6d\x77\x9b\xee\x19\x07\xf1\x87\xba\xee\x50\xd9\xcd\x4f\x00\x00\x00\xff\xff\x53\x7c\x82\x7f\x0f\x01\x00\x00")

func static_css_ie_css_bytes() ([]byte, error) {
	return bindata_read(
		_static_css_ie_css,
		"static/css/ie.css",
	)
}

func static_css_ie_css() (*asset, error) {
	bytes, err := static_css_ie_css_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "static/css/ie.css", size: 271, mode: os.FileMode(420), modTime: time.Unix(1420838733, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _static_css_print_css = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x8d\x3d\x4f\xc6\x30\x0c\x84\xf7\xf7\x57\x9c\x32\x76\x68\x76\x44\x58\x58\x40\x82\x0d\xc4\x5c\xb5\x0e\xb5\x70\x3e\x88\x5d\xa0\xff\x9e\x36\x0c\xbc\xe3\x63\x3f\x77\xe7\x07\xbc\x91\xcc\x25\x11\xac\xe0\xbe\xa4\x3a\xa9\x8e\x78\xd5\x83\x57\x56\x44\x96\xfe\x59\x28\x72\x26\xd4\xc6\xd9\xa0\xb6\x0b\xe9\x78\xc1\x80\xc7\x54\x4b\xb3\x2b\x77\x53\xce\xef\x07\x13\x62\x11\x29\xdf\x27\x3d\xbc\x3c\x3f\xa1\x34\xd0\xe7\xc6\x5f\x93\x50\xb6\x9b\x33\x7b\x2b\x9c\x3f\xb0\x36\x8a\xc1\xf9\xbf\xce\x95\xc8\xd4\xf7\x95\x71\x56\x75\x48\xb4\xf0\x14\x5c\xbf\x38\x34\x92\xe0\xfe\x4d\x07\xdb\x2b\x05\x67\xf4\x63\xbe\xeb\xfe\x0e\x83\xbf\xfc\x06\x00\x00\xff\xff\xe3\x03\x4b\x88\xd5\x00\x00\x00")

func static_css_print_css_bytes() ([]byte, error) {
	return bindata_read(
		_static_css_print_css,
		"static/css/print.css",
	)
}

func static_css_print_css() (*asset, error) {
	bytes, err := static_css_print_css_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "static/css/print.css", size: 213, mode: os.FileMode(420), modTime: time.Unix(1420838733, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _static_css_screen_css = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x56\x5d\x6f\xec\x34\x10\x7d\xcf\xaf\x18\xad\x84\x04\x55\x76\xb3\xdb\x5e\xae\xd0\x96\xf2\xc2\x0b\x48\xf0\x86\xc4\x23\x72\xe2\x49\x32\xac\x63\x67\x6d\x67\xcb\x82\xf8\xef\x1c\x3b\xfb\xd5\xd2\xfb\xda\x4a\x75\xd6\x1f\xe3\x99\x73\xec\x33\x9e\x56\x77\xf4\x3b\x9b\xc6\x0d\x4c\xd1\xd1\x8f\x6e\x18\x55\x08\xab\x82\xee\xe8\x67\x4b\xb1\x97\x40\xad\x18\xa6\xa3\x9b\x28\xf4\x6e\x32\x9a\x9e\xbd\xc4\x3c\xe1\x69\x50\x62\x29\xc4\xa3\xe1\xb0\xa2\xaf\x9d\xa7\x86\x6d\xf4\xca\xc8\xdf\x27\x03\x19\x46\xe7\x63\xf8\x26\xfb\xcb\xfd\x1b\x9f\x53\x10\xdb\x61\xcc\xd4\x3a\x63\xdc\x73\x1a\xfd\xf4\xdb\xaf\xbf\x10\x1c\xf1\x7e\x92\x83\x32\x70\xb7\x4d\x7b\xbf\x37\x62\x77\xd4\x7b\x6e\x9f\x16\xd5\x1c\xb0\x67\x8e\xa1\x0a\x8d\x67\xb6\xab\x26\x84\x05\x0d\xac\x45\x3d\x2d\xe6\xa9\x92\x46\xef\xfe\xe4\x26\x8a\xb3\x0b\xf2\x6c\xb0\x70\xd9\xb7\xa0\x78\x1c\xf9\x69\x11\xf9\xaf\x58\xe5\xbd\xd5\x0f\x74\x57\x15\x38\x0c\x04\x62\xfa\xb6\xa4\xd5\xaa\xfa\xe2\x5f\xc7\x43\xe5\xa7\xfa\x58\xdd\xaf\x36\xab\x75\x85\x61\xa8\x9a\xf9\xe4\x96\x8d\xf3\xbc\xc4\xec\x6a\xf3\x02\xe7\x69\xb9\xf2\x1c\x38\x56\x7f\x4c\x51\x8c\x44\xc1\xb1\x05\x84\x4f\xa1\xfb\x38\x98\x92\x6a\xa7\x8f\x25\x69\x39\x94\x14\x46\x05\x12\x6a\x1c\x0d\xc7\x92\x5c\x9d\xb8\x94\x24\xad\x57\x03\x97\x45\xbf\x29\xa9\xbf\x47\x7b\x40\xfb\x84\x06\xc8\xfd\x67\x90\x86\x0f\xe3\x9a\xdd\x7e\x72\x91\xd3\x19\xc0\x56\xc1\x4d\x5d\x7b\x7c\x1b\xef\xec\x71\x40\x47\x6b\x00\x09\xb0\x95\xae\xa4\x46\x92\x69\xe3\x34\x6c\x35\x03\x85\x6e\x11\x99\x61\x27\x03\x96\xc5\xc2\x70\x57\xeb\x92\xf6\x40\x85\x3f\x35\x8c\x65\x11\x06\x65\x60\x1a\xa2\x97\x1d\xe7\x5f\x67\x61\x1c\xa6\x3a\x7d\x00\x23\x02\xed\x41\xf9\xb2\xc0\xc4\x04\x2f\x65\x16\x07\x63\x42\xa7\x10\x58\xd5\x70\xe9\xd0\x9f\xd0\x8c\x94\x45\x2b\x6c\x74\x48\x6c\x5b\xe7\x11\xdd\xa8\x3a\xa1\x31\xdc\xb1\xd5\x65\x11\x55\x6d\x12\x4e\x35\xa6\x3b\x45\x80\xf9\xb0\x62\xeb\x1c\xb6\x40\x46\x0a\xfe\xa2\x4f\x5d\x34\x6c\x50\x3e\x4a\x93\xb6\xa8\x20\x3a\xef\xb4\x07\x05\x02\x9a\xa3\x12\x13\x12\xc5\x9a\x75\x8a\xdb\x4d\x38\x27\x88\xb2\xbb\x38\x4f\x4e\x81\x95\x92\xd7\xfc\xdb\x79\x07\x56\xc5\xc0\x16\x64\xac\xc2\x05\xb9\x29\x8e\x13\x22\x27\x25\x80\xf3\x2c\xb5\x44\x7e\x18\x94\x3f\x02\xaf\xe0\xa2\x90\x22\x7e\x07\x04\x93\x16\x87\xf3\x00\x0e\x47\xff\x14\x94\xa6\x3b\xb1\x5b\x5a\x3f\x62\x30\xe2\x3e\x20\xfe\xd3\xa8\x76\x1e\x21\x4f\x83\xd6\x21\x03\x70\x05\x3d\x23\xeb\xce\x13\xcb\x80\x04\xdb\xd2\x66\xbd\xfe\x2a\x4d\x1d\x38\x11\x55\x66\x89\xc4\xeb\xe0\xb3\x56\x81\x93\x8c\x1f\x8b\x7f\x8b\x8b\xa4\xef\xef\x3f\x46\xd3\x99\x6d\x42\xb0\xec\x59\xba\x1e\x64\x36\x2f\x71\x7d\x7a\x7f\x5c\xb3\xe8\x4e\xc8\x02\x8e\x33\x6d\xde\x92\x75\xaf\xcf\xec\xf3\xfb\x63\xcb\x2a\xcf\xd0\x66\x1d\xc0\x91\x31\x6a\x0c\xc0\x77\xee\x5d\x45\xb2\xc4\x23\xd1\x9c\x95\x73\x0b\xfc\xbb\xf7\x07\x7e\xcd\xcb\x9c\x7e\x99\x42\x7a\x5f\xcf\xaa\x34\xdc\x5e\x05\xfc\x7c\xd2\x82\x45\xa2\x2b\xf3\x96\x88\x07\xd1\xda\xbc\xbc\x8e\x87\xf5\xfb\xb3\xda\xdf\xbe\xa7\x99\x53\xee\x85\xab\x5a\xce\xe8\x36\xeb\x87\x0f\x80\xb7\xad\x19\x8f\x25\xde\x99\xfd\x56\xb5\xf9\xc5\xba\xc2\xbd\xac\xdd\x4c\x65\xa3\xcc\xa3\xc1\x3d\xa4\xf2\x4a\x8b\xc5\xe3\xed\xf0\x7f\x59\xf0\xf0\x01\x2f\x87\x4a\xa5\xe7\x26\x0b\xde\x80\xb5\xd9\x7c\x40\x76\xbe\x2e\x29\x97\x52\xf2\xa2\x76\x9c\xeb\xc9\xdb\x35\x24\xff\xd7\x84\xef\xb5\x92\xbc\xae\x1d\x99\xb8\x96\x30\x1a\x75\xdc\xce\xb7\x97\xa8\xff\x17\x00\x00\xff\xff\xc5\x7f\x58\xa5\xab\x09\x00\x00")

func static_css_screen_css_bytes() ([]byte, error) {
	return bindata_read(
		_static_css_screen_css,
		"static/css/screen.css",
	)
}

func static_css_screen_css() (*asset, error) {
	bytes, err := static_css_screen_css_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "static/css/screen.css", size: 2475, mode: os.FileMode(420), modTime: time.Unix(1420838733, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _static_js_hier_js = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x51\x4d\x6b\xeb\x30\x10\x3c\xc7\xbf\x42\x4f\x27\x19\x82\xfe\x80\xf0\xe1\xf1\x08\xe4\x90\x77\xcb\xad\x94\xa2\xd8\x6b\x22\xaa\xae\x8d\xb4\x0a\x0d\xc1\xff\xbd\xfa\x70\x6c\x17\x4a\x73\x90\x94\xdd\x9d\xd9\x99\xb1\xe9\x99\xa0\xfb\x08\x43\xcf\x8e\x06\x1c\x6b\x9a\x86\xf1\x80\x1d\xf4\x06\xa1\xe3\xf5\x5c\x65\x8f\x49\x55\x55\x7a\x4b\x7f\xc7\xf6\xcd\x81\x1f\x6c\x20\x33\xa0\x8f\xcd\x3e\x60\x9b\xde\xc2\x93\xa6\xe0\xf7\x6c\x6d\xd7\xec\x51\xb1\xf8\xbb\x69\xc7\xae\xa6\xcb\x5c\x18\xac\x55\xb9\xda\x0f\x8e\x89\x67\x79\x05\xa9\x79\xf4\x4f\x33\x0f\x2f\xd0\x7c\xcb\x51\x3b\x40\x3a\x58\xf8\x88\x57\xda\xb0\x4b\x36\x4a\xaf\xb5\xda\xfb\x93\xf1\x24\xdb\x01\x49\x1b\xf4\x82\xe7\x0e\xaf\xf3\x64\x5a\x7b\x71\xa0\xdf\x55\xb5\x9b\xb2\x88\x72\x2e\x0c\x6c\x5e\x5a\x57\x3b\x07\x14\x1c\xaa\xc5\x80\x0f\x97\xe2\x10\x92\xeb\x92\xc6\xfc\xff\xa5\x3c\xe4\x4d\xdb\x00\xaf\x6a\xa1\xdc\x96\x4b\xb8\x3c\xcb\x28\x62\xe3\xd9\x01\x46\x2e\x72\x01\x0a\x68\x62\x60\x3d\x14\xec\xba\xee\x27\x4c\xaf\xe3\xa0\x4a\x2a\x9f\xb9\x49\x83\x08\xee\x78\xfe\x7f\x62\x69\x51\xec\x6d\x28\x64\x0c\xfb\xa0\xdb\xab\xd8\x7c\xad\x25\x91\xe4\x0e\x6c\x44\x75\x43\x1b\x52\xaa\xb2\x8d\x19\x11\xcc\x19\x0b\x3e\x8c\x09\xc2\x6b\x55\xe6\xc1\x4a\x0f\xf4\x97\xc8\x99\x4b\x20\x10\x3c\x1b\xe4\x7b\x16\x39\xd7\x91\x2c\xe7\x0c\x9f\x14\x89\x3d\xcd\xf5\x8d\x5c\x3d\x8e\x80\xdd\xbf\xab\xb1\x9d\x00\x9b\x80\x53\xfd\x2d\x85\x5f\x83\xaa\xa6\xea\x2b\x00\x00\xff\xff\x13\x21\xd7\x26\xbe\x02\x00\x00")

func static_js_hier_js_bytes() ([]byte, error) {
	return bindata_read(
		_static_js_hier_js,
		"static/js/hier.js",
	)
}

func static_js_hier_js() (*asset, error) {
	bytes, err := static_js_hier_js_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "static/js/hier.js", size: 702, mode: os.FileMode(420), modTime: time.Unix(1420847430, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templates_new_ticket_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x54\xc1\x6e\xdb\x30\x0c\x3d\xaf\x5f\xc1\x09\x3d\x6c\x40\x6b\x61\xd7\x42\xf6\xa5\x2b\xd0\x1d\xb6\x16\x6d\x2e\x03\x06\x0c\x8a\xcd\x56\x4a\x65\x39\x90\x94\x6e\x41\x90\x7f\x1f\x69\xa9\x89\x87\x65\x40\x72\x91\x64\x91\x8f\x8f\x7c\x14\xad\xde\x7f\xbe\xbb\x9e\x7d\xbf\xbf\x01\x93\x7a\xd7\x9c\xa9\xbc\x01\x28\x83\xba\xe3\x03\x1d\x93\x4d\x0e\x9b\xdb\x2f\x37\x0f\x70\x75\x05\xdf\xf0\x17\x24\xdb\xbe\x60\x52\x32\x5b\xb2\x97\xb3\xfe\x05\x02\xba\x5a\xc4\xb4\x76\x18\x0d\x62\x12\x90\xd6\x4b\xac\x45\xc2\xdf\x49\xb6\x31\x0a\x30\x01\x9f\x6a\x21\x63\xd2\x14\x83\xaf\x64\x6c\x03\xa2\xaf\xd8\x5a\x22\xd1\x8d\x5d\xa6\x29\x74\xa1\x5f\x75\xbe\x15\x10\x43\xbb\x0f\xb0\x88\xd2\x58\x0c\xd5\x82\xc0\x4a\x66\x97\x13\xa2\x10\xbc\xd3\x49\xff\x03\x57\xf2\xad\x7c\x35\x1f\xba\x75\x89\x68\x3e\x35\xd7\x01\x75\x42\xd0\xe0\x27\x2a\xd0\x7d\x76\x78\x1a\x42\x0f\xba\x4d\x76\xf0\x14\x3c\x9b\x25\x79\x0a\xe8\x31\x99\xa1\xab\xc5\xfd\xdd\xe3\xac\x94\xc9\xc2\xea\x39\xcb\xf7\x4e\xa5\x40\x2b\x5f\x98\x46\x39\x3d\x47\x07\x14\x8a\xb2\x66\x79\x45\x33\xe3\x4d\xc9\xd1\x40\x69\x92\x53\x76\xee\x1a\x65\xfd\x72\x35\xad\x51\x80\xd7\x3d\xee\x90\xe4\xdc\x71\x7c\x39\x12\xfc\x97\x86\xe0\xc4\x42\xeb\x61\x92\x88\x0e\xdb\xf4\x16\x79\x74\x3e\xdb\x6c\x82\xf6\xcf\x08\x15\xc3\xe2\x76\xfb\xa3\xd4\x54\x2a\x1b\x96\x2c\x02\xbc\x6a\xb7\x22\xcc\x66\x53\x6d\xb7\xa2\x19\x37\x25\xb3\x8d\x43\xa0\xef\x18\xc9\x3c\x32\x93\x1c\x99\x31\x77\x7f\x45\x3d\x7b\x1c\xf7\x23\xb2\x2e\x00\x18\x7c\x6b\x38\xef\x5a\xdc\xf2\xb3\x89\x6b\xdf\xfe\x0c\x18\x07\xb7\xe2\x9c\xe2\x87\x64\x6c\xbc\x00\x5e\x2b\x6e\x66\xb5\xb7\x7d\x14\x63\xe8\x03\xd5\x89\xe6\xf2\xf2\x72\x5a\x56\x56\xe6\x9c\x99\x2f\xe0\x3c\xae\xe6\x99\x1d\xae\x6a\xa8\x72\xc6\x47\x28\x36\xc2\xb3\x6a\xe5\x78\xaa\x72\xd0\x3a\x1d\x63\x2d\x8c\xed\x30\xd0\xe4\xd9\xae\x43\x7f\x50\xce\x7d\x95\xa2\x79\xd8\x9d\x8f\x90\x75\x0a\x3c\xb1\x85\x3d\xc6\xa8\x9f\xe9\x29\x7d\xcd\x87\xc3\x6c\xfc\xa6\x35\xcd\x5c\xe1\xdb\x81\xc8\xab\x58\xfe\xa6\x2b\x62\xca\x32\x59\xe5\x73\x3a\x25\xd4\x8e\xde\xd2\x9c\x14\xa1\x1f\xf3\x67\x19\x60\xc9\x4d\xcf\x3f\x80\x3c\xf7\x34\xdc\xe3\x0f\xf1\x4f\x00\x00\x00\xff\xff\x1d\x36\x5c\x1d\x28\x05\x00\x00")

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

	info := bindata_file_info{name: "templates/new_ticket.html", size: 1320, mode: os.FileMode(420), modTime: time.Unix(1420847558, 0)}
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
	"static/#hier.css#": static_hier_css_,
	"static/css/ie.css": static_css_ie_css,
	"static/css/print.css": static_css_print_css,
	"static/css/screen.css": static_css_screen_css,
	"static/js/hier.js": static_js_hier_js,
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
	"static": &_bintree_t{nil, map[string]*_bintree_t{
		"#hier.css#": &_bintree_t{static_hier_css_, map[string]*_bintree_t{
		}},
		"css": &_bintree_t{nil, map[string]*_bintree_t{
			"ie.css": &_bintree_t{static_css_ie_css, map[string]*_bintree_t{
			}},
			"print.css": &_bintree_t{static_css_print_css, map[string]*_bintree_t{
			}},
			"screen.css": &_bintree_t{static_css_screen_css, map[string]*_bintree_t{
			}},
		}},
		"js": &_bintree_t{nil, map[string]*_bintree_t{
			"hier.js": &_bintree_t{static_js_hier_js, map[string]*_bintree_t{
			}},
		}},
	}},
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

