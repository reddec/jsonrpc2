// Code generated by go-bindata. (@generated) DO NOT EDIT.

//Package internal generated by go-bindata.// sources:
// template.gotemplate
// python.gotemplate
// js.gotemplate
// ts.gotemplate
// method_doc.gotemplate
package internal

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

var _templateGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x53\x5d\x4f\xdc\x30\x10\x7c\xf7\xaf\x18\xe5\x4e\xa2\x45\x24\xfd\x78\x3c\x81\x54\x04\xa5\xa5\xa2\xa5\x2d\xf4\x09\x21\xc5\x24\x26\xe7\x36\xe7\x44\xb6\xaf\x14\xd9\xfe\xef\xd5\xda\xc9\x1d\x77\x05\xa4\xe6\x25\xf6\x78\x67\xbd\x3b\xb3\x9e\xc0\xb9\xe2\x42\xe8\xdf\xb2\x12\xc5\x17\xbe\x10\x21\x30\xf6\x00\x3a\xea\x16\x0b\xa1\x6c\x42\x35\x57\x8d\xc0\x54\xaa\x5a\xfc\xd9\xc3\x74\x21\xec\xbc\xab\x31\x3b\x40\xf1\xc3\x88\xfa\x73\xdc\x9a\x10\xd8\x2e\xae\x9c\x9b\x16\x1f\x84\x12\x9a\xdb\x4e\x17\xdf\x96\xbc\x1d\xe3\x43\xb8\x7e\x31\x79\xfa\x18\x1e\x6d\x77\x27\x34\x3c\xb4\xe8\x5b\x5e\x09\x64\x45\x86\x2c\x83\x87\x51\xfc\x97\xa8\xb8\x11\x21\xbc\x44\x0e\xe7\x06\xce\x58\x25\x3c\x6e\xa5\x36\xf6\x4c\x2a\xea\xc3\xb9\x1c\x42\xd5\xff\x57\x3b\x9b\x90\x24\x4f\x17\x4f\xb9\xb6\xae\x25\x70\x17\x29\xc5\x0c\xe5\x73\xec\x92\xed\xe2\xbb\xb0\x4b\xad\x4c\x8a\x1c\x32\x25\xec\xf2\xbe\x17\x14\xe3\xdc\x9d\xb4\xf3\x91\x55\x1c\xea\x26\xa9\x7a\xa8\x9b\x25\x5d\x68\x66\x8c\x79\x7c\xed\x8c\xb4\xb2\x53\xf0\x20\xe3\xe0\x41\x7c\x78\xe6\xf3\xd5\xe7\x37\x7f\x51\x92\x2d\x25\xb8\x6e\xa2\x0c\x21\x30\x4f\x9d\x47\x3c\x04\xc4\x0d\xd7\xcd\x30\x15\xf0\xb1\x5e\x02\xce\x7b\x13\xc2\xb0\x1e\x4a\x86\x7f\xa0\xb6\x73\xe3\x3f\x87\xbc\xc5\xb4\x38\xee\xaa\xc3\xba\xd6\x21\xb0\xb2\x2c\x6f\xb8\x99\xb3\x6a\xa9\x5b\xe4\x1f\xb1\x73\xd4\x29\x2b\x94\xcd\x29\xcf\x0c\xbc\xef\x5b\x59\x71\xea\xea\xd5\x4f\xd3\xa9\x1d\xe4\x79\xcd\x2d\xcf\x6f\xa4\xe2\xfa\x1e\xef\x72\x64\x24\xef\x2a\x61\x86\xfd\xfd\xf7\xe7\x27\xcc\x31\x00\xc8\x88\xa3\xfb\x2a\xc3\x0c\xd9\xdb\xe2\x75\xb6\x97\x60\x59\x13\xf2\x66\xd8\x25\x55\x63\xcc\x73\x56\x8d\xe4\x9e\x6b\xbe\x30\x14\x7e\x75\xcd\x02\xa3\xdb\xca\xb2\xdc\x1c\xae\x7f\x44\x6d\xbb\x8a\xb7\xd1\x8e\xd9\xc1\xca\xc7\xb3\x11\x34\x98\x16\xa7\x8b\xbe\xd3\x34\x3a\x93\x34\x70\x2b\xc6\x20\x29\x73\x8e\xb4\x5b\xc3\xa7\xe6\xc2\xea\x65\x65\xa3\x4f\x9f\x4c\xf4\x3d\x19\x8e\xd5\xf8\x33\xff\x98\xed\xdb\xce\xaf\x73\xa6\x8c\x27\x52\xb4\x71\xf6\xc9\xf2\xe2\x92\x37\xa3\xdb\x6b\x77\xe9\xe0\x91\x47\x06\x8a\xdc\x70\x3e\x87\x68\xe9\x7d\x92\x46\x4d\xc7\x2c\x15\xb8\xd1\xdd\xb1\xb8\x95\x2a\xce\x6d\x0a\xda\x66\x27\x49\xd7\xeb\xbf\x01\x00\x00\xff\xff\x1d\x67\xdc\x81\xa1\x04\x00\x00")

func templateGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_templateGotemplate,
		"template.gotemplate",
	)
}

func templateGotemplate() (*asset, error) {
	bytes, err := templateGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template.gotemplate", size: 1185, mode: os.FileMode(420), modTime: time.Unix(1584625238, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pythonGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\x41\x8f\xd3\x30\x10\x85\xef\xf9\x15\xa3\xa8\x52\x5a\x11\x2c\xe0\x58\x29\x87\x15\x20\x4e\x20\x21\xc4\x69\x55\x45\x26\x9e\x2d\x86\xc4\x0e\x33\x0e\x2a\xb2\xfc\xdf\x91\x5d\x27\xdb\xa4\x97\xed\xa5\xb1\xdf\x7b\x33\xf6\xe7\xd1\xc3\x68\xc9\x01\xe1\x9f\x09\xd9\x71\x51\x74\xbd\x64\x06\xef\xc5\x37\xa4\xbf\xba\x43\xf1\x45\x0e\x18\xc2\xb1\x00\x00\x28\xcb\x32\xfd\xdf\xc8\xef\xed\x30\xa0\x71\x21\x2c\x86\xf4\xa1\xf0\x09\xda\x56\x1b\xed\xda\x76\xcf\xd8\x3f\xd5\xf0\x43\x32\xb6\x13\xf5\xd0\x40\xe5\xbd\xf8\x60\xbb\x07\xa5\x28\x84\xea\x70\x2d\x1e\x7f\xd1\x29\xda\xec\x9a\x03\x5b\x55\x2b\x68\xe0\xed\x6d\x1f\x83\x17\xd7\x6a\x95\x1a\xdd\x57\xd3\x0a\x5e\xc5\xc0\xbc\x4d\xe8\x26\x32\xcf\x6a\x51\x78\x4f\xd2\x9c\x11\x76\xda\x28\xbc\xd4\xb0\x1b\xd0\xfd\xb4\x0a\x8e\x0d\x88\xef\x8c\xea\x73\x5a\x72\xbe\x64\xec\xe9\x7d\xf6\x64\x3e\xa9\xb5\xf7\xaf\x61\x53\x48\xd2\x39\x56\x99\xcd\x0f\x74\xe6\x10\xea\x18\x97\x74\xce\x59\xef\xd1\xa8\x10\x6e\x0e\x3e\x73\xbe\xb2\x9e\xc3\x6b\xd4\x5b\x1f\x21\x8f\xd6\x30\x42\xb3\xbc\xa6\x18\x2d\xbb\xfd\x33\xd3\x1a\x7e\xb1\x35\x8d\x5f\x32\xa9\x46\xdc\xa3\xb1\x2b\xe1\x08\xe5\x3b\xf1\xa6\xac\xd7\xf2\xb5\x79\x52\xbd\xdf\x89\x4f\x68\x90\xa4\xb3\x24\xbe\x4e\xb2\x9f\x2f\x16\xc2\x36\xa6\x53\x24\xf7\x9e\xdf\xe7\xb0\x31\x8d\x92\xe4\xc0\xd1\xf8\xf8\x52\x76\x2b\x72\x11\x64\x62\x77\x5a\xd5\x0d\x87\x65\x29\x99\x31\xcd\xf7\x15\x8e\xb0\xbf\x17\x69\x94\xff\x7a\x2b\x55\xe2\x95\xd5\x88\x62\x7f\x17\xae\x90\xc8\x52\x05\xc6\x3a\xd0\x66\xce\xd5\xf0\xf1\xd2\xe1\xe8\xb4\x35\xfb\xbc\xf5\x98\x9d\xa7\xc3\x76\xd6\x16\x03\x21\x4f\xbd\xab\x4e\x71\xe6\xd2\xc9\xff\x07\x00\x00\xff\xff\x21\x46\x5d\x4b\x81\x03\x00\x00")

func pythonGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_pythonGotemplate,
		"python.gotemplate",
	)
}

func pythonGotemplate() (*asset, error) {
	bytes, err := pythonGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "python.gotemplate", size: 897, mode: os.FileMode(420), modTime: time.Unix(1581586275, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jsGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x55\x4d\x6f\xdb\x38\x10\xbd\xeb\x57\xcc\x1a\x01\x24\x39\x8e\x9c\xdd\xa3\x03\x1f\x82\xec\x62\x3f\x80\x4d\x5c\x24\x3d\x15\x45\xc0\x88\x63\x9b\xad\x4c\xaa\x43\x2a\xb1\xa1\xf2\xbf\x17\xa4\x28\xeb\xc3\x4e\x6f\xd5\x45\x10\xe7\xbd\x99\xa7\xc7\xe1\x10\xf7\xa5\x22\x03\x79\xc1\xb4\x86\xba\xce\x1e\x91\x5e\x45\x8e\xd9\x3d\xdb\xa1\xb5\x7f\x11\x29\x02\xdc\x1b\x94\x5c\x43\xf3\x55\x47\x00\x00\xb9\x92\xda\x50\x95\x1b\x45\xc9\x0e\xb5\x66\x1b\x9c\x41\xae\x38\xce\x80\xa3\x61\xa2\xd0\x69\x40\xba\x47\x57\x25\x52\xe2\xc2\x70\x09\xf1\x02\x62\xb8\x84\xc0\x4a\x6f\x8e\x28\xb3\x15\x3a\xf3\xa0\xa5\x4f\x35\x8a\x84\xbc\xb0\x6c\x2b\x34\x71\x1b\xd9\x28\xfa\xf9\x6f\x04\x25\xf3\xe9\xd4\xbf\x7b\xf1\x3b\xb5\xdb\xa1\x34\xd6\xfa\xc0\x74\x3a\x8f\x1a\xe4\x1c\xee\x08\x99\x41\x90\xf8\x06\xb7\xab\x7f\x61\xcb\x24\x2f\x90\xc0\xa8\xd3\xf4\x59\xcb\x29\x09\xd7\x85\xd8\x6c\xcd\x3f\x01\x9d\x88\x35\x70\x5c\x0b\x89\x3c\x85\x9c\x49\x20\x34\x15\x49\x28\x49\xed\x84\xc6\x13\x27\x5f\x98\xc6\xe7\x8a\x0a\x58\x42\x5c\xd7\xd9\x9f\x2a\xbf\xe5\x9c\xac\x8d\x67\xa7\xb9\x97\x20\xab\xa2\xe8\xbb\xec\x5d\x7a\x0e\xfc\x36\xd5\xcd\x38\x2c\x38\x2c\xe1\xf7\x93\xe5\x33\xf9\xc7\x4b\xad\xdf\x51\x5d\x13\x93\x1b\x84\x0b\x21\x39\xee\x67\x70\xb1\x43\xb3\x55\x1c\x16\x4b\xc8\x3e\x6a\xe4\xff\xfb\x4f\x1d\x5c\xed\x6c\x0f\xb8\x33\xae\xbb\x37\xd3\x07\x99\xf7\x50\xce\x5c\xf8\x0e\x95\x34\xc2\x14\x68\x6d\x52\xd7\x57\x30\x2a\xcc\x68\xe3\xaa\xb6\x94\x5b\xda\x68\x6b\xeb\x5a\xac\x61\x63\x02\x0a\xae\xad\x9d\x41\x5d\xa3\xe4\x2e\xe4\x28\x61\xdf\xc2\x5a\xda\x39\x18\xb6\x27\x61\x6f\x4c\x98\xd6\x99\x9c\x15\x45\x12\x0f\x75\xf9\x2d\xe9\x78\xee\x99\x7c\xd1\x4a\x52\x99\x4f\x60\x01\x93\x3f\xb2\xeb\xc9\x6c\x18\x6e\xd8\x3e\x5a\xd7\x17\xd9\xdf\x28\x91\x98\x51\x94\x7d\xa8\x58\xd1\xfe\x81\xb5\x63\x9a\xf0\x94\x20\x45\xe2\xde\x3c\x0b\x9e\xa4\x23\x50\xc9\x88\xed\xb4\x03\x7e\xfa\x45\x26\x7d\x3e\x16\xb4\x69\xda\x36\x42\x88\x45\xcd\xa1\xe9\xc9\x3b\x6d\x4a\xc1\xe1\x72\xd0\x76\xc1\xea\x63\xb8\xed\xad\xae\x13\x82\xf3\x8d\xea\x19\x10\x7e\xeb\xe7\xf5\xc7\x06\xd6\x68\xf2\xed\xca\xff\x3d\x2c\x47\x1b\xd2\x10\x17\x30\x59\x3d\x3c\x3e\x8d\x6c\xdd\x22\xe3\x48\x7a\x31\xa2\xb8\x27\xbe\x53\xd2\xa0\x34\x57\x4f\x87\x12\x63\x58\x40\xcc\xca\xb2\x10\x39\x33\x42\xc9\xb9\xdb\xe4\x78\x98\xcb\x0e\x3f\x5f\x14\x3f\x2c\xe0\xbf\xc7\x87\xfb\x4c\x1b\x12\x72\x23\xd6\x87\xc4\x89\xef\x1c\xec\x5c\x10\x6b\x48\xde\x39\x80\xe9\x48\x5b\xd3\x93\xab\x66\x72\x64\x84\x5a\x15\xaf\xf8\x1e\xf9\xe8\x5a\xcf\xa0\xb4\x37\x6a\xed\xc8\x47\x42\xe7\x5f\x53\xc2\x53\x92\x6e\x98\x0c\x93\x0c\xb5\xff\x46\xa8\x33\xf5\x75\xac\xd5\x6c\x49\xbd\xf9\xd9\xe9\xef\x8c\xc4\xa1\xb4\x61\xa6\xd2\x6e\xfe\xfb\xf1\xdf\x2d\x3d\xe1\xde\x0c\xa4\x8d\xb4\x71\x66\xd8\x51\x9c\xa3\xb9\x3d\x48\xd2\x9b\x68\xa0\x24\x46\x57\x29\x06\x21\x3d\xe1\x7d\x45\xe7\x2f\xb9\xc4\x91\x32\x9f\x23\x3b\xde\x67\xbd\xb5\x70\xb5\x75\x0b\xbe\xc8\x59\xd5\xa1\xb3\x3d\x96\x50\x57\x85\x39\x5e\x54\x3f\x02\x00\x00\xff\xff\xf5\x5c\xbf\xbd\x6e\x07\x00\x00")

func jsGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_jsGotemplate,
		"js.gotemplate",
	)
}

func jsGotemplate() (*asset, error) {
	bytes, err := jsGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js.gotemplate", size: 1902, mode: os.FileMode(420), modTime: time.Unix(1586348282, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tsGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\x5f\x8f\xdc\xb6\x11\x7f\xdf\x4f\x31\x5d\x18\x58\xc9\x5e\xeb\xdc\x3e\xea\xbc\x2e\x0c\xd7\x29\x52\x34\x8e\x9b\xbb\x20\x0f\x87\xc3\x85\x27\x8e\x76\x99\xe3\x92\x0a\x49\xdd\xdd\x42\xd1\x77\x2f\xf8\x47\xff\x28\x6d\x6c\x04\x01\xda\x7d\xd2\x8a\xc3\xe1\xcc\x6f\x7e\x33\x9c\x11\x3e\x57\x52\x19\x28\x38\xd1\x1a\x9a\x26\xbb\x42\xf5\xc8\x0a\xcc\x3e\x91\x23\xb6\xed\x47\xa5\xa4\x02\x7c\x36\x28\xa8\x06\xff\xaf\x59\x01\x00\x54\xf5\x3d\x67\x05\x28\x24\x54\x0a\x7e\x82\x42\x52\xcc\x41\xd4\xc7\x7b\x54\x97\x8b\x12\x14\x0d\x61\x5c\xe7\x40\xc4\xe9\x72\xe5\x44\x0a\x29\xb4\x51\x75\x61\xa4\x4a\x8e\xa8\x35\xd9\x63\x0e\xda\x28\x26\xf6\xdb\x89\xc6\xed\x64\x77\x1a\x6c\xb0\x3f\x5d\x57\xa8\x12\x2b\x0b\xaf\x60\x93\xc3\x06\x5e\x41\x50\x95\x5e\xf6\x52\xe6\xc0\x74\xe6\x84\x76\x4e\x6f\xb4\x12\x94\xc3\xae\x3b\xc6\xaf\xb7\xab\x76\xb5\x6a\x1a\x45\xc4\x1e\x81\x62\xc9\x04\x33\x4c\x0a\xdd\xb6\xab\x00\x1b\x13\x06\x55\x49\x0a\xb4\xd0\x5d\x9f\x2a\xf4\xb0\x05\xfb\x9a\xe6\x35\xf8\xcd\xd9\x95\x73\xf3\x1b\x86\x9c\xda\xed\x7e\x35\xbb\x26\xfb\xb6\xcd\xed\x13\xfc\x06\xe6\x54\xa1\x2e\x14\xab\x4c\x2f\xf0\x1a\x50\xd0\xd6\x9a\xd1\x34\xfe\x69\x75\x71\x61\x5d\x76\x87\x6b\x53\x97\xe5\x6a\xb5\x1a\x8c\x50\x55\xf1\xf1\x19\x8b\xda\xf4\x61\x2a\x08\xe7\x09\xa3\x03\x8e\x15\x39\x71\x49\x68\x07\x73\x9a\xc3\x67\x25\x8f\x4c\xe3\x5b\x79\xff\x0b\x16\xe6\xdd\xa5\x75\xda\xd3\xe1\x49\x47\xda\x2a\xc5\x1e\x89\x41\xd0\xb2\x78\x40\xf3\xf7\x1c\x7e\xc2\xfb\x2b\xf7\x7c\x39\x59\x2f\xa4\x10\x58\x18\x26\xf6\xb0\x83\x92\x70\x8d\xd3\xf5\x9e\x13\x15\x0a\xca\xc4\xfe\x43\x90\x97\x22\x87\xf7\x4a\x91\xd3\xdb\x24\x85\xdd\x3b\x48\x1e\x25\xa3\xe9\x3b\xd8\xc1\xcd\xed\x19\x0d\x85\x54\x0a\x39\xb1\x7b\x61\x07\x02\x9f\xe0\x3b\x52\xbd\xed\xbc\xbd\x49\x28\x31\x24\x07\xef\x9b\xd3\x69\x55\x6e\x21\x41\xa5\x66\xaf\x6f\xdf\x25\xe9\x02\x33\x67\x67\xd6\x8a\xf7\xf8\x05\x64\x5a\xbf\x8d\xe8\x93\x28\xfe\x00\xe8\x23\x42\xbb\x93\x1d\x82\xb0\x03\xf2\x44\x98\xe9\xd8\xeb\x30\xfa\xb6\xfc\x84\x48\x91\x26\x23\x76\xb3\x12\x92\x20\xd3\x83\x91\x1d\x88\x4e\x18\x4d\xc7\xb9\xe2\xf9\xae\xe4\x93\xc3\xc9\x65\x73\xf2\x33\xe1\xd6\xaf\x13\xe0\x33\xd3\x46\x77\x11\x01\x85\xbf\xd6\xa8\x0d\x3c\x31\x73\x00\x46\xe1\x45\xc3\x68\xfb\xf3\xe8\xd0\xb6\x7f\xe2\x68\xa0\xac\x4d\xad\x30\x44\x20\x72\x2e\x49\x14\x6a\xc9\x1f\x71\x0b\x0a\x7b\xc0\x63\xb3\x22\xf3\x35\x9a\xc4\x06\xea\x26\xda\x7b\x3b\x36\x61\xf4\x6c\xe1\xc9\x34\x0a\x9a\x04\xb0\x47\x6b\x0a\x4d\xad\x04\x24\x1e\x4d\x6f\x6a\x7a\x39\x8e\x5b\x17\xe2\x10\xbf\x18\xea\x21\x5e\x3d\xe1\xc7\x0e\x3c\x1d\x18\xc7\x3e\x04\x1d\xf5\x63\xe4\xfd\xe9\x23\x7c\x7a\x5c\xe6\x70\xf4\x90\xcc\x12\x24\xab\x6a\x7d\xe8\x36\x8e\x7c\xf4\x78\xa4\x0b\xe1\xe9\xd9\xe1\xb3\x36\x36\x2b\x80\x33\x92\x58\x8a\x71\xe4\x1b\xec\xc0\xa8\x7a\x54\x46\x2d\x07\xe2\xdd\x46\x9d\xa2\xb3\xbc\xc4\x22\xaf\xc7\x7c\x6e\xa1\x64\x82\x70\x1e\x6f\x9f\x1b\x31\xaa\x2d\x53\x7b\xfd\x49\x99\x14\xe8\xee\xac\x1d\x24\xe7\x48\x27\xc5\x00\xee\x37\x84\xf1\x69\x6a\x2d\x68\x2c\xb8\xd4\xf8\xa7\x6a\x0c\xf7\x95\xd5\xd9\xd8\x7a\xd5\x2e\x68\xb6\x00\x2b\xd4\xd3\x80\xcf\x11\xf6\x11\xb5\x37\xd9\xbf\xae\xbe\xff\x94\x55\x44\x69\x74\x35\x30\xa6\x0a\x14\xc4\x14\x07\x48\x30\xe6\x03\x84\x1a\x24\x39\x66\x0e\xbc\x64\x5d\x3a\x27\xc0\x29\xeb\x0a\x43\xbe\xde\xc2\x8c\x7f\xab\x58\x89\x01\x43\xf4\x83\x65\x4b\x9c\xdf\x7b\x34\x96\xc5\x19\xa3\x91\x12\x47\x57\xa2\x1f\x96\xec\x9a\x69\xa1\xc8\xd1\xe0\xb2\x22\xb7\x81\xe8\x87\x9b\x37\xb7\x56\xe0\xac\xad\x11\xc9\x7b\x8e\x76\x74\x9e\x70\xbc\xa8\x3a\x67\x66\x99\x19\xf5\x14\xf3\xcc\xd5\x9c\x15\x98\xbc\xd9\xc2\x9b\x71\xdd\xaa\xb2\x52\xaa\x8f\xa4\x38\x24\x49\xe9\xe2\x5e\x26\xe9\xbc\x76\xcd\xd2\x33\x2a\x5c\x7d\x16\x7d\xa1\x54\x05\x75\xe3\x2a\x3d\xc8\x25\x5f\x53\xa9\x87\x4c\x0f\xd5\xbe\xdf\xef\xab\x4c\xad\x78\x04\x75\xa0\xae\xd5\x4b\xe7\x39\x0b\xe3\x4c\x90\x15\x8a\x33\xa9\xe5\xad\xef\xb5\x4c\xcb\x4f\x24\x90\x84\x52\x17\x87\x7c\xf9\xd0\xbe\x44\xe0\x99\x63\x2d\x25\xff\xd2\x9d\xbd\xc4\x4b\x7f\xb4\x05\x2c\x89\x73\xe2\xab\x4d\x6f\xbf\xca\xd6\xbe\xf8\xfc\x5f\xd9\xda\x3f\xa5\xcb\xd7\xea\x52\x45\x1c\x99\xd6\x51\xaa\x4b\xad\xd9\x45\x32\xc9\xcb\x5a\xb8\x56\x1c\xe9\xb4\x01\xb2\xeb\xb1\xbf\xf6\x5d\xe6\x20\x5b\xae\xc0\xa1\xd7\xb2\x39\xed\xda\xce\xac\x54\xf2\x38\x6f\xa5\x1e\x09\xaf\x51\x4f\xd2\x72\x26\x53\x70\x24\x2a\x99\x36\x24\xda\x00\x2a\x15\xd2\xc4\xb7\x5b\x9b\xa2\x07\x02\x9c\x61\x74\x73\xa6\x18\x24\x37\x77\x43\xcf\x33\x8f\x75\x17\x42\xa5\x26\x0d\x51\x3a\x0c\x2e\xbe\x87\xaf\xa4\x36\xf1\x4c\xf0\x3f\xee\x6d\x4b\x34\xc5\xe1\x33\x51\xe4\x68\xef\xa9\xa9\x5b\x47\x34\x07\x49\x73\x58\x7f\xfe\xfe\xea\x7a\xbd\x9d\xac\x1d\x90\x50\x54\x3a\x5f\x20\xf5\xe6\x83\x14\x06\x85\x79\x6d\x67\x30\x3b\x05\x92\xaa\xe2\xac\x70\xa1\xb9\xf8\x45\x4b\xb1\x99\xaa\x6a\xa7\x7f\xef\x25\x3d\xe5\x9d\x27\x03\x9a\x71\x30\xfd\xc5\x1a\xba\x48\xeb\x44\x5f\xf1\xb6\x63\xa7\xa2\xce\xdc\x66\x63\x26\x67\xdc\x8c\x1b\x71\x2b\xa5\x0d\x31\xb5\xb6\x83\xac\x9b\x63\x87\x57\xd7\xf8\x6c\x16\x19\x1c\xea\xb9\xb7\xc9\xca\x5b\x67\x93\x74\x34\xbf\x5e\xbc\x7c\xb9\x1a\xcd\xf5\x1f\xe4\xf1\x88\xc2\x8e\x98\x2f\x5f\x5e\xac\x7e\x7f\xfa\x87\x66\x9a\xc6\x77\x77\x43\xe0\x2f\xa3\x15\x0c\x1c\xcb\x47\x33\xe8\xe5\xca\xef\xbf\xb8\x80\x0f\x0a\xad\x9c\x75\xf7\xfd\xe7\x6f\xe1\x40\x04\xe5\xa8\xc0\xc8\xf9\xa9\xd9\x8c\xa4\xf7\x44\xe3\x5d\xad\x38\x74\x24\x83\x1d\x6c\x9a\x26\xfb\x87\x2c\xde\x53\xaa\xda\x76\x93\xce\x48\x56\x29\x69\xa4\x2d\x94\xf6\xc8\x1f\x7f\xf8\x77\xaf\x24\x4d\x33\xb7\x56\x48\x3e\x00\xaa\x9f\x98\xeb\x86\xdc\x4a\x1c\xa9\x82\x68\x84\xf5\x93\x5e\xe7\x8b\xaf\xf5\x3a\x3f\xd3\xa9\x0c\xa8\xec\xac\x15\xc3\x38\x3d\x18\x33\xaf\xad\xf7\x0a\xc9\xc3\xef\xb5\x55\xee\xd8\x83\x31\xd5\xa2\x3d\x76\x21\xb6\x94\x62\x49\x6a\x6e\xbe\x6c\x66\x28\x55\xe3\x9a\xf1\x47\x4d\x8d\xba\x2a\x4b\x1d\xd8\xc1\x5f\xfb\xeb\xa1\xfb\xac\xf2\x82\x09\x8a\xcf\x5b\x78\xe1\x33\x1f\xf2\x1d\x64\x3f\x6a\xa4\xdf\xb9\xbf\xdd\xb7\x12\x4b\x62\x70\x9f\x44\x82\xdc\x88\xc7\xf6\xbd\xe5\x32\xf4\x35\x6a\x90\xb2\x8c\x82\xdf\xa0\x16\x86\x19\x8e\x6d\x9b\x0c\x9f\x64\xfa\x83\x89\xda\xdb\x53\xbb\x2d\xef\xd5\x5e\xb7\x6d\xd3\xb0\x12\xf6\x26\x48\xc1\x9b\xb6\xdd\x42\xf8\x06\xd3\x34\x76\x4b\x20\x6b\x0e\xfe\x6f\xf4\xfd\x26\x88\x8e\x8a\xe1\x60\xd3\x0f\x3e\x65\xa7\x1b\x16\x7a\xb4\x64\x34\x24\xdd\xdd\xb9\xba\x3b\x0d\xe0\xda\xe6\xba\xaa\x8a\x35\xe4\xb0\xfe\x5b\xf6\x26\x2a\x97\x6b\x7f\x9e\x5b\x6d\x9a\x17\xd9\x3f\x51\xa0\x22\x46\xaa\xec\x3f\x35\xe1\x9d\xbf\x6d\x1b\x6f\x63\x6e\x4b\x38\x56\xe0\xb3\xb9\x63\x34\x49\x23\xa1\xca\x55\x3a\x2b\x78\xf3\x67\x43\x1a\xde\xdd\x8e\x6f\x35\x20\x1a\xbe\x80\x60\xc7\xac\xfe\x4b\x19\x4c\xea\x53\xef\xc8\x08\xe8\x81\x99\xaf\x7a\x6a\x42\xd4\x73\xdb\xe5\xf3\x5f\x0a\x42\x5c\x14\xfe\xda\x7f\x4c\x1a\x42\x4e\xc4\x69\x7e\xf9\xd9\x59\x6c\x3a\x00\x0f\xe9\x97\x75\xca\x32\x46\xb7\x6e\x7e\xf3\xe5\x8e\x95\x27\xfb\x36\x4d\x57\x93\x8b\x65\xe3\x1a\xd7\x0d\x30\xe1\xb4\x9e\xbf\x60\xae\x99\x28\x7e\xc2\x7b\x7f\xcf\x58\x51\x3f\xd8\x65\x61\xee\xdc\xc2\xe8\x5d\x21\xe9\xf4\x45\x34\x3c\x8e\x1a\xd3\x00\x93\x93\x55\xa8\x6b\xde\x8f\x26\xed\x7f\x03\x00\x00\xff\xff\x2d\xc2\x8a\x35\x56\x16\x00\x00")

func tsGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_tsGotemplate,
		"ts.gotemplate",
	)
}

func tsGotemplate() (*asset, error) {
	bytes, err := tsGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ts.gotemplate", size: 5718, mode: os.FileMode(420), modTime: time.Unix(1586615821, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _method_docGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\x4d\x6e\x32\x31\x0c\x86\xf7\x39\x85\x25\x58\xa1\x8f\x1c\x00\xe9\x5b\xa0\x56\x45\x54\xf4\x9f\x03\x4c\x0a\x26\x44\xca\x24\x28\xc9\xa8\xad\x92\xdc\xbd\xca\xcf\xcc\x74\x50\x17\x65\x13\xfc\x8e\xfd\xda\x7e\x3c\x03\xef\xe7\x74\x23\xf5\x3b\x93\x74\x83\x0a\x0d\x73\xda\xd0\x97\x8e\x49\xa0\x0f\xe8\xce\xfa\x18\x23\x21\xde\xd7\x80\xde\xe8\xb6\x45\xe5\x92\xb8\x80\xa2\xad\xa0\xf9\x83\x49\x43\x16\xf0\x8a\xae\x33\xca\xe6\x82\xde\xb0\x68\xfb\xaf\x0b\xa6\x1c\xef\x3f\x84\x3b\xf7\x55\x74\x6d\xb8\x8d\x91\x2c\x60\x6d\x78\x97\xfa\xda\x15\x21\x01\x9e\xb5\x15\x4e\x68\x05\x01\x1e\x59\x8b\x10\x20\xd5\x43\x20\x61\x39\xfc\xc2\xf4\x21\xde\x2f\xc1\x30\xc5\x11\xe6\x42\x1d\xf1\xf3\x1f\xcc\x99\xe1\xb0\xfa\x0f\x34\x46\x12\x12\x87\xac\xc7\x08\x39\x60\x86\xd3\x64\x9e\xe3\xa6\x0a\x4f\x17\x1b\x63\xfd\x5f\x47\x86\x62\x8d\x2a\x81\xf2\xbe\xbc\xbf\x75\x93\xfa\xc0\x64\x9e\x33\xf5\xac\x0b\xee\x7a\xd1\xc2\x40\x70\xdb\x5e\xb4\x49\x84\x67\xb3\x7c\x9e\xa1\xb0\xb6\x24\xde\x8b\xd3\x0f\x3f\xba\xb5\x6f\xce\x74\x07\x97\xf7\xb8\xb7\x99\x4b\x01\x02\xf5\x5c\x23\x9a\xe9\x73\x4d\x66\xf4\x2c\x8e\x77\x02\xe5\xd1\x56\x3e\x74\xcf\x78\x4f\x63\xdc\x3e\x7d\x18\xba\xc0\x49\x18\xeb\x76\x42\x21\xa4\xcc\x09\x99\x25\xa0\xb4\x69\xfc\xa6\x69\xb8\x26\x2e\x0d\x38\xd9\xee\x16\x4f\x42\xe5\xbb\x96\xa4\xeb\xea\x42\xf6\x3b\x00\x00\xff\xff\xb0\x28\xe0\x93\xb2\x02\x00\x00")

func method_docGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_method_docGotemplate,
		"method_doc.gotemplate",
	)
}

func method_docGotemplate() (*asset, error) {
	bytes, err := method_docGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "method_doc.gotemplate", size: 690, mode: os.FileMode(420), modTime: time.Unix(1586438099, 0)}
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
	"template.gotemplate":   templateGotemplate,
	"python.gotemplate":     pythonGotemplate,
	"js.gotemplate":         jsGotemplate,
	"ts.gotemplate":         tsGotemplate,
	"method_doc.gotemplate": method_docGotemplate,
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
	"js.gotemplate":         &bintree{jsGotemplate, map[string]*bintree{}},
	"method_doc.gotemplate": &bintree{method_docGotemplate, map[string]*bintree{}},
	"python.gotemplate":     &bintree{pythonGotemplate, map[string]*bintree{}},
	"template.gotemplate":   &bintree{templateGotemplate, map[string]*bintree{}},
	"ts.gotemplate":         &bintree{tsGotemplate, map[string]*bintree{}},
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
