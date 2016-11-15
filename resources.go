// Code generated by go-bindata.
// sources:
// resources/config.yml
// DO NOT EDIT!

package main

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

var _resourcesConfigYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x52\xcd\x6a\xdc\x30\x10\xbe\xeb\x29\x06\xa5\x87\x04\x82\x1f\xc0\x97\x42\x37\x14\xb6\x87\xb4\xb0\xdb\x53\x29\x46\x2b\x8f\xd7\x6a\xec\x91\x90\x46\x4d\x8c\x57\xef\x5e\x24\x7b\x5d\x36\xb4\x3a\xcd\xcf\x37\x9a\x6f\xbe\x99\xbb\x3b\x08\xec\xa3\x66\x0f\xda\x52\x67\xce\x71\x50\x6c\x2c\x41\x67\x06\x14\xe2\x0e\x9e\xb0\x33\x64\x72\x28\x40\x67\x3d\xb0\x57\x14\x32\x86\xce\x80\x4a\xf7\xc0\x93\x43\xb0\x1d\x7c\x39\x7c\x7d\x3e\xe8\x1e\x47\x55\x89\x1c\x6b\xae\x48\x6c\x46\xe5\x6a\x01\xb9\x91\xa1\x73\x0d\xf2\x50\x0c\x29\x00\x0c\x31\x9e\xd1\xd7\x20\xf7\xc4\x39\x40\x71\x3c\x15\xff\xf3\x60\x55\x89\x9c\xac\x1d\x50\x51\x0d\xf2\x93\xb5\xc3\x82\x19\x86\x1a\x24\x99\xe2\x29\xef\xd5\x54\x83\xfc\x31\xcf\xd5\x9e\x08\xfd\x71\x72\x98\xd2\xcf\x9c\xb3\xa7\x5f\xa8\xb9\x06\x79\x9b\x93\x79\xb0\x7d\x07\xf8\x86\x3a\x32\x6e\x0a\xbc\x1a\xee\x41\xda\xc8\x4f\xc6\x4b\xb0\x2e\x4f\xfd\xb8\x65\x6d\x64\x17\x39\x14\x65\xee\xc3\xc3\x82\xe6\x1e\xa1\xdd\x24\xca\x3a\xc8\x05\xd7\x64\x18\xa9\x11\x25\x30\x8e\x2e\xeb\x50\x89\x77\xa9\x85\xd8\xb3\x1a\x31\xa5\x2a\xbc\x9a\x8e\x0b\xb3\xa3\x5d\x7b\x95\xdf\x97\xf6\xd1\xbf\x13\x19\x0c\x41\x8b\x0e\xa9\x45\xd2\xd3\xbd\xaa\x5e\x2a\x05\xc1\xa1\x36\x9d\xc1\xb6\x10\xf9\xe0\xb1\x93\xf0\x82\x53\xc6\xfe\xad\x7c\xd8\x78\x6c\xf5\x06\x43\x0d\xec\x63\xd9\xf8\x71\xa5\xbb\xec\xfb\x5f\x4c\xc4\x66\x35\xd7\xd9\x6a\xb8\x2c\x1b\x8e\x9a\x61\x1d\xea\xc2\xf6\xbb\x73\xe8\x77\x6a\xc4\x61\xa7\x02\xa6\x04\xb3\x00\x98\x67\xaf\xe8\x8c\x50\x7d\xf3\xd6\xa1\x67\x83\x21\x25\x01\xe5\xfd\x56\x1e\x36\x4d\xea\x6c\xe6\x95\x5d\xb6\x5b\xca\xde\x7f\xfe\xfe\x38\xcf\x48\x6d\x4a\xa2\xb4\xa8\x76\xbd\x19\x5a\x8f\x74\xc1\x37\xf6\x4a\xf3\xe1\xca\x79\x69\x96\xd6\x1b\xb8\x95\xb8\x57\x01\xf4\x5a\xf8\x08\x6c\x57\x45\xa7\xf5\x76\x43\x0e\x11\x06\xbe\xdd\xca\xb5\x42\x14\xa3\xd9\x72\xa1\xc9\xd8\xe5\xe6\xf3\x70\x52\xfc\x09\x00\x00\xff\xff\xe6\x4a\x68\x4e\x6f\x03\x00\x00")

func resourcesConfigYmlBytes() ([]byte, error) {
	return bindataRead(
		_resourcesConfigYml,
		"resources/config.yml",
	)
}

func resourcesConfigYml() (*asset, error) {
	bytes, err := resourcesConfigYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/config.yml", size: 879, mode: os.FileMode(420), modTime: time.Unix(1479199471, 0)}
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
	"resources/config.yml": resourcesConfigYml,
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
	"resources": &bintree{nil, map[string]*bintree{
		"config.yml": &bintree{resourcesConfigYml, map[string]*bintree{}},
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
