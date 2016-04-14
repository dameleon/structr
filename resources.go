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

var _resourcesConfigYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x52\xc1\x6e\xdb\x30\x0c\xbd\xfb\x2b\x08\x67\x87\x16\x28\xfc\x01\xbe\x0c\x58\x8a\x01\xd9\xa1\x1b\x90\xec\x34\x0c\x86\xa2\xd0\xb1\x56\x5b\x32\x24\x7a\xad\xe1\xf8\xdf\x47\xca\x8e\x86\x04\x9b\x4e\xe4\xe3\xa3\x48\x3e\x72\xb3\x81\x40\x7e\xd0\xe4\x41\x3b\x5b\x9b\xf3\xd0\x2a\x32\xce\x42\x6d\x5a\xcc\xb2\x0d\x3c\x63\x6d\xac\x11\x28\x40\xed\x3c\x90\x57\x36\x08\xc7\x9e\x01\x95\x6e\x80\xc6\x1e\xc1\xd5\xf0\x65\xff\xf5\x65\xaf\x1b\xec\x54\x91\x09\x56\x5d\x99\x58\x75\xaa\x2f\x33\x90\x42\x9c\x55\x42\xbe\x8f\x46\xce\x90\xb1\x84\x67\xf4\x8c\xed\x2c\x09\x60\x87\xee\x18\xfd\xcf\xad\x53\x11\x39\x3a\xd7\xa2\xb2\x0c\x7d\x62\x6b\xe1\xb4\x2d\xbb\xd6\x44\x4f\x79\xaf\x46\x76\x7f\x4c\x53\xb1\xb3\x16\xfd\x81\x8b\xcf\xf3\x4f\x89\xb9\xe3\x2f\xd4\xc4\xc1\xdb\x58\x2e\x83\xed\x6a\xc0\x77\xd4\x03\x61\x52\xe0\xcd\x50\x03\xb9\x1b\xe8\xd9\xf8\x1c\x5c\x2f\x53\x3f\xa5\x28\xe3\xfd\x40\x21\x2a\xf3\x10\x1e\x17\x36\x35\x08\xa7\x24\x91\xe8\x90\x2f\xbc\x4a\x68\x56\x75\x98\x03\x61\xd7\x8b\x0e\x45\x76\x17\x5a\x1a\x7b\x61\x6b\x9e\x8b\xf0\x66\x6a\x8a\x9d\x1d\xdc\x5a\x2b\xfe\xbe\x94\x1f\xfc\x9d\xc8\x2c\x1d\x17\xee\xd1\x9e\xd0\xea\xf1\x41\x15\xaf\x85\x82\xd0\xa3\x36\xb5\xc1\x53\x6c\xe4\x83\xc7\x3a\x87\x57\x1c\x85\xfb\x37\xf3\x31\xf5\x91\xf2\x0d\x86\x92\x37\x3b\xc4\x8d\x1f\xd6\x76\x97\x7d\xff\xab\x93\x2c\x59\xd5\x75\xb6\x12\x2e\xcb\x86\x19\x87\x75\xa8\x0b\xb9\xef\x7d\x8f\x7e\xcb\x76\xbb\x55\x81\xa7\x84\x89\x59\xd3\xc4\x97\x71\x46\x28\xbe\x79\xc7\x61\xe2\xea\xf3\xcc\xb8\xbc\xdf\xca\x43\xd2\xa4\x14\x53\x56\x76\x49\xb7\x24\xde\x7f\xfe\xfe\x38\x4d\x3c\x0d\xff\x14\x4b\x14\xdb\xc6\xb4\x27\x8f\xf6\x82\xef\x9c\xad\x69\x7f\xed\x79\x29\x36\xaf\x37\x70\x2b\x71\xa3\x02\xe8\x35\xf1\x09\xc8\xad\x8a\x8e\xeb\xed\x06\x81\x2c\x06\xba\xdd\xca\x35\x23\x8b\x46\x95\x62\xa1\x12\xee\x72\xf3\x32\x5c\x9e\xfd\x09\x00\x00\xff\xff\xe6\x4a\x68\x4e\x6f\x03\x00\x00")

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

	info := bindataFileInfo{name: "resources/config.yml", size: 879, mode: os.FileMode(420), modTime: time.Unix(1460621234, 0)}
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
