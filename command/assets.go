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

var _assetsDockerfilePython = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x52\x4d\x6f\xdb\x30\x0c\xbd\xeb\x57\x10\x35\x50\x6c\x07\x4b\x4d\x06\x74\x40\x6f\xd9\x72\xc9\xda\x7d\x20\x4b\xb1\xcb\x2e\xb2\xc5\xd8\x44\x15\x49\x93\xe8\x04\xae\xe7\xff\x3e\xd8\x71\xb2\xac\xab\x2f\x26\x1f\x1f\x09\xbe\x47\x65\x22\x83\xa5\x2f\x9f\x30\x6e\xc9\x62\x1e\x5a\xae\xbd\x13\x03\xfa\xd1\x87\x36\x52\x55\x33\xbc\x29\xdf\xc2\xfc\x66\x76\x9b\xcf\x6f\x66\xef\xe1\x53\xe3\x02\x12\xdc\xeb\x83\xde\x79\xf6\x23\x77\x53\x53\x82\xe4\xb7\x7c\xd0\x11\x81\x12\x44\xb4\xa8\x13\x1a\x68\x9c\xc1\x08\x5c\x23\x7c\x5e\x6d\xe0\x81\x4a\x74\x09\xe5\xd8\x54\x33\x87\x3b\xa5\x7c\x40\x97\x7c\x13\x4b\x94\x3e\x56\xca\x1e\x29\x49\xed\x88\xf3\x29\x91\xa1\x0e\x22\x13\x5d\x67\x70\x4b\x0e\xe1\x4a\x07\xbe\xea\x7b\x91\xc1\xca\x25\xd6\xd6\x82\x0e\x0c\x41\x97\x4f\xba\xc2\x24\xc5\xfa\xf1\xcb\x80\xe4\x15\x32\x34\xc1\x68\x46\xb8\xbe\xbe\x40\xaa\xa8\x0d\x42\xde\x0e\xe8\x4f\x01\x00\xe7\x1a\x4d\xf3\xf2\x76\x84\x1a\x26\x9b\xe0\x68\xca\xf4\xcb\x03\x05\xd0\x91\xf4\x1c\xca\x26\x5a\xa8\x88\xa7\x19\xaf\x7c\xcf\x96\x8a\x59\x95\x1b\xdc\x83\xa5\x22\xa2\x36\x96\x1c\xde\x9e\x80\xe2\x79\x7e\x0a\xd3\x2f\x4b\x8c\xef\xce\x69\xb2\x43\x28\xba\x2e\x6a\x57\x21\xc8\x4d\xd4\x7b\x4a\x72\x61\x8c\x77\x49\x2e\x02\xcb\x6f\x93\xdc\xbe\xff\x47\xef\x85\x82\xae\x93\x7d\x2f\xba\x0e\x9d\xe9\x7b\x21\x32\x78\x3c\x7a\x11\x28\x8c\x2d\x83\x92\x33\x3d\x3f\xd9\x32\x54\x2f\x9c\x0d\x2d\xba\xfd\x48\x1f\xe5\xe6\x0f\xe3\xdd\xd2\x9d\x52\x51\x1f\x64\x45\x5c\x37\x45\x93\x30\x96\xde\x31\x3a\x96\xa5\xdf\xa9\xb6\x6d\x1a\x35\x36\xe6\xd3\x7c\x8c\x6a\xa7\x13\x63\x54\x05\xb9\x97\x25\xf8\x0d\x85\x4e\xf5\xdf\x4d\x2f\x2e\x7d\x7c\x19\xc3\xb1\x17\xcb\xe5\xa0\x68\x11\xcb\x9a\xf6\xd8\xf7\xa0\x8c\x66\x2d\x7e\x7c\x5d\xdf\x2f\x57\xeb\x29\x13\x19\x7c\xc0\xad\x8f\x78\xde\xff\x3b\x63\x48\xf2\x3f\x23\x8f\xac\x89\x34\x59\xf8\xc2\xaf\x53\xf0\x27\x00\x00\xff\xff\x8f\x6d\xf1\xa8\x23\x03\x00\x00")

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

var _assetsEntrypointGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x93\xe1\x6f\xda\x3a\x14\xc5\xbf\xfb\xaf\x38\x2f\x44\x7a\x50\x29\x49\xe9\x93\xde\x93\xa8\x90\x5e\xbb\x55\x8c\x6d\xb4\x53\x61\xfd\x52\xa1\xca\x4d\x6e\x8c\xd5\xc4\x8e\x6c\x87\x16\x05\xfe\xf7\x29\x24\xcb\xe8\xa8\xb6\x7d\xc2\x9c\xfb\xbb\xd7\xe7\xd8\x4e\xef\xaf\xe8\x51\xaa\xe8\x91\xdb\x15\x02\x2a\x59\x8f\xf5\x40\xca\x99\x4d\xa1\xa5\x72\x81\xd0\x7b\xe5\x9d\x2e\x36\x46\x8a\x95\x43\x3f\x1e\xe0\xec\x74\xf8\x6f\x70\x76\x3a\xfc\x0f\x1f\x4b\x55\x90\xc4\x27\xfe\xcc\x73\xed\x1a\x76\xb1\x92\x16\x56\xa7\xee\x99\x1b\x82\xb4\x30\x94\x11\xb7\x94\xa0\x54\x09\x19\xb8\x15\x61\x36\x5d\xe0\xb3\x8c\x49\x59\x0a\xf7\x4d\x2b\xe7\x8a\x51\x14\xe9\x82\x94\xd5\xa5\x89\x29\xd4\x46\x44\x59\x83\xd8\x28\x97\x2e\x68\xff\x84\xc5\xaa\x60\x3d\xf6\xe5\x62\xf1\x61\xec\xf9\x93\x9b\x7a\x51\x47\x18\xf9\xf5\xca\x63\x4c\xa6\xb8\xbf\x87\xdf\xc3\x18\xa7\x58\x2e\xcf\xeb\x0d\x15\x03\xe8\x85\x62\xd4\x39\x59\x2a\x3b\x6a\x88\x31\x3c\xae\x36\xde\x21\x69\x88\x27\x5a\x65\x1b\x2c\x6e\x2f\xee\xa6\xf3\x87\xc9\xcd\xc3\xdd\xd5\xed\x7c\x7a\x73\x3d\xf6\xfb\x42\x63\x4d\xc6\x4a\xad\xb0\x45\x5c\x3a\x04\x29\xfe\x41\x90\xc0\x83\x87\x2d\xea\x9c\x36\x12\x3a\x8a\xc4\x80\x51\x66\xe9\x37\xf3\x86\xb5\xb3\x78\xa5\xe1\xcd\xc9\x61\xd2\x0d\x1f\xc1\x3f\xa2\xbd\x9a\x5d\xf3\x0c\x9e\xdf\x9f\x4c\x67\xb3\xab\x57\x93\x8e\x70\x08\x99\xe7\x34\xf0\xf6\x79\x53\x6d\x40\x90\x0a\x7e\xf5\xff\xe8\x6c\xd4\xf7\x7b\xc1\x70\xb0\x3b\x47\xa2\x5f\x19\x20\xb5\x96\x46\xab\x9c\x94\x1b\xc1\xa7\x7a\xc7\x84\xe2\xac\xbe\xc9\xe0\x05\x3e\xb1\x44\x2b\x62\xac\x69\x98\x2a\xeb\x78\x96\x61\xee\xa8\xb0\x23\x8f\x55\x95\x4c\x11\xb6\xea\x6e\xc7\xaa\xca\x70\x25\xe8\x50\x6a\x1a\xab\x2a\xdc\xed\x6a\x3e\xdc\x53\xa4\x92\xe6\x37\xb3\xd4\x31\x42\x43\x90\x43\xe0\x10\x46\x61\x18\x7a\xec\x27\xa1\xeb\x6b\xf9\x4b\x4a\xb5\x21\xcc\x63\x23\x0b\x77\x60\xa9\xb5\xd0\x94\x9b\xea\xaf\x7d\xb4\xb5\xa3\x41\xd1\x09\x64\x0a\xdb\xc8\xd2\xaa\xbf\x1d\x84\x5c\x93\xca\x51\x5a\x42\xce\x9f\x08\xba\x79\xdd\x09\xa5\xbc\xcc\x1c\x84\xce\xb8\x12\x70\x64\x1d\x62\x9d\xe7\x5c\x25\x21\x4e\xa2\xfd\x5e\xf5\x49\x75\x6e\x3a\x97\x7f\xe2\xaf\x3b\xa7\xe6\x09\x07\x84\xc9\xf5\xd7\x19\x7f\xa2\x54\x66\x84\xe5\x12\xdb\x6d\xab\xbf\x29\x5e\xce\xdf\xe7\x6f\xe9\x07\x62\xf7\x21\xd4\xda\xf7\x47\xdc\x58\x9e\xe8\xcb\x52\x66\xc9\x85\x11\x76\xb7\x63\x80\xd0\x4d\xbc\xaa\x3a\x2a\x75\x3e\x7f\x50\xc1\xba\xbd\xbb\x7d\x75\x9f\x26\x95\x5d\xb0\x6f\x01\x00\x00\xff\xff\x31\xe2\xb2\x96\x8b\x04\x00\x00")

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

var _assetsEntrypointPython = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x91\x5f\x4f\xdb\x30\x14\xc5\xdf\xfd\x29\xce\xdc\x08\xb5\x0f\x4e\x28\x0f\x7b\x08\xaa\x34\xd8\x90\xe8\xc6\x9f\x8a\x76\x93\x10\x03\x64\x92\xdb\xc6\x5a\xb0\x3d\xdb\x2d\xad\x42\xbe\xfb\x94\x10\x0a\x6c\x9a\xc4\x53\x94\x7b\x4e\xce\xef\xde\x93\xde\x87\xe4\x4e\xe9\xe4\x4e\xfa\x02\x82\x96\xac\xc7\x7a\x20\x1d\xdc\xc6\x1a\xa5\x83\xb0\x9b\x50\x18\xdd\x4e\x3f\x1b\xbb\x71\x6a\x51\x04\xf4\xb3\x01\xf6\x76\x87\x1f\xf1\x75\xa9\x2d\x29\x7c\x93\x0f\xf2\xde\x04\xd3\xda\x66\x85\xf2\xf0\x66\x1e\x1e\xa4\x23\x28\x0f\x47\x25\x49\x4f\x39\x96\x3a\x27\x87\x50\x10\x4e\xc7\x33\x9c\xa8\x8c\xb4\xa7\xb8\xfd\xa8\x08\xc1\xa6\x49\x62\x2c\x69\x6f\x96\x2e\xa3\xd8\xb8\x45\x52\x3e\x59\x7c\x72\xaf\x82\xe8\x5e\x62\x5b\x58\xd6\x63\x93\x83\xd9\xf1\x88\x27\xce\x98\x90\xc4\x76\x43\x7a\xd5\x9c\x91\x46\xcd\x9c\xb3\x9c\xb2\xb2\xa1\x8b\x35\x26\x97\xb3\xe3\xf3\xb3\xdb\xc3\xef\xe3\x93\x2f\xb7\x07\x17\xe3\x83\xbd\xdb\xf3\xc9\x6c\x3a\xe2\x62\x8d\xe1\x2e\xc4\x2f\x0c\x4f\x39\xa3\x95\x2c\xc1\xa3\x7e\x9b\x04\xa5\x55\x80\x18\x70\xec\xec\xe0\xe7\x5f\xda\x4a\xb9\xb0\x94\x25\xe9\x95\x78\xb6\x31\xa6\xe6\xb8\xba\x42\xd4\xc3\x08\xbb\xb8\xbe\xde\x6f\x8e\xd4\x0c\xa0\x35\x65\x68\xaa\x65\x73\xc5\x98\x23\x99\x1b\x5d\x6e\x30\xb9\xfc\x71\x74\x31\x1d\x9f\x9f\x8d\x5e\x88\x3e\xc8\xb2\x84\x28\xf1\x88\xa6\x2b\xee\x13\x24\xc9\x82\xe3\x11\x0b\x47\x16\xe2\x08\xfc\x26\x1a\xf6\xa3\xc7\xab\x1b\x71\x3d\x68\xe6\x41\xaa\x12\x62\x38\x60\x94\x15\x06\x7c\x4a\x01\x93\xf6\x6f\x61\x45\xce\x2b\xa3\x53\x44\x5b\x12\x67\x6f\x41\x2f\x4a\x27\x94\x26\x93\x6f\xc6\x6c\x6e\x1c\x08\x4a\x23\xaa\x3e\xa5\x7b\x69\x3f\xea\x89\xe1\xa0\xde\x47\x6e\x9a\xcb\xb6\x4c\xd2\x2b\xe5\x8c\xbe\x27\x1d\x52\x44\xc4\x19\xf0\xaa\xfe\x88\x58\x6e\x34\xb1\x6e\xc9\x71\x87\x9f\x06\xb2\x3e\xe5\x5d\x71\x82\xe0\xe8\xf7\x52\x39\x6a\x52\x7c\x1c\xd6\xe1\x75\x8b\x56\xd9\x97\x82\xdc\x3f\xd6\xa6\xdb\xaa\x72\x52\x2f\x08\x71\x07\xa8\xeb\x0e\x58\x55\x71\x5d\x73\xd6\x3e\x58\x55\x91\xce\xeb\xfa\x79\x99\x43\x9a\x1b\x47\x98\x66\x4e\xd9\xb0\x5d\x69\x1b\xf5\x24\x3f\xa9\xef\xca\xfb\x5f\xd0\x7b\x22\xfe\x04\x00\x00\xff\xff\xd6\x1d\xe1\xea\x87\x03\x00\x00")

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

