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

var _tsGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\x5f\x6f\xdc\x36\x12\x7f\xdf\x4f\x31\xb7\x08\x60\x29\xd9\xc8\xbe\x7b\x94\xb3\x0e\x82\x20\x3d\xf4\x70\x4d\x73\xb5\x8b\x3c\x18\x86\x4b\x8b\xa3\x5d\xd6\x32\xa9\x92\x94\xed\x85\xa2\xef\x7e\xe0\x1f\x49\x14\xa5\x4d\x82\xa2\xc0\xdd\x3e\x69\xc9\xe1\x70\xe6\x37\xbf\x19\x0e\x89\xcf\xb5\x90\x1a\x8a\x8a\x28\x05\x6d\x9b\x5d\xa2\x7c\x64\x05\x66\x1f\xc9\x03\x76\xdd\x07\x29\x85\x04\x7c\xd6\xc8\xa9\x02\xf7\xaf\x5d\x01\x00\xd4\xcd\x5d\xc5\x0a\x90\x48\xa8\xe0\xd5\x01\x0a\x41\x31\x07\xde\x3c\xdc\xa1\x3c\x5f\x94\xa0\xa8\x09\xab\x54\x0e\x84\x1f\xce\x57\x56\xa4\x10\x5c\x69\xd9\x14\x5a\xc8\xe4\x01\x95\x22\x3b\xcc\x41\x69\xc9\xf8\x6e\x33\xd1\xb8\x99\xac\x4e\xbd\x0d\xe6\xa7\x9a\x1a\x65\x62\x64\xe1\x15\x9c\xe4\x70\x02\xaf\xc0\xab\x4a\xcf\x07\x29\xbd\x67\x2a\xb3\x42\x5b\xab\x37\x9a\xf1\xca\x61\xdb\x6f\xe3\xe6\xbb\x55\xb7\x5a\xb5\xad\x24\x7c\x87\x40\xb1\x64\x9c\x69\x26\xb8\xea\xba\x95\x87\x8d\x71\x8d\xb2\x24\x05\x1a\xe8\xae\x0e\x35\x3a\xd8\xbc\x7d\x6d\xfb\x1a\xdc\xe2\xec\xd2\xba\xf9\x03\xc3\x8a\x9a\xe5\x6e\x36\xbb\x22\xbb\xae\xcb\xcd\x17\x7c\x01\x7d\xa8\x51\x15\x92\xd5\x7a\x10\x78\x0d\xc8\x69\x67\xcc\x68\x5b\xf7\xb5\x3a\x3d\x35\x2e\xdb\xcd\x95\x6e\xca\x72\xb5\x5a\x8d\x46\xc8\xba\xf8\xf0\x8c\x45\xa3\x87\x30\x15\xa4\xaa\x12\x46\x47\x1c\x6b\x72\xa8\x04\xa1\x3d\xcc\x69\x0e\x9f\xa4\x78\x60\x0a\xdf\x88\xbb\xdf\xb1\xd0\x17\xe7\xc6\x69\x47\x87\x27\x15\x69\xab\x25\x7b\x24\x1a\x41\x89\xe2\x1e\xf5\xdb\x1c\x3e\xe3\xdd\xa5\xfd\x3e\x9f\xcc\x17\x82\x73\x2c\x34\xe3\x3b\xd8\x42\x49\x2a\x85\xd3\xf9\x81\x13\x35\x72\xca\xf8\xee\xbd\x97\x17\x3c\x87\x77\x52\x92\xc3\x9b\x24\x85\xed\x05\x24\x8f\x82\xd1\xf4\x02\xb6\x70\x7d\x73\x44\x43\x21\xa4\xc4\x8a\x98\xb5\xb0\x05\x8e\x4f\xf0\x13\xa9\xdf\xf4\xde\x5e\x27\x94\x68\x92\x83\xf3\xcd\xea\x34\x2a\x37\x90\xa0\x94\xb3\xe1\x9b\x8b\x24\x5d\x60\xe6\x6c\xcf\x46\x56\x03\x7e\x1e\x99\xce\x2d\x23\xea\xc0\x8b\x3f\x01\x7a\x40\x68\xbb\xb3\x45\x10\xb6\x40\x9e\x08\xd3\x3d\x7b\x2d\x46\x3f\x96\x1f\x11\x29\xd2\x24\x60\x37\x2b\x21\xf1\x32\x03\x18\xd9\x9e\xa8\x84\xd1\x34\xcc\x15\xc7\x77\x29\x9e\x2c\x4e\x36\x9b\x93\xdf\x48\x65\xfc\x3a\x00\x3e\x33\xa5\x55\x1f\x11\x90\xf8\x47\x83\x4a\xc3\x13\xd3\x7b\x60\x14\x5e\xb4\x8c\x76\xbf\x05\x9b\x76\xc3\x57\x85\x1a\xca\x46\x37\x12\x7d\x04\x22\xe7\x92\x44\xa2\x12\xd5\x23\x6e\x40\xe2\x00\x78\x6c\x56\x64\xbe\x42\x9d\x98\x40\x5d\x47\x6b\x6f\x42\x13\x82\x6f\x03\x4f\xa6\x90\xd3\xc4\x83\x1d\xcc\x49\xd4\x8d\xe4\x90\x38\x34\x9d\xa9\xe9\x79\x18\xb7\x3e\xc4\x3e\x7e\x31\xd4\x63\xbc\x06\xc2\x87\x0e\x3c\xed\x59\x85\x43\x08\x7a\xea\xc7\xc8\xbb\xdd\x03\x7c\x06\x5c\xe6\x70\x0c\x90\xcc\x12\x24\xab\x1b\xb5\xef\x17\x06\x3e\x3a\x3c\xd2\x85\xf0\x0c\xec\x70\x59\x1b\x9b\xe5\xc1\x09\x24\x96\x62\x1c\xf9\x06\x5b\xd0\xb2\x09\xca\xa8\xe1\x40\xbc\x5a\xcb\x43\xb4\x97\x93\x58\xe4\x75\xc8\xe7\x0e\x4a\xc6\x49\x55\xc5\xcb\xe7\x46\x04\xb5\x65\x6a\xaf\xdb\x29\x13\x1c\xed\x99\xb5\x85\xe4\x18\xe9\x04\x1f\xc1\xfd\x81\xb0\x6a\x9a\x5a\x0b\x1a\x8b\x4a\x28\xfc\x4b\x35\xfa\xf3\xca\xe8\x6c\x4d\xbd\xea\x16\x34\x1b\x80\x25\xaa\x69\xc0\xe7\x08\xbb\x88\x9a\x93\xec\x5f\x97\x3f\x7f\xcc\x6a\x22\x15\xda\x1a\x18\x53\x05\x0a\xa2\x8b\x3d\x24\x18\xf3\x01\x7c\x0d\x12\x15\x66\x16\xbc\x64\x5d\x5a\x27\xc0\x2a\xeb\x0b\x43\xbe\xde\xc0\x8c\x7f\xab\x58\x89\x06\x4d\xd4\xbd\x61\x4b\x9c\xdf\x3b\xd4\x86\xc5\x19\xa3\x91\x12\x4b\x57\xa2\xee\x97\xec\x9a\x69\xa1\x58\xa1\xc6\x65\x45\x76\x01\x51\xf7\xd7\x67\x37\x46\xe0\xa8\xad\x11\xc9\x07\x8e\xf6\x74\x9e\x70\xbc\xa8\x7b\x67\x66\x99\x19\xf5\x14\xf3\xcc\x55\x15\x2b\x30\x39\xdb\xc0\x59\x58\xb7\xea\xac\x14\xf2\x03\x29\xf6\x49\x52\xda\xb8\x97\x49\x3a\xaf\x5d\xb3\xf4\x8c\x0a\xd7\x90\x45\xdf\x28\x55\x5e\x5d\x58\xa5\x47\xb9\xe4\x7b\x2a\xf5\x98\xe9\xbe\xda\x0f\xeb\x5d\x95\x69\x64\x15\x41\xed\xa9\x6b\xf4\xd2\x79\xce\x42\x98\x09\xa2\x46\x7e\x24\xb5\x9c\xf5\x83\x96\x69\xf9\x89\x04\x12\x5f\xea\xe2\x90\x2f\x6f\x3a\x94\x08\x3c\xb2\xad\xa1\xe4\xdf\xfa\xbd\x97\x78\xe9\xb6\x36\x80\x25\x71\x4e\x7c\xb7\xe9\xdd\x77\xd9\x3a\x14\x9f\xff\x2b\x5b\x87\xaf\x74\xf9\x58\x5d\xaa\x88\x81\x69\x3d\xa5\xfa\xd4\x9a\x1d\x24\x93\xbc\x6c\xb8\x6d\xc5\x91\x4e\x1b\x20\x33\x1f\xfb\x6b\xc6\x32\x0b\xd9\x72\x05\xf6\xbd\x96\xc9\x69\xdb\x76\x66\xa5\x14\x0f\xf3\x56\xea\x91\x54\x0d\xaa\x49\x5a\xce\x64\x8a\x0a\x89\x4c\xa6\x0d\x89\xd2\x80\x52\xfa\x34\x71\xed\xd6\x49\x31\x00\x01\xd6\x30\x7a\x72\xa4\x18\x24\xd7\xb7\x63\xcf\x33\x8f\x75\x1f\x42\x29\x27\x0d\x51\x3a\x5e\x5c\x5c\x0f\x5f\x0b\xa5\xe3\x3b\xc1\xff\xb8\xb7\x2d\x51\x17\xfb\x4f\x44\x92\x07\x73\x4e\x4d\xdd\x7a\x40\xbd\x17\x34\x87\xf5\xa7\x9f\x2f\xaf\xd6\x9b\xc9\xdc\x1e\x09\x45\xa9\xf2\x05\x52\x9f\xbc\x17\x5c\x23\xd7\xaf\xcd\x1d\xcc\xdc\x02\x49\x5d\x57\xac\xb0\xa1\x39\xfd\x5d\x09\x7e\x32\x55\xd5\x4d\xff\xde\x09\x7a\xc8\x7b\x4f\x46\x34\xe3\x60\xba\x83\xd5\x77\x91\xc6\x89\xa1\xe2\x6d\x42\xa7\xa2\xce\xdc\x64\x63\x26\x66\xdc\x8c\x1b\x71\x23\xa5\x34\xd1\x8d\x32\x17\x59\x7b\x8f\x1d\x87\xae\xf0\x59\x2f\x32\xd8\xd7\x73\x67\x93\x91\x37\xce\x26\x69\x70\x7f\x3d\x7d\xf9\x72\x15\xdc\xeb\xdf\x8b\x87\x07\xe4\xe6\x8a\xf9\xf2\xe5\xe9\xea\xeb\xb7\x7f\x68\xa7\x69\x7c\x7b\x3b\x06\xfe\x3c\x9a\x41\xcf\xb1\x3c\xb8\x83\x9e\xaf\xdc\xfa\xd3\x53\x78\x2f\xd1\xc8\x19\x77\xdf\x7d\xfa\x11\xf6\x84\xd3\x0a\x25\x68\x31\xdf\x35\x9b\x91\xf4\x8e\x28\xbc\x6d\x64\x05\x3d\xc9\x60\x0b\x27\x6d\x9b\x7d\x56\xef\x28\x95\x5d\x77\x92\xce\x38\x56\x4b\xa1\x85\xa9\x93\x66\xc7\x5f\x7f\xf9\xf7\xa0\x23\x4d\x33\x3b\x57\x88\x6a\xc4\x53\x3d\x31\xdb\x0c\xd9\x99\x38\x50\x05\x51\x08\xeb\x27\x95\xaf\xf3\xc5\x71\x33\x71\xa4\x55\x19\x61\xd9\x1a\x3b\xc6\xfb\xf4\x68\xce\xbc\xb8\xde\x49\x24\xf7\x5f\xeb\xab\xec\xbe\x7b\xad\xeb\x65\x8b\xcc\xcc\xcc\x58\x8a\x25\x69\x2a\xfd\x6d\x43\x7d\xb5\x0a\xcb\xc6\x9f\x35\x36\x6a\xac\x0c\x7b\x60\x0b\x7f\x1f\x4e\x88\xfe\x65\xe5\x05\xe3\x14\x9f\x37\xf0\xc2\x25\x3f\xe4\x5b\xc8\x7e\x55\x48\x7f\xb2\x7f\xfb\xe7\x12\xc3\x63\xb0\xaf\x22\x5e\x2e\xa0\xb2\x19\x37\x74\x86\xa1\x4c\x8d\x52\x86\x54\xf0\x05\x1a\xae\x99\xae\xb0\xeb\x92\xf1\x55\x66\xd8\x98\xc8\x9d\xd9\xb5\x5f\xf2\x4e\xee\x54\xd7\xb5\x2d\x2b\x61\xa7\xbd\x14\x9c\x75\xdd\x06\xfc\x33\x4c\xdb\x9a\x25\x9e\xaf\x39\xb8\xbf\xd1\x13\x8e\x17\x0d\xea\xe1\x68\xd3\x2f\x2e\x6b\xa7\x0b\x16\xda\xb4\x24\xb8\x27\xdd\xde\xda\xd2\x3b\x0d\xe0\xda\xa4\xbb\xac\x8b\x35\xe4\xb0\xfe\x47\x76\x16\x55\xcc\xb5\xdb\xcf\xce\xb6\xed\x8b\xec\x9f\xc8\x51\x12\x2d\x64\xf6\x9f\x86\x54\xbd\xbf\x5d\x17\x2f\x63\x76\x89\xdf\x96\xe3\xb3\xbe\x65\x34\x49\x23\xa1\xda\x16\x3b\x23\x78\xfd\x57\x43\xea\xc7\x6e\xc2\x83\x0d\x88\x82\x6f\x20\xd8\x33\x6b\x78\x2c\x83\x49\x89\x1a\x1c\x09\x80\x1e\x99\xf9\x6a\xa0\x26\x44\x6d\xb7\x99\x3e\xfe\x58\xe0\xe3\x22\xf1\x8f\x1c\x5a\x08\x4f\x46\x1f\x9b\xf1\x45\xb3\x3f\xdc\xfa\xff\x0e\xc1\xfe\x19\x0a\xbe\xf8\xb7\x2f\xc2\x0f\x17\x10\x12\xc7\x0e\xc4\x15\xce\x5c\xea\xa6\x37\xe9\x31\x89\xb3\xde\xa4\x8c\xd1\x8d\xbb\x09\xba\x3d\x59\x79\x30\xc3\x1e\xcd\x49\x3c\x6d\x17\xfc\x76\xe9\x70\x8d\x5f\x66\x17\xae\x8b\xc1\x4b\xed\x6c\xd6\x18\xfa\xd6\xbe\xdb\x7e\xed\x08\x96\xa8\x9a\x4a\xbf\xcd\x43\xb1\xa0\x09\x36\x07\xa9\x51\xe4\xae\xa4\xc7\xcf\xd2\x2b\xc6\x8b\xcf\x78\xe7\x8e\xd4\x71\x41\xe6\x7d\xd8\x40\x30\x66\xcc\x9e\x0c\x44\xf7\xe4\x60\x7b\x4f\x07\x2b\xeb\x2c\x1d\x4e\xd8\xff\x06\x00\x00\xff\xff\x16\x7a\x65\x98\x41\x17\x00\x00")

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

	info := bindataFileInfo{name: "ts.gotemplate", size: 5953, mode: os.FileMode(420), modTime: time.Unix(1586617743, 0)}
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
