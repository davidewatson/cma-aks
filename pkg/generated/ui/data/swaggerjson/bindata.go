// Code generated by go-bindata.
// sources:
// assets/generated/swagger/api.swagger.json
// DO NOT EDIT!

package swaggerjson

import (
	"github.com/elazarl/go-bindata-assetfs"
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

var _apiSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5d\x5f\x73\xdb\x36\x12\x7f\xcf\xa7\xd8\xe1\xdd\xcc\xa5\x33\x8e\x9d\xe6\x5e\x6e\xfc\x74\x1a\xdb\x97\x6a\xd2\x38\x1e\xcb\x69\x1f\xae\x1d\xcd\x8a\x5c\x49\xa8\x41\x80\x05\x40\x39\xca\x8d\xbf\xfb\x0d\x00\x52\x22\x29\x92\xa2\xe9\xc4\x66\x13\xf5\xa1\x96\x48\x60\xb1\xff\xf0\xc3\xee\x02\x50\xfe\xf7\x02\x20\xd0\x77\xb8\x58\x90\x0a\x4e\x21\x78\x73\xfc\x3a\x38\xb2\xcf\x98\x98\xcb\xe0\x14\xec\x7b\x80\xc0\x30\xc3\xc9\xbe\x3f\xe3\xa9\x36\xa4\xe0\x3d\x0a\x5c\x90\x82\xd1\xbb\x09\xfc\x44\x3c\xb1\x1f\xaf\xc6\xae\x2b\x40\xb0\x22\xa5\x99\x14\xb6\xc3\xea\xf5\xf1\x8f\x19\x4d\x80\x20\x94\xc2\x60\x68\x36\x84\x01\x02\x81\xb1\xa3\xfc\x9e\x85\x4b\x24\x0e\xbf\x90\xa0\xcf\x0c\xb3\x1e\x00\x41\xaa\xb8\x7d\xbf\x34\x26\xd1\xa7\x27\x27\x0b\x66\x96\xe9\xec\x38\x94\xf1\x89\xc6\x58\xa7\x62\xf1\x2a\x14\xa1\x39\x09\x63\x7c\x85\xb7\x7a\xdb\x8f\x62\x64\xae\x67\xbc\xf2\x24\xff\xbd\xb0\x4f\x6c\xcf\xc0\xb5\xb9\x7f\x01\x70\xef\x84\xd5\xe1\x92\x62\xd2\xc1\x29\xfc\xd7\xf3\xe9\x06\xcb\x99\xb6\x5f\x6c\x8f\xdf\x5d\xdb\x50\x0a\x9d\x96\x1a\x63\x92\x70\x16\xa2\x61\x52\x9c\xfc\xa1\xa5\xd8\xb6\x4d\x94\x8c\xd2\xb0\x63\x5b\x34\x4b\xbd\xd5\xf8\x09\x26\xec\x64\xf5\xe3\x49\xe8\x15\x5e\xd4\xd8\x82\x8a\x0a\xb4\xec\xa7\x71\x8c\x6a\x6d\x65\xfd\x95\x71\x0e\x8a\x8c\x62\xb4\x22\x30\x4b\x02\x6d\xd0\xa4\x1a\xe4\x1c\x10\x32\x62\x80\x22\x02\x66\x34\xdc\xa6\x33\x0a\xa5\x98\xb3\x05\xcc\xa5\x82\x50\x0a\x41\xa1\x61\x2b\x66\xd6\x1b\x3d\x02\x04\x32\x21\xe5\x58\x1e\x47\x76\x8c\xb7\x64\x32\x37\x28\x36\x52\xa4\x13\x29\x34\xe9\x12\x6f\x00\xc1\x9b\xd7\xaf\x2b\x8f\x00\x82\x88\x74\xa8\x58\x62\x32\x2f\x29\x10\xf2\x12\x59\x83\xe0\x4e\x37\x80\xe0\xef\x8a\xe6\xb6\xc7\xdf\x4e\x22\x9a\x33\xc1\x2c\x05\x6d\x8d\x8f\xb7\x7a\xcb\xd8\x35\x25\x7c\x1d\x94\xfa\xde\xbf\xa8\xfb\x7c\x5f\x90\x20\x41\x85\x31\x19\x52\x5b\x7b\xf9\xff\x2a\xbc\xe7\x1e\xeb\xfe\x1e\xb5\xca\x75\x89\x31\x59\xd5\x5b\x43\xe4\xca\x37\x12\x66\x04\x5c\xca\x5b\x8a\x20\x4d\x8e\xab\x24\x98\xeb\xf9\x67\x4a\x6a\x5d\x7d\xa5\xe8\xcf\x94\x29\xb2\x56\x98\x23\xd7\x54\x79\x6d\xd6\x89\x63\x4c\x1b\xc5\xc4\xa2\x28\xfe\xfd\xd1\x7e\x71\x42\x45\x11\x09\xc3\x90\xeb\x63\x4c\x92\x29\x8b\xf6\x08\x77\xb3\x24\x18\x25\xc9\x38\x72\xce\x33\xba\x1a\xc3\x28\x0c\x49\xeb\x41\x4a\x64\x48\xa0\x30\x1d\x24\xba\x71\x0d\x37\x22\xe1\x70\x45\x4a\x50\xeb\x3b\xa9\xba\x98\xe9\x2a\x6b\xfa\x57\x10\x4b\xa7\xb3\x0d\xff\xdd\x9c\x70\x52\xe8\xf1\x2c\x12\x6e\x3e\xff\x5e\x80\x13\x83\x8b\x2a\x90\xe4\xcb\xe7\xb6\xf3\xef\x2f\x2a\x4a\x0a\x22\xe2\x64\xa8\x1d\xe0\x7d\x9b\x2d\xa0\xb7\x80\xf5\xb9\x6b\x3a\x4c\xbc\x2e\xf1\x36\x14\xc8\xfe\x75\x89\x06\x98\x2e\x42\xf6\x3f\x34\xd8\x8e\x16\xb9\x23\xd2\x46\xc9\xf5\x20\x27\xce\x01\xb4\xff\x02\xe8\x76\x00\xed\x6f\x11\xb4\x93\x74\x4f\x48\x9e\x28\xb9\x62\x36\x27\xea\x04\xda\x67\x8a\x70\xa8\xa0\x5d\xe2\xed\x49\x40\x7b\x26\xa3\x1d\xb3\x7b\x8f\xa8\x7b\x53\x70\x08\xa3\xd2\xaa\x3f\x7c\x01\x99\xdf\xeb\x45\x17\x89\xfb\x7b\xd5\x8b\x82\xc2\xaa\x79\xe0\x09\xa6\x46\xea\x10\xb9\xf5\xe6\x42\x4e\xb8\xd7\x01\x49\xe0\x8c\x6f\x93\x90\x22\x99\x66\x3f\xbc\x70\x9d\x32\x4e\x47\xf5\x5d\x06\xe0\x92\x4d\x6c\x7e\x27\xde\xd9\x24\xfe\xf3\x3a\x2a\x67\xda\xf4\xab\x5a\x20\xd8\xbe\x36\x6d\xce\x68\xe9\x4e\xc5\x88\x9f\xed\x80\xc3\x72\xcc\x32\x73\x4f\xe2\x8e\x87\x80\xf0\x10\x10\x0e\x41\xac\xef\x33\x20\x6c\x45\x44\x0b\xcb\xd4\x0f\x12\x85\x8c\x08\x42\x99\x96\x7c\xbf\x19\x0a\x2f\x65\x44\x67\xd5\xd6\x83\xc2\xc3\x0d\x87\x43\x49\xfb\x0f\x95\xda\xa1\x48\x74\xc0\xf8\x03\xc6\x0f\x19\xe3\xbb\x27\xfd\x0e\xf0\xb7\x09\xbf\xee\x06\xe3\x13\xdb\x6b\x98\xb9\x7f\x91\xb5\xef\x24\xb9\x2a\x8a\xfc\xbc\x09\x55\x9a\x2c\x14\x46\x3d\x03\x08\x5c\x21\xe3\xae\x02\x90\x91\x29\xef\x08\x77\x8a\x2a\x3e\x66\x3d\x07\xe6\x94\xbb\x0c\x1e\x62\x8a\x43\x4c\x71\x88\x29\x0e\x31\xc5\xb7\x18\x53\x64\xf0\xdd\x09\xba\x33\x3c\x1c\x66\x2c\x51\x66\xee\x3b\x89\x26\xca\x42\x3f\x4f\x3c\xb1\x3d\x9c\xf7\xa0\x38\x22\x55\x02\xb2\xae\xc0\xc4\x5c\xaa\xd8\x79\x1a\xe0\x4c\xa6\x06\x30\x61\xa0\x49\xad\xf6\x85\x12\xbf\x78\x0a\xe3\x2d\x81\x81\xb9\xe5\x96\xc7\x5e\x2e\xd9\xc7\x3e\x9b\x63\x88\x05\x6e\xb6\x07\x01\x4b\x7b\x4f\xa3\x77\x93\x49\x42\xe1\xe8\xdd\x64\x2c\xb4\x41\x11\xd2\x5b\x25\xd3\xa4\x68\xcb\x1c\xa8\xe4\xec\x0f\x0a\xb7\x4b\x5e\x90\x28\x6b\x0d\xc3\x2a\xca\xcd\x27\x47\x49\xdd\x15\xb0\x3b\x2a\xbd\xcb\x8f\x82\x5a\xec\x15\x85\x08\x67\xe1\x38\xa9\xd7\x8a\xa7\xf7\xf0\x31\x72\x29\xc1\xb6\x86\x97\x13\x83\x22\x42\x15\x4d\xcf\xdf\x4c\x57\x6f\x8e\x80\x4c\x78\xfc\x43\xfd\x90\x31\x13\xd3\x3f\x53\x14\x86\x99\x75\xd3\xd0\x4c\x18\x5a\x94\x3c\x16\x20\xf0\x8e\x99\xbd\xfe\xe7\x9b\x06\xc6\xde\x33\xc1\xe2\x34\x06\x91\xc6\x33\x52\x56\x05\x2c\x63\x55\xc3\xcb\x88\xe6\x98\x72\xa3\x6d\xb0\xf7\x99\x94\x2c\xb2\xb8\x83\xf8\xbb\xb2\x3a\x45\x6a\x70\xde\x60\x81\xde\xc9\x8e\x22\xaa\x1b\x2c\x28\x4d\xf1\x3a\x5f\xb9\xf1\x1e\xd9\xdb\x3f\x6e\xa9\x51\x7f\xed\xee\x81\x0b\xb0\x7d\x6b\xad\xb3\x42\x9e\xf6\xf4\x3a\x5c\x80\xef\xdd\x45\xa7\x56\x78\x40\xe5\x1d\xf5\xc4\xf5\x83\x04\x99\xd2\x60\x96\x68\xf2\x4d\xd1\xb5\x4c\xad\xa9\x42\x34\xb4\x90\x8a\x7d\x26\x50\xa4\x65\xaa\xac\x31\xad\xde\x57\x8c\xee\x20\x94\x42\x4b\xce\x22\x34\x14\xc1\x8c\x71\xbe\x09\x03\x72\xed\xb7\xec\xc3\x8d\x3e\xa7\x8a\x0a\x8f\x2e\x65\xf4\x8c\x33\x37\xcf\x4b\x5c\x39\xa2\x65\xda\x3e\xf9\x1c\xb2\x0c\x69\x17\xc3\x15\xf7\xa7\xeb\x79\xc3\x4f\x5f\x95\x37\xfc\xf4\x40\xde\x6a\xd7\xdc\xca\x6a\x52\xb3\xfa\x3d\xc2\xfe\x0b\x66\xa6\xbb\x8b\x79\x77\x37\xb0\x00\x6e\x70\x01\x52\x78\xfc\x66\x06\x14\x25\x52\x33\x23\x55\xc3\xbc\xb5\x43\x86\x32\x8e\x99\xe9\x3d\xe2\x12\xf5\x72\xb3\x64\x30\x03\x19\xb9\xc6\xe1\x8c\x22\x9a\x6a\x83\xa6\x9f\xaf\xff\xba\x24\xb3\xb4\x16\x54\x20\xa4\x71\xa3\x5a\x8a\x70\x87\x1a\x42\x4e\x28\xe0\x6e\x49\x02\x66\x29\xe3\x0d\x4c\xd8\x57\xd1\x34\xea\xcb\xc0\x39\x1a\x37\xd9\x1c\x99\x06\x31\xe5\xa3\xec\x98\x79\x95\x1d\x64\x21\x21\xd5\x14\x39\x34\x93\x71\xc2\x38\xd5\x8f\x98\xbd\x54\xbd\xc6\x3b\xcb\x3a\xbb\xa1\xea\xe9\x27\x1c\x8d\xf5\xf1\x5e\xf4\xaf\xb2\xce\xc0\x8c\x37\x93\x1f\xcf\xa7\xaf\x27\xa0\x52\x21\x98\xb0\x6e\xbb\x6f\xf6\xf9\x88\xce\x61\x6f\x86\xca\x13\x52\x2b\x16\xd2\x28\xf4\x45\xd8\x47\x4c\xbe\x90\x33\x12\xc6\x66\xa8\x7d\x27\xc2\x99\xa3\x30\x8e\xe0\x25\xde\xe2\xa9\xab\xa5\x9c\x37\x04\x33\xd9\x60\x9a\x42\x45\xfd\x67\x9e\x1f\x70\xe2\x88\x64\x83\xe6\x45\x84\x6e\x11\x8a\x9d\x3c\xe8\x55\xe7\xdd\x6c\xb6\xae\xd6\xb6\x42\x17\x81\x00\x5a\x9d\x17\x56\xd1\x97\xf4\xe9\x14\xb8\xc4\x08\x66\xc8\x6d\xe0\xa2\x7e\x08\x1a\x4d\xb5\xad\x06\x3c\xc6\x40\x59\x11\xab\xaf\xb2\xea\x4a\x5b\x0d\xc1\xad\xaf\x2d\xf5\x1d\xa8\xb6\xe2\xd4\x30\xab\xf2\x92\x4f\xdf\xb1\x1a\x0a\x41\xf5\xa3\x55\x2b\x31\x7d\x07\x6d\xa9\xcf\x74\x8b\xe1\xac\x87\x6d\x7d\xc2\x7a\x59\xaa\xc9\x5f\x58\xb2\xde\x66\xb1\xa0\xe0\x85\x75\x7e\x95\xcd\xfe\x73\x32\xc8\xf8\xd8\x50\xfc\x18\xc7\xea\xa9\x8a\xf1\x79\xa5\x14\x5c\xaf\xf5\x47\x07\x75\x7b\x47\xf0\x17\xc2\xa6\x31\x69\x8d\x8b\x7e\x63\x8d\xa2\xc8\xa5\xa9\xc8\x6b\x4a\x01\xe5\x4b\x67\x7b\xd9\xd9\xde\x41\xeb\xb9\xbe\x6f\x8f\xeb\x17\xae\xb3\xb9\x05\xd0\xdd\x66\xb3\x1f\x3b\xea\xa4\xca\x40\xeb\x91\xd0\x6c\x41\xf1\x1d\x9b\xbd\x7f\x9f\x26\x5a\x16\xae\x6c\x88\x83\xbf\x0e\xca\x5f\x87\xec\x2a\x93\x2a\x6f\x4d\xea\x09\x48\xa4\x71\xa9\x4e\x15\x4c\x6e\x46\x37\x1f\x27\xd3\x8f\x97\x93\xab\x8b\xb3\xf1\x7f\xc6\x17\xe7\xc5\xca\xdc\xd5\xf5\x87\x5f\xc6\x93\xf1\x87\xcb\xf1\xe5\xdb\xe2\xf3\xeb\x8f\x97\x3b\x8f\x2e\xce\x3e\x5c\x9e\x8d\x7f\xae\x3c\x9e\xdc\x7c\xb8\xba\xaa\x3c\xbb\xb8\xbe\xfe\x70\x5d\x7c\x70\x7e\xf1\xf6\x7a\x74\x7e\x71\x9e\x4b\xbd\x29\x7d\x06\x59\x4d\xc5\x6d\x90\x37\x73\xba\x55\xe8\x2b\xd8\x6d\x76\x0a\x97\xd2\x80\x26\xf3\x9b\x80\x57\x50\x14\xe9\x14\xdc\xf2\x58\x78\xe2\xac\x41\xc0\x44\xc4\x42\x34\x54\xba\x0f\x64\xf1\x66\x46\x76\xdd\xf1\xe1\x4e\x74\xec\x08\x66\xba\xf0\xb4\xb2\x2f\xad\x64\x96\x68\xe9\x90\xc8\xc9\xf8\x7b\xb9\x1a\xe6\x29\xe7\x6b\x48\x35\xce\x38\x65\xa4\xb7\x3a\xcd\xc8\x6f\x1f\xd4\x0c\x81\x06\xb4\x8c\x09\xee\xa4\xba\xb5\x04\x31\x34\x6c\x45\x7c\x9d\x71\x1d\x49\x41\x79\xd6\x97\xf1\x72\x04\x3a\x0d\x97\x80\x3a\xdb\x60\xc8\x17\xd5\x18\x1d\xa7\x2e\x7b\x8a\x08\xb4\x9c\x9b\x3b\x54\x19\x57\xb9\x49\x3d\x4b\xf9\xb7\x6e\x9a\xf3\x57\xd8\x22\x47\xc7\xb9\x81\x27\xe2\x3e\xb6\x52\x88\xd1\x8a\x01\xa9\xf0\xfa\x71\x04\x72\xb7\xf1\x34\xf2\x6f\xad\x64\xb2\xea\xbe\xb6\xb1\x84\x72\x0a\xb2\x0a\x91\x36\x66\x35\x52\x91\xb3\x01\xcc\x53\x11\x7a\xdc\x60\x66\x5d\x1b\x58\xd4\x14\xdc\x1e\x83\xd5\x5c\x86\xd5\xb2\x40\x77\x94\x73\x61\xab\x0b\xbc\xcf\xd1\x20\x9c\x91\x68\x04\xb1\xb0\x36\xc6\xde\x8b\x64\x3b\xf1\x79\x43\x82\x58\x8e\xd6\x5c\xe6\xbb\x1f\x58\xb3\x97\xbb\xb9\x59\x47\xb6\x6a\x33\xbc\x06\x06\x33\x27\xd0\xbe\x71\x39\xa9\xb1\xf1\x02\xf2\x5b\xfb\xd7\xa7\x31\x2e\x79\xd9\x84\x9a\x75\x59\x4c\x55\x94\xbc\x3c\x3b\xf5\xa5\xdc\x26\x6b\xa2\x52\x58\xde\x76\x0a\x98\xa1\xb8\xda\xbe\x51\xfa\x4e\x5b\x03\x05\x42\xf7\x7b\x8a\xec\x19\xb7\xf5\x59\x4e\xb9\x8c\xfc\xd5\xe5\x70\x65\xeb\xfd\xbc\x97\x9b\xb5\x2e\x91\xd5\x3b\x3c\xcf\x59\x79\xdd\x3d\x11\xb2\xb9\x20\xd6\x58\x52\xb1\x0d\xa2\xdd\x92\x4d\xe7\x7b\x4b\x57\x19\x01\x87\x51\xcd\x08\x92\x8f\x03\x3a\xa1\x90\xcd\xb3\x1f\xab\xe8\xa1\xe3\xd2\x78\xcf\xa1\xec\x62\x6a\xb0\x91\xea\x15\x70\x76\x4b\x80\xb7\x0d\x7e\x7e\xfb\x2f\xfd\xe8\xa2\xea\x6a\x5b\x90\x7b\x97\xce\x48\x09\x32\xd4\x30\x9c\x43\x98\xde\x26\xcd\x57\x9c\x96\xf5\xe0\xdd\xa4\xd1\x90\x05\x36\x96\x6c\xb1\x9c\x66\x67\xc4\x18\x6f\x29\xa9\xcf\xa4\xe4\x84\xa2\xa9\xa4\x5e\xfb\xba\xad\x16\x5b\x88\x0b\x7e\x1a\x35\x64\x18\x64\x6c\x1c\x33\x9d\xe3\x4c\xb1\xb0\xb7\x51\x7c\xf7\x6c\xb6\x55\x2a\x97\x5d\x5d\xda\x6f\x09\x3f\xc2\x97\xe5\xed\x53\x2b\xf6\x0e\x75\x11\x5b\x7c\xc9\x8e\x69\x17\xff\x90\x6e\x28\x79\xef\xfe\xe8\xcc\x7e\xbf\x2c\x24\xab\xcd\x46\x88\x5c\x05\xa6\x9a\xe3\xe4\xcc\x40\xbe\xff\xdf\xcd\x34\x35\x3f\x1f\x30\x3c\xd3\x9c\xc9\xb4\x1c\xfc\x58\xf7\xcb\x7e\x45\xa0\x09\xea\xeb\x53\xcc\x4e\x9e\x3e\x29\xe5\x90\xbb\x46\x6e\x51\x67\xdb\x35\xc2\x81\xac\x95\x0d\x00\x21\x23\xfa\xba\x81\x56\xaf\x9d\xdd\xfd\xc1\xcb\x5b\xbf\xd1\x6f\x03\xcd\x28\x2a\x6e\x2b\x82\x14\xbf\x89\xb3\x54\x29\x12\x86\xaf\x41\x0a\xf7\x3f\x2a\x6c\xd8\xba\xac\x8e\x73\x79\x47\x11\x30\x61\x61\xfe\x08\xee\x18\xe7\x3e\xe5\xbb\x63\x66\x09\xb1\x4d\x63\xcc\x12\x05\xfc\x08\x24\x8c\x5a\x03\xc7\x61\x24\x04\xbe\xf8\xfb\xf0\x2a\x47\xfb\x3d\xdf\xe1\xcd\xfd\x0a\x2c\xd7\xdc\xbd\x76\xf0\xec\x4f\x21\xb4\x41\x73\x8b\x4e\xea\xae\x98\x0e\x4f\x13\xe3\x4a\x15\xc0\xd7\x1e\xf4\x5a\xdb\xe5\xa2\x6d\x09\xfa\x6a\x53\x7a\x77\xd5\xda\x3f\x5f\x7f\xae\x5e\x4b\x7e\xa0\x85\x2a\x97\xde\xbe\x05\x33\xf5\x46\x77\xb7\x97\xb8\xb9\x94\xe0\x80\x2d\x91\x92\xbb\x83\x3a\x4d\x7b\xd6\x35\x65\x81\x2f\x70\xd4\xe3\x66\x97\x13\x3f\xd4\xae\x75\x77\xcb\x8c\xce\x92\x1e\x6f\x9b\x69\x34\xf9\xc3\xb7\xe3\x06\xfd\x03\xc6\xc2\x9e\xdc\x03\x67\x53\xf9\xb6\xc7\xb7\xa0\xc5\xfc\x6a\xce\xd7\xc5\xbc\x4c\x6f\x5d\x8a\x2b\xcb\xba\x8b\x43\xfd\xe7\x45\x99\xc2\x8e\x5d\x4b\xe7\x6e\x07\x68\xcf\x39\x6c\xce\x33\xbb\x75\xfb\xc3\xbb\x86\xc3\x8d\x5e\x8e\x29\xab\x3d\xe6\xd5\x62\x9f\xfd\xa7\xc5\xea\x39\xcb\x0f\x00\x15\x5b\x76\x37\xd2\x5c\xaa\x4d\xc5\xa2\x14\x7a\x94\x0d\x54\xbd\x08\xf7\x9c\x29\x41\x7b\x1d\xf9\x29\x43\x59\x4d\xa8\xc2\x25\xe8\xea\x19\x87\x9a\x10\xa1\x92\xb1\x4c\xed\x72\xd7\x4b\x0b\xa2\x78\x7c\xd3\x51\x79\xd2\xf5\xd2\x48\x83\xdc\xa6\xb0\x4c\x51\xe4\x8f\x45\x76\x83\xee\xdd\xdb\xa3\xc3\x9b\xe4\x77\x4b\xc2\xb6\x52\x8a\xbb\xd2\x3b\xa8\xb4\xfd\xe3\xee\xb5\xd0\x87\xea\xb2\x5b\xbd\xb3\x0b\xa2\xec\x60\x7d\x8e\x2c\x75\x88\x52\xbe\x0c\xf3\x05\x76\xcf\x06\xb3\xab\xb5\xdf\x5c\x03\x01\xd2\xdd\x7d\x88\xcc\x6e\x5f\x70\x13\xa2\x2c\xf1\x13\xec\x42\xb4\x0c\x78\xd8\x86\xf8\x42\xdb\x10\xf5\x73\xb7\xdf\x3e\x44\x67\x53\x0e\x75\xd1\xd8\x53\x7e\xcf\xa1\xf0\xaf\x5d\x7a\xdf\xdc\x52\xa3\x4f\x86\x94\x40\x7e\x2e\xc3\xc2\x35\xb5\xca\x65\xbc\xf7\x52\x51\x76\xbe\xaa\xdb\x3f\x19\xf0\xc0\x5f\xf9\xb7\xdc\xbc\xb8\x7f\xf1\xff\x00\x00\x00\xff\xff\xf6\x3b\xe8\x24\xbf\x60\x00\x00")

func apiSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_apiSwaggerJson,
		"api.swagger.json",
	)
}

func apiSwaggerJson() (*asset, error) {
	bytes, err := apiSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "api.swagger.json", size: 24767, mode: os.FileMode(420), modTime: time.Unix(1539119179, 0)}
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
	"api.swagger.json": apiSwaggerJson,
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
	"api.swagger.json": &bintree{apiSwaggerJson, map[string]*bintree{}},
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
