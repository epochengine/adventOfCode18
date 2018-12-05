// Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// resources/frequency_input.txt
// resources/suit_ids_input.txt
package main

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

var _resourcesFrequency_inputTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x57\x5d\xce\xe4\x30\x08\x7b\xcf\x55\xa2\x48\x0d\xf9\x6b\xee\x7f\xb1\x11\xb6\xa1\x23\xed\xb7\xd3\xa6\x29\x01\x63\x0c\x6d\xfd\x94\x66\x4f\x69\x7d\x95\x66\xa5\x9d\xd2\x66\x69\xfd\xc5\xd5\x2a\x6d\x97\xda\x77\xa9\xbb\x54\x3c\xac\xfd\x2d\x75\x96\x76\x4b\x1b\xc5\xdf\xad\x17\x4f\x2e\x36\x77\x2b\xb5\xbb\x2d\x7f\xb6\xb1\x7f\x95\xfa\x96\xea\x87\xf8\xda\x03\x6b\xad\x97\x4a\xeb\x37\xf7\x75\x2b\xad\x77\xec\xac\x9d\x96\x7b\x87\x13\x17\x0f\x61\x65\xe3\x38\x3e\x6e\x38\xc8\x74\xdd\x65\xe1\xc0\xd5\xee\x21\xf0\x38\x38\x5a\xdd\xa1\x89\xb3\x70\xf9\x96\xd6\xa7\xfb\xd6\xb6\xaf\xbb\x33\x5b\x67\x8e\x52\xed\x29\x75\xd0\xf6\x8a\x1f\xb8\x3d\x60\xb9\x3a\x52\xfd\x72\x9f\xdf\x1b\x1d\xb4\xc3\x8d\xbe\x7a\x85\x43\xd7\x6f\xfc\x10\x99\xa7\x54\x33\xa1\x10\x71\xf5\x8b\x14\x54\xae\x63\xa5\x27\xac\x38\x7a\x33\x04\xa1\x7c\x33\x07\x11\x27\x4f\x65\x7a\x8c\x01\x2a\x99\x58\xea\x4a\x2b\xbc\x5f\x84\xf5\x49\xd4\xf8\x68\x7c\x46\xfc\x7e\x3a\x52\x0e\xf7\x12\x01\x00\x5e\x2f\xd5\x2e\x91\xf2\x33\x08\xf0\x12\xe8\x2f\x2e\x17\x40\xc2\x96\x45\xd7\xe8\xff\x20\x00\x2f\x02\x07\x7c\x0f\x77\xf2\xd0\xb6\x75\xfe\xe4\x6d\x27\xeb\xb0\xe9\xa5\xc7\x9b\x2e\xbd\x30\xd7\xf9\xbf\xa2\x08\x9f\x3a\xb3\xb3\xf8\x92\xe1\x2c\x40\x2e\xd7\x1a\xdf\x21\xb5\x82\x3d\x7e\xf7\x12\x56\x02\x4e\x66\x32\x41\x4d\x64\x64\x06\x86\xb0\x04\x2e\xfe\x68\x92\x73\x0c\xc5\x63\xb0\x45\x94\xba\xf3\xc8\xa9\x74\xf9\xfe\x24\xc7\xcc\xe4\x90\x88\xe3\xce\x2c\xe6\x68\xf3\xe0\x0e\x8a\xce\x24\x8f\x53\x8b\x58\x63\x65\x88\x20\x6f\xae\xea\x6a\xd0\x93\xf7\xbf\x8c\x1d\x13\x59\xba\x7f\xa9\x7e\x4a\xb3\x23\xe2\xf2\x6c\x80\x69\xb4\x67\x4c\xb7\x21\x4b\xa0\xb1\x91\x78\x4d\xd5\x30\xb9\xf2\x46\x05\x34\xd5\xb9\x0a\xf5\x8a\xfb\x21\x1e\xa2\xb0\x0a\x5c\x55\x27\x22\xb2\x38\x6a\x86\x08\x35\x9a\xff\xfa\x43\x75\xaa\x9b\x40\x4a\x42\x8c\xdc\xb8\x7f\xb4\x65\xaa\x13\xec\x7a\x75\x80\x49\x56\x6e\xa4\x90\x74\x01\x10\x0a\x06\x51\x78\xb2\x58\xdf\xd5\x26\x33\xa9\x47\xd2\x17\x1a\x1f\x70\x6c\xa4\x94\xe1\x9d\x83\x08\xa4\x55\x38\xed\x30\xf8\x13\xd4\x06\xf5\x0f\x18\x33\xb0\x7c\xc5\x9e\xc8\x37\xff\xe1\x7d\xe4\xb2\x41\x19\x9b\x3d\xa9\x0e\xe0\xe6\xa6\x3a\x8d\x8f\x00\x43\x7a\xb4\x93\x73\x5e\x20\x2f\x0d\x77\x59\x61\x46\xdc\xd7\x57\x90\x92\x0a\xa8\x5b\x63\x08\x88\xcb\x6f\x59\x7c\xfc\x63\x9a\x2a\x68\xe4\x70\xd9\xa5\xb5\xa1\x1c\xdf\x52\x47\x2f\xf5\xaa\x0c\x86\x72\x33\xe8\x0c\x95\x37\xb0\xef\xaa\xfb\x43\x94\x17\xe3\x58\x51\x40\xd1\x77\x20\x4a\xac\x85\x29\x61\xd9\x28\xe5\xf8\x91\x4a\x4b\xdc\x9f\x50\x34\xf7\x7f\x8a\xb3\x0f\x91\x8b\x3a\x51\x9e\x5f\xa5\xf7\x89\x54\x50\x09\x78\x06\xf3\x63\x2a\x0f\x49\xb0\x8d\xec\x46\x95\xd0\xa0\xcc\x09\xbb\x0d\xd2\x84\x35\x6f\xcc\x65\xdd\x94\x8c\xc1\x1a\xb6\x91\x72\xdd\x95\x4d\x4f\xe1\x28\x6d\x6c\x26\xcd\x59\xc0\x4e\x60\xaa\x18\x53\x7d\x79\x9b\x08\xbd\x39\x54\xcd\xb0\xb1\x54\x97\x5e\x5f\x72\x7e\x2e\xd2\xee\x32\x4e\xb7\xf7\xc2\xc7\xba\xdd\x1d\x94\xd6\xa5\x35\x8f\xf1\x22\x7c\xd7\xbe\x41\x2f\xdf\x6e\x67\x66\x07\xc3\x1e\x1e\x72\x52\x96\xe4\x85\xfa\x13\xeb\x61\x29\xeb\x93\xfc\x3f\xaa\xca\x27\x64\x1f\x35\xad\x77\x76\x12\xbc\x31\x70\x0a\xac\xf2\x11\x38\x87\xc8\x91\x3b\x9b\x2c\xdf\x4c\xdb\xcc\xde\x74\x98\x23\x15\xd7\x26\x9d\x7a\x7a\x40\xe0\xd5\x60\xf0\xfb\x91\xd6\xb2\x29\x03\xa6\x91\x04\xdc\x7f\x1d\x71\x44\x65\x81\x8f\x52\xfe\x15\x73\x49\x1d\xa9\x10\xec\x4d\x4b\x65\x4a\x4d\xb9\x0a\x58\x2c\x3b\xd1\x94\xbe\x21\xe7\x48\x37\x03\x34\x1d\xf9\x75\x3b\xb0\x10\x40\x35\x8b\x10\x7a\xe6\x04\x9c\x90\x5a\x62\x79\x89\x3a\x4f\xf6\x76\x4d\x61\x43\xb3\x12\x05\x71\x42\x28\xdf\x0c\x5c\x00\xef\x52\x67\x88\xdc\xa0\x86\xa8\x29\xf7\x11\xb9\x42\x14\x52\xed\x93\xee\x25\xc4\x23\x27\x37\x0d\x8d\xec\xf3\xea\x9c\x9a\x6f\x34\x1d\xa8\x80\xed\xf9\x3a\xbe\x1a\x85\xda\xc6\xf9\xfa\x83\x65\x03\xed\x3c\x4e\xd3\x90\x14\xaa\xaa\xdb\x4f\x0d\x2e\xc6\x5e\xc2\x4c\x86\x18\xc2\x10\xb9\x36\x65\xf6\x32\x88\x9d\xad\x87\xa3\xa7\x44\xcc\x6b\xfa\x92\x24\x16\x0d\x2a\xe6\x9a\xec\xa8\xa0\xcc\x9b\x73\x87\xda\x30\x9b\x2f\xc9\x70\xb2\x4e\x28\x6a\x2b\x7b\x49\x5b\x1f\x9f\xa9\x8c\x52\x36\x96\xf7\xf8\xea\xcd\x24\x51\x37\x29\x1a\x4a\xb1\xc4\xd6\xc1\xa3\x31\x7c\x49\x2f\x6a\x0c\xaa\x43\x13\x6e\x8d\x49\x47\x23\x2c\xee\x63\x1e\x8c\xf1\x32\x35\xf2\xc9\x0c\x8a\xfb\x28\x28\x82\xa3\x58\x3e\x39\x6f\x26\xbc\xd3\x5a\x8b\xc6\x11\x90\x66\x2b\xc2\xdf\x2b\x5e\xd7\xab\x16\xa4\xa8\xa3\x19\x2b\xa6\x7e\xbe\xe6\xff\x4d\xc4\x5e\xd7\x60\xd3\xfd\x2f\xce\x1a\x2c\x57\x5f\xd4\xfc\xb4\xbe\x66\x32\xbf\x8f\x8e\x6c\xce\x6c\xbc\x21\x55\x18\xc8\x38\xa7\x2d\x34\x35\x88\xd0\x42\x8a\x8d\xb9\x6f\xc6\xd1\xa2\x1f\xcd\x76\xc1\x70\xf5\x04\xb4\x1d\x55\xe3\xe8\x84\x24\x1a\xd6\x0a\x9e\xb7\x91\xac\x11\x95\x38\x82\xbd\x9c\xf0\x4c\x9f\x07\x53\x23\x8d\x1b\x24\x1d\x63\xb8\xd0\xa8\x5a\xe3\x13\x25\xc6\x05\xd5\xa0\x68\xc9\x4f\xba\x1d\xdf\x14\xd1\x05\x2d\xe7\x42\x18\x5e\xdf\x10\xaa\xc9\xd9\x77\x8c\x05\x43\xce\x23\x37\xce\xf1\x79\xd0\x5b\xb5\xf5\xda\x73\x40\xc8\xb9\x99\x35\x35\x24\x89\x84\x9e\x9f\x69\x1a\x31\xc5\x6d\x06\x61\xcf\x19\xe5\x17\x00\x00\xff\xff\xc8\x6e\x9b\xaa\x81\x0e\x00\x00")

func resourcesFrequency_inputTxtBytes() ([]byte, error) {
	return bindataRead(
		_resourcesFrequency_inputTxt,
		"resources/frequency_input.txt",
	)
}

func resourcesFrequency_inputTxt() (*asset, error) {
	bytes, err := resourcesFrequency_inputTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/frequency_input.txt", size: 3713, mode: os.FileMode(420), modTime: time.Unix(1543842776, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourcesSuit_ids_inputTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x99\xdb\xb6\x46\x20\x10\xc7\xef\xbf\xb7\x24\x3a\x50\x42\x8a\x3c\xfd\x5e\x3a\x4e\xc9\xbe\xfe\xad\xff\x34\xd5\x34\x07\xb8\x3e\x76\x49\xe6\x7e\x3c\xd6\x5d\x28\x3a\x58\xcc\x36\x73\xb3\xab\x3b\x7f\x5c\x4f\x8b\x80\x88\xa1\x07\x99\x80\x4a\xd5\xbc\x99\x63\x8d\xaa\xcd\x23\xe1\xd0\x89\xa0\x41\x49\x8e\x2e\xab\x70\x89\x82\x41\xcd\x07\x8b\xd1\x66\xb6\x80\x74\x42\xdc\xa9\xd0\xb6\x17\xaa\x39\x19\xa4\x68\x33\x87\x47\xf6\x41\x77\xf0\xb0\xf7\x06\xbd\x6a\x7c\x39\xbf\x3b\x24\xa2\x41\x04\x90\x79\xd0\x06\x3d\x14\xeb\xae\x3b\x0a\x0d\x72\x3d\xf1\x64\x90\x3f\xc8\x26\xb4\xc7\xb5\xac\x33\x28\x31\x6a\x6e\x99\xd3\xd1\xe2\x0e\x20\x29\x29\xf4\x10\xa1\x88\xd8\x34\x4b\x32\x33\x80\x0a\x83\xba\x50\x6d\xe6\x16\xd7\x5c\xae\xb5\x7a\xc4\x1e\x15\x3b\x7f\x1c\x4f\xac\x3c\x8d\xb3\xf6\x10\x05\x84\xfd\x5a\xa8\x72\xbe\x53\x74\x40\x73\x76\xe3\xcc\x68\xca\x1e\xda\xb8\x2f\x8f\x54\xf0\x70\x01\x6b\xed\x64\x9e\xbc\x1b\xa6\xdc\x17\x7b\x0c\xa2\xd6\x69\xe8\x69\xe1\xa4\xde\x32\x63\xfb\xf9\xdb\x5f\xb7\xec\x4f\xa3\x23\x3f\xae\xb7\x80\x66\x87\xe8\xfb\x52\xe8\x95\x54\x7d\x44\x6b\x30\x68\x03\x12\x8d\xab\x1c\xac\x50\xd4\xd4\x97\xd2\x65\x37\x78\xb9\x65\x49\xba\x07\x69\xa1\xec\xb3\x96\xac\xd7\xaa\x6e\xf9\x66\x3f\x12\xd7\x32\xeb\x3e\x54\x01\x70\x43\xd5\x30\x58\x55\xba\xc1\xe7\x5e\x24\x83\xfa\xe5\xfc\xb5\x8e\x4b\x23\xa2\xc8\xbc\x44\x37\x76\x87\xd6\x2a\x7a\xb1\x50\xb4\xf3\xaa\xb3\x44\x22\x3a\x8f\xcd\xcd\x78\x46\xdb\x78\xac\x5b\xe5\xfc\xa5\xa7\xc5\x44\xe7\x97\x96\x1b\x0e\xa9\xa8\x62\xec\xea\x58\xbe\x65\x96\x0e\xca\x04\x95\xf1\x68\x04\x67\xa8\x5e\x06\x91\x47\x2b\xda\xae\xa8\xf2\xd9\x66\x78\xbf\xaf\xa5\x88\x28\x63\x1b\x1e\xd2\xa0\x3a\xf0\xfa\x76\xfe\xde\xc5\x0a\x0c\x2e\xef\x2c\x7a\x78\xc4\xf4\x8f\xeb\x2e\x20\x9d\xdd\x60\xa5\x41\x0c\x3c\xa4\xec\xea\xf0\x8f\x6b\x03\x0c\xe2\xb4\x16\x0d\x2a\xe4\x11\x06\xfb\x1a\x1e\x34\x56\xc1\xd6\xa5\xf7\x85\xa1\x87\x2e\x2d\x5f\xe0\xbe\x3c\xba\x9c\x4a\x83\xd3\x50\x13\x38\xc3\xf3\xeb\x2a\x99\x77\x63\x35\xc4\x21\x9c\x10\xcd\xc9\xfc\xf5\x52\xf6\xfa\xa5\x98\x64\xd0\x08\x45\xef\xe6\x5a\x3a\x7b\x88\x20\xca\x59\xf4\xb6\x95\xaa\xeb\xbe\xf3\xe1\x78\xac\xa3\x50\x54\x86\xc0\x0e\x28\x66\x51\x2d\x14\x65\x95\x1b\x3e\xb0\x59\x19\x00\x1e\xcd\xf1\x78\x69\x46\x22\x14\x8e\x3d\xa0\x52\x85\x2a\xe7\x27\x60\x10\xe4\x0d\x0b\x6e\xb9\x77\x6b\xe5\x42\x8f\xe8\x70\x97\x06\x53\x16\x15\xc5\x5a\xf1\x7d\x49\xf0\x1c\x50\x74\x1e\x15\xc9\xa1\x0e\x36\x32\xf7\xdd\x73\xf2\xc3\xfb\x2d\x2f\x6e\x2d\x24\x14\x45\xad\xfb\xe2\xce\xe0\x19\x10\x75\xa8\x2b\x43\xd4\xc7\x86\x4b\x29\x3c\x6e\xf9\x4c\x6e\x5c\x8d\x0a\x1b\xd0\x38\xb1\xab\xbb\x01\xea\x13\x62\x37\xbb\x88\x53\x8d\x51\x35\xbf\x0a\x7d\xca\xbd\x42\x51\xa5\x20\x3a\x82\x41\x92\x4f\x43\x64\xd5\xba\xf6\xcb\x47\xc9\x4e\x1e\x8a\xc1\xe2\x63\x33\x43\x44\x5a\x12\x5d\x96\x00\x60\x30\xc6\xc6\xf3\x1c\xb6\x8f\x1c\x95\xd6\xc2\x43\xb9\x96\x8b\x0d\x82\x52\x6f\x33\x80\x1e\x40\x7f\x79\xa8\x04\x01\x88\x80\x0c\xf0\xac\xb5\x0f\x63\x43\x35\x39\x37\xc2\x73\x90\x00\xf1\xf1\xe0\x4b\x75\x1a\x1d\x34\x78\xda\xef\x83\x72\xed\xc1\xee\xd0\x9a\xd1\x98\xb7\x2c\x61\x3b\x87\x73\x4a\xa9\x8b\x6f\x3a\x28\xb5\x63\x0e\xd6\xd2\x0f\x5a\xf6\xc6\x7d\xc5\x5b\x5e\x7b\xa1\x48\xed\xe1\x16\xf3\xbc\x57\x5d\xc5\x5a\x5b\x5c\x8b\x5c\xef\x7d\xd9\xe2\xbe\x96\x6b\x6f\x47\xaf\xe1\x6c\xb8\xce\xdf\xc2\x21\xda\xbe\x0f\x8a\x79\x84\xd9\x35\x55\x28\x57\x58\xdc\x6d\x25\x3a\xe9\x60\xb9\x37\xc8\xb3\x4a\x79\x37\xd0\x57\xd3\xee\xaa\xde\xa3\x1a\x43\xa7\x67\x6e\xe8\xfc\x92\x3c\xec\x8b\x64\x1e\xdd\x20\xd0\xe0\xea\x9b\x8a\x7c\xcb\x0f\x1a\x96\xaa\x75\xdc\x4c\xef\xd0\x4a\xbc\xea\x68\xbc\xaf\xab\x91\x01\x6e\x76\xd9\xbc\x96\x4e\x48\x39\x55\x38\xa8\xd5\xa9\xb8\x43\xec\xeb\x78\xdd\x23\xba\x1d\x12\x31\xcf\x9f\x2e\xcf\x7f\x3f\x73\x1a\x91\x61\x97\x2c\x22\x6a\x87\x5b\x26\x31\x6f\x44\xa4\x83\xc1\x0d\xb6\x73\x0e\x4d\xd4\xe7\x8d\x3b\x96\x51\xd9\xc8\x00\xfa\x51\x6d\xa9\xfb\x4a\x48\x34\xf6\x45\x50\x54\xf5\x09\x8d\xfd\xd8\xc7\x5a\xf9\x75\x1a\x37\x81\xef\xcb\x65\xec\x3d\xae\x05\x52\x65\x19\xbd\xde\x8d\x25\x1b\xb4\x65\x2b\x02\x2f\x05\xaa\x38\x46\xdb\xe1\x13\xd1\x5d\xa6\xe5\x7c\xf2\x32\xa9\x7a\xaf\x3a\x2c\x46\xf6\xe5\xfc\xb4\x0b\x03\x54\xda\xb7\x22\x53\xe9\xa1\xf5\x2a\xa5\x65\x23\xc1\x8e\xf5\x69\xec\x11\xad\x0c\x03\x34\x01\xd5\xfc\x18\x3c\x52\x3e\xbc\xde\x6b\x71\x3d\x75\x8d\xab\xa4\xdd\xf1\x9b\xcb\xd3\xe0\x0e\x4d\xd5\xbe\x24\x50\xf5\x39\x55\xde\x3a\x34\x30\xff\xb5\xfa\x16\xa3\xd3\x70\x50\xbf\x52\xab\x5f\x54\xd8\xed\x75\x95\x72\xfb\x2c\xbe\x83\xb9\xe9\xd0\x55\xad\x88\xa2\x4b\x70\x63\x72\x48\x9a\x6a\x50\x85\x03\xc2\x55\xb7\xc4\x5f\x6b\xd9\x4d\x3f\xc8\x40\x74\x24\xd5\xc2\x8e\xa8\x52\x41\x65\x8a\x41\x75\x68\xe4\x5e\xf2\x18\x1c\x7f\xfc\x28\x91\x81\xf7\x75\xbf\x63\x43\xb3\x4b\x01\x0f\xc3\x9c\x92\x46\x4e\xe7\xe1\x91\xee\x6b\xfe\x67\x5f\x2a\x22\xde\xa9\x7a\xf2\x3d\x1b\xa5\xad\x3f\x60\x41\xc4\x61\x42\x44\x92\xcc\xeb\xd7\xd7\x03\xf8\x2d\x25\x20\x4c\x7e\x12\xb8\x41\x14\x1d\x96\xac\xba\x33\xea\xa1\x8a\x9d\x3f\x8e\xa6\x45\x9e\xb0\xb7\xf9\xd8\xd7\x38\x58\xf3\xa0\x3b\xa2\x9b\xbd\x27\x5f\x5e\xa8\xec\xf6\x2e\xbe\x52\xc7\x96\x58\xb5\xd7\xb2\xbe\x59\x3a\xab\x4a\x34\x0f\x5e\x75\xfd\x1b\x51\x13\xda\xba\x9b\x5d\xcb\xf9\xe3\xbe\xdc\x1c\x1e\xbd\xdd\x78\x54\xdd\xba\x0b\xdf\x3a\x8a\x84\xfa\x45\x96\xfb\xc2\xe9\x7b\x94\x7a\xf5\xd8\x95\x41\x12\x50\x8f\xa7\x7f\xef\x0b\x4d\x95\xf3\x34\xe7\xde\x54\x61\x69\x31\xde\x8e\x8d\x2d\x77\x42\x4d\xb1\xdf\x68\x9c\x06\xf7\x5f\x60\xae\x80\xee\x32\xe6\x89\x33\xb8\x96\x2a\x37\x72\x3e\x4d\xc5\x98\x0c\xb2\x57\x66\xdb\xaa\xb6\x67\x73\xa8\x4f\x13\x7d\xe9\xc6\x15\xbe\xce\xd5\x9f\x23\x76\x81\x63\xb7\x9c\x1b\xce\xf8\xcc\x29\xfe\xbe\x65\x39\xcc\x0e\x0d\x15\xba\x77\x70\x29\xeb\xf9\x13\xef\xe4\xa0\xe0\x44\xbf\x9c\xaf\xe8\x0d\xed\x81\x24\x33\x01\x03\xc2\xf5\x76\x83\x41\x15\x51\xbf\xae\x1e\x7c\x1a\xaa\xc5\x8d\xd2\x0b\xd8\x97\xc8\xd3\xe8\x3e\xb6\x23\x6a\x60\x6e\xfe\x82\x68\xae\xf7\xa5\xab\x00\x50\x7d\xf1\x21\x68\x0f\xc8\xbe\x7b\x1b\xe5\x51\x1f\x3a\x87\xaa\xfb\xaa\x4f\x7e\x71\x75\x79\xab\xfb\x79\xa7\xba\x9f\x54\xd9\x0a\x80\x23\x7c\x5b\x0e\x06\xd9\x2b\xf7\xf2\xe6\x4c\x34\x53\xfb\x35\x71\x28\x4a\x8b\x4a\xe4\x03\x5b\x65\xe7\x41\xe7\x20\xaa\xbe\xf7\x88\x68\x8a\x06\xf7\x32\x00\x7c\x8f\x3d\x15\x1e\x0e\xe8\xb5\xaf\x03\xb4\x58\x2e\x63\xfb\x4a\x64\xfb\xf1\x30\xdf\x55\x6f\x74\x48\x38\x64\x5a\xe9\x6b\x72\xc5\x97\x5d\x5d\x0f\x90\x6b\x53\xa7\x56\xe7\xc0\xf4\xab\x15\x09\x9f\x74\xc4\xb1\xee\xb2\xe9\xc6\x86\xcb\x98\x3f\xc3\x14\xf0\x2a\x88\x12\xb6\xdf\x32\xa3\xb5\xfa\x38\x53\x4d\xbe\x84\x06\x04\x7f\x49\xb8\xce\x3c\xa6\x14\xc1\xa5\x79\x17\x8e\xfc\x55\x64\x2f\x10\x0e\x2a\xb3\x48\xb2\x02\xd5\xf1\xf1\xf9\x88\x94\xdf\x7b\xdd\xfb\xea\xdf\x13\xa2\x7d\x8d\x15\x01\xf5\xfe\x34\xb0\x6e\x7a\xa8\xa7\xf3\x51\xed\x10\x2d\x01\x29\x70\x1a\x69\x28\x9e\xe1\x14\x70\xe6\xb6\xa7\xf1\x52\x8c\x37\xc8\xa0\x8a\x5c\xe9\x7d\x59\x7c\x57\x2a\xe3\xd0\x45\x07\x3b\xb4\x52\x25\x7b\x86\xfd\xb3\x0a\x00\xd1\xcb\xc6\xbe\xfe\x02\x00\x00\xff\xff\x30\xdf\xb5\xbf\x5e\x1a\x00\x00")

func resourcesSuit_ids_inputTxtBytes() ([]byte, error) {
	return bindataRead(
		_resourcesSuit_ids_inputTxt,
		"resources/suit_ids_input.txt",
	)
}

func resourcesSuit_ids_inputTxt() (*asset, error) {
	bytes, err := resourcesSuit_ids_inputTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/suit_ids_input.txt", size: 6750, mode: os.FileMode(420), modTime: time.Unix(1543882016, 0)}
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
	"resources/frequency_input.txt": resourcesFrequency_inputTxt,
	"resources/suit_ids_input.txt":  resourcesSuit_ids_inputTxt,
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
	"resources": &bintree{nil, map[string]*bintree{
		"frequency_input.txt": &bintree{resourcesFrequency_inputTxt, map[string]*bintree{}},
		"suit_ids_input.txt":  &bintree{resourcesSuit_ids_inputTxt, map[string]*bintree{}},
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
