// Code generated by go-bindata.
// sources:
// assets/html/layout.html
// assets/html/table.html
// DO NOT EDIT!

package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _assetsHtmlLayoutHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x52\xcb\x8e\xdb\x30\x0c\xbc\xe7\x2b\x58\x9d\x1b\x69\x03\x14\x6d\x51\xc8\x7b\xe8\xe3\xdc\x1e\xf6\xd2\x23\x2d\xd1\x90\xb6\x7a\x18\x26\x9d\x6c\x60\xe4\xdf\x0b\x3f\xb2\xc9\xf6\xd6\x93\x39\x1c\x70\x86\x1c\x6b\x9a\xc0\x53\x17\x0b\x81\x4a\x78\xae\xa3\x28\xb8\x5c\x76\x3b\xfb\xee\xfb\xcf\x6f\x4f\xbf\x7f\xfd\x80\x20\x39\x3d\xee\xec\xfa\x01\xb0\x81\xd0\xcf\x05\x80\xcd\x24\x08\x2e\xe0\xc0\x24\x8d\x1a\xa5\xdb\x7f\x56\xf7\x54\xc1\x4c\x8d\x3a\x46\x3a\xf5\x75\x10\x05\xae\x16\xa1\x22\x8d\x3a\x45\x2f\xa1\xf1\x74\x8c\x8e\xf6\x0b\x78\x0f\xb1\x44\x89\x98\xf6\xec\x30\x51\x73\xb8\x0a\x49\x94\x44\x8f\xd3\x04\xfa\x69\xae\xe0\x72\xb1\x66\xed\xad\x7c\x8a\xe5\x0f\x0c\x94\x1a\xc5\x72\x4e\xc4\x81\x48\x14\x84\x81\xba\x46\x05\x91\x9e\xbf\x18\x93\xf1\xc5\xf9\xa2\xdb\x5a\x85\x65\xc0\x7e\x06\xae\x66\xd3\xd5\x22\x7b\x3c\x11\xd7\x4c\xe6\x83\xfe\xa4\x1f\x8c\x63\x7e\xd3\xd6\x39\x16\xed\x98\xd5\x7f\xd8\x39\x5f\x9e\x59\xbb\x54\x47\xdf\x25\x1c\x68\xf1\xc2\x67\x7c\x31\x29\xb6\x6c\xda\x31\x65\x34\x0f\xfa\xa3\x3e\x2c\x76\x0b\x7e\xe3\x63\xcd\x35\x64\xdb\x56\x7f\xde\xac\x99\x9c\xc4\x5a\xc0\x25\x64\x6e\xd4\x06\xb7\xc5\x00\xac\x8f\xc7\x2b\x37\x07\x8d\xb1\xd0\xf0\xca\xce\x3f\xee\x70\xa5\x97\xfc\xee\x28\x80\xfb\x7c\x6f\x13\x26\x1c\xee\xe6\xfb\x57\xe7\xb1\xfd\x57\xc1\x9a\xfe\x06\x66\xb1\xaf\xd5\x9f\x6f\x5a\xd6\xf8\x78\xdc\xce\x30\xdb\xe2\xeb\xa1\xeb\x7d\xd6\xac\xcf\x6b\x37\x4d\x40\xc5\xcf\x83\x7f\x03\x00\x00\xff\xff\x4f\x4c\xf7\xe1\x98\x02\x00\x00")

func assetsHtmlLayoutHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsHtmlLayoutHtml,
		"assets/html/layout.html",
	)
}

func assetsHtmlLayoutHtml() (*asset, error) {
	bytes, err := assetsHtmlLayoutHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/html/layout.html", size: 664, mode: os.FileMode(420), modTime: time.Unix(1515204438, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsHtmlTableHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\x8c\xc1\x09\x42\x41\x0c\x44\xef\x5b\xc5\xb0\x05\xb8\x0d\x84\xf4\xb2\x92\x28\x82\xe4\xb0\xe4\x36\xa4\x77\x89\xca\xbf\xcd\x9b\x07\x8f\x84\xf9\xe3\x15\x8e\x99\xfb\xfe\xf6\x89\xaa\x21\xdf\xa9\x24\xce\x8e\xa7\xe3\x56\x25\x79\x14\xe4\xc5\x90\x34\x25\xdb\xac\xb4\x56\x1e\xd6\xf7\xca\xa3\x7f\x90\xf5\xeb\x0c\x12\x1e\xd6\xe5\x4f\x00\x00\x00\xff\xff\x9d\x5d\xaa\x8c\x6f\x00\x00\x00")

func assetsHtmlTableHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsHtmlTableHtml,
		"assets/html/table.html",
	)
}

func assetsHtmlTableHtml() (*asset, error) {
	bytes, err := assetsHtmlTableHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/html/table.html", size: 111, mode: os.FileMode(420), modTime: time.Unix(1515200017, 0)}
	a := &asset{bytes: bytes, info: info}
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

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
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
	"assets/html/layout.html": assetsHtmlLayoutHtml,
	"assets/html/table.html": assetsHtmlTableHtml,
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
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"assets": &bintree{nil, map[string]*bintree{
		"html": &bintree{nil, map[string]*bintree{
			"layout.html": &bintree{assetsHtmlLayoutHtml, map[string]*bintree{}},
			"table.html": &bintree{assetsHtmlTableHtml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
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

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
