// Code generated by go-bindata.
// sources:
// assets/Dockerfile
// assets/Dockerfile-go
// assets/Dockerfile-python
// assets/entrypoint
// assets/entrypoint-go
// assets/entrypoint-python
// DO NOT EDIT!

package command

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

var _assetsDockerfile = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x92\x51\x6f\xd3\x3e\x14\xc5\xdf\xfd\x29\xae\xd2\xfd\xa7\x3f\x13\x8d\x3b\x1e\x78\xe8\xe8\x44\x47\x3b\x11\xa0\x49\x94\x05\xc4\xb4\x4e\x93\x97\xdc\x36\x66\x89\x6d\x6c\x67\x5b\x15\x85\xcf\x8e\xe2\x94\xb5\x0c\xc4\x43\x14\x5f\xfb\x77\xee\x3d\xc7\xf2\x80\x0c\x60\x26\xb3\x3b\xd4\x2b\x5e\x22\xe9\xca\x77\x52\x6d\x34\x5f\x17\x16\xfe\xcf\x5e\xc0\xab\xd1\xf1\x6b\xf8\x50\x0b\x85\x1c\x3e\xb2\x07\x56\x49\x2b\x1d\x96\x16\xdc\x80\x91\x2b\xfb\xc0\x34\x02\x37\xa0\xb1\x44\x66\x30\x87\x5a\xe4\xa8\xc1\x16\x08\x8b\x20\x85\x4f\x3c\x43\x61\xd0\x77\xa2\xc2\x5a\x35\xa6\x54\x2a\x14\x46\xd6\x3a\x43\x5f\xea\x35\x2d\x7b\xc4\xd0\x8a\xdb\xe1\xb6\xf0\x55\xa1\xc8\x80\x34\x4d\x8e\x2b\x2e\x10\xbc\x5b\x66\xd0\x6b\x5b\x72\x9e\x44\x0b\x68\x1a\xff\x8c\x19\x0c\x2a\xb6\xc6\xb6\x25\x8b\x69\x10\xa6\xd3\x20\x9c\x27\xcf\xad\xc2\x9b\xbb\xed\xca\xff\xe6\x4e\xde\xae\x2b\xc6\x4b\x3f\x93\xd5\x29\x21\xf3\xf0\x0b\xa4\xf3\x64\x01\xf7\xf6\x78\x34\x72\xe5\x6c\x7e\x16\x4c\xc3\x9b\xf3\x24\x0a\xd3\x79\x38\x03\x21\x05\x17\x16\x35\xcb\x2c\xbf\x47\xd2\x34\x0f\xdc\x16\xe0\xbf\x4f\xd3\x38\xd6\xf2\x71\xd3\xb6\x4e\xd6\xd5\x37\x71\x12\x7d\xbd\xec\xbc\xb5\x2d\x69\x1a\x14\xb9\xfb\xef\x04\x17\xcf\x15\x17\xff\x94\x84\x72\x9f\x0f\xa3\xbf\xc2\x4f\xf4\x54\xd9\x5f\x78\xf2\x39\x04\xcc\x0a\x09\xde\x34\xfb\x5e\x73\x8d\xe3\x71\x77\xf1\xd0\x80\x23\x60\xe9\xb9\x1e\x4b\xef\x04\xda\x13\x0f\x4e\x4f\x81\xa2\xcd\x28\x53\xb6\xfb\xfc\x4c\x8a\x95\x9f\xd3\xd1\xb1\xea\xe8\x3f\x27\xc5\x1b\xc5\xf7\x47\xc5\x97\x71\xd0\x7b\x9b\xb8\xbe\x70\x78\x08\x4b\x02\x00\x10\xc4\x71\x94\xa4\x93\x83\x66\x87\x0c\x8e\x28\xdd\x23\xaa\xbb\x9c\x6b\x18\x2a\xf8\x41\x7d\xc5\x15\xdd\x9d\xf4\x01\xae\xd6\xa5\xbc\x65\xe5\xf5\x52\x70\x91\xe3\xe3\xb0\xd6\xe5\xe4\x60\xd7\x8d\x6a\x29\x2d\x55\x1b\xc5\x97\xc2\xea\xda\x58\xcc\x87\x85\x34\x76\x72\xd0\xf4\xb3\xff\x1b\x1f\xb5\x2e\xe1\xb6\xbf\xe2\xca\xe5\xdb\x4f\x65\xb1\x52\x25\xb3\x08\x1e\x53\xd6\x83\xfe\x76\x77\x9b\xfd\x43\xed\xf7\xc9\x74\x36\x03\x14\x56\x6f\x94\xe4\xc2\xfa\xa6\x00\x67\x81\xcc\xc3\x34\xb9\x8c\xa3\x20\x4c\xe1\xaa\x7b\xaa\x85\xf7\x12\xbc\xde\xdd\x6f\xb8\x77\xfd\x34\xf9\x67\x00\x00\x00\xff\xff\xd9\xcc\x02\xe5\x7b\x03\x00\x00")

func assetsDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_assetsDockerfile,
		"assets/Dockerfile",
	)
}

func assetsDockerfile() (*asset, error) {
	bytes, err := assetsDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/Dockerfile", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsDockerfileGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x92\x5f\x6f\xd3\x30\x14\xc5\xdf\xfd\x29\xae\x56\x54\xf1\x47\x89\xb7\x3d\x80\xb4\xb7\x42\xd1\x28\x1b\xdb\x54\x0a\xbc\xf0\x72\xe7\xdc\x38\xd6\x1c\xdb\xb2\x6f\x56\x2a\x2b\xdf\x1d\x35\x89\xba\x8a\x09\xc4\x4b\x94\xdc\xfc\x4e\x4e\x7c\xce\x9d\x89\x19\x2c\xbd\x7a\xa0\x58\x1b\x4b\x85\xf6\x62\x3f\xf9\xe0\xc3\x2e\x1a\xdd\x30\xbc\x54\xaf\xe0\xfc\xf4\xec\x6d\x71\x7e\x7a\xf6\x0e\x3e\x77\x2e\x90\x81\x2b\xdc\x62\xeb\x79\x64\x37\x8d\x49\x90\x7c\xcd\x5b\x8c\x04\x26\x41\x24\x4b\x98\xa8\x82\xce\x55\x14\x81\x1b\x82\x2f\xab\x0d\x5c\x1b\x45\x2e\x51\x39\x88\x1a\xe6\x70\x21\xa5\x0f\xe4\x92\xef\xa2\xa2\xd2\x47\x2d\xed\x88\x24\xd9\x1a\x2e\xa6\x87\x32\x34\x41\xcc\x44\xce\x15\xd5\xc6\x11\x9c\x60\xe0\x93\xbe\x17\xeb\x6f\x37\x80\x81\x0b\x4d\x0c\x5d\xa8\x90\x09\xe6\x73\xf8\x29\x00\xe0\x30\x37\x2e\x31\x5a\x0b\xc5\x6e\x18\x75\x6c\x6c\x02\xed\x2d\x3a\x0d\xda\x30\xa8\x2e\x5a\xd8\xee\xc9\x9c\x23\x3a\x4d\x50\x6e\x22\x3e\x9a\x54\x2e\xaa\xca\xbb\x54\x2e\x02\x97\x77\xa8\x1e\x50\x53\xea\x7b\xc8\xb9\x1c\xae\xe4\xaa\xbe\x7f\x6e\xd7\x05\x1d\xb1\xa2\xbd\xdd\x7c\x7e\x98\x2a\x4b\xe8\x9e\xe0\xd8\x42\x11\x6b\x90\x8f\x18\xa5\x42\xd5\x90\xc4\xc0\x12\xa3\x6a\xcc\x23\x25\xf9\x7a\x7c\x63\xcd\xfd\x30\xb7\x26\x71\x92\x42\x7c\xbc\xf9\x0e\x97\xb7\x77\x8b\xcd\x27\x90\x15\x32\x0a\x31\x83\xd5\x74\x3a\x6d\xda\x96\x86\x3c\xda\x87\xca\x44\x28\x02\xbc\x18\x59\x79\x6f\x8e\x9c\x87\xd3\x16\xe9\x1a\x0a\x7f\x0c\xc8\x41\x3f\x34\x92\x2e\xa4\x8c\xb8\x2d\xb5\xe1\xa6\xbb\xef\x12\x45\xe5\x1d\x93\xe3\x52\xf9\x56\xf2\x10\x4d\xa1\xcc\xa8\x90\x2d\x26\xa6\x38\xc9\x9f\x5c\x9a\xd6\x57\xf0\xe6\xd7\x73\x0b\x31\x05\x27\x8e\xca\x1c\xcb\xdf\xf7\x99\xb3\xa9\x0f\xf1\x5f\xfa\x55\x1b\x7c\xe4\x3b\xe4\xa6\xef\xc5\x62\xb9\xdc\x67\xbf\x18\x43\xea\xfb\xc3\xb7\x53\x54\x32\xe7\xbf\xa8\x72\x26\x9b\xe8\x3f\xe4\x6b\x0a\x3e\x19\xf6\x71\x37\xaa\x86\x9f\xfc\x71\xbb\xbe\x5a\xae\xd6\xff\x64\xc5\x0c\xde\x53\xed\x23\x1d\xba\xf8\xca\x14\x52\x29\xfe\x5c\xa7\x91\x9a\xa0\x69\x79\x87\x65\x3a\xce\x64\xbc\xf9\x1d\x00\x00\xff\xff\x27\xb4\x25\xaf\x91\x03\x00\x00")

func assetsDockerfileGoBytes() ([]byte, error) {
	return bindataRead(
		_assetsDockerfileGo,
		"assets/Dockerfile-go",
	)
}

func assetsDockerfileGo() (*asset, error) {
	bytes, err := assetsDockerfileGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/Dockerfile-go", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsDockerfilePython = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x52\x4d\x6f\xdb\x30\x0c\xbd\xeb\x57\x10\x35\x50\x6c\x07\x4b\x4d\x06\x74\x40\x6f\xd9\x72\xc9\xda\x7d\x20\x4b\xb1\xcb\x2e\xb2\xc5\xd8\x44\x15\x49\x93\xe8\x04\xae\xe7\xff\x3e\xd8\x71\xb2\xac\xab\x2f\x26\x1f\x1f\x09\xbe\x47\x65\x22\x83\xa5\x2f\x9f\x30\x6e\xc9\x62\x1e\x5a\xae\xbd\x13\x03\xfa\xd1\x87\x36\x52\x55\x33\xbc\x29\xdf\xc2\xfc\x66\x76\x9b\xcf\x6f\x66\xef\xe1\x53\xe3\x02\x12\xdc\xeb\x83\xde\x79\xf6\x23\x77\x53\x53\x82\xe4\xb7\x7c\xd0\x11\x81\x12\x44\xb4\xa8\x13\x1a\x68\x9c\xc1\x08\x5c\x23\x7c\x5e\x6d\xe0\x81\x4a\x74\x09\xe5\xd8\x54\x33\x87\x3b\xa5\x7c\x40\x97\x7c\x13\x4b\x94\x3e\x56\xca\x1e\x29\x49\xed\x88\xf3\x29\x91\xa1\x0e\x22\x13\x5d\x67\x70\x4b\x0e\xe1\x4a\x07\xbe\xea\x7b\x91\xc1\xca\x25\xd6\xd6\x82\x0e\x0c\x41\x97\x4f\xba\xc2\x24\xc5\xfa\xf1\xcb\x80\xe4\x15\x32\x34\xc1\x68\x46\xb8\xbe\xbe\x40\xaa\xa8\x0d\x42\xde\x0e\xe8\x4f\x01\x00\xe7\x1a\x4d\xf3\xf2\x76\x84\x1a\x26\x9b\xe0\x68\xca\xf4\xcb\x03\x05\x38\x0c\xdc\xb2\x89\x16\x2a\xe2\x69\xc4\x2b\xdf\xb3\xa5\x62\x56\xe5\x06\xf7\x60\xa9\x88\xa8\x8d\x25\x87\xb7\x27\xa0\x78\x9e\x9f\xc2\xf4\xcb\x12\xe3\xbb\x73\x9a\xec\x10\x8a\xae\x8b\xda\x55\x08\x72\x13\xf5\x9e\x92\x5c\x18\xe3\x5d\x92\x8b\xc0\xf2\xdb\xa4\xb6\xef\xff\x91\x7b\x21\xa0\xeb\x64\xdf\x8b\xae\x43\x67\xfa\x5e\x88\x0c\x1e\x8f\x56\x04\x0a\x63\xcb\x20\xe4\x4c\xcf\x4f\xae\x0c\xd5\x0b\x63\x43\x8b\x6e\x3f\xd2\x47\xb9\xf9\xc3\x78\xb6\x74\xa7\x54\xd4\x07\x59\x11\xd7\x4d\xd1\x24\x8c\xa5\x77\x8c\x8e\x65\xe9\x77\xaa\x6d\x9b\x46\x8d\x8d\xf9\x34\x1f\xa3\xda\xe9\xc4\x18\x55\x41\xee\x65\x09\x7e\x43\xa1\x53\xfd\x77\xd3\x8b\x43\x1f\x1f\xc6\x70\xeb\xc5\x72\x39\x28\x5a\xc4\xb2\xa6\x3d\xf6\x3d\x28\xa3\x59\x8b\x1f\x5f\xd7\xf7\xcb\xd5\x7a\xca\x44\x06\x1f\x70\xeb\x23\x9e\xf7\xff\xce\x18\x92\xfc\xcf\xc8\x23\x6b\x22\x4d\x16\xbe\xf0\xeb\x14\xfc\x09\x00\x00\xff\xff\x76\x42\x1b\xca\x22\x03\x00\x00")

func assetsDockerfilePythonBytes() ([]byte, error) {
	return bindataRead(
		_assetsDockerfilePython,
		"assets/Dockerfile-python",
	)
}

func assetsDockerfilePython() (*asset, error) {
	bytes, err := assetsDockerfilePythonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/Dockerfile-python", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsEntrypoint = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x8f\xbd\x4e\x2b\x31\x14\x84\x7b\x3f\xc5\xdc\xec\x2d\xa0\x88\x1d\x28\x28\x82\x68\xa0\xe2\xaf\x4a\xba\x28\x85\xe3\x3d\xbb\xb6\x94\xd8\x96\x7d\x56\x21\xb2\xfc\xee\x68\x77\x43\x89\x44\x39\x9a\x6f\x3e\x69\x9a\x7f\xea\xe0\xbc\x3a\xe8\x6c\xb1\xa4\x41\x34\xa2\x01\x79\x4e\x97\x18\x9c\xe7\x29\xbe\x84\x78\x49\xae\xb7\x8c\x1b\x73\x8b\xfb\xd5\xdd\x03\xde\x06\x1f\xc9\xe1\x5d\x9f\xf5\x29\x70\x98\xb0\xad\x75\x19\x39\x74\x7c\xd6\x89\xe0\x32\x12\x1d\x49\x67\x6a\x31\xf8\x96\x12\xd8\x12\x3e\x5f\xb7\xf8\x70\x86\x7c\x26\x39\x8d\x2c\x73\x5c\x2b\x15\x22\xf9\x1c\x86\x64\x48\x86\xd4\xab\xe3\x8c\x64\x75\x72\xbc\xbc\x06\x19\x6d\x14\x8d\x70\x1d\x76\x3b\xfc\x6f\xf0\x84\x15\xf6\xfb\xc7\x51\xeb\x05\x40\x5f\x64\x30\xbe\x10\x9d\x13\x82\x8c\x0d\x58\x3c\x53\x17\x12\x61\x63\x92\x8b\x8c\x0d\x53\xcc\xeb\x85\x28\x25\x69\xdf\x13\xe4\x5c\xcf\x6d\xad\xd7\x4d\x29\xb2\xd6\x11\x92\xb5\x8a\x52\xc8\xb7\xb5\xfe\xf8\x7e\x13\xfd\x45\xf1\x1d\x00\x00\xff\xff\xae\x78\xf8\x79\x68\x01\x00\x00")

func assetsEntrypointBytes() ([]byte, error) {
	return bindataRead(
		_assetsEntrypoint,
		"assets/entrypoint",
	)
}

func assetsEntrypoint() (*asset, error) {
	bytes, err := assetsEntrypointBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/entrypoint", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsEntrypointGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x92\x5f\x4f\xdb\x3c\x14\xc6\xef\xfd\x29\x9e\x37\x8d\xf4\x02\x52\x12\xfe\x48\x9b\xd4\xa9\x17\xb0\xa1\xae\xdb\x0a\x13\xed\xb8\x41\x08\x99\xf4\xc4\xb1\x48\xec\xc8\x76\xca\x50\xc8\x77\x9f\x12\x77\x81\xd2\x69\xda\x95\xed\xe7\x3c\xc7\xe7\x77\x7c\x3c\xfa\x2f\xb9\x97\x2a\xb9\xe7\x36\x47\x44\x35\x1b\xb1\x11\x48\x39\xf3\x54\x69\xa9\x5c\x24\x74\xaf\x7c\xd4\xd5\x93\x91\x22\x77\xd8\x4b\xf7\x71\x7c\x78\xf4\x2e\x3a\x3e\x3c\x7a\x8f\x2f\xb5\xaa\x48\xe2\x2b\x7f\xe4\xa5\x76\xde\xbb\xcc\xa5\x85\xd5\x99\x7b\xe4\x86\x20\x2d\x0c\x15\xc4\x2d\xad\x50\xab\x15\x19\xb8\x9c\x30\x9f\x2d\xf1\x4d\xa6\xa4\x2c\xc5\x7d\x52\xee\x5c\x35\x4e\x12\x5d\x91\xb2\xba\x36\x29\xc5\xda\x88\xa4\xf0\x16\x9b\x94\xd2\x45\x9b\x43\x5c\xe5\x15\x1b\xb1\xef\xa7\xcb\xcf\x93\x20\x9c\x5e\x76\x9b\xae\x85\x71\xd8\xed\x02\xc6\x64\x86\x9b\x1b\x84\x23\x4c\x70\x88\xdb\xdb\x0f\x5d\x41\xc5\x00\xfa\x49\x29\xba\x3e\x59\x26\x19\x5b\x5e\x9d\x5e\xcf\x16\x77\xd3\xcb\xbb\xeb\xf3\xab\xc5\xec\xf2\x62\x12\x1e\x31\x4a\x73\x8d\x60\x41\x0e\x53\x8d\x35\x19\x2b\xb5\x1a\x23\xdc\xf1\x06\x8c\xd6\xbc\x40\x10\xee\x4d\x67\xf3\xf9\xf9\xd6\x2d\x3b\x66\x08\x59\x96\xb4\xbf\x4d\x76\xb2\x45\x36\x94\x25\xb5\x96\x46\xab\x92\x94\x1b\x23\x3c\x9e\x84\x27\x41\x17\x57\x6b\x7f\xe8\xc9\xbd\x7b\xa6\xac\xe3\x45\x81\x85\xa3\xca\x8e\x03\xd6\x34\x32\x43\xbc\x51\xdb\x96\x35\x8d\xe1\x4a\xd0\x6b\xc9\x27\x36\x4d\xdc\xb6\x9d\x3f\xee\x5d\xa4\x56\x7e\x2d\x2c\x0d\x1e\xa1\x21\xc8\x21\x72\x88\x93\x38\x8e\x03\xf6\x46\x18\xf2\x36\xfe\x33\xca\xb4\x21\x2c\x52\x23\x2b\xf7\x0a\x69\x83\xe0\xc3\x3e\xfa\x77\x8e\xdf\x13\x78\x7b\x51\x72\x00\x99\xc1\x7a\x59\x5a\xf5\xbf\x83\x90\x6b\x52\x25\x6a\x4b\x28\xf9\x03\x41\xfb\x9f\xb5\xa2\x8c\xd7\x85\x83\xd0\x05\x57\x02\x8e\xac\x43\xaa\xcb\x92\xab\x55\x8c\x83\xa4\xaf\xd5\xbd\xd4\x40\x33\x50\xfe\x0b\xdf\xf0\x4e\xdd\x28\x11\x11\xa6\x17\x3f\xe6\xfc\x81\x32\x59\x10\x6e\xf1\xfc\xec\xd5\x3f\x48\x67\x8b\x4f\xe5\xae\xfa\x22\x0d\x9f\xa1\x93\x58\x57\x85\x01\x1e\x75\xaa\xcf\x6a\x59\xac\x4e\x8d\xb0\x6d\xcb\x00\xa1\x7d\x5b\x4d\xb3\x13\x1a\xf8\x5e\x5c\xd1\x7a\x33\xb3\x3e\xda\x77\x91\xc9\xa1\xa1\x5f\x01\x00\x00\xff\xff\xf5\x0b\x49\xe3\xff\x03\x00\x00")

func assetsEntrypointGoBytes() ([]byte, error) {
	return bindataRead(
		_assetsEntrypointGo,
		"assets/entrypoint-go",
	)
}

func assetsEntrypointGo() (*asset, error) {
	bytes, err := assetsEntrypointGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/entrypoint-go", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsEntrypointPython = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x90\xdf\x6f\xda\x3e\x14\xc5\xdf\xfd\x57\x9c\xaf\x89\x2a\x78\x70\x5c\x5a\xe9\xfb\xc0\xc4\xc3\x36\x55\x1a\xfb\x89\x06\x9a\x34\x31\x2a\xb9\xe1\x42\x2c\x05\xdb\xb3\x9d\xb4\x28\xe4\x7f\x9f\x12\x28\x8c\x4d\x93\xfa\x14\xf9\x9e\x93\xcf\xb9\xe7\xf6\xfe\x93\x0f\xda\xc8\x07\x15\x72\x08\x2a\x59\x8f\xf5\x40\x26\xfa\x9d\xb3\xda\x44\xe1\x76\x31\xb7\xa6\x9b\xbe\xb5\x6e\xe7\xf5\x26\x8f\xe8\x67\x03\xdc\x5c\x0f\xff\xc7\xfb\xd2\x38\xd2\xf8\xa0\x1e\xd5\xd6\x46\xdb\xd9\xe6\xb9\x0e\x08\x76\x1d\x1f\x95\x27\xe8\x00\x4f\x05\xa9\x40\x2b\x94\x66\x45\x1e\x31\x27\x7c\x9a\xcc\xf1\x51\x67\x64\x02\xa5\xdd\x4f\x79\x8c\x6e\x24\xa5\x75\x64\x82\x2d\x7d\x46\xa9\xf5\x1b\x59\x1c\x2c\x41\x6e\x75\x14\xc7\x47\xea\x72\xc7\x7a\x6c\xfa\x7a\xfe\x6e\xcc\xa5\xb7\x36\xca\xd4\xed\xc8\x54\x6d\x8d\x51\xd2\xce\x39\xa3\x4a\x15\xe0\x49\xbf\x13\xa0\x8d\x8e\x10\x03\x8e\xab\x2b\xfc\xf8\x43\xab\xb4\x8f\xa5\x2a\xc8\x54\xe2\xd9\xc6\x98\x5e\x63\xb1\x40\xd2\xc3\x18\xd7\x58\x2e\x5f\xb5\x3b\x1b\x06\xd0\x13\x65\x68\x2f\xc5\xd6\x9a\xb1\xe9\xf7\x6f\x77\x5f\x67\x93\x2f\x9f\xc7\xe7\xa0\x10\x55\x51\x40\x14\xd8\xa3\x6d\xcc\x83\x84\x94\x1b\x8e\x3d\x36\x9e\x1c\xc4\x1d\xf8\x7d\x32\xec\x27\xfb\xc5\xbd\x58\x0e\xda\x79\x4e\x6a\x05\x31\x1c\x30\xca\x72\x0b\x3e\xa3\x88\x69\x77\x73\x54\xe4\x83\xb6\x66\x84\xe4\x94\xc4\xd9\x65\xd0\x59\x39\x0a\x85\xcd\xd4\xc5\xf8\xf7\x2e\xb7\x17\x5d\x4e\x71\x64\x2a\xed\xad\xd9\x92\x89\x23\x24\x37\xe3\xe4\x96\xb7\xba\xa9\x0e\x8f\xae\xeb\xc1\x3d\x39\xc6\xce\x22\xb9\x30\xe2\x1d\x1b\x82\xe0\xe9\x67\xa9\x3d\xb5\x84\x90\xc6\xa7\x88\x73\x8c\xd3\xee\x7c\x16\xff\x97\xb3\x85\xd7\xb5\x57\x66\x43\x48\x8f\xf8\xa6\x39\xc6\xd5\x75\xda\x34\x9c\x75\x1f\x56\xd7\x64\x56\x4d\xf3\xbc\xca\x1b\x5a\x5b\x4f\x98\x65\x5e\xbb\x78\x5a\xe8\x84\x3a\xc8\x07\xf5\x45\xbc\x7f\x81\x5e\x82\xf8\x15\x00\x00\xff\xff\xd5\x38\x57\x39\x43\x03\x00\x00")

func assetsEntrypointPythonBytes() ([]byte, error) {
	return bindataRead(
		_assetsEntrypointPython,
		"assets/entrypoint-python",
	)
}

func assetsEntrypointPython() (*asset, error) {
	bytes, err := assetsEntrypointPythonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/entrypoint-python", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"assets/Dockerfile": assetsDockerfile,
	"assets/Dockerfile-go": assetsDockerfileGo,
	"assets/Dockerfile-python": assetsDockerfilePython,
	"assets/entrypoint": assetsEntrypoint,
	"assets/entrypoint-go": assetsEntrypointGo,
	"assets/entrypoint-python": assetsEntrypointPython,
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
		"Dockerfile": &bintree{assetsDockerfile, map[string]*bintree{}},
		"Dockerfile-go": &bintree{assetsDockerfileGo, map[string]*bintree{}},
		"Dockerfile-python": &bintree{assetsDockerfilePython, map[string]*bintree{}},
		"entrypoint": &bintree{assetsEntrypoint, map[string]*bintree{}},
		"entrypoint-go": &bintree{assetsEntrypointGo, map[string]*bintree{}},
		"entrypoint-python": &bintree{assetsEntrypointPython, map[string]*bintree{}},
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

