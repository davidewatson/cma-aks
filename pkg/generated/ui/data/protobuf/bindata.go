// Code generated by go-bindata.
// sources:
// api/api.proto
// DO NOT EDIT!

package protobuf

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/elazarl/go-bindata-assetfs"
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

var _apiProto = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\x51\x73\xdb\x38\x0e\x7e\xd7\xaf\xc0\xf8\xe5\xd2\x9b\xc4\x4a\xd3\xdd\xbb\x9d\xf8\x72\x73\x3e\xa7\xd7\x7a\xd2\x26\x99\x3a\xdb\x9d\x7d\xf2\xd0\x14\x2c\xf3\x2c\x91\x2c\x09\xd9\x71\x3b\xf9\xef\x37\xa4\x28\x59\x92\xe5\xb4\xdd\xee\xc3\xf9\xa1\x95\x49\x00\x02\x3e\x00\x1f\xe0\xc4\x31\x4c\x94\xde\x19\x91\xae\x08\x2e\xce\x5f\xfe\x02\x33\x96\xdb\x42\xa6\x30\xbb\x9e\xc1\x24\x53\x45\x02\xb7\x8c\xc4\x06\x61\xa2\x72\x5d\x90\x90\x29\x3c\x20\xcb\x81\x15\xb4\x52\xc6\x0e\xa3\x38\x8e\xe2\x18\xde\x09\x8e\xd2\x62\x02\x85\x4c\xd0\x00\xad\x10\xc6\x9a\xf1\x15\x56\x37\xa7\xf0\x11\x8d\x15\x4a\xc2\xc5\xf0\x1c\x4e\x9c\xc0\x20\x5c\x0d\x5e\x8c\x9c\x89\x9d\x2a\x20\x67\x3b\x90\x8a\xa0\xb0\x08\xb4\x12\x16\x96\x22\x43\xc0\x47\x8e\x9a\x40\x48\xe0\x2a\xd7\x99\x60\x92\x23\x6c\x05\xad\xfc\x7b\x82\x15\xe7\x09\xfc\x1e\x6c\xa8\x05\x31\x21\x81\x01\x57\x7a\x07\x6a\xd9\x14\x04\x46\xc1\x69\xf7\x59\x11\xe9\xcb\x38\xde\x6e\xb7\x43\xe6\x1d\x1e\x2a\x93\xc6\x59\x29\x6a\xe3\x77\xd3\xc9\xeb\xdb\xd9\xeb\xb3\x8b\xe1\x79\x50\xfa\x55\x66\x68\x2d\x18\xfc\x54\x08\x83\x09\x2c\x76\xc0\xb4\xce\x04\x67\x8b\x0c\x21\x63\x5b\x50\x06\x58\x6a\x10\x13\x20\xe5\x9c\xde\x1a\xe1\x70\x3b\x05\xab\x96\xb4\x65\x06\x9d\x99\x44\x58\x32\x62\x51\x50\x0b\xb3\xca\x45\x61\x5b\x02\x4a\x02\x93\x30\x18\xcf\x60\x3a\x1b\xc0\xbf\xc7\xb3\xe9\xec\xd4\x19\xf9\x6d\xfa\xf0\xf6\xee\xd7\x07\xf8\x6d\xfc\xe1\xc3\xf8\xf6\x61\xfa\x7a\x06\x77\x1f\x60\x72\x77\x7b\x3d\x7d\x98\xde\xdd\xce\xe0\xee\x3f\x30\xbe\xfd\x1d\x6e\xa6\xb7\xd7\xa7\x80\x82\x56\x68\x00\x1f\xb5\x71\x11\x28\x03\xc2\xa1\x89\x89\x87\x6e\x86\xd8\x72\x61\xa9\x4a\x97\xac\x46\x2e\x96\x82\x43\xc6\x64\x5a\xb0\x14\x21\x55\x1b\x34\xd2\x55\x82\x46\x93\x0b\xeb\xb2\x6a\x81\xc9\xc4\x99\xc9\x44\x2e\x88\x91\x3f\x3a\x88\x6b\x18\x39\x91\xf7\x82\xaf\x18\x66\xf0\x11\x25\x7e\x16\x0c\xfe\x91\x6f\xca\xa7\x7f\xa5\x39\x13\xd9\x90\xab\xfc\x9f\x51\x64\x77\x92\xd8\x23\x5c\xc1\x40\x1b\x45\xea\xd5\x60\x14\x45\x9a\xf1\xb5\xf3\x80\xe7\x8c\xad\xed\x28\x8a\x44\xae\x95\x21\x18\xa4\x4a\xa5\x19\xc6\x4c\x8b\x98\x49\xa9\x82\x03\x43\xaf\x39\x18\xd5\x62\xfe\x3b\x3f\x4b\x51\x9e\xd9\x2d\x4b\x53\x34\xb1\xd2\x5e\xb4\x57\x2d\x8a\xca\x5b\x38\x49\x8d\xe6\xc3\x94\x11\x6e\xd9\xae\xbc\xe6\xf3\x14\xe5\x3c\x58\x19\x06\x2b\x43\xa5\x51\x32\x2d\x36\x17\xd5\xcd\x0b\xb8\x82\x2f\x11\x80\x90\x4b\x75\xe9\x9f\x00\x48\x50\x86\x97\x30\x98\x64\x85\x25\x34\xf0\x9e\x49\x96\xa2\x81\xf1\xcd\x0c\xde\x62\xa6\xdd\xe3\xfd\x74\x30\xf2\xc2\x9b\xb2\x6b\x2e\x61\xb0\x39\x1f\xbe\x1c\x9e\x87\x63\xae\x24\x31\x4e\x95\x49\xf7\x91\x2c\x77\x56\x3b\xe8\x06\x79\xf7\x29\x4c\x76\x09\x03\x57\xf0\xf6\x32\x8e\x53\x41\x19\x5b\x38\xb0\xe3\x0a\xff\x98\xe7\xec\x8c\xad\x6d\x43\x07\x5d\x46\x2e\x61\x70\x98\xa2\x20\xf4\xe4\xfe\xf3\xff\xe0\x23\xa1\x91\x2c\x9b\x27\x8a\xdb\xca\xb1\xef\x7d\x67\x82\x96\x1b\xe1\xd1\x74\xb1\x28\x83\xc0\x16\xaa\x20\xf8\x06\xb0\x9e\x22\x00\xcb\x57\x98\xa3\xbd\x84\xb7\x0f\x0f\xf7\xb3\x51\xf7\xc4\x1d\x70\x25\x6d\xe1\x4f\x06\xa1\x71\xdd\xdb\xe2\xff\x5a\x25\xbd\x19\x6d\x54\x52\xf0\x63\xf7\x4f\xa3\x28\xb2\x68\x36\x82\x63\xed\x53\x19\xaa\xeb\x47\x91\x65\x4e\x7f\x23\x3c\xd3\x31\xe0\xa5\x84\xbf\x37\x9a\xc3\xc4\x20\x23\xac\xf4\x4e\x5a\x5f\xdf\xdb\xf4\x05\x18\xa4\xc2\x48\xdb\xb9\xfa\x80\x3a\xdb\xbd\x68\xe4\xba\xae\x4b\x5f\xf7\x43\xa6\xc5\xd0\x61\x5c\x55\xdb\xfe\xa3\x0b\x82\x4b\x18\xf8\xce\xd8\xbc\x8c\x83\x3f\x83\x96\xcc\x42\x25\x3b\x27\xf4\xd7\xfd\xf1\x53\x48\x6e\x2b\x30\x83\x64\x04\x6e\x4a\x9a\xb0\xc4\xa8\xb0\x8e\x5a\xeb\x28\x1d\x05\x80\x20\x0b\xeb\x62\x81\x5c\xc9\xa5\x48\x3d\x8b\x70\x25\x25\x72\x12\x1b\x41\xbb\x1a\x89\x37\x48\x35\x0c\xfb\xe7\x36\x06\xfb\xf3\x3f\x0e\x40\x8a\xcf\x03\xd0\x1b\x69\x82\x19\x12\xf6\xe4\xef\xda\x5f\xd4\x8e\xb7\xbe\xb6\x7d\x6f\x5d\xfd\x71\xf7\x83\x27\xdf\x1d\x41\x9d\x2b\x06\x99\xb0\xe4\xf2\x14\x14\x6d\x4f\x0a\xde\x39\x91\x93\xf6\xf7\x63\xa9\x70\x77\x7f\x76\x3a\x62\xe7\xe3\xd7\x23\x2a\x8c\xac\xf8\xd0\x13\xaa\xc9\x7d\x6b\x06\x86\x60\x5a\x80\xeb\xcc\x46\xba\xde\x20\x85\xad\x63\xda\x10\x3f\xd9\x1f\x1f\x04\x19\xce\xff\xb4\x00\x83\xbb\x3d\xb1\x3d\x45\x51\x8e\xd6\xba\x71\xd6\xa5\x81\x3d\xa1\xdc\xb2\x1c\xab\xf5\xa5\xea\x32\x52\xb0\xc0\x3d\xcb\x60\xe2\x85\xdd\xb2\x20\x53\x3f\x04\xe0\x0a\x5e\x8e\x2a\x0b\x0f\xab\x20\xeb\x46\x71\x35\xcb\x3d\x0e\x5e\xa2\xf5\xea\xfb\x20\x37\xd3\xc8\xf7\x4a\x57\x70\x31\x3a\xea\xad\x07\xaa\x41\x80\x2b\xf4\x3b\x86\x32\x7e\x8d\x6b\xba\xbd\x65\xb6\xe9\xb4\xdb\x9b\xfc\x86\xe7\x16\x29\xb4\x14\x95\x4c\xa4\x32\x50\xeb\x83\x00\x12\x24\x26\x32\xdb\x45\x22\xa8\x82\x41\xab\x95\xb4\x58\x46\x54\x5e\x4e\x09\xf3\x5a\xb0\x1b\x42\x8b\x70\xbe\x05\xed\x4c\xa9\xb5\x5b\xd4\xf4\xb3\x58\x4f\x0c\x26\x28\x49\xb0\xcc\x3a\xbd\x4f\x05\x9a\x5d\xbd\x47\x35\xa9\x64\xfc\xb9\x30\xd8\x94\xe6\x8d\xe7\xe3\xce\x76\xc0\x9e\xda\x96\xa7\x42\x96\xc4\xbc\xb3\x84\xf9\x21\x9c\x4d\x70\xae\x3d\x9e\xcf\x42\xd4\xa5\xb6\x66\x8e\x19\xb9\x05\xb5\xf1\xee\xbf\xd8\x12\x0c\x52\x6e\x80\x93\x51\xbb\xef\xc1\x29\x70\xdc\x0f\x81\x74\x48\xb7\x7b\x87\x27\xaa\xc8\x92\x16\x54\x0b\xac\xfc\x0c\xdd\xd3\x57\x78\xb3\x7a\xc2\x39\xd5\x66\x99\x86\xb8\xc2\x08\x3c\x9e\xaf\x40\xa3\xf0\xe5\xf8\xf5\x0f\xa5\x34\x28\xbd\xeb\x25\x78\xd4\xae\x4d\x93\xbe\x7e\x38\xf4\xb9\x29\xb4\x77\xe6\xba\xd3\x0c\xcd\xe0\x45\xd2\xf2\xa1\xa7\x75\x7a\x4a\xe0\x62\xd4\x57\x44\xb6\x05\x74\x8f\x76\x0d\xf4\xab\x3e\xa7\x1b\xc5\xfc\xff\xed\x7a\x8f\x7e\x63\x53\x22\x55\x2d\x4a\xee\xf1\x88\xb9\x86\xfc\x15\xfc\x74\x9c\x96\x5b\x4c\xde\xdb\xb9\x35\xbd\x9f\x41\x26\xd6\x08\x6c\x6d\xbf\x3a\x47\xaa\xd9\xab\x96\x70\x53\x2c\xd0\x48\x24\x6c\x69\xad\x7f\xb1\xf3\x4a\xc8\x43\xd6\xd4\x76\x1b\xfb\x57\x06\xd0\xf8\x66\xe6\x3d\x66\xae\xf5\x4b\xd0\xbe\x61\xb2\x08\x0b\x6f\xc7\xfb\xee\x58\x89\x74\x35\x67\x1b\x26\x32\xb6\x10\x99\xa0\x5d\x09\x55\xc3\x93\x25\x5b\x18\xc1\x03\xb5\x17\xb6\x33\x41\x91\xb6\xca\xac\xe7\x41\xe8\x0a\x7e\x1e\x45\x0e\xe7\xa0\xcb\xdb\xdc\x55\x84\x5f\xca\xdc\x85\xe1\xd4\x9b\x89\xab\x52\x73\xc0\x64\x5f\x5a\xb8\x68\x3d\x4d\xbc\x91\xf1\xfd\x14\xc6\x9c\xa3\x6d\x81\xca\xb4\x9e\x77\x4a\xd6\xa9\x3d\xa0\x64\x92\x6a\x3d\x76\xa0\x47\xa5\x40\xb3\x74\x9d\xde\x3d\xb3\x76\xab\x4c\xf2\x8c\xa6\xae\x44\x0e\x3a\xae\x2f\x59\xed\x60\x7c\xe6\xae\x19\x31\x98\xa0\xec\x54\x6f\xa6\xca\xc4\x3f\x37\x0c\x16\x85\x68\x93\xf5\x37\xce\x82\xaa\xf1\xa5\x25\xff\xe7\xa1\xd4\xa8\x42\x77\xa8\x70\x7c\x33\xab\xee\xdf\xb8\x6b\x10\xe1\xdb\xbc\x94\x2e\x0b\xe5\x88\x29\x48\x70\x29\xa4\x5b\xa8\x69\xa7\xd1\xff\xd6\x91\x45\xbe\x70\x25\xb9\xac\x0d\x95\x2f\xac\xf3\xde\x7d\xdf\x7e\x55\x0c\x70\xc9\x06\xf7\xf8\xd7\xd4\x02\xbd\x7d\xd8\xf5\xcc\x7b\x72\x32\x23\x26\x13\x66\x92\xf9\xf5\xc5\x7c\x73\x71\x0a\x48\x7c\xf8\xa2\x6b\xc8\x8b\xd6\x48\x05\x43\xef\x85\x14\x79\x91\xf7\x05\x02\x27\x09\x2e\x59\x91\x91\xcf\xca\x67\x34\x6a\x6f\x52\x48\x7a\x75\x01\xb9\x90\xf3\x4f\x05\x93\x54\xb6\xd8\xab\xb6\x65\xf6\xf8\x03\x96\xd9\x63\xd3\xf2\x4f\x8d\x95\x39\x8e\xdd\x00\x6d\x32\x91\x2b\xe1\x59\xb9\xf6\x37\x46\xec\x7e\xbf\x2f\xa7\x6f\x1c\x43\x39\x6a\x5d\xd5\x57\xda\xd5\x4c\x3f\xd4\xeb\x8e\xe5\x25\x28\x8d\xa6\x2c\x5e\xb7\xc8\xde\xdd\x1c\x59\xb0\x2a\x53\x3d\x3f\x3b\x0e\x92\x4f\x2c\x05\x55\x4e\xf8\x54\xb8\x2d\x56\x2b\x2b\x48\x99\x5d\x37\x77\xa9\xa0\x06\xad\xb6\x6b\xc1\x19\x5a\x31\xbb\xaa\xab\x48\x10\x70\x95\xe7\x82\xfa\xac\x94\x37\x07\x75\xd0\x43\xae\x64\x10\x7d\xa8\x3c\x43\x26\x61\xbb\x42\xe9\x7b\xb3\xd7\xac\x13\x9e\xbb\x51\x87\x07\x85\x70\xed\x0e\xd5\xb2\xec\xeb\xae\xae\x3f\x9c\x27\xa5\xde\x4f\x2d\xbd\x8f\xfb\x0c\xa7\x9e\x66\x93\x72\x42\xe6\x5a\x64\x78\xe0\x83\x6a\xe0\xf3\x73\xcb\xce\xa4\xd4\x30\x7b\x9e\x6f\xe8\xf1\xea\xf2\x0a\xfe\xd6\xd2\xba\xcf\x18\xb9\xcc\x81\xa0\x12\x84\x52\xb0\x24\xcd\x18\x4c\x21\xfd\xdf\x3f\xc3\x0c\x6b\x52\x67\xa5\x78\x05\x7f\xaf\xca\x36\xea\x84\xd4\x28\x0a\x7f\xd5\x53\x2b\x21\x9a\x79\xf3\x57\x6e\xb5\xb3\x45\xff\x0b\x00\x00\xff\xff\x40\x1a\x74\x1b\xb5\x17\x00\x00")

func apiProtoBytes() ([]byte, error) {
	return bindataRead(
		_apiProto,
		"api.proto",
	)
}

func apiProto() (*asset, error) {
	bytes, err := apiProtoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "api.proto", size: 6069, mode: os.FileMode(420), modTime: time.Unix(1536094325, 0)}
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
	"api.proto": apiProto,
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
	"api.proto": {apiProto, map[string]*bintree{}},
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

func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
