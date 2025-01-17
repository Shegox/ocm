// Code generated by go-bindata. (@generated) DO NOT EDIT.

// Package jsonscheme generated by go-bindata.
// sources:
// ../../../../../../../../resources/component-descriptor-ocm-v3-schema.yaml
package jsonscheme

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _ResourcesComponentDescriptorOcmV3SchemaYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x1a\xef\x6f\xdb\xb8\xf5\xbb\xfe\x8a\x87\x4b\x01\x39\x4d\x64\x37\x29\x3a\xe0\xfc\x25\xc8\x7a\x18\x50\x6c\x77\x39\xb4\xb7\x7d\x58\xea\x15\xb4\xf4\x64\xb3\x47\x91\x1e\x49\x39\x71\x7b\xfd\xdf\x07\x92\xa2\x7e\x59\xb2\x65\x3b\xed\x36\xdc\xe5\x4b\x4c\xea\xfd\xe2\xe3\xfb\x2d\x3d\xa3\xc9\x14\xc2\xa5\xd6\x2b\x35\x9d\x4c\x16\x44\x26\xc8\x51\x8e\x63\x26\xf2\x64\xa2\xe2\x25\x66\x44\x4d\x62\x91\xad\x04\x47\xae\xa3\x04\x55\x2c\xe9\x4a\x0b\x19\x89\x38\x8b\xd6\x2f\x09\x5b\x2d\xc9\x55\x18\x3c\x73\xb0\x35\x5a\x1f\x95\xe0\x91\xdb\x1d\x0b\xb9\x98\x24\x92\xa4\x7a\x72\xfd\xe2\xfa\x45\x74\x75\x5d\x90\x0e\x03\x4f\x90\x0a\x3e\x85\xf0\xee\xf5\x8f\xf0\xda\x33\x83\x1f\x4a\x66\xb0\x7e\x09\x15\x46\x4a\x39\x35\x08\x6a\x1a\x00\x64\xa8\x89\xf9\x0f\xa0\x37\x2b\x9c\x42\x28\xe6\x1f\x31\xd6\xa1\xdd\x6a\x52\x2f\x8f\x01\x6b\x94\x8a\x0a\x6e\x91\x13\xa2\x89\x83\x96\xf8\xef\x9c\x4a\x4c\x1c\x39\x80\x08\x42\x4e\x32\x0c\xab\x65\x81\xe7\x76\x48\x92\x58\x31\x08\xfb\x59\x8a\x15\x4a\x4d\x51\x4d\x21\x25\x4c\xa1\x7d\xbe\xaa\x76\x0b\x0a\x86\x9a\xff\x0d\xf0\x4c\x62\x3a\x85\xf0\x6c\x52\x3b\x51\xa5\xea\x9f\x6a\x9c\x0b\xb6\x7b\x50\x25\x32\xf2\x88\xc9\x3b\xcc\xd6\x28\x3d\x2a\x23\x73\x64\x6a\x0f\xa6\x03\xf2\x28\x2b\x29\xd6\x34\x41\xb9\x07\xc9\x83\x85\x41\xd0\x64\x53\xdc\x03\x91\x92\x6c\x1c\x4d\xaa\x31\x2b\x65\xe8\x97\x20\xf4\x84\x7a\xef\x73\xc0\x0d\x11\x96\x17\xeb\x7d\xfa\x77\xf4\x95\x96\x94\x2f\xbc\xa2\x0d\xf6\x14\x3e\x7f\xe9\x53\xfc\x8a\x68\x8d\xd2\x18\xd3\xbf\xd6\xf7\x2f\xa2\xef\x67\x17\xcf\x3c\x73\x45\x17\x9c\xe8\x5c\x6e\x71\x08\xe7\x42\x30\x24\x03\xac\x26\x00\x68\xdc\x7f\x43\x0f\x4e\x50\x47\x24\x23\x8f\x7f\x43\xbe\xd0\xcb\x29\x5c\xbf\x7a\x15\xb4\x24\xbb\x27\xd1\xa7\xd9\x7d\x44\xa2\x4f\x46\xc2\xe7\xa3\xfb\xf1\xac\xb5\x75\xfe\xdc\xef\x7d\xbe\xbe\xfc\x32\x9a\x34\x1e\x7f\xe8\x40\xf9\x60\x70\xce\xcd\x61\x03\x00\x9a\x20\xd7\x54\x6f\x6e\xb5\x96\x74\x9e\x6b\xfc\x2b\x6e\x9c\xa8\x19\xe5\xa5\x5c\x5d\x52\x19\xe6\xa3\xfb\xe8\xc3\x85\x17\xc4\x6f\x9e\xdf\x38\xd2\x0d\x1b\x76\x34\xcf\x40\x93\x5f\x91\x43\x2a\x45\x06\xca\x3e\x30\xf1\x04\x08\x4f\x80\x24\x1f\x73\xa5\x31\x01\x2d\x80\x30\x26\x1e\x80\x70\x10\x2b\xa7\x5f\x60\x48\x12\xca\x17\x10\xae\xc3\x4b\xc8\xc8\x47\x13\xb4\x38\xdb\x5c\x5a\x54\xbb\x1e\x67\x94\x17\xbb\x9e\xd7\x92\x2a\xc8\x90\x70\x05\x7a\x89\x90\x0a\x43\xd5\x10\x71\xea\x57\x40\x24\x1a\x56\xc6\x54\x68\xd2\x94\x57\x79\x81\xaf\xc6\xd7\xe3\x97\xf5\xdf\x51\x2a\xc4\xc5\x9c\xc8\x62\x6f\x5d\x07\x58\x77\x41\x5c\x8d\xaf\xfd\xaf\x12\xac\x06\x5f\xfe\x6c\xa0\xd5\x95\xbd\x9e\xdd\x8c\x5e\xfc\x76\x7f\x15\x7d\x3f\x7b\x9f\x3c\x3f\x1f\xdd\x4c\xdf\x8f\xeb\x1b\xe7\x37\xdd\x5b\xd1\x68\x74\x33\xad\x36\x7f\x7b\x9f\xd8\x3b\xba\x8d\xfe\x19\xcd\x8c\xc1\xfb\xdf\x9e\xe4\x40\xe0\x73\xcf\xf1\x62\x54\x7f\x70\x61\x89\x34\x76\x2c\x64\xe1\x54\x2d\xcb\xef\x32\xbd\xde\x50\x51\x78\xff\xc6\xf8\x91\x9a\xc2\xe7\xee\xb8\xd3\x65\xca\x21\x7c\x71\xa6\xb8\x12\x8a\x6a\x21\x37\xaf\x05\xd7\xf8\xa8\x0f\x89\x4a\x06\xaa\x2f\x0a\x59\x0a\xed\x18\x51\x3b\xa3\x88\xe9\xdb\x6e\xde\x84\xb1\xbb\xb4\xe2\xd2\x93\x05\x5a\xa8\x55\x70\x6c\xcb\x59\xc8\x3a\x27\x0a\xff\x2e\x59\x58\x05\xb9\x2d\x91\xcd\x5f\x01\x56\xdf\xea\x8c\x4d\xee\xaf\x11\xc7\x7e\x24\xab\x15\xe5\x8b\x81\xa8\x00\xc8\xf3\x6c\x0a\xf7\x61\x2e\xd9\xcf\x44\x2f\xc3\x4b\x08\xd5\x92\x5c\xbf\xfa\x53\x94\xd0\x05\x2a\x1d\xce\x82\x16\x9d\x43\x29\x5b\x1d\x2f\xa8\xd2\x72\x63\xa8\xdf\xbd\x7e\x53\x2e\x67\xe6\x0e\x48\x1c\xa3\x52\x03\xeb\x0a\xa3\x19\x0b\x05\xa9\x90\x05\x2a\x2a\x18\x99\x15\x3e\x6a\xe4\x26\x87\xa8\xf3\x3d\xc6\x12\x00\x2c\xa8\x5e\xe6\xf3\xdb\xdd\xbc\x77\x5a\x9b\x5d\x1a\x13\xa8\x5d\xa8\xdd\x49\x8f\xb2\xc6\xb6\xda\x9c\x80\xa5\xfa\x0b\x46\x7b\xd0\x8d\x95\xee\x86\x88\x45\x96\x51\xbd\xcb\x27\xb8\xe0\x78\x8a\x5e\x4e\x3c\xf7\x4f\x82\xa3\x33\x0c\x25\x72\x19\xe3\x0f\xa5\xc3\x1d\x20\x8e\xa9\x3e\xca\x45\x51\x59\x94\x6b\x43\xa1\x5c\x38\x13\x3a\xa0\x88\xd9\x12\x7c\x78\xb0\x2b\x50\xf0\x51\x4b\xf2\xa6\x00\xd8\x53\xf9\x6d\xd1\x79\x82\x3a\x75\xc0\x75\x1c\x51\xca\xd6\xdd\xd8\xae\xf9\xe6\x2e\x6d\x86\xbf\x4e\x2a\x0e\x2f\xdc\x0f\x58\xf7\xd8\x01\xe0\xa6\x37\xf2\xc0\x01\x80\x8b\x66\xef\x56\x18\x1f\x60\x46\x4b\xa2\x96\xb7\x6c\x21\x24\xd5\xcb\xac\x32\x2e\x21\x33\xc2\xa8\x22\x86\xd1\xf6\x63\x5b\xd8\x1e\xd9\xb5\x34\x18\xee\x2c\x9f\xbb\x85\x18\x50\x71\x77\x43\x04\xb5\xa2\xfa\x40\x25\x91\x1d\x1a\x30\xab\x0c\x13\x4a\x7e\xf1\x3e\x77\xb8\x4e\xc8\xc9\x87\x73\x5b\xa5\x1c\x15\x54\x33\xb7\xfc\xb2\x44\x07\xe4\x12\x8c\x48\x6d\x59\x5a\xaa\x05\x6a\xfd\xce\x4e\xfd\x1d\x1b\xa7\x9c\x89\x96\xcb\x92\xde\x91\x7a\xdb\xdb\x81\x39\x7e\x7b\x9c\xbc\xf2\x9b\x1d\xcd\x57\x27\x66\xc3\x9e\xac\x0f\x2a\x19\xbf\xf5\x09\x6a\x6f\xa6\x27\x26\x99\xa1\x44\x1e\xa3\x6d\x39\x60\x54\x8d\x46\x98\x88\x09\x3b\x2f\x12\x44\x5f\xd6\xf1\xa1\xf3\x1d\x32\x8c\xb5\xd8\xd7\x63\xf7\x46\xda\x83\x62\xa1\x2d\x66\x0b\xb1\x8f\x3d\x68\x79\xce\xa1\x8d\x78\xe7\x20\xe3\xf4\x11\x4a\x47\x7f\xdc\x7b\xfe\x4e\x11\x76\xa5\x4f\x38\x03\x12\xeb\x9c\x30\xb6\x99\x56\x9c\x22\xeb\x79\x0f\x13\x50\x2b\x8c\x29\x61\x20\xd1\xc0\xc7\x96\xc9\xff\x6f\xc6\x3d\x22\x9d\xb6\x9d\x53\x70\x6c\xa7\xd3\x42\xa1\x3c\x67\x6c\x40\x3e\xac\x3b\xb2\xb5\x52\xe7\x3d\x55\x40\x3c\xb0\xf6\xf6\x04\xd4\xa1\x03\x3d\x38\xb3\xf8\xd6\x87\x2b\x2a\x97\xc5\x38\x20\x57\x1a\x32\xa2\xe3\x65\xcd\x0d\xd4\x56\x09\xb7\x5d\x86\x33\x9b\x08\x6b\x5b\xf5\xba\xe2\x8f\xca\xae\x3c\x95\x8b\xc1\x6a\x0b\xaa\x36\x42\x84\xf6\x18\xb1\x57\x08\x47\xac\x6a\x3e\xdc\x25\x0c\x2e\xf5\xad\x09\x98\x9e\xd0\x74\x6e\x92\x13\x56\x76\x3b\xff\x8b\xf5\xa7\x88\xe9\x9f\x99\x18\x5e\x80\xda\xd3\xfd\x85\x32\x54\x1b\xa5\x31\x3b\x1c\xf7\xae\x8b\xe1\xd7\x8e\x0b\x22\xa6\x6f\x32\xb2\x38\xa9\x03\xb4\x4b\x6a\xa8\xbc\xf5\x99\xed\x49\x5a\xc3\xfa\x24\xc1\x5b\x4a\x93\xcd\x9e\x59\x4f\xa5\xce\x13\x0e\xc6\xc8\xc6\x7b\xdc\x69\xe7\x81\xb0\x10\x29\x84\xaa\xcb\x4f\xfb\xaa\xd3\x5b\x73\x80\x66\xa9\x60\xca\xd3\x8c\x70\x9a\xa2\xd2\xed\xba\xb4\xc5\xf4\xc8\xe2\xd7\x69\xc6\x85\x66\xe7\x28\x4e\x02\x05\x5a\xec\xe1\xd8\x36\xd4\x6d\x76\x0e\xc2\xb3\xd2\x44\x2e\x50\x63\x02\xb1\xe0\xba\x2c\x7e\x7a\xc9\x2b\xfa\x69\xe7\x59\xcc\x73\xa0\x1c\xe6\x1b\x8d\xca\xf3\x98\x1b\x65\xb7\xe9\xf2\x3c\x9b\xfb\x57\x2b\x7d\x2e\x7b\x82\xb9\xa4\x94\x61\x95\x09\x4f\xb5\x98\x0e\x09\x2b\xeb\xf1\xac\xfa\xf4\xe2\x9f\xd7\xd5\x01\x7a\x49\x34\x50\x65\xcf\x6e\xd4\x4f\xb9\x7d\xf6\x9d\x79\xa8\xbe\x83\x84\x4a\x5b\x3d\x6f\x7a\xef\xc3\xeb\xed\xee\x89\xfc\xeb\x2b\x28\xec\xae\xed\x67\xbb\x8d\xb3\x69\x98\xd6\xdf\xe1\x81\xea\x65\xa1\x9a\x38\x97\x12\xb9\xae\x0a\x14\xa8\x5e\xd5\xee\xd2\x92\x0f\xad\x6f\x8b\x9a\xe7\x94\x57\x6f\xf5\xca\xbe\x4b\x89\x7f\x54\x3f\xfb\x73\x89\xbd\x8c\xa7\x2c\x39\xfa\xca\x86\x5a\x42\xfd\x36\x69\x3c\x00\xa8\xc6\x5f\x27\xb8\x62\xee\x27\xdb\x27\x26\x6e\x23\x4c\xa9\xe8\x7c\xc7\x14\x3b\x00\x58\x20\x47\x49\xe3\xff\xe2\x04\xba\x90\xc0\x0d\xa1\x8b\xc5\xb7\xf6\xd9\xa7\x19\xf7\xfc\xce\x7c\xba\xba\x38\xb7\xff\xb5\x5c\xba\x61\xa2\xdf\xaa\x30\x6f\x7e\x2a\x72\xa8\x05\x7e\x15\x7b\x3a\x74\x32\xa6\x76\x0d\x96\x9b\x29\xd8\xce\x7f\x52\x1a\xdb\x86\xd2\x67\xe2\xa2\x32\x34\xcb\xda\x94\xcc\x9b\x97\x3e\xf6\xa4\xc5\x04\xe2\x89\x5a\xe2\xd6\x4b\xab\xda\x9b\x39\x57\xb8\x3f\x11\x1f\xd9\xec\xac\xaa\x81\xce\xe1\xf4\xb7\x3a\xe5\x1d\x2f\xbc\xab\xa1\x51\x38\x04\xa1\x5d\xf2\x0c\x42\x6a\x85\xdc\x30\x08\x5a\xe6\x52\xb7\x74\x13\x37\x57\xf4\x1f\x55\x6c\x8d\x20\xfc\x95\xf2\xa4\xf8\x59\xff\xea\x2c\x72\x66\x15\x06\x4d\x13\xa8\xd0\x1b\xb6\x59\x37\xf5\x5a\xc3\x96\x8d\x5b\x1f\xee\x95\xdf\xe5\x5d\xba\xc7\x4a\xa4\xfa\x81\x48\xac\x1e\xd8\xaa\xd3\xc8\xd4\x4b\x3f\x16\x5c\xe9\x29\x84\xe5\xf7\x78\xb5\xf3\xf8\x13\x38\xe4\x4e\x85\x19\x90\xb0\xeb\x33\x8a\x61\x5f\x89\xb5\xee\xbf\xff\x2a\xb7\x3e\x95\x08\xe1\xcc\x57\xc3\x6c\x73\x09\x0f\x08\x82\xb3\x4d\xf1\x79\x90\x6d\x1a\x05\xc7\x86\xe3\x77\xfb\x4c\xf1\x76\xa1\x7c\x63\x70\xc2\xd7\x6d\x25\x8d\xf0\x3f\x01\x00\x00\xff\xff\x02\xfb\x97\xa2\x6f\x29\x00\x00")

func ResourcesComponentDescriptorOcmV3SchemaYamlBytes() ([]byte, error) {
	return bindataRead(
		_ResourcesComponentDescriptorOcmV3SchemaYaml,
		"../../../../../../../../resources/component-descriptor-ocm-v3-schema.yaml",
	)
}

func ResourcesComponentDescriptorOcmV3SchemaYaml() (*asset, error) {
	bytes, err := ResourcesComponentDescriptorOcmV3SchemaYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../../../../../../../resources/component-descriptor-ocm-v3-schema.yaml", size: 10607, mode: os.FileMode(436), modTime: time.Unix(1668066250, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
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
	"../../../../../../../../resources/component-descriptor-ocm-v3-schema.yaml": ResourcesComponentDescriptorOcmV3SchemaYaml,
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
// AssetDir("foo.txt") and AssetDir("nonexistent") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"..": &bintree{nil, map[string]*bintree{
		"..": &bintree{nil, map[string]*bintree{
			"..": &bintree{nil, map[string]*bintree{
				"..": &bintree{nil, map[string]*bintree{
					"..": &bintree{nil, map[string]*bintree{
						"..": &bintree{nil, map[string]*bintree{
							"..": &bintree{nil, map[string]*bintree{
								"..": &bintree{nil, map[string]*bintree{
									"resources": &bintree{nil, map[string]*bintree{
										"component-descriptor-ocm-v3-schema.yaml": &bintree{ResourcesComponentDescriptorOcmV3SchemaYaml, map[string]*bintree{}},
									}},
								}},
							}},
						}},
					}},
				}},
			}},
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
