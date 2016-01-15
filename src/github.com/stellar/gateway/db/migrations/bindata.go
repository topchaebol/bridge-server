// Code generated by go-bindata.
// sources:
// mysql_01_init.sql
// DO NOT EDIT!

package migrations

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

var _mysql_01_initSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\xd1\xc1\x6b\xfa\x30\x14\xc0\xf1\x7b\xfe\x8a\x77\xac\xfc\x7e\xc2\x14\x1c\x03\xf1\x90\x69\xb6\x95\xb5\x51\x6a\x7a\xf0\x64\x62\xf3\xe6\x02\x9a\x4a\xfa\xea\xdc\x7f\xbf\x16\xd9\x74\x95\x9d\xf3\x79\xc9\x97\xbc\x7e\x1f\xfe\xed\xdd\x36\x18\x42\xc8\x0f\x6c\x9a\x09\xae\x04\x28\xfe\x98\x08\xd0\x4b\xf4\xa4\x82\xf1\x95\x29\xc8\x95\x5e\x43\xc4\x00\xb4\xb3\x1a\x9c\xa7\x68\x30\xe8\x81\x9c\x2b\x90\x79\x92\x00\xcf\xd5\x7c\x1d\xcb\x66\x3e\x15\x52\xfd\x6f\x5d\x45\x86\xea\x4a\xc3\xd1\x84\xe2\xdd\x84\x68\x70\x77\xf1\x67\x50\xd6\xa1\xc0\x0b\x18\xdd\x77\x41\xbd\xd9\x3b\x22\xb4\x6b\x43\x1a\x6c\xd3\x48\x6e\x8f\x5d\x53\x14\x88\xb6\x6b\x66\xe2\x89\xe7\xc9\x95\xdb\xa1\xdd\x62\xd0\xb0\x71\xdb\x36\x7e\xd8\xc4\xdc\x18\xf4\x47\xdc\x95\x07\x0c\xeb\x93\x6d\x28\xe1\x89\x7e\x3f\x16\xb0\xaa\x77\x74\x3e\xfd\xae\x1e\x8e\x46\xb7\x57\x2d\xb2\x38\xe5\xd9\x0a\x5e\xc5\x0a\xa2\xf6\xc7\x7a\xac\x07\x42\x3e\xc7\x52\x4c\xd2\xcf\x78\xc9\xd3\x9f\x91\xe9\x0b\xcf\x96\x42\x4d\x6a\x7a\x7b\x18\x33\x76\xbd\x90\x59\xf9\xe1\xd9\x2c\x9b\x2f\xfe\x5a\xc8\x98\x7d\x05\x00\x00\xff\xff\xd7\x8c\x1e\xc9\xc0\x01\x00\x00")

func mysql_01_initSqlBytes() ([]byte, error) {
	return bindataRead(
		_mysql_01_initSql,
		"mysql_01_init.sql",
	)
}

func mysql_01_initSql() (*asset, error) {
	bytes, err := mysql_01_initSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mysql_01_init.sql", size: 448, mode: os.FileMode(420), modTime: time.Unix(1452874929, 0)}
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
	"mysql_01_init.sql": mysql_01_initSql,
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
	"mysql_01_init.sql": &bintree{mysql_01_initSql, map[string]*bintree{}},
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
