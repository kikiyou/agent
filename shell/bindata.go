// Code generated by go-bindata.
// sources:
// shell/linux_json_api.sh
// DO NOT EDIT!

package shell

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

var _shellLinux_json_apiSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x7c\x6b\x6f\x24\xc7\x75\xe8\xe7\xe9\x5f\x71\x5c\xec\x55\xcf\x2c\xd9\x6c\x0e\xb9\x4b\x49\x4b\x8d\xae\x57\xdc\x5d\x69\x7d\x77\xb9\xc4\x3e\x7c\x71\xc1\xa1\x86\x35\xdd\x35\x33\x25\x76\x57\xb7\xaa\xaa\x39\x4b\x0d\xe7\x5b\x90\x28\x0e\x04\xc4\x40\x00\x27\x81\x83\xd8\x01\x02\x18\x41\x80\x20\x48\x60\x1b\x11\xfc\x6f\xbc\x52\xfc\x2f\x82\x7a\xf4\x6b\x1e\xd4\xee\x3a\x46\x02\x98\x12\xb8\xdd\x55\xa7\x4e\x9d\x3a\x75\xea\xbc\xea\x34\x37\xbe\x17\x0c\x29\x0b\x86\x58\x4c\x1c\x67\x90\x61\x2e\xc8\x5d\x16\x1d\x73\xca\x64\xbb\x03\x33\x07\x60\x3a\xa1\x31\x01\x4e\x70\x04\x11\x96\xf8\x00\xa2\xd4\x01\x00\xd8\x20\xe1\x24\x05\x9f\x01\x72\x55\x3b\x82\x2b\x10\x24\x02\x9f\x03\x12\x41\x1f\x05\x7d\xf5\x83\x82\xb1\xea\x90\x1c\xfc\x08\x50\x9f\xa1\x03\x3d\x76\x71\xa8\x03\x10\xa5\x8c\x1c\x38\x73\xc7\xc1\x3c\x1b\x84\x38\x9c\x10\x3b\x3f\xe6\xd9\x61\x9a\x24\x98\x45\x3d\xb7\x1d\x9a\x27\xf0\x2f\x54\x7b\xc7\x71\x00\x38\x11\x79\x2c\x7b\x6e\xdb\xad\x20\x15\xee\x2b\xc0\xd3\x73\xf0\x3e\xba\xff\xf1\xc3\x23\x98\x65\x6a\x45\x80\x4e\xd0\x1c\x8e\x9e\x7e\xd8\x85\xbe\xa6\xa3\xfa\x29\x00\x66\xd0\x47\x38\x8a\x78\x1f\xdd\x81\x3e\x42\xe0\x76\x01\xf5\xd1\x16\xa0\xa5\x11\xe6\x07\xf5\xd1\x64\x3a\x90\x97\x19\x29\x47\xec\xbe\xce\x08\x35\xc7\x76\x39\x64\xef\x3b\x87\x24\x58\x9c\x97\xe0\xb7\x55\x0b\xcc\xd7\x0f\x98\x2f\xb4\xdf\x3f\xba\x57\xae\xf0\x14\xcd\xbd\x5a\xf7\x15\x68\x01\x50\x5b\xe7\x1d\x1d\xb8\x22\x98\x6f\x05\xf3\xe0\xe0\xf8\xe0\x9e\xa7\xd9\x4b\x47\x70\x02\xfe\x17\x80\x5c\xc3\x68\x04\xa7\x07\x00\x72\x42\x98\xd9\xc5\xd9\xdc\x01\x20\xb1\x20\xe6\xd5\x42\xc1\x15\x34\x85\xc9\x01\x18\x51\xb5\xbd\x43\xcc\xa2\x29\x8d\xe4\x44\x6f\xaf\x03\x66\xfa\x10\x4b\x08\x32\x9e\x86\x01\x23\x32\x88\xc8\x85\xa6\xf0\xba\x2d\xdc\x6d\x6c\x19\x65\x92\xf0\x11\x0e\xc9\xc2\xbe\x2d\x32\x08\x41\x1f\xc9\x97\x0a\xc8\x6c\xd4\x4a\x00\x5e\x00\x74\x77\x40\xb1\x19\xc1\x7c\x35\x07\x97\x78\xb7\xd5\x67\x41\x9f\x19\xee\x59\x88\x05\x36\xcc\x1d\x47\xc9\x70\xca\x06\x38\xcb\x62\x1a\x62\x49\x53\x26\xac\xa8\x97\xa2\x3c\x9d\x10\x4e\xa8\x00\xf0\x87\x20\x12\xcc\x65\x28\x63\x98\x44\x91\x24\x49\x06\x90\x5c\x8a\xcf\x63\xb8\x10\x23\x99\x45\xc0\xc6\x94\xbd\xac\x71\xcb\x7f\x70\x07\xbc\x19\x1d\xb5\x63\xc2\xc6\x72\xd2\x76\x77\x3b\xbd\xde\x4e\x07\x66\x40\x99\x90\x38\x8e\x49\xd4\x43\x23\x1c\x0b\x82\x0e\x60\x6e\x76\xae\xd1\x27\x79\x6e\xba\x2a\xc6\x98\x75\x57\xef\x8a\xe5\x35\xae\xf5\xd1\x90\x32\xcc\x2f\x2d\xef\xdd\xae\x16\xe6\x26\x44\x9c\x9a\xa5\x16\x30\xbb\x2b\x60\x4a\x22\x34\xfb\xcb\xb7\xfa\x16\xcd\xb7\xd0\xdc\xc8\xa5\x96\x36\x74\x82\xc0\x9d\x19\xb6\xdd\xf8\x3f\x73\xb5\x35\xab\x59\x9e\xe5\x03\xca\x46\x69\x21\x73\x25\xa3\x83\x5c\x70\xbd\x87\xb1\x08\xb3\xbc\x9c\xe8\x0a\xca\x8e\x8a\xa7\x76\xf7\x8b\x15\xd6\xd6\xa1\x04\x04\xbc\x62\x70\x8d\xbc\xd9\x02\x79\xf3\xeb\xc8\x93\x84\x09\x7a\x41\x06\xea\x18\x10\x21\x88\x58\xa6\x56\x11\x94\x09\xc0\x2f\x53\xc8\x68\xb4\x95\x0b\xc2\xb7\xb2\x30\xcb\xb7\xb8\x10\x5b\x17\xe2\x8b\x2d\x25\x5b\x56\x11\x77\xa3\x2b\x10\x29\x97\xe0\x73\x06\xfe\xf9\x5e\xe3\xcc\x4f\x94\x3a\xf7\x19\x74\x6f\x37\x55\x41\x7d\xd5\xe6\xdc\xcd\x9e\x3c\x78\xd6\x43\x77\x0a\xbd\x59\x3f\x74\x19\x8d\x8a\x83\xb2\x42\x0f\xa9\xed\x45\x8a\xc2\xa6\x62\x5c\xa5\xb2\x34\x68\x98\xe5\x37\x2c\xba\xbd\x75\x30\x5c\x08\x0b\x72\x6b\x1d\xc8\x85\xf8\xc2\x82\xdc\x5e\x3b\x53\x12\x95\x34\xed\x1b\x9a\xd0\x7c\x0b\x2d\x42\xbf\xad\xa0\xa9\x43\x5a\x6c\x9d\xd6\x9e\x67\xd3\x09\x0d\x27\x20\x08\x13\x29\x17\x67\x4a\x81\x2a\xfd\xa9\xa7\xe3\x44\xe6\x9c\x3d\x93\x9c\xb2\x71\xef\xac\x00\x31\x06\x16\x27\x91\x7e\x50\x48\x4e\x00\xa9\xf9\x2b\xe0\x00\x9d\x77\x77\x94\x40\x7d\xaf\xb7\xd8\x35\x47\x70\x7a\x0a\xb5\x49\xec\x22\x9a\x50\x1b\x1b\x37\xcf\xbb\x3b\x73\xb8\x82\x30\x97\xca\x3c\x7b\xe0\x81\x3f\x82\xfd\xa2\x25\x84\x5d\xbf\x7a\xee\xfa\xb7\x0c\x55\x4a\x50\x63\x63\xc5\xe3\x35\x94\x85\x29\x27\xaf\x4d\xda\x88\xa7\x89\x1a\xd0\x5b\x22\x4f\xe3\x51\xdc\x44\xf3\xe6\x32\x8a\x21\x1b\x1b\x37\x8f\x27\x97\x82\x86\x38\x9e\xc3\xd2\x42\xf6\xe0\x0a\xaa\x95\xdc\x5e\x65\x92\xb4\x51\x32\xe6\xab\xf2\x4b\xd0\xc9\xaa\xcd\x2d\xec\xd7\x06\x0c\x2f\xe1\x18\xe7\x31\x1c\xa6\xf1\xf0\x12\xda\x13\x29\xb3\x3b\x41\x10\xaa\xb7\x6d\x1a\x6d\xe3\xbc\xb3\x05\x2c\x05\x4e\xc7\x13\x29\xd4\xe1\x25\xfc\x82\x44\x70\xd0\xd1\xd2\x91\x4b\x1a\xd3\x2f\xb4\x2e\x2c\x84\xe4\xf8\xe9\xfd\x1f\x0e\x9e\x3f\x79\x7e\xf7\x51\x6f\xa7\x78\x7d\x78\xef\xd1\x7d\xfd\x46\x25\xe1\x1a\xba\xb7\xe3\x94\xde\xd8\xc9\x49\xd5\x0e\x7e\x2c\x61\x17\x4e\x4f\x2b\xcf\x0c\x3e\x26\x52\xb1\x18\x64\x2a\x71\x0c\x87\xc7\x2f\x40\x48\x2c\xa9\x90\x34\x14\x5b\x10\x51\x11\x62\x1e\x51\x36\xd6\x40\x9e\xd2\x7c\x1e\x64\x9c\x8c\xe8\xcb\x6d\x8d\xe1\xf0\xf8\x45\xaf\x7d\xa6\xdd\x39\x06\x9e\x08\x3e\x0d\xb3\xbc\x2f\x82\x20\xf3\xac\x95\x56\xe8\xce\x3a\x1a\x56\x53\xea\xce\x0e\x8f\x5f\x9c\xec\x9d\xce\x61\x03\x7e\x90\x0b\x33\x3b\x8d\x62\xa2\x27\x97\x34\x21\xdb\x8e\xa5\xed\x10\xc7\x61\x1e\x63\x49\x16\x28\x34\x40\x0a\xa6\x62\x05\xc0\x28\xe5\xf0\xc3\xbb\x8f\x5e\xdc\x07\xaa\xdc\x45\x3d\xcd\xf7\x4f\xe7\xa8\x5c\x2c\x40\x4c\x24\x20\x33\xc6\xd5\xff\x6c\xba\x7a\x04\xd2\xfd\xca\xaf\x5c\x3d\xb5\x9a\x34\x17\x78\x4c\x40\x50\x16\x12\x98\x12\x88\xb1\x90\x10\x4e\x48\x78\x4e\x22\x43\x8b\x46\x7e\xef\xe1\x83\x07\x66\x47\x5c\xf5\xdb\x77\xcb\x2d\x42\x0b\x40\x75\x32\x2c\x98\x7e\x5e\x84\x7b\xf1\xec\xee\xc7\xf7\x7b\xed\xee\xce\xce\xce\xcd\xb6\x5b\x0d\xf5\xdd\x72\xae\x4e\x50\x6b\xdf\xbc\xdd\x09\xba\x3b\xa8\xee\x78\x13\x06\xa8\xcf\x0f\x8f\x5f\xdc\x01\xb7\xc2\x79\x03\xa0\x3f\xec\x0f\x51\xb1\xe4\xa7\x24\x21\xc9\x90\xf0\x1a\xb3\x95\x7f\xdc\xd8\x1a\xa1\xb9\xac\x00\x18\x79\x69\xd7\x6f\x56\x5f\x93\x4c\xe4\xd6\x16\x52\x89\x28\x72\x0d\x1b\xec\x74\xff\x0f\x53\x09\x43\x32\x4a\x39\x31\x78\x94\x90\xe1\x31\xa6\xcc\x20\x14\x31\x21\x19\x74\x8d\x5a\x2b\x25\x1b\xb9\xe5\xf3\x66\x17\xd9\x60\xa0\x50\xbd\x7a\xa1\xb5\x15\x22\xad\x67\x79\xca\x06\x13\x2a\x64\xca\x2f\x8b\x63\x34\xe6\x24\x3b\x4c\x22\xed\x3e\x29\x7d\xab\xde\x95\x88\x4a\x4c\xe3\x7a\xbb\x7a\x57\xed\x0a\xc7\xa3\x74\xdc\xf3\x82\x0b\xcc\x83\x38\x1d\x07\xaa\xc5\x73\x00\x58\xae\x58\xf6\x64\xf4\x88\x32\x22\x7a\x5e\x77\xc7\x53\xf8\x37\xe0\x71\xca\xe4\x64\x0b\xee\xe1\xcb\x2d\x78\x4e\x13\xb2\x05\x9f\xa4\x42\x32\xac\x9e\x24\x1e\x6f\x81\x36\xc8\x06\xb6\x0a\x4b\x2c\x5d\xe0\x27\x6e\x03\x31\x1c\x3e\x7d\x72\x04\xae\x25\x43\x1b\xab\x6a\x90\x25\x1a\x7c\xb6\x30\xa8\x01\x5f\xba\xc8\x33\x10\xd0\x03\x84\x0e\xf4\x56\xb6\x29\xf4\x60\xff\x00\x28\x7c\xd0\x83\xa3\x07\x07\x40\x37\x37\x3b\x1a\x42\x80\x4b\x01\x01\x3a\x58\xf2\xee\xd0\xac\x8f\xbe\xf9\xc9\x2f\x7f\xf7\x93\x7f\xef\x23\x28\x7c\x68\x04\xca\xc3\x01\x64\x03\x94\x25\x67\xb9\x8f\xbe\xfd\xab\x5f\x7c\xf3\xe5\xaf\xaa\x21\xfb\xeb\x00\x13\x22\xd4\x59\xab\x20\x6f\x23\x40\x63\xc2\x44\x3e\x6c\x2b\x13\xbc\x05\xc8\x44\x8c\xea\x69\x8c\xb6\x40\x74\x96\xdd\x05\x65\xa6\x2b\x5f\xd0\xd3\x8f\x95\x91\x3e\xa9\x59\xe8\xd3\xd5\xe6\x39\xe7\x9c\x30\x39\xe0\x38\x29\xa4\x06\x4f\xcf\x95\x70\x58\x1b\x8d\xa7\xe7\xca\xf8\x86\x58\xd6\x1a\x43\x2c\xcf\x6a\xf2\x75\x56\x89\x97\x6a\x4e\x48\xf2\x90\x8d\xd2\x07\x34\x26\x3d\x64\xb4\x63\x42\x12\xe5\x6d\x22\x23\x09\x4f\xc9\x88\x70\xc2\x42\x22\xee\xe8\x77\x28\xf5\x90\x72\xfd\xef\x80\x35\x20\x5f\x84\x84\x49\x4e\xc3\xed\x30\x4d\x82\xdd\x9d\xee\x6e\xb0\x73\x3b\xd8\x7d\x3f\x48\x70\x96\x51\x36\xf6\x15\x66\x8b\xd8\x97\xa9\x9f\xe6\x32\xcb\xa5\x9f\x8e\xfc\x11\x27\xc4\xb7\x81\x71\x60\x67\x78\x40\x49\x1c\x59\xdc\xe2\x4e\x10\x4c\xa7\xd3\xed\x73\xc2\x19\x89\xb7\x53\x3e\x0e\xa2\x34\x0c\xee\xa5\x61\x9e\x10\x26\x35\x19\xc1\x88\xc6\x44\x5c\x0a\x49\x12\xa1\xd7\xb0\x2d\x5f\x4a\xa7\x5a\x5d\xef\xcc\x35\x3c\x01\xb7\xb6\x5e\xb8\x82\x52\xba\xbd\xc7\x24\x79\xae\xf4\x4b\xff\xea\x31\x49\x1e\x70\x42\xfa\x57\x1f\xe5\xa3\x11\xe1\xa2\x7f\x75\xa8\xe2\xf9\xc8\x3b\x2b\xb7\xaa\xc0\xa2\x30\x98\x1d\xa8\x1c\x6c\xe5\x5a\x6a\x4d\x65\xfc\xb8\xb6\xbb\x1b\x74\x77\x76\x6f\x75\x4a\x8f\xd2\x7a\x9d\x6d\xd5\xe5\xb7\xdd\xdb\x9b\xee\x7b\x9b\x6e\xb7\xdb\xe9\x68\x38\xb0\x80\xf8\x02\xd3\x18\x0f\x63\x62\xa1\x6b\x80\x05\x3e\x98\x23\x25\x46\x2b\x45\x25\xa2\xe2\x5c\xb5\x4a\x5a\x45\x68\x4b\xbe\x78\x34\x02\xff\x78\x52\x1c\xc2\xa6\x83\xdc\x47\x8a\xa7\x03\xc3\xd4\xc5\x74\x42\x1f\x09\xfa\xc5\x52\xca\xa0\x5a\x5d\x3d\x2b\x60\x97\x52\x36\xdf\xaa\x03\xdf\x68\x26\x05\x54\x73\x92\xe6\x4c\x92\x05\x17\x77\x21\x70\x3a\x69\x78\xb3\xab\x0f\x4b\x94\x86\xe7\x84\xaf\x0d\x47\x90\x3a\x89\x61\xca\x24\xa6\x8c\x70\xd1\x43\x6e\xdb\x8c\x80\x4c\x94\x6a\x89\x8e\xda\x8a\x2b\x1d\xab\x64\xdc\xa3\x07\x73\xaf\xa3\x06\x2a\x35\x45\x95\x4d\x77\x2b\x14\xd6\xa6\x17\xf8\x8b\x54\x42\x2d\x4a\xb3\x13\xc8\x34\x53\x7a\x6c\x39\x0c\xca\x12\x92\x98\x18\xc8\xf7\x4d\xec\xa3\x9b\x7d\xd5\xfe\xc6\xf1\x8f\x7f\x01\x21\x53\x74\x50\xf4\xdd\xb1\x50\xa8\xac\x40\xc1\xf2\xb0\x11\x2d\xd7\x94\xe0\xd6\xff\x48\xd4\x94\x90\xe4\xc6\x77\x85\x4d\xf5\x98\xe8\xf6\xb5\x31\x51\x69\x9f\xdf\x34\x36\x8a\xd2\x29\x8b\x53\x1c\x0d\x24\xc7\x4c\x8c\x08\x1f\x70\x2c\x4d\x96\xcf\x69\x69\xfd\xd3\x6b\x07\xe2\x52\x04\x61\x8c\x85\xd0\x89\xa0\x9b\x1d\xa7\x95\xa5\xa2\xe7\xb6\xdb\xe0\xce\x36\x34\xd0\xc9\xcd\xd3\x39\xf8\xd0\x85\x4e\xc7\x69\x29\x87\xad\xe7\xce\x4c\x87\x9b\xa5\xe2\x74\xee\x38\xad\xcf\x44\xca\x06\x46\x41\xf6\xd0\x0c\x29\xf4\x4a\xdc\x8a\x1c\x91\x75\x25\xcd\x20\xe5\x4c\x3a\xad\x28\x75\x5a\xad\x21\x16\x44\x6d\x63\xcf\x6d\x17\x8f\x80\xdc\x72\x18\xea\x38\x4e\xab\xb5\x01\x23\xca\x22\xe3\x2d\x69\xc3\x0c\xe9\x08\x86\x97\x92\x08\x28\xd6\x45\x22\xeb\x50\x51\x51\x4d\xea\xb4\x5a\x94\x75\x7b\x6e\x5b\xa7\xba\x9a\xcb\x44\x6e\x31\x1f\x0a\x2a\x47\x3d\xe0\x2f\x07\x1a\xb1\x9d\x77\xaa\x7c\x2b\x0c\x82\x84\x29\x8b\x9c\x56\xab\xf0\xa5\x74\xa7\x76\xb6\x40\x28\x92\xab\x75\x6a\xc7\x4b\x4f\xbc\xfb\xfb\x4c\x3c\xb6\xd1\x44\x44\x47\xd6\x92\x41\xbb\x58\x2b\xa8\x3d\xec\xe8\x39\xcc\x18\xb5\x57\x94\xed\x82\x0f\x94\x75\x3b\x16\x43\x98\xb2\x0b\xc2\x25\x34\x46\x81\x4c\xe1\xff\x7e\x64\x86\x9e\xd7\xc6\x1a\x3c\x10\x80\xd6\xd2\xaf\x85\xa1\xb1\xe1\x6e\xed\x0d\xfa\xd5\x12\x95\x7c\xbb\xe5\x5c\xc8\xe0\xa5\x23\xa0\x12\xa8\x00\x96\x9a\x45\xea\x10\x20\xa6\x4c\x6f\x98\x0e\x6f\xbf\x07\x95\x0c\x40\xaf\x07\xae\x06\x39\x3d\x75\x5a\x2d\x1d\xc6\xb6\x14\x1e\x1c\x45\x80\x41\x5b\x61\x45\x95\xc6\x44\x19\x81\xf6\x0f\x9e\x3d\x39\x52\xe2\x90\x60\x29\x29\x1b\x2b\x56\xad\x27\x77\x0b\x39\xad\xd6\x88\x2a\x79\x54\x07\x4c\xad\x3b\x4e\x85\x09\x57\x34\xa2\x74\xf8\x19\x09\x25\xbc\x63\xf5\xa9\x4c\x41\x84\x9c\x28\x22\xcc\x51\xac\x23\x5b\x93\x6c\x1a\x13\x46\x38\x8e\x1b\xf9\xb0\x51\xce\x42\x1d\x4b\x46\x54\x64\x31\xbe\x54\x51\x81\xce\x48\x02\xc4\x69\x88\x63\x78\xde\x73\xbb\xb5\xd7\x7b\x6a\xa7\x9e\x07\xfb\x3b\xea\x7f\xbd\x49\x55\xdf\x27\xb5\xbe\x1b\x0b\x7d\x8f\x8b\xbe\x1b\xfb\x3b\x8d\x8e\x67\xba\xa3\x6a\x3d\x39\x01\xf7\x1e\x7c\x08\x3b\x70\x7a\x0a\xef\xd8\xe5\x8e\xc0\xbb\x11\x41\x84\x2f\x05\x78\xe0\xde\x2b\x01\x3f\x59\x0d\x38\x49\x73\xae\x21\x3f\x29\x21\x1f\xaf\x86\x4c\x28\xcb\x95\xc8\x79\xe0\x3e\x5e\x98\xfe\xea\xaa\xc0\xaf\x9e\x56\x8d\x57\xe1\x94\xf1\x44\x6b\x18\xcd\x21\x15\x7d\xe6\x81\xfb\xcc\x01\x98\x2b\x2e\xc7\x62\xf8\x94\xc4\x04\x0b\xd2\xcc\x3c\x0e\x07\xdc\x34\x83\x1f\x89\xe2\x1a\x85\xe8\xb8\x1b\x05\x81\x07\xf6\x05\xb9\x41\xe0\x29\xf6\xe4\x56\x49\xe9\xe1\xfa\x05\x7c\xfe\x1a\xe3\x52\xd1\x3b\x33\xde\x57\x45\x09\xb8\x1a\x81\x72\x64\x27\x36\x96\x29\x30\x17\xef\x7a\xca\x4c\x89\xc4\xc0\x2e\xab\x80\xa8\xb2\xf6\xa6\xbf\x34\xfb\xd6\xda\x77\xe7\x7a\x5e\x9d\x00\xe1\x03\x05\xd1\x73\xdb\x91\x56\x18\x55\x8a\x14\xfa\xe8\xdb\x7f\xfb\xfa\xdb\xaf\xff\xfe\xdb\x3f\xff\xf2\x9b\x9f\xfe\xb3\xb1\x42\x6e\x2a\x8c\xb5\xfc\xed\xaf\xbf\xfe\xe6\xa7\xff\xf1\xea\x2f\xbf\xb2\xed\x05\x51\xa6\xf7\x85\x9e\xd6\x74\x81\xdb\xae\x4b\xaf\x3b\x6b\xd2\x7c\x63\xfb\xe6\xbc\x03\x66\xd8\x37\x3f\xfd\xea\xd5\x8f\x7e\xfe\xea\x6f\x7e\x51\x04\x3c\x1a\x75\x8d\xcc\xbe\x76\x0d\x57\x9d\x1e\x9a\x0e\x94\xa2\x5c\xe3\x0f\x56\x0c\x51\xee\xa3\x06\x5c\xf4\x2a\xea\x16\xd5\x9b\x01\x1d\xb5\xdd\x5b\xbd\xde\x8e\x92\x28\xf7\xbd\xe2\xa1\xbb\x5b\x3e\xed\xe9\xec\xbd\x0e\xcc\x9b\xb7\x3a\x75\xef\x23\x22\x17\xb4\x76\xf7\x51\x3a\x8e\x9c\xe0\x48\x14\xa9\xea\x5b\xb6\x71\xca\xa9\x24\xa2\x04\x7e\xaf\x00\xa6\x4c\xf9\x7a\xe3\xea\x56\xaa\x5b\xba\xa5\x15\x97\x55\xf3\x5e\xe9\x50\x2e\x65\xbc\x5f\xc7\xaf\xa4\x99\xbe\xfd\x6a\xf8\x94\x2b\xa3\xb0\x35\x01\x97\x20\x51\xad\x55\x90\xe8\x4c\x27\x5a\xc3\x94\x8d\xe8\xb8\xd6\x53\x34\x69\x54\x44\x3e\xc9\xe5\xc3\x42\x8d\xf7\xce\x68\x06\x3c\xcd\x25\xd1\xc6\xed\xbd\x6d\xfd\x5f\x71\x84\x18\x57\xc7\x66\xfb\x66\x44\x2e\xa0\x7d\xf2\x69\x1f\x4e\x37\x3b\xdb\x37\x83\x7e\x37\xc8\x3c\x85\x4c\xf2\xda\x2c\x92\xab\xa6\xa8\x31\x73\xa4\x26\xd5\x91\x18\x79\x29\x09\x67\x38\x7e\x98\xf5\xce\x5c\x03\x04\x9b\x62\xa2\x5c\xd3\xe4\x92\x66\xdb\x69\x46\x58\xc4\x84\x8a\xf9\xe0\xfb\x9c\x88\x34\xbe\x20\xbc\x5b\x6f\xae\x62\x26\x9f\x29\x57\xcb\x29\x9c\x67\x49\x12\xed\x3f\xcf\x16\xd6\x36\xd7\xfe\x59\x3d\x61\xaa\x06\xce\xfa\xe8\xdb\xdf\xfc\xf8\xd5\x57\x3f\xd7\x47\xc9\xc4\xe0\xae\xc2\x51\x6c\x7e\x66\x5b\xdd\x36\xac\xe2\xcd\x95\x66\x3f\xe8\x21\x86\x4d\x9a\x45\x82\x87\xd0\x6f\x6f\xdf\xec\x77\x5c\xc5\x9f\xb1\xd7\x51\x92\x01\xcb\x3e\xa2\x3e\xe6\x15\x01\x6a\xa6\x82\x37\x15\x01\xaa\xf5\xd5\x57\x3f\xfb\xed\x6f\x14\x89\x73\x58\xe3\x47\x6a\x1f\x12\x5f\x8c\x17\xf3\x3e\x0b\x62\xf2\xba\x81\x7d\x2d\xdf\x73\x98\x72\xed\x84\x54\x29\x9b\x10\x3c\x1b\xff\xa4\xbc\xc8\x7d\x86\x59\xae\x6c\x67\x75\x3b\xea\x36\xc6\x83\x4f\x3e\x57\x46\xa2\x96\xe4\x6e\xe2\xef\x9a\x5c\x72\x23\xdd\x63\x83\x6b\x33\x81\x5a\x20\xbe\x18\xaf\x8b\x8d\xbb\x83\x84\x32\xc5\x80\x22\x3e\xee\xde\xec\xee\xec\x74\x02\xaf\x49\x87\x67\x1c\xfb\xdb\x8b\xd0\xbb\xd7\x41\x77\x97\xc0\xf7\xd6\x81\xeb\x7b\xd1\x5a\x48\x59\x3f\xf8\xab\xb7\x6d\x3c\x26\xd1\x80\xb2\x81\x0a\x6f\x96\x15\xe8\xe1\x93\x47\x2f\x1e\x1f\x3d\xeb\xed\xed\xec\x54\x1a\x73\x0a\xfe\x64\xe9\x5e\xaa\x16\x62\x37\x22\xa5\x32\xb6\x1e\xf1\x34\x59\x56\x86\xd3\x09\x61\xcd\x20\xfa\x2d\x82\xe2\x84\x24\xba\x38\x21\xb2\x57\xb6\x46\xba\xb5\xa6\x47\x65\x4a\x4e\x93\xca\x42\xf0\xa7\xd0\x85\xee\xee\xbb\xdb\x3b\xdb\x3b\xdb\x5d\xe8\x76\x77\xbb\xdd\x26\x90\x3e\x54\x9e\x76\x48\xbd\xaa\x67\xf9\x16\xae\xd2\xf6\x68\x0e\xb5\x9b\x48\x1b\x1e\xda\xf8\xaf\x79\x4d\x3d\xaf\x2e\xfa\x6b\x38\x25\x07\xaf\xcf\x3d\xf0\xb6\xbc\x26\x29\xd7\xdc\x62\xaf\xbe\xc7\x4e\x48\x92\xf2\xcb\x86\x1f\xb9\x60\x04\x6d\xba\x6a\xf5\xc2\xf4\xa5\xea\x77\x2c\xae\x0b\x68\x21\x47\xb2\x78\x17\xdf\x5c\xe4\x5b\xad\x83\x11\x39\x4d\xf9\xf9\x20\x4c\x19\x23\x61\x23\xdb\xc3\x88\x54\x5b\x5b\x53\x19\xb6\x65\xad\x82\x11\x29\xaf\x83\xab\xd7\x33\xed\xb2\xd1\xcf\x6b\xcd\xea\x75\x8d\x29\x73\x00\xdc\x6a\x5a\xf0\x99\xcc\xed\xed\x7e\xa9\x0c\xea\xd5\x0f\xee\xed\xb2\x16\xc1\xb5\x73\x17\xaf\x76\x4e\xa5\xc4\x16\x10\x2c\x17\x54\x34\x52\x19\x15\x1f\xca\x44\x85\x49\x43\x19\xab\xdd\xdc\x91\xeb\xeb\x23\x5c\xb3\xc2\xa5\xd2\x92\xb5\xc5\x11\x46\xcb\x0c\xd2\xd1\x20\xcc\xf2\x41\xa8\x74\x4d\xb9\x19\x85\x02\x3a\x7e\x51\x28\xea\xea\x0c\x29\x4d\x9d\xa4\x11\x89\x41\x79\x87\x6b\x55\xb5\x29\x8a\xa8\x69\x6c\x8b\xab\xa1\xae\xcd\x99\x0e\x31\x63\x2a\x28\x1c\x12\x18\xa5\x39\x8b\x74\xe9\x52\x71\x07\x68\x73\xb6\x58\x9c\xe7\xc3\x9c\xc9\x5c\x27\x6d\x3f\xcf\x89\xd0\x6c\x0b\x6e\x75\xf7\xf6\xf6\xdf\x0d\x74\xde\x36\xc9\x63\x49\xb3\x98\xf8\x34\x13\x7e\x2e\x54\x93\xae\xb6\x52\x9d\x36\xf9\x61\x42\x6c\x3c\x14\x69\xac\xcc\x6e\x86\xe5\x44\x85\x76\xc6\x7d\x81\x11\x8d\x89\xd3\x7a\x76\xf8\xf4\xe1\xf1\xf3\xe3\xbb\xcf\x3f\xe9\x9d\x45\x94\x6b\xaf\xdf\x6d\x2b\x07\x2f\xa6\xec\x1c\xfc\x11\xb8\x3b\x9d\x33\xa7\x75\xf8\xe4\xe8\xc1\xc3\x8f\x07\x1a\xd0\xad\x06\x21\x4d\xcc\x40\xb9\xcf\x3a\xf4\x5d\x61\x08\x5b\x0a\xa2\xd6\xa8\x5e\xcf\x9c\xd6\x0a\x31\x6f\xad\x90\xdb\x16\xcb\x13\x7b\x47\xf1\x90\x1d\x6a\xca\x7b\x67\xc5\xf6\xfb\x0c\x3c\xb7\xe7\x81\x5b\xa3\xee\xcc\x69\x59\xdd\xef\x9d\x78\x8e\xd3\x2a\x13\xcc\x35\x18\xe8\x3b\xad\xab\x46\x09\x9a\x09\x6d\x9d\x96\xc9\xd8\x00\x40\x4b\x71\x34\x26\xc7\x94\x8d\x95\x21\xb5\x6b\x00\xff\xf3\x10\x76\xc1\x2d\xb2\x00\x06\xb6\x76\x06\xfc\x07\x41\x79\x0e\x80\xb0\x48\x11\xde\x33\x36\x0d\x82\x4f\xb9\x94\x01\x28\xe7\x1c\xda\xda\xf0\x2d\x2c\xcc\xeb\xf5\xba\x9d\x59\x35\x0a\xcd\xcb\x4b\x14\xa4\xeb\xbb\x52\x21\xcd\x29\xf1\x2c\x05\x5e\x91\xfb\x63\xe3\xb2\x3a\x01\x01\x2a\x66\x06\x7d\x60\x0c\x8d\x1d\xfb\xef\x0a\x7e\xba\xed\xf6\x0a\x6a\xfc\x6e\xa7\x18\x53\x18\x53\xf3\xaf\x5b\xb1\xc6\x69\x99\x6c\x06\xac\x18\x5f\x38\x2c\x0a\xa6\xcc\x68\xb4\x5a\xf5\x72\x2e\x74\x8a\x54\xe3\x88\x6a\xc6\x33\x62\xf6\xa5\x3c\xdc\x22\xe8\xcf\xb7\x4e\x83\xf9\x69\x30\x5e\x7f\xb2\xb3\x64\xb7\x1e\x3a\xb5\x36\x94\xcc\x47\x58\x62\xa7\x65\x2f\x2f\x7a\x28\x4b\x76\x21\xa6\x42\x22\xa7\xa5\x7a\x7a\x48\xb9\x46\xa6\xb3\xa3\xa4\x76\x23\x65\xf1\x25\x58\x8f\x4c\x0f\x56\x7b\x74\x81\x39\xc5\xc3\x98\xc0\x04\x0b\xc0\xf6\x8c\x3b\xad\x0d\x9d\x80\x13\x93\x34\x8f\x23\x98\x60\x16\xc5\x04\x42\x2c\x88\x00\x5d\x5e\x05\x6a\x32\x9b\xf8\x29\x6b\x8d\x1c\xc3\xa5\x5a\x69\x62\xa1\x18\x9c\x56\x6b\x43\x48\xcc\x65\x31\x3d\x65\x63\x43\x41\xca\x4c\xaa\xe7\x96\x02\x89\x52\xe6\x95\x20\x26\x99\xb4\xab\xbb\x85\x4d\x53\xf5\xdc\xb6\x55\x30\x65\xd9\xa4\xc4\x34\x56\x53\x6e\xde\xaa\x65\xa6\x37\x77\xb5\x44\x98\x88\xba\xe5\xd5\x4c\xa4\x6a\x6e\x55\x66\x12\x67\xd9\x91\xc9\x6e\xd5\xcd\xe4\x1c\x16\xc0\x74\xd6\xb9\x76\x7f\xb0\x0c\xa1\xd4\x67\x01\xb3\xbf\x06\x26\xab\xd0\xbc\x67\x41\x16\x20\xd4\x16\xe7\xa2\x00\xea\xee\xac\x86\xe2\x44\xf3\xb2\x04\xdb\x5d\x0d\x96\x17\xf1\xbd\x81\xba\xb5\x1a\xca\x38\x22\x25\xd4\x3e\xb8\xdd\x77\x57\x43\x4e\xb1\x0c\x27\xe6\x0c\x1a\xd8\xf7\xb5\xa3\xd1\x04\xb3\x5e\x61\xab\xb5\x91\xe0\x73\x02\x22\xe7\x3a\xd1\xc8\x49\x92\x5e\x14\x25\x02\x3a\xcf\x87\x59\x64\xcf\x3d\x65\x80\x39\xc7\x97\x8e\x3d\x36\x26\x25\xae\xb6\x7b\x6d\x42\xbc\xa5\x8b\x4e\x5a\xad\x0d\x96\x1a\x31\xd2\x46\xa6\x42\xb0\x72\x8c\x31\x3f\x1c\x27\xbf\x47\xb9\x58\x42\x92\xd7\x28\x17\xbb\x05\x57\xaf\x55\x23\xf6\x86\x45\x61\xeb\x2f\x39\xd6\xc2\xf6\x17\xef\x30\x56\xdd\x73\xd4\xf0\x5e\x5f\x23\x56\x03\xbc\xbe\x52\xac\x06\xd8\xac\x17\xbb\x86\xd0\xb7\x88\x27\x38\x09\x09\x93\x03\x1c\x86\x69\xce\xe4\x20\x4e\xc7\x74\xd5\x65\x63\x99\x2c\xc4\x42\xc6\xe9\x18\x7c\x09\x7b\xfb\xdf\xbd\x2f\x75\x5a\xab\x3d\x6a\x2c\x60\x4d\x0c\xb5\x00\x53\x44\xe7\x65\x3c\xb5\x74\x23\xd5\x47\x11\x96\x55\x92\xe8\x36\x02\xe4\xee\xab\x5f\xef\xaa\x5f\xef\xa9\x5f\xef\x5b\x1e\xd5\xc6\xd5\x52\x48\x6f\xc4\xb4\x88\x96\x4c\xda\xa8\x7e\xe0\x3e\x93\x84\xc3\xff\x4f\x73\x0e\x4f\x15\x0c\x1c\x63\x21\xa6\x29\x8f\x00\x3e\xb9\xff\xf4\x7e\x05\xab\x99\x1b\x51\x51\xf4\xf7\x3c\xef\xad\x51\x15\xb8\xaa\x2a\x75\xe3\x16\xe9\x46\x3f\x8c\x69\xe7\xa0\x2a\xa4\x66\xba\x90\xba\x36\x33\x5a\x28\x07\xac\x21\xb2\x90\x65\x4d\x3b\x86\x85\xa1\x2b\x12\x09\x8d\x11\x0f\x8f\x1e\x3c\x69\xc8\x88\x09\x37\x35\xcc\xe0\x82\x70\x41\x53\xd6\xbf\xb2\xfe\x3e\x89\x06\x61\x4c\x09\x93\xa2\xde\x24\x62\x7c\x41\x44\xff\x2a\x17\x24\x1a\xd8\x88\x6f\x92\x27\x98\xf5\xaf\xf4\x1d\x7d\x3d\x6a\x1a\x28\x61\xa6\x17\x24\xaa\xfa\x34\x1d\xa2\xd4\x51\x51\xb3\xf2\x7c\x55\x6d\x6d\x11\xf4\xa1\x32\xbe\x5d\x28\x57\x2f\x03\x58\xe4\x15\x6f\x4c\x87\xb3\x8b\x92\x34\x6b\x48\xd2\xea\xa4\x84\x50\xa1\x7c\x1e\xab\xb5\xf3\x5a\xdc\x57\x88\x42\x51\x37\xa6\x38\x26\xcb\x4a\x0d\x21\x71\x78\x9e\x5e\x10\x3e\x8a\xd3\xe9\x82\xe7\xdf\xdd\xbb\xf5\xfe\xce\x7e\x30\x49\xa7\x7e\x94\xfa\xd4\x57\xde\x8c\x8f\xe3\xd8\x57\xf8\xfd\xcf\xd2\xa1\xf0\x47\x29\xd7\x2d\x3a\x0d\xb2\x81\x99\x98\x12\xee\x77\xf7\xde\xed\xbe\xbb\x57\x9f\x5a\x3f\xaf\x2c\x44\x59\x13\x7a\xae\xc9\x8c\x02\x90\x85\x0c\x19\x29\x52\x64\x86\x51\xf5\x9e\x70\x92\x9a\x0e\x45\xae\xc4\xc3\xfa\xe4\xa6\xc5\x74\x2f\xa5\x40\x2d\xa7\x9e\xe9\x6a\x07\x7f\x4a\x23\x52\x0c\xd0\xf1\x8b\xb6\x95\xaa\x01\x3e\x4b\x87\x10\x51\x4e\x42\x99\xf2\xcb\x6d\x38\x9c\x60\x36\xd6\x17\x5c\x82\xe8\xbc\xe6\xa5\x3a\x71\xa6\x68\x62\xbb\x60\xff\xd3\x27\x47\xcf\xef\x7e\xd4\xf3\x02\x22\xc3\xc0\xa2\xf5\x6a\x9d\xf7\x1e\x3e\xad\x75\x6e\x47\xa6\x4f\x64\x69\x1a\x0f\xaa\xfe\x0b\xcc\x03\xdd\x66\xeb\xbc\x4a\x9a\xb5\x97\x0c\x8a\xd4\x70\x82\x39\x0e\x25\xe1\xdb\x70\x97\xb1\xf4\x52\xf5\x5c\x02\x23\x4a\x78\x31\xbf\x34\x04\x49\x3c\xec\xb9\xed\x5a\x11\x9e\x44\x9d\x02\xd7\xc7\xf4\x82\x30\xc0\x20\x24\x27\x38\x81\x74\x54\x72\x41\xfb\x80\x5b\x40\x5e\x86\x71\x1e\x11\x60\x29\xf3\x4b\x7e\xd8\x3e\x4e\xb2\x18\x87\xc4\xa2\x9a\x4e\xa8\x24\x22\xc3\x21\xa9\xc8\x12\x30\xa5\x72\xa2\xf0\x1b\x9a\x75\xf7\x96\x66\xae\x75\x4f\x30\xbb\x34\xad\x42\x17\xd6\x2a\xc6\x5a\x7c\x43\x32\xa6\x8c\x29\xaf\x35\x1d\x01\xc1\xe1\x44\x4f\x6b\x96\x54\xde\xff\x85\x31\xc1\x4c\x9f\x85\x81\x26\xca\x66\xb9\xcc\x4f\x2d\x0c\xd3\x0e\x6f\xad\x06\xd3\xfc\xb8\x56\x9c\x00\xb9\x33\x05\x31\x47\x70\xb5\x64\x2d\x5d\x52\x66\x57\x7d\xaa\x2f\x75\xfd\x44\xb9\x65\xe0\x7d\xda\x76\xaf\xfa\xe2\xe6\x86\xfa\x75\x72\x72\x07\xc7\x2c\x4f\xee\x9c\x0e\x4e\x37\x7b\x1d\x6f\x15\xa2\x22\xb8\xf4\x39\x19\x93\x97\x99\x4f\x5e\x4a\xc2\x22\x12\xe9\x0f\x90\xc4\x66\x00\xfa\xe3\xa3\x37\x1b\xf7\x29\x04\x41\x65\xa0\xcc\x47\x49\xda\x50\x5d\xb3\xc3\x8a\x67\x24\x5a\xdc\xe9\x92\x19\x6a\x4b\xe4\x04\x4b\x30\xb1\x01\x65\x46\x02\xaa\x9d\xe1\x39\xf3\x33\xcc\xa5\x00\xab\x30\xcd\x8e\x9a\x1a\x02\x73\xfd\xab\x47\x6f\xa9\x60\x66\x6a\xf6\x4e\x89\x8d\x3e\x58\x94\xe9\xdb\xe1\x12\x87\xc5\x59\x1e\x31\xc0\xc2\x5e\x72\x4f\x55\xd0\x53\x2a\x3c\x20\x2f\xb3\x98\x86\x54\xc6\x97\x0b\x32\x10\xa7\xe9\x79\x9e\x0d\x78\xce\x74\xfd\xd3\x1b\xca\x80\xde\x4b\x65\x89\x56\xc8\x42\x7d\xef\x53\xf0\xaa\x75\xb7\xfd\x59\x77\x6b\x77\xde\x7f\xb6\x09\x9d\x9b\xfd\x67\x9b\x9e\x3d\x50\xc5\x8f\xb9\x88\xd7\x5f\x21\xcd\xf4\x04\x2b\xea\xc8\xaf\x13\xc3\x06\x48\x59\xed\x5d\xff\xd1\x22\x3f\xd2\x55\x73\xeb\x88\xd7\xa5\xe4\xa3\xae\xbf\x0f\x7e\xe4\x81\xd7\x59\x8d\xe3\xb3\x74\x38\x88\x28\xaf\x23\xa9\x53\xbd\x70\x43\xaa\xeb\xa1\x9c\x25\x4c\x76\xbd\x7a\xfe\x3a\xd6\x6b\x97\x0d\xb6\x5c\xba\x1c\x50\xc8\xc7\x12\x92\xe0\xa6\xde\x37\x25\x27\x37\x03\xf8\x40\x45\xc0\xa6\x76\x81\x30\xf9\xe1\x1a\xf7\x56\x5f\x8a\xfb\xa3\x06\x32\x35\x81\x21\xe9\x9d\x77\x1a\x5c\xaf\x71\x73\x0e\x4b\x03\x56\xce\x60\x4b\x7e\x17\x96\x43\x9d\x35\xaf\x2b\x4e\xe6\x73\x92\x64\x29\xc7\xfc\xd2\x1c\x8c\x82\x15\xe5\x91\xb4\xda\x9b\x24\x59\xcf\x6d\x27\xe7\xea\xa1\x03\x57\x57\x40\x5e\x52\x09\xdd\x02\xcb\xdd\x28\x02\x1c\xc7\xea\x60\xab\x73\xa5\x2c\x76\xa9\x49\xad\x5d\x5a\x36\x6f\x85\xa5\x2a\x72\x58\xc8\x9d\x59\xa3\x65\x24\x67\x41\xab\xc2\xd5\xd2\x21\x83\x0f\x91\x3b\x53\x24\xcd\xd1\x5b\x51\x52\xb3\xa9\xab\x69\xb9\xf7\xf0\xa9\xde\xf8\x15\xd4\x7c\x58\xcd\x7d\x8d\x48\xd4\xf0\x35\x2c\xeb\xdb\x63\xad\xad\x53\xeb\x34\xe5\x10\x79\xa2\x64\x6c\xdb\xa8\x2d\xf2\x92\x0a\x29\x3a\xdb\xf0\x90\x09\x5d\x01\x34\x21\x05\xa4\x4e\x82\x0e\x89\x9c\x12\xc2\x6a\xfa\x74\x44\x2f\x88\x2e\x6c\x07\x23\x82\x5a\x9b\xaa\x61\x56\xbd\x6e\x57\x16\xb6\xd0\x65\x0a\x21\xd4\x3e\xa0\x30\x0b\x2e\x7d\x20\xf0\x63\xf0\x73\xb5\x74\x05\x38\x47\xb0\xfb\x61\x10\x91\x8b\x80\xe5\x71\x6c\x2d\xcc\x46\xa5\x06\x96\x76\x7b\x01\xe0\x7a\x03\xd4\x6e\x2b\x35\xb8\xd9\x99\xdd\x9e\x77\xda\xdb\x9b\xfa\x62\xd5\xce\x0b\xfd\xbd\x00\xd5\x19\x5b\x28\xfb\x94\x11\xf8\x00\xdc\xb6\x56\x51\xbe\x55\x63\x5d\xf0\xfd\x88\xc4\x34\xa1\x92\xf0\xde\x1d\xd0\x0e\x52\xa6\x62\x87\xa8\xf4\x57\x9e\x98\xfc\xa8\x61\x4e\x1c\x6b\x9f\x7f\xf1\xd4\x58\xa7\xd4\xba\x6a\xe2\x0e\x1c\x2b\x2f\x5a\xd8\x54\x2c\x50\x26\x53\x5d\x9c\xe4\x2c\x8a\x89\xdd\xfc\x45\x8f\x7f\xc5\xf5\xc3\xea\x98\xba\x9e\x56\x78\xf5\xe5\x9f\xd6\x62\xd4\xb5\x9f\xc9\xea\xea\x8e\x9f\xfc\xf2\xb5\xbf\xc5\xd5\xe0\xff\xf8\xda\xdf\xe1\xea\xda\x91\x2f\x17\x2b\x71\xaf\x03\x7f\xf5\xe3\x5f\x2c\x56\xe8\x5e\x07\x7e\xf8\xf8\x9e\x05\x5f\xc7\x14\xcd\x98\x51\xca\xdb\xb4\xb7\x7f\x40\x3f\xe8\x1d\x3d\x38\xd8\xdc\xa4\xb6\xc0\x76\xd4\x46\x37\x04\xa0\x2d\xb8\xae\xb0\xde\xa5\x1d\xe8\x2c\x7d\x22\xbc\xc8\xf4\x3e\x5a\x4f\xa8\x49\x93\x2f\x77\x5e\xfb\x99\x71\xf3\x2a\x68\xab\xcf\x82\xc6\xcd\x1c\xac\xba\x9d\x33\x31\x71\xa2\x64\x3a\xe5\x21\xa9\x84\x4a\xc5\x6e\x8e\x98\xe2\xf2\x43\xbb\xe5\x40\x49\x59\x87\x69\x58\x6b\x9c\x86\xba\x6d\x45\xec\xe4\x00\x28\x5c\x8f\x28\x23\x87\x69\xce\x96\x6f\xe4\x55\xaf\xd2\xda\xae\x46\x08\x7e\x5c\xdd\x28\x21\xb7\x31\x14\x81\x3f\x96\xd0\xad\xd2\xc6\x00\x6b\xef\xf9\x0d\xd6\x06\x8f\x6a\x57\x7c\xcd\xb4\x9a\xb2\x32\xf5\x12\xe3\xe2\x23\x5b\x54\xfb\xf4\xbc\xf8\xa8\xb6\x5e\x5b\xee\xee\x2d\x57\x96\x97\x95\x40\x19\xa7\x29\xa7\xb2\xfc\x70\xf7\xb6\xc9\xcf\x40\xb3\xb0\xe7\x75\x13\x33\x4b\x1f\xd2\x9d\xac\x00\x2a\xae\xd0\xf2\xec\x8f\xba\x14\x38\xcd\xe5\x9b\xd5\x02\xcb\xff\xae\x5a\xe0\x34\x97\x6f\x56\x0c\x2c\xdf\xa2\x18\x38\xcd\x65\x55\x0d\xac\x66\x04\x5f\x99\x8b\xd7\xac\x07\x56\xa3\xcf\xeb\xc3\xff\xa0\x15\xc1\xd5\x6c\x7f\xb4\x25\xc1\xca\xb9\x28\x52\xd0\xd7\x24\x9f\xab\x74\x5d\x43\xf7\xd3\x11\xb4\xdd\xbd\x0f\x7a\xb7\xde\x7f\xbf\xa3\xfd\x94\xe7\x97\x19\xe9\x21\xe3\xa1\xa2\x83\xa6\xa1\x30\x7f\x6f\xa0\x82\x52\x4f\x8b\x30\xf5\xef\x80\x6a\x7f\x58\xa3\x18\x64\xbf\x4e\x58\xba\x43\x28\xcb\x7f\x26\x69\x4d\x4f\xee\x57\x45\x0a\x1e\x7c\x50\x77\x82\xea\x3a\xce\xdc\x19\x16\x69\x1d\x77\xb6\x61\x56\x3f\xaf\x15\x79\xd9\xd2\x81\xe5\xea\xae\x92\x51\x63\x22\x09\x93\x60\xb0\xaf\xfe\x1b\x02\xdf\xcd\xac\x75\x0c\xfa\x83\x31\xa5\x53\x66\x91\x5f\x37\xff\xbe\xb1\x01\xbf\xfd\xf5\x5f\xfc\xee\xef\xfe\xe1\x9b\xbf\xfe\x97\xff\xfc\xb3\x7f\x7a\xf5\xab\x7f\xfd\xe6\x57\x5f\xbf\xfa\xd1\xcf\x1c\x21\xaa\x6a\x20\x93\xee\x8e\x00\x05\x69\x26\x83\x51\xca\xc4\x05\x25\xd3\xe0\xe8\x7e\x20\x44\x10\x61\x89\xb5\x0d\x2c\xb3\xdf\xd0\xa8\x1e\xba\x6e\x48\x60\x13\xd7\xff\x9b\xab\x8a\xb4\x8d\x7b\xf6\xe8\xfe\xfd\x63\x9b\x3c\xb1\xfa\x79\x47\x73\x4f\xb1\xef\x47\xcb\xec\xfb\xf6\x6f\xff\xc4\x19\xb1\x43\xfb\x67\x38\xdc\x2e\x72\x1c\x77\x56\x34\xe8\xaa\x93\xa2\xf2\xd1\xf9\xaf\x00\x00\x00\xff\xff\xc0\x82\xff\xe3\xb6\x47\x00\x00")

func shellLinux_json_apiShBytes() ([]byte, error) {
	return bindataRead(
		_shellLinux_json_apiSh,
		"shell/linux_json_api.sh",
	)
}

func shellLinux_json_apiSh() (*asset, error) {
	bytes, err := shellLinux_json_apiShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "shell/linux_json_api.sh", size: 18358, mode: os.FileMode(493), modTime: time.Unix(1508753189, 0)}
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
	"shell/linux_json_api.sh": shellLinux_json_apiSh,
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
	"shell": &bintree{nil, map[string]*bintree{
		"linux_json_api.sh": &bintree{shellLinux_json_apiSh, map[string]*bintree{}},
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

