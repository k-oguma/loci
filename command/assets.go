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

var _assetsDockerfile = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\x71\x6f\xdb\xb6\x13\xfd\x5f\x9f\xe2\xa0\xe4\x57\xfc\x16\x4c\x92\x93\x01\x1b\x90\x2e\xc1\xd4\xda\xdd\xb4\x2e\xb2\x20\x6b\xd9\x82\xba\x08\x68\xea\x2c\xdd\x2c\x91\x1c\x49\x39\xf1\x04\xed\xb3\x0f\x92\x55\xdb\x2d\xb6\x62\xfb\xc3\xb0\x8e\x7c\xf7\xee\x3d\xf2\x8e\x67\xce\x19\x4c\x25\xdf\xa0\x5e\x53\x85\x4e\x1f\xbe\x96\x6a\xa7\xa9\x28\x2d\xfc\x9f\x7f\x01\x57\x93\xcb\xaf\xbd\xab\xc9\xe5\x37\xf0\x63\x23\x14\x12\xbc\x65\x4f\xac\x96\x56\x0e\xd8\xac\x24\x03\x46\xae\xed\x13\xd3\x08\x64\x40\x63\x85\xcc\x60\x0e\x8d\xc8\x51\x83\x2d\x11\xee\xa2\x0c\x7e\x22\x8e\xc2\xa0\x3f\x24\x95\xd6\xaa\xeb\x20\x90\x0a\x85\x91\x8d\xe6\xe8\x4b\x5d\x04\xd5\x1e\x62\x82\x9a\xac\x37\x06\xbe\x2a\x95\x73\xe6\xb4\x6d\x8e\x6b\x12\x08\xee\x8a\x19\x74\xbb\xce\x79\x93\xce\xef\xa0\x6d\xfd\x57\xcc\x60\x54\xb3\x02\xbb\xce\xb9\x0b\xa3\x38\x0b\xa3\x78\x96\x7e\x2a\x15\xbe\xdd\x8c\x5f\xfe\x6f\xc3\xce\x77\x45\xcd\xa8\xf2\xb9\xac\x6f\x1d\x67\x16\xdf\x43\x36\x4b\xef\x60\x6b\x2f\x27\x93\x21\x9c\xce\x5e\x45\x61\xfc\xf8\x26\x9d\xc7\xd9\x2c\x9e\x82\x90\x82\x84\x45\xcd\xb8\xa5\x2d\x3a\x6d\xfb\x44\xb6\x04\xff\x87\x2c\x4b\x12\x2d\x9f\x77\x5d\x37\xa4\xf5\xf1\x63\x92\xce\x7f\x7d\xe8\xb5\x75\x9d\xd3\xb6\x28\xf2\xe1\xff\x98\xb0\xf8\x34\x63\xf1\xd9\x94\x58\x9e\xe2\xe3\xf9\xdf\x82\x07\x0b\x69\x78\x1f\x2d\x1e\xe7\x8b\xc7\x38\xbc\x9b\x41\x45\xa2\x79\x76\x0e\x34\xa1\xb2\x1f\x78\xd2\x9f\x63\x40\x5e\x4a\x70\x43\xfe\x7b\x43\x1a\xaf\xaf\xfb\x1b\x81\x16\x06\x04\x2c\xdd\x81\x7c\xe9\xbe\x84\xee\xa5\x0b\xb7\xb7\x10\xa0\xe5\x01\x53\xb6\xff\xf9\x5c\x8a\xb5\x9f\x07\x93\x4b\xd5\xa3\x0f\x12\x0e\x95\x92\x9d\xa2\xd3\x52\xc9\x43\x12\xed\x45\xdf\x0c\xbc\xf0\xe2\x05\x2c\x1d\x00\x80\x28\x49\xe6\x69\x76\x73\xde\x1e\x21\x67\x17\x41\x70\x82\xa8\x37\x39\x69\xf0\x14\xfc\x19\xf8\x8a\x54\x70\xdc\xd9\x1b\x78\x57\x54\x72\xc5\xaa\xf7\x4b\x41\x22\xc7\x67\xaf\xd1\xd5\xcd\xf9\x91\x2d\xd0\x52\xda\x40\xed\x14\x2d\x85\xd5\x8d\xb1\x98\x7b\xa5\x34\xf6\xe6\xbc\xdd\xd7\xfe\xdf\xf5\x45\x37\x38\x1c\xf9\x15\xa9\xc1\xdf\xd1\xd5\x19\x44\xc2\x58\x56\x55\xf0\x5a\xd6\xb5\x14\x10\x26\x19\x28\xc6\x37\xac\x40\xe3\x0f\x06\x99\xb2\x5e\x81\x16\x1a\x95\x33\x8b\x47\x89\x1f\xd6\x69\x24\xf0\x76\xc3\x52\x63\xa9\x32\x50\x90\x05\xde\xe8\x0a\x9e\x7a\x88\x69\x72\x09\x8d\xf8\x83\x14\x70\xce\x78\x89\xa0\x36\x85\xd7\x2b\xa1\x02\x9e\xb7\xeb\x15\x54\xb4\x2a\xec\xe6\xca\x9f\x78\x39\x6e\x61\xad\x11\x8b\xaa\xb1\x5f\x0d\xd1\xa1\xa2\xae\xc1\xd3\x6b\x08\xb6\x4c\x07\x15\xad\x86\x2b\xab\xc8\x58\x13\x5c\x9c\x18\xa9\x98\x28\x1a\x56\x20\x18\x85\x9c\xd6\xc4\x4f\xfc\xb4\xed\xaa\x92\x7c\x03\xee\xb8\xe4\xc2\x69\x9b\x1d\x39\x74\x23\x2c\xd5\x68\xc0\x4a\x58\x21\x34\xfd\xc4\x93\x00\x8b\xc6\x1a\xdf\x09\xd3\xef\xe1\x7e\x96\x2e\xa2\x79\x7c\x64\xdc\xa2\x36\x24\xc5\x3f\x32\xe2\xd0\x8b\xf9\x47\x62\x68\x0d\x7e\xa6\xd9\x96\x8c\x1f\xe6\xb9\x14\xa6\xef\x63\x3f\x19\x11\x63\x83\xfd\x87\xf3\x6f\x5b\xcd\x44\x81\x9f\x27\xdd\xcf\x16\x8c\x12\xff\xdd\xe9\x9e\x4c\xc1\x68\x77\xff\xb2\xed\xdd\x86\xd3\x69\x4f\x1a\x6a\x5e\xd2\x16\xbb\x0e\x82\x9c\x59\xe6\xfc\x32\x4f\xdf\x4e\xa3\x74\x8c\x0e\x14\x3d\x1a\x85\xd5\x3b\x25\x49\x58\xdf\x94\x30\xf4\xb1\x33\x8b\xb3\xf4\x21\x99\x47\x71\x06\xef\xfa\x87\xb0\x74\xbf\x04\x77\xdf\xe2\x1f\xc1\xdd\xf7\x07\xae\xbf\x02\x00\x00\xff\xff\xba\x68\xdb\x23\xde\x05\x00\x00")

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

var _assetsDockerfileGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\xdf\x6b\xdb\x3e\x14\xc5\xdf\xf5\x57\x1c\x9c\x50\xda\xef\x17\x5b\x6d\x1f\x36\x28\xf4\xa1\x2c\x25\xcb\xfa\x23\x25\xcd\xba\x87\x6d\x14\x55\xbe\xb6\x45\x6d\x49\x48\x72\xb2\xa0\xfa\x7f\x1f\x71\x42\x12\xd8\x06\x7d\xb3\xef\x3d\x9f\x7b\xef\x39\x68\xc0\x06\x18\x19\xf9\x4a\xae\x50\x35\xa5\xa5\x61\xeb\xca\x27\x63\x57\x4e\x95\x55\xc0\xb1\x3c\xc1\xf9\xe9\xd9\x87\xf4\xfc\xf4\xec\x23\xbe\xb4\xda\x92\xc2\x8d\x58\x8a\xc6\x84\x8d\x76\x5e\x29\x0f\x6f\x8a\xb0\x14\x8e\xa0\x3c\x1c\xd5\x24\x3c\xe5\x68\x75\x4e\x0e\xa1\x22\xdc\x4d\xe6\xb8\x55\x92\xb4\xa7\xac\x87\xaa\x10\xec\x05\xe7\xc6\x92\xf6\xa6\x75\x92\x32\xe3\x4a\x5e\x6f\x24\x9e\x37\x2a\xa4\xdb\x9f\xcc\x56\x96\x0d\x58\x8c\x39\x15\x4a\x13\x12\x2b\xe4\xab\x28\x29\xe9\x3a\x76\x7d\xff\x84\xf1\xf4\xe1\x6a\xfe\x19\x3c\x17\x41\xb0\xd9\xd7\x7b\x08\x1b\xd2\x92\x02\x5a\x9b\x8b\x40\x38\x3a\xc2\x0f\x06\x60\x57\x57\xda\x07\x51\xd7\x48\x57\x28\x4d\x2d\x74\xb9\x97\xb8\x06\xa9\x2b\xc0\x17\xc2\xf1\x5a\xbd\x70\x61\x03\xaf\x95\x0f\x9e\xff\xc7\xd8\x00\x93\x2d\x59\xaa\xa6\xa1\x7e\x57\xf3\x9a\x2b\x87\xd4\x62\xb8\x39\x83\xbf\x28\xbd\x1f\x27\x5b\x57\x23\xf5\xb7\x48\xcd\xa1\x80\xf7\x7c\x1f\x81\xbf\xe0\xdc\x89\x65\x56\xaa\x50\xb5\x2f\xad\x27\x27\x8d\x0e\xa4\x43\x26\x4d\xc3\x83\x13\x0b\xe5\x53\xa9\x36\x04\x6f\x84\x0f\xe4\xb6\xf8\x7e\x4b\xd5\x98\x1c\xff\xff\xfa\x73\x05\x8b\x91\x74\xde\x75\xec\x20\xbd\x05\x39\xaf\x8c\x5e\xa7\xb7\x37\x24\xe0\x2d\x49\x55\x28\x89\x6d\x1f\xa6\xc0\xd8\xf4\x1e\xd7\x33\x2f\x93\xc3\xe1\x17\xc3\xf5\x67\xb2\x3f\x81\x64\x65\x90\x3c\x52\xc0\xd8\x60\xf8\x74\x3d\x7b\x9c\x4c\xef\x0f\xfa\xdf\x77\x45\x5c\x22\x11\x7a\x95\xe0\x27\xde\xde\x40\x0b\x51\x23\x19\x1e\x8f\x27\x77\x77\xd7\xcf\xe3\xe9\xf3\x56\x75\xb9\x93\xf7\x3e\x4e\x92\xbf\x39\xd9\xbc\x9b\xb5\x91\x18\x55\x81\x6c\xde\xa7\x95\x8d\xcd\xa4\xb1\xc6\x85\x07\x11\xaa\xae\x63\x57\xa3\x11\x62\xcc\xae\x9c\xac\xd4\x82\xba\x6e\x97\x92\x77\x92\xc7\xf8\x0f\x2a\x46\xaa\x3d\xbd\x03\x9f\x91\x35\x5e\x05\xe3\x56\x1b\xaa\x3f\xf2\xdb\x74\x76\x33\x9a\xcc\xde\xa5\xfd\x1d\x00\x00\xff\xff\x30\x98\xba\x6b\x80\x03\x00\x00")

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

var _assetsDockerfilePython = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x53\x5d\x6f\xda\x4a\x10\x7d\xf7\xaf\x18\x19\x14\x85\x2b\x2d\x1b\xb8\x52\xae\x14\x29\x0f\xb9\x09\x6d\x68\x13\x40\x7c\x44\x8a\xda\x26\x5a\xdb\x83\x3d\xca\xb2\xbb\xdd\x5d\x3b\x25\xc0\x7f\xaf\xb0\xcd\x87\x50\xeb\x17\x7b\xe7\x9c\x39\x7b\xe6\x8c\xdc\x08\x1a\x70\xa7\xe3\x37\xb4\x73\x92\xc8\xcc\xd2\x67\x5a\x05\xdb\xea\xad\x36\x4b\x4b\x69\xe6\xe1\x3c\x6e\x41\xf7\xa2\x73\xc9\xba\x17\x9d\xff\xe0\x4b\xae\x0c\x12\x7c\x15\xef\x62\xa1\xbd\x2e\xb9\xd3\x8c\x1c\x38\x3d\xf7\xef\xc2\x22\x90\x03\x8b\x12\x85\xc3\x04\x72\x95\xa0\x05\x9f\x21\x3c\xf6\xa7\xf0\x40\x31\x2a\x87\xed\xb2\x29\xf3\xde\x5c\x71\xae\x0d\x2a\xa7\x73\x1b\x63\x5b\xdb\x94\xcb\x8a\xe2\xf8\x82\x3c\xab\x0f\x6d\x93\x99\xa0\x11\xac\x56\x09\xce\x49\x21\x84\x46\xc4\x6f\x22\xc5\x70\xb3\x09\x1a\xd0\x57\xce\x0b\x29\x41\x18\x0f\x35\xe0\xda\xc1\x78\x36\xd8\x56\x58\x8a\x1e\x72\x93\x08\x8f\x70\x76\x06\xdf\x03\x00\xd8\xd7\xa9\xee\x64\x4b\xa8\x06\xaf\x5f\xcc\x90\x01\x61\x49\x74\xe1\x43\x52\xd4\x49\x59\x82\x05\x48\x8a\x2c\x8a\x44\x92\xc2\xcb\xb2\x50\x89\xfd\xe1\x91\x14\x45\x1f\xdd\x5d\x93\xfb\x29\xc9\xe3\xbf\xfb\xa3\x93\xe5\xe7\xde\x8d\x5d\x00\xb3\x73\xe0\x85\xb0\x5c\x52\xc4\x85\xf1\x5c\x92\xf3\x8e\xff\x13\x04\x0d\x98\x55\xe6\x0d\x99\x72\xa6\xad\xb5\xbd\x6f\x96\x9b\xd4\x8a\xa4\x42\x0f\x49\x98\x25\xaa\xa2\x64\xc7\xb9\x95\xc0\xdc\xe4\xa1\x4c\xdb\x5d\x71\x6e\xc5\x7b\x3b\x25\x9f\xe5\x51\xee\xd0\xc6\x5a\x79\x54\xbe\x1d\xeb\x05\x5f\x2e\xf3\x9c\x97\xad\xac\xbe\x00\x2d\x5f\x08\xe7\xd1\xf2\x88\xd4\x29\x04\x6b\x88\x84\xcb\x82\xd5\x0a\x55\xb2\xd9\x04\x47\xeb\x29\xd0\x3a\xd2\xea\x64\x3d\xe0\x0c\xc6\x34\xa7\x78\x17\x76\x4d\x0b\x7a\x83\x27\x18\x3d\x4f\xef\x87\x83\xd7\xff\x67\xfd\x87\xbb\xd7\x9b\x71\xff\xa6\xfb\x3a\x1c\x4d\x27\x10\xb2\x5f\xd0\xb9\x00\xf6\x06\x9d\xc7\xb0\x1c\x69\x74\x33\xbd\xbf\x0e\xb9\xd5\xda\xf3\x76\xe9\x69\xeb\xee\xaa\xb9\xad\x87\x87\x54\x31\xce\x34\x84\x13\xf4\x30\xaa\x6e\x6b\x3e\xf5\xc6\x93\xfe\x70\x70\xcc\x29\x84\x84\xb0\x79\x5e\xaa\x00\x29\xf2\xc0\x5a\x7f\xc5\x0b\xb2\x3e\x17\xb2\x0a\xe1\x94\x3a\x7a\xae\xe5\xaf\x0f\x72\xf5\x92\x24\xac\x61\xfb\x2b\x84\x8e\x03\xe7\x69\x08\x6b\x48\x2d\x1a\x60\x3d\x08\x5f\x76\xae\xce\x9b\xeb\x6f\x2f\xec\x47\x6b\x8b\x7a\x41\x12\x58\xa7\x75\x64\x44\x15\xbb\x84\x6e\x87\x83\x4f\xfd\xcf\xb3\x71\xaf\xcc\xe7\x3a\x64\x0c\x95\x88\x24\x32\x97\x09\x8b\x49\x08\x27\xb7\x3b\x60\x05\x34\xf7\xf6\x0e\x9a\x15\x4f\xea\x58\xc8\x23\x7c\xbf\xce\xdf\x01\x00\x00\xff\xff\x50\xf3\xe1\x63\x19\x04\x00\x00")

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

var _assetsEntrypoint = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x92\x5f\x6b\x9c\x40\x14\xc5\xdf\xe7\x53\x9c\x68\x1e\x12\xd8\xd5\x4d\x02\x2d\xa4\x14\x4a\x0b\x85\xed\x9f\x97\x26\x6f\x9b\xd0\x8c\x7a\xd5\xa1\x3a\x33\xcc\x5c\xb3\x59\xc4\xef\x5e\x74\x4c\x97\xb6\x20\x24\x4f\x7a\xaf\x67\xee\xef\x9e\xe3\xf4\x7d\x41\xa5\xd2\x84\x28\x93\x9e\xa2\x61\x10\xf1\x49\x9a\x29\x9d\x66\xd2\xd7\x22\x16\x31\x48\xb3\x3b\x58\xa3\x34\x4f\xe5\x27\x63\x0f\x4e\x55\x35\xe3\x2c\x3f\xc7\xe5\xe6\xe2\xcd\xfa\x72\x73\xf1\x16\x5f\x3a\x6d\x49\xe1\xab\xdc\xcb\xd6\xb0\x99\xb4\xb7\xb5\xf2\xf0\xa6\xe4\xbd\x74\x04\xe5\xe1\xa8\x21\xe9\xa9\x40\xa7\x0b\x72\xe0\x9a\xf0\x7d\x7b\x8b\x6f\x2a\x27\xed\x29\x99\x0e\xd5\xcc\xf6\x3a\x4d\x8d\x25\xed\x4d\xe7\x72\x4a\x8c\xab\xd2\x26\x48\x7c\xda\x2a\x5e\xcf\x45\x62\x6b\x2b\x62\x21\x62\xfc\xe8\x34\x32\x2a\xcd\x48\xd1\x9e\x65\xd3\xc0\x33\x59\xbf\xfa\xb7\x9c\x45\x3e\x77\xca\xf2\x73\x53\xea\x62\xee\x88\x38\xf4\xa0\x34\x78\x5c\xde\xb8\x82\x5c\x82\x6d\x09\xa9\x0f\xb0\xd2\xc9\x96\x98\x9c\xc7\x68\xa8\x52\x8f\xa4\x57\xa0\x27\xca\x3b\xa6\xc9\xcd\x51\x21\xe2\x89\x4d\xb2\x48\xf0\xd9\x38\xd0\x93\x6c\x6d\x43\x2b\xb0\x41\x41\x59\x57\x05\x40\xe0\xae\xe0\xba\x99\xb8\x57\x5c\xe3\x61\x4c\xff\x01\xd2\x8f\x43\x45\x7c\x1c\x3b\xed\x5a\x11\x43\xc2\xd7\xd4\x34\x90\x79\x4e\xde\x27\xc2\x13\x63\x4d\xa2\xef\xb3\xc6\xe4\xbf\x10\x91\x7e\x8c\x90\x0c\x83\xe8\x7b\xd2\xc5\x30\x08\x55\x62\xb7\xc3\x69\x8c\x93\xf7\xd8\xe0\xfe\xfe\xdd\x38\x5a\x0b\x4c\xeb\xe3\xf4\x83\x28\xd5\xf1\xb4\x75\x64\xa5\xa3\xbf\x27\x1c\x3f\x87\x14\x7f\xce\xd9\x06\x15\xe5\xb5\xc1\x9a\x10\xdd\xd1\xee\xea\xaa\xfd\x18\x82\xde\xce\xf1\xdf\x8c\xa9\x5e\xdf\xd1\xae\x8d\x44\xdf\x3b\xa9\x2b\x42\x12\x34\xb3\xe4\x79\x44\xd4\xf7\xc9\x30\x44\xe2\x6c\x7a\x9e\xff\xa1\xff\xbf\xc5\x12\x7e\x91\xfb\x6a\xe2\xec\x3b\xfc\xb3\x25\xdb\x37\xe1\x7e\x2d\xb8\x0e\x8a\x97\xaf\xb0\xc0\x5e\x82\xbe\x18\x17\x5e\x7e\x07\x00\x00\xff\xff\x71\xc3\x35\xb8\x1e\x04\x00\x00")

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

var _assetsEntrypointGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x53\xc1\x6e\xdc\x36\x10\xbd\xf3\x2b\x5e\x25\x03\xb5\x03\x8b\x4a\x62\xa0\x05\x5c\xe4\xe0\xb4\xe9\x76\xdb\x3a\x2e\xbc\x6e\x2e\x8e\x91\x70\xa5\x11\x45\x98\x22\x05\x92\x5a\xdb\x90\xf5\xef\x05\x25\xd9\xbb\x58\xfb\xd0\x8b\x44\x3e\x0e\x67\xde\x9b\x79\x4c\x7f\xc8\xd7\xca\xe4\x6b\xe1\x6b\x96\xb2\x14\x64\x82\x7b\x68\xad\x32\x21\x93\x76\x44\x7e\xb5\xed\x83\x53\xb2\x0e\x38\x2c\x8e\xf0\xfe\xed\xbb\x9f\xb2\xf7\x6f\xdf\xfd\x8c\x3f\x3b\xd3\x92\xc2\x5f\xe2\x4e\x34\x36\x4c\xb1\x57\xb5\xf2\xf0\xb6\x0a\x77\xc2\x11\x94\x87\x23\x4d\xc2\x53\x89\xce\x94\xe4\x10\x6a\xc2\xf9\xf2\x0a\x7f\xab\x82\x8c\x27\x3e\x5e\xaa\x43\x68\x4f\xf3\xdc\xb6\x64\xbc\xed\x5c\x41\xdc\x3a\x99\xeb\x29\xc4\xe7\x8d\x0a\xd9\xbc\xe1\x6d\xdd\xb2\x94\xb1\x14\x97\x9d\xc1\x9a\x2a\x1b\xab\x18\x1f\x84\xd6\xf0\x81\x5a\x7f\xbc\xbf\x9d\x83\x7c\xe1\x54\x1b\x9e\x40\x61\xca\x19\x61\xe9\x84\x41\x19\x84\x48\xde\xba\x92\x1c\x9f\x84\x8c\xbd\xc0\xd8\x8c\xa7\x04\x95\x2d\x3a\x4f\x1e\xd6\xc0\x75\xc6\x28\x23\x11\xc8\x07\x8f\xca\x3a\x96\x62\x61\x21\xda\x56\xab\x42\x04\x65\x8d\xe7\x58\x56\x10\xe6\x01\xad\x70\xa2\xa1\x40\xce\x23\x36\x46\xaa\x0d\x99\x63\xd0\x3d\x15\x5d\xa0\xb1\x2b\xdb\x08\x96\x8e\x1a\x48\x94\x1c\xbf\x5b\x07\xba\x17\x4d\xab\xe9\x18\xc1\xa2\xa4\x75\x27\x27\xa2\x13\xa1\xe3\x48\x63\x02\xee\x54\xa8\xf1\x3d\x0e\xf2\x3b\x84\x8f\x49\x59\xba\x4d\x3b\x6a\x96\x14\x20\xe0\x6b\xd2\x1a\xa2\x28\xc8\x7b\xce\xfa\xbe\xa4\x4a\x19\x42\x42\x66\x93\x0c\x03\xfb\xe7\xec\xea\x8f\x0f\xc9\xc1\xe2\x22\x2e\xa2\x39\x4e\x0f\xe2\x2a\x61\x25\x15\x3a\xb2\xcf\xee\xb1\x58\x9e\x9f\x7f\xfa\xb6\xb8\xf8\xf6\xe5\xd3\xe5\x6a\x79\xf1\xf9\xc3\xc1\xa1\x54\x4d\x43\xc8\xf4\x11\x53\x15\xae\xaf\x91\x19\x24\x07\xfb\x71\x09\x6e\x6e\x7e\x89\xd4\x0c\x03\x68\x23\x34\x92\xf9\xe6\x51\xc2\x2a\xc5\x1c\x89\xd2\x1a\xfd\x80\xab\xcb\xb3\x2f\xcb\xd5\x5e\x05\x8b\x0d\x39\xaf\xac\xc1\x23\x8a\x2e\x20\xab\x70\x82\xac\x44\x82\x04\x8f\x88\x2e\xf3\xb9\xb4\x79\x2e\x8f\x58\xdf\x93\x29\x87\x81\xed\xc8\x9b\x8d\x11\x25\x52\x51\x5b\x64\x84\xe4\x2b\x5d\x9f\x9c\x34\xcb\xd9\x32\xab\xe8\x84\xd3\xaf\x74\xdd\x24\x4c\xaa\x80\xc2\x9a\x4a\x49\x64\x99\xd4\x76\x2d\xf4\x68\x54\x1e\x3f\xfe\x34\xcf\xa5\x6d\x6f\x25\x57\x86\x57\x56\x6b\x7b\x77\x49\xa5\x72\x54\x04\x8f\xe0\x3a\x62\x7d\xaf\x2a\xf0\x39\xf1\x30\xb0\xbe\x77\xc2\x48\xda\x85\x46\x12\x49\xdf\xf3\x61\x48\xd8\xe1\xf8\xdf\xf2\xee\x7b\xd2\x9e\x9e\xa3\xa4\x1d\x67\x97\x05\xf0\x9c\x73\x9e\xb0\x3d\x60\xf7\xde\xbe\xee\xc9\x28\xaf\xc8\x5e\x4d\x96\xde\x55\xdd\xf7\xf9\x1b\x24\xaa\x7a\xb2\xbb\xf2\xe6\xc7\xf0\x64\xd7\xce\x13\x1a\x71\x4b\xb0\xd3\x43\x2e\xa9\x12\x9d\x0e\x90\x56\x8b\xf9\x1d\xa0\xb0\x4d\x23\x4c\xc9\x13\xbc\xc9\x47\x3a\xb1\x0d\x53\xa1\xdd\x2e\x3c\x23\xff\xaf\x09\xb3\xa3\x08\x8b\xcf\xff\x9e\x8b\x5b\xaa\x94\x26\xdc\xdc\xe0\xf1\x71\xc6\x5f\x05\x3f\xae\x7e\x6b\x5e\xc3\x77\xc0\x67\x2f\x46\x8c\xc5\x6a\x0c\x98\x48\x2f\xec\xc7\x4e\xe9\xf2\xcc\x49\x3f\x0c\x0c\x90\x76\x52\xd8\xf7\x2f\x8e\x9e\x79\x6e\xa3\xb2\xcd\x3c\x98\xf1\x74\x54\x53\xa9\x17\x53\xfa\x2f\x00\x00\xff\xff\xc3\x68\x4c\xe6\x7d\x05\x00\x00")

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

var _assetsEntrypointPython = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x52\x5f\x6b\xdb\x4e\x10\x7c\xbf\x4f\x31\x3f\x29\x84\x04\x2c\x29\x7f\xe0\x57\x48\xc9\x43\x29\x29\x75\xdb\xb4\x21\x31\x85\xe2\x04\x72\x96\x56\xd2\x81\xb4\x77\xbd\x5b\x39\x36\xc2\xdf\xbd\x48\xb6\xeb\x92\xf6\xe9\xb8\xb9\xb9\xd9\xd9\xd9\x8d\xff\xcb\x16\x86\xb3\x85\x0e\xb5\x8a\x55\x0c\x62\xf1\x6b\x67\x0d\x4b\xe2\xd6\x52\x5b\x1e\xd1\xf7\xd6\xad\xbd\xa9\x6a\xc1\x49\x7e\x8a\x8b\xb3\xf3\xff\x93\x8b\xb3\xf3\x37\xf8\xd4\xb1\x23\x83\xcf\xfa\x45\xb7\x56\xec\xc8\x9d\xd5\x26\x20\xd8\x52\x5e\xb4\x27\x98\x00\x4f\x0d\xe9\x40\x05\x3a\x2e\xc8\x43\x6a\xc2\xed\x74\x86\x2f\x26\x27\x0e\x94\x8e\x9f\x6a\x11\x77\x95\x65\xd6\x11\x07\xdb\xf9\x9c\x52\xeb\xab\xac\xd9\x52\x42\xd6\x1a\x49\x76\x97\xd4\xd5\x4e\xc5\x4a\xc5\xb8\xef\x18\x0b\x2a\xed\x50\x85\x83\xe8\xa6\x41\x10\x72\x61\xf2\xfa\xba\x23\x85\xdc\x1b\x27\x7b\x50\x73\xb1\x43\x54\xbc\xc5\x60\x18\x32\x98\xb7\xbe\x20\x9f\x6e\x1b\x19\xf3\xc0\x18\xc8\x5e\xa0\xb4\x79\x17\x28\xc0\x32\x7c\xc7\x6c\xb8\x82\x50\x90\x80\xd2\x7a\x15\xe3\x6e\x8c\x0d\xda\xb9\xc6\xe4\x5a\x8c\xe5\x90\x62\x5a\x42\xf3\x1a\x4e\x7b\xdd\x92\x90\x0f\x18\xc2\xa9\xcc\x92\x78\x02\x5a\x51\xde\x09\x8d\xc9\x1c\x18\x2a\x1e\xfb\x20\x5d\xa4\xf8\x60\x3d\x68\xa5\x5b\xd7\xd0\x04\x62\x51\xd0\xa2\xab\xb6\x66\xb7\xa6\x26\x83\x95\x2d\xf0\x62\xa4\xc6\xf3\x30\xd0\x67\xe8\x30\x88\xaa\xf8\x20\x3b\xf6\x5d\x91\x40\x23\xd4\xd4\x34\xd0\x79\x4e\x21\xa4\xaa\xef\x0b\x2a\x0d\x13\x22\xe2\x65\xb4\xd9\xa8\xbb\x77\xb3\x8f\xd7\x51\xe6\xad\x95\x2c\x75\x6b\xe2\xe5\xb0\x29\x57\x47\x03\x1e\x29\x5a\xea\x06\xd1\xd1\xc9\xf8\x00\xc3\x46\x90\x9c\x46\x38\x3e\xc6\xe3\xab\xb7\xa5\xf1\xd2\xe9\x86\x78\x99\xec\x69\xaa\xef\x89\x8b\xcd\x46\xfd\x51\xd5\x79\x72\xda\xd3\x50\x79\x76\x73\x7f\x7b\xbd\x12\xf2\xad\xf2\xa4\x0b\xcb\xcd\x1a\x77\x3f\xbe\xdf\xdc\x3f\x4c\xbf\x7d\xbd\xde\xeb\x36\x36\xd7\xcd\xe9\xbf\xa4\x76\xe3\x1f\xa4\x28\xaf\x2d\x12\x42\xf4\x48\xf3\xcb\xcb\x76\xba\x5b\x8c\x87\x61\xde\x57\x8f\x34\x6f\x23\x65\x4a\xcc\xe7\x03\xc7\xd3\xcf\xce\x78\x6a\x89\x25\xa4\xb2\x12\x3c\x3d\xbd\x1d\xe2\x63\x05\x38\xe3\x7e\x6f\x55\xe2\xff\xa2\xaa\xd2\xa8\xbe\xf7\x9a\x2b\x42\xba\x2b\xb2\xaf\x1e\xf5\x7d\xba\xd9\x44\xea\x64\x3c\x0f\x86\xf7\xe7\xaf\x00\x00\x00\xff\xff\xa4\xee\xb5\x7b\x82\x03\x00\x00")

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

