// Code generated by go-bindata.
// sources:
// files/cf-mgmt.sh
// files/cf-mgmt.yml
// files/pipeline.yml
// files/security-group.json
// files/vars.yml
// DO NOT EDIT!

package generated

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

var _filesCfMgmtSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\xd4\x2f\x2d\x2e\xd2\x4f\xca\xcc\xd3\x4f\xcd\x2b\x53\x48\x4a\x2c\xce\xe0\x2a\x4e\x2d\x51\xd0\x4d\xe5\x4a\x4e\x51\x48\xce\xcf\x4b\xcb\x4c\xd7\x2d\x4a\x2d\xc8\xe7\x4a\x4e\xd3\xcd\x4d\xcf\x2d\x51\x50\x71\x76\x8b\xf7\x75\xf7\x0d\x89\x77\xf6\xf7\xf5\x75\xf4\x73\xe1\x02\x04\x00\x00\xff\xff\x4b\xdb\x02\xbb\x43\x00\x00\x00")

func filesCfMgmtShBytes() ([]byte, error) {
	return bindataRead(
		_filesCfMgmtSh,
		"files/cf-mgmt.sh",
	)
}

func filesCfMgmtSh() (*asset, error) {
	bytes, err := filesCfMgmtShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/cf-mgmt.sh", size: 67, mode: os.FileMode(493), modTime: time.Unix(1518199076, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _filesCfMgmtYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xcc\x4b\x6a\xc4\x30\x0c\x00\xd0\xbd\x4e\x21\xb2\xae\x9b\xbd\x2f\x53\x84\xab\x38\x22\xfe\x21\xc9\xa1\xa1\xf4\xee\x25\xc3\x30\x30\xeb\x07\x2f\x84\x00\xa3\x90\x6f\x5d\x6b\xc4\x22\x6d\xfe\x00\x48\xa5\xcc\x5f\xca\xd6\xa7\x26\x8e\x80\xe8\xd7\xe0\x88\xdf\x3d\x1d\xac\xe1\xc1\x80\xf8\x64\xfc\x55\x1e\xdd\xc4\xbb\x5e\x11\x87\x9c\xdd\xa9\x18\xeb\x29\x89\x6d\x4d\x5b\xa8\xb9\xfa\x07\x3a\xe5\x88\x4b\x21\x67\xf3\xe5\x0f\x40\xda\x98\x6e\x77\x1e\xb0\x51\xe5\x88\xa9\xb7\x4d\x72\xb8\x37\x00\x9d\xed\xb6\x41\xbe\xbf\xc9\x9a\x64\x75\xb2\xe3\x35\x7f\xda\x0e\xff\x01\x00\x00\xff\xff\x0a\x5a\x24\x71\xc6\x00\x00\x00")

func filesCfMgmtYmlBytes() ([]byte, error) {
	return bindataRead(
		_filesCfMgmtYml,
		"files/cf-mgmt.yml",
	)
}

func filesCfMgmtYml() (*asset, error) {
	bytes, err := filesCfMgmtYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/cf-mgmt.yml", size: 198, mode: os.FileMode(420), modTime: time.Unix(1512503607, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _filesPipelineYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x98\xdf\x6f\x9b\x30\x10\xc7\xdf\xf9\x2b\xee\x0f\x28\x8d\xf6\xb0\x17\xde\xa2\x86\x56\x91\xf2\xa3\x4a\xb2\x4d\xd3\x54\x59\x2e\x5c\xa8\x37\x83\x99\x6d\x32\x45\x55\xff\xf7\x89\xd0\x42\x62\x9c\x86\x92\x66\x95\x32\xde\x90\x7c\xdc\x9d\xbf\x1f\xdd\xf9\x6c\x89\x4a\x64\x32\x40\xe5\x39\x2e\x24\x34\x46\x0f\x02\x91\x2c\x59\xe4\x4a\x4c\x85\x03\xa0\xd7\x29\x7a\x10\x31\xed\x00\x14\xa6\x9e\x03\x00\x90\x49\xe6\xc1\xe3\x63\xc4\x34\xc9\x2d\x49\x26\xd9\xd3\xd3\x66\xe5\x5e\xd2\x24\x78\xf0\x20\xa6\x4a\xa3\x2c\xdd\x7e\xfa\x1c\x97\xee\x34\x8b\xb1\xf2\x07\x8f\x2c\xd1\x28\x57\x94\x6f\xac\x9e\x1c\xe7\xa7\xb8\xdf\x4e\x48\x22\xd5\xe8\x0a\x19\x29\x07\x20\xe5\x34\xc9\x53\x70\x21\x42\x6d\x66\x0b\xa0\x25\x8b\x22\x94\x1e\x68\x99\xe1\xc6\x4c\x53\xf5\xcb\x74\x02\xb0\x64\x7c\x77\xaf\xbd\x80\xf5\x72\x53\xd5\x0b\x96\x6e\x1c\xc5\xfa\x72\x1d\xf3\x8d\x69\x4a\x25\x8d\x55\xb1\x6d\x80\xf9\xf7\xf9\xc2\x1f\x93\xc1\x74\xdc\x1f\x4e\x72\x09\xd4\x5a\x69\x8c\x49\x28\x62\xca\x92\x67\x0d\x00\xbe\xcc\xfd\x19\x19\x0e\x72\x83\x4c\xa1\x24\x2c\x2c\x97\x6e\xfb\xf3\xf9\xb7\xe9\x6c\xb3\x96\x52\xa5\xfe\x08\x59\x2d\x5e\x4d\x27\xd7\xc3\x1b\x32\x18\xce\x3c\xb8\xec\x15\xf9\xbd\x2c\x8d\x86\xfe\x64\x41\xe6\xfe\xd5\xcc\x5f\xe4\x3f\x07\x9c\x61\xa2\x89\xc2\x40\xa2\x2e\x3d\x8c\xa6\x37\x64\xe4\x7f\xf5\x47\xb9\x09\x17\x11\xe1\xb8\x42\x5e\x05\xb8\x26\xe3\x9b\xf1\x82\x5c\x4d\xc7\xe3\xfe\x64\xb0\xab\x8c\xa9\xb9\xc2\x20\x93\x4c\xaf\xdd\x48\x8a\x2c\x3d\x5e\xfe\xba\xbf\x8e\x44\x8d\x84\x29\x52\x09\x85\x2a\xc5\xa2\xc4\x0d\x71\x49\x33\xae\x8f\x87\x93\x0b\xaa\x14\x86\x1e\xfc\xb0\x87\xbe\xdb\x22\x78\x30\x78\x47\xb2\x22\x79\x40\xac\x92\x68\x88\x1c\x8f\x68\x6d\x1b\xb3\xa2\xb1\xee\x2f\xbd\xdd\x18\x1d\xa5\x8a\xd2\xb6\x32\x96\xd3\xc6\x4d\x25\x5b\xe5\xdf\xc5\x06\x1b\x10\x32\xcb\x29\xf7\x7c\xd7\xf0\x58\xb2\x44\xeb\x58\xd9\x4e\xa9\x9a\x50\x25\x3a\x85\x72\xc5\x02\x74\x69\x10\xa0\x3a\x09\xae\x5a\x84\x0e\x51\x85\xc8\x10\xa7\xc2\xf2\x40\xe5\xfb\x15\x94\xe9\xe4\x00\xb0\x57\x62\x77\xec\xb6\xd8\xed\xd5\xa9\x36\x12\xa6\x34\xc0\x96\xc5\x75\x01\x59\x1a\xbe\x70\xcc\x15\x69\xd6\x1d\xcb\x90\x67\x4d\x6c\xd0\xbf\x25\xdb\x31\x78\x48\x53\x52\x0b\xd4\x76\xa6\x2c\x24\x34\x07\x8f\x63\x60\xee\x43\xf7\x96\xa1\xe4\xff\x20\xdb\x6e\x2c\x31\x91\x3d\x97\x4e\x5b\x64\xc5\x7f\xaf\xd7\x9b\x19\xa2\xa3\x52\x51\xd9\xd5\xc6\x4a\xa5\x68\x69\xef\x8d\xa6\x51\x3d\x59\xd3\x38\x6b\x7c\xa7\x69\x97\x16\x1d\xed\xa8\x7f\x67\x42\xd3\x7f\x50\x86\x55\xa0\xb3\xa6\x79\x0c\xa6\x67\x85\xec\x9c\xde\xfc\x48\x62\x07\x76\xb1\xff\x81\xa4\x21\xc7\xee\xc1\xa4\x31\xd0\xbd\xcf\x25\xe6\xf4\xf8\x01\x83\x8b\x25\x85\xb3\xe6\x78\xd2\x36\x5b\xa9\x68\x43\xdc\xb6\xc5\x1e\xbe\xc6\xdb\x82\x9c\x35\xc5\xf6\x78\xcc\xe6\xca\x94\xe0\x54\x33\x91\xb8\x0a\xa3\x18\x13\xfd\x2e\xef\xce\xc5\xe5\xb0\x76\x44\x36\xac\x49\x6b\x4e\x1f\xca\xb3\x21\x97\x57\xc8\xbe\x15\x99\x45\x83\xea\xe6\xce\x91\x26\x59\xda\xaa\x6f\x9a\x95\x7a\x51\x9f\x90\x8e\xea\xa7\xb6\xdc\xba\x52\xdc\xba\xbf\xd7\xf4\xf9\x1b\x00\x00\xff\xff\x51\x31\x39\x03\xb3\x1d\x00\x00")

func filesPipelineYmlBytes() ([]byte, error) {
	return bindataRead(
		_filesPipelineYml,
		"files/pipeline.yml",
	)
}

func filesPipelineYml() (*asset, error) {
	bytes, err := filesPipelineYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/pipeline.yml", size: 7603, mode: os.FileMode(420), modTime: time.Unix(1540913983, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _filesSecurityGroupJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8a\x8e\xe5\x02\x04\x00\x00\xff\xff\x44\xd2\x68\x70\x03\x00\x00\x00")

func filesSecurityGroupJsonBytes() ([]byte, error) {
	return bindataRead(
		_filesSecurityGroupJson,
		"files/security-group.json",
	)
}

func filesSecurityGroupJson() (*asset, error) {
	bytes, err := filesSecurityGroupJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/security-group.json", size: 3, mode: os.FileMode(420), modTime: time.Unix(1510432475, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _filesVarsYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x8f\x41\x4b\x03\x31\x14\x84\xef\xfb\x2b\x06\x7a\x5e\xbc\xf7\x56\x6c\x05\x41\xb4\x88\x9e\x43\xcc\xbe\x4d\x1f\x24\x79\x21\xef\xad\x65\xff\xbd\xec\xae\x15\xbd\x7a\xca\xf0\x65\x32\x33\x89\x6c\xae\x51\x15\x37\x35\xde\x63\x37\xcb\xd4\x10\xd9\xb0\x30\x4c\x8d\x3b\x9d\xd5\x28\xbb\x41\xb2\xe7\x72\x73\x84\x11\x1b\xc7\xc6\xbb\x49\xa9\x39\x1e\xf6\xd8\x61\x91\xf0\x21\xc8\x54\x0c\x57\xb6\x0b\x2a\xb5\xcc\xaa\x2c\x05\x26\x08\x8d\xbc\x11\xa4\x45\xbd\xd3\xea\x03\x69\x57\xbd\xea\x55\xda\xfa\xfc\x78\x3a\xbf\x9e\xee\x0f\x6f\xa7\x23\x7a\xbc\x2b\x21\x24\xa6\x62\x4e\x29\x34\x32\xf4\xb8\x99\x21\xe3\x7f\xba\xfe\xc4\x2d\x85\x1b\xc0\x77\xfe\x28\x0d\x93\xf7\x61\x13\xdb\xaf\xba\x34\xf8\xea\x7e\x8f\xfc\xd9\x60\x82\x0f\x2e\xeb\xb9\x98\xd0\x43\x4a\x9a\x51\x88\x06\x1a\xc0\xcb\x42\x2e\x11\x4f\xc7\xc3\xb9\x4b\x12\x5d\xa2\x4f\x4a\x7b\x3c\x3e\x3f\xbc\x60\x87\x24\x31\x2e\xd7\x2b\x5d\x1b\xc3\xd8\xe7\x98\x0d\x41\x72\xf6\x65\x50\x70\x81\x5d\x08\x95\x2b\x25\x2e\xd4\x7d\x05\x00\x00\xff\xff\xc0\x62\xea\xa1\xb0\x01\x00\x00")

func filesVarsYmlBytes() ([]byte, error) {
	return bindataRead(
		_filesVarsYml,
		"files/vars.yml",
	)
}

func filesVarsYml() (*asset, error) {
	bytes, err := filesVarsYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/vars.yml", size: 432, mode: os.FileMode(420), modTime: time.Unix(1527696617, 0)}
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
	"files/cf-mgmt.sh": filesCfMgmtSh,
	"files/cf-mgmt.yml": filesCfMgmtYml,
	"files/pipeline.yml": filesPipelineYml,
	"files/security-group.json": filesSecurityGroupJson,
	"files/vars.yml": filesVarsYml,
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
	"files": &bintree{nil, map[string]*bintree{
		"cf-mgmt.sh": &bintree{filesCfMgmtSh, map[string]*bintree{}},
		"cf-mgmt.yml": &bintree{filesCfMgmtYml, map[string]*bintree{}},
		"pipeline.yml": &bintree{filesPipelineYml, map[string]*bintree{}},
		"security-group.json": &bintree{filesSecurityGroupJson, map[string]*bintree{}},
		"vars.yml": &bintree{filesVarsYml, map[string]*bintree{}},
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

