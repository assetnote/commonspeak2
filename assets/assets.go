// Code generated by go-bindata. DO NOT EDIT. @generated
// sources:
// data/.DS_Store
// data/filters/numerical-parameters.txt
// data/filters/string-parameters.txt
// data/sql/.DS_Store
// data/sql/github/deleted-files-test.sql
// data/sql/github/deleted-files.sql
// data/sql/github/nodejs-routes-test.sql
// data/sql/github/nodejs-routes.sql
// data/sql/github/rails-routes-test.sql
// data/sql/github/rails-routes.sql
// data/sql/github/tomcat-routes-test.sql
// data/sql/github/tomcat-routes.sql
// data/sql/github/words-with-ext.sql
// data/sql/hackernews/subdomains.sql
// data/sql/http-archive/directories.sql
// data/sql/http-archive/subdomains.sql
// data/sql/http-archive/technologies.sql
// data/sql/http-archive/words-with-ext.sql
package assets

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

var _dataDs_store = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x98\x4d\x6a\xc3\x30\x10\x85\xdf\x28\x2e\x18\xba\xd1\xb2\x4b\x5d\x21\x37\x10\x41\x3d\x81\xa1\xeb\x42\x0c\xe9\xc2\x89\xd3\x9f\x74\xed\x53\xf4\x6a\xbd\x4d\x29\x42\x2f\xb4\x25\x76\x77\x6d\xdd\xf0\x3e\x30\x9f\x61\x2c\x79\xd0\x42\x9a\x11\x00\x5b\x1d\xda\x25\xe0\x01\xd4\x28\xb6\xfc\x32\x42\xcd\xe7\x04\x47\x5f\xe4\xc1\x79\x8e\x65\xb3\xee\xb7\xfb\xf1\x59\x66\x4a\xce\x7d\x81\x47\xdc\xa3\xeb\x36\x27\xf9\x33\xb2\xed\x53\x6a\x0f\x4f\x6b\xe0\xf5\xe5\xed\xe6\x6b\xa4\x9d\x88\xec\xef\xa6\x66\x7b\x6e\x1e\x76\x5d\xbf\xdb\x94\x55\x13\x42\x08\x21\x7e\x19\x9e\x3e\xf5\xe5\x5f\x27\x22\x84\x98\x1d\x79\x7f\x08\x74\xa4\x87\x62\x63\xdc\xd1\xd5\xa7\x31\x9e\x0e\x74\xa4\x87\x62\xe3\x77\x8e\xae\xe8\x9a\xf6\x74\xa0\x23\x3d\x14\x73\xd3\x32\x36\x1f\xc6\x3f\x1f\x9b\x17\xf3\x74\xa0\xe3\xcf\xac\x8d\x10\xff\x9d\x45\x91\xcf\xe7\xff\xf5\x74\xff\x2f\x84\x38\x63\xac\x4a\x4d\x5a\x7d\x73\x1d\xe5\x58\x08\xdc\x1e\x07\xb0\x10\xc0\x48\x11\xe0\xca\x65\xe1\x15\x3e\xe2\x2a\x04\x84\x98\x19\xef\x01\x00\x00\xff\xff\xf8\x96\xbb\x9f\x04\x18\x00\x00")

func dataDs_storeBytes() ([]byte, error) {
	return bindataRead(
		_dataDs_store,
		"data/.DS_Store",
	)
}

func dataDs_store() (*asset, error) {
	bytes, err := dataDs_storeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/.DS_Store", size: 6148, mode: os.FileMode(420), modTime: time.Unix(1534152160, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataFiltersNumericalParametersTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x50\x4d\x92\xec\x20\x08\xde\x73\x9b\x77\xa2\x2e\x5a\xd1\xe6\x25\x11\x0b\x30\xd3\xb9\xfd\x14\xea\x62\x56\xf0\x29\xdf\x0f\x70\x06\x73\x2c\xe5\xc5\x19\xe8\xa6\xe6\xd1\xd8\x85\xea\xaf\xae\xf2\x7d\x02\x1e\xf4\xc0\x29\xb5\xe2\xfb\xa4\xc0\x45\x74\x5c\xd1\xb8\x74\x4e\xd1\x24\x69\x37\xa9\xa1\xb3\xb4\xc0\xdc\x32\x7d\x41\x29\x89\x66\xc0\x94\x64\x2c\xe1\x8e\x95\x80\xdb\xcd\xbe\x46\x5d\x0e\x6a\x11\x40\x1d\xa8\x65\xc8\xe8\x04\x5d\xe5\x3f\xa5\x39\xdf\xa4\x75\x95\xc2\x13\x0c\x23\x8d\x8a\x75\xc7\xcc\x6c\x69\x98\x6d\xcf\x42\x94\x5f\x63\x86\xb9\xae\xc5\x38\xe8\x59\xae\x9d\x74\x7e\x55\x95\xd1\xa7\xc6\xf0\xcf\x76\x4f\xe8\x54\x45\xd7\xa4\xb2\x2d\xb5\x95\x9d\x5b\xdd\xeb\x15\xd6\xeb\x6f\xe6\x30\xd9\x31\x3e\x62\x7b\x39\x8d\x17\x3b\x47\x85\x37\xa6\x63\x32\xf1\x3c\xa3\xfe\x03\x4c\xce\xf7\x52\x48\x92\x09\x32\x3e\x06\x8d\x7e\xe6\xb9\xcc\xc6\xbc\x6c\x95\xdb\x75\x53\xcd\xa9\xff\x06\x00\x00\xff\xff\xa5\xa2\x2a\xa0\x9e\x01\x00\x00")

func dataFiltersNumericalParametersTxtBytes() ([]byte, error) {
	return bindataRead(
		_dataFiltersNumericalParametersTxt,
		"data/filters/numerical-parameters.txt",
	)
}

func dataFiltersNumericalParametersTxt() (*asset, error) {
	bytes, err := dataFiltersNumericalParametersTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/filters/numerical-parameters.txt", size: 414, mode: os.FileMode(420), modTime: time.Unix(1534152261, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataFiltersStringParametersTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x91\x41\xb2\x1b\x21\x0c\x44\xf7\x7d\x9b\x24\x77\xc8\x0d\xb2\x9e\x92\xa1\x3d\xa6\x2c\x10\x91\x34\x4e\x26\xa7\x4f\x0d\xfe\x0b\xef\x1a\xaa\x1f\x7a\x85\xa6\xdb\xab\x55\x3a\x86\x74\x42\x4a\x36\x1b\xf8\x15\x74\xdc\xac\x9e\x98\xf4\x2e\xda\xc6\x13\xca\xbd\x85\x4a\xb6\x17\xb7\xa4\x77\x4c\x99\x74\x84\x1e\x3b\x8a\xf5\x69\x83\x23\xb7\x6f\x1f\xf9\xfb\x47\xfe\x01\xb5\x7d\x97\x9b\x72\xcb\x73\x12\xce\x69\x90\x5a\x9d\x11\x38\x29\x8e\x6e\x23\x1f\x50\x2b\xa2\x44\xda\x93\x03\x29\x7b\xe0\x08\xfa\xdb\xcd\xcb\x83\x0b\x5e\x3e\xe5\xc1\xf2\x84\xcc\xf6\xc2\xef\x83\x7e\xe2\xde\x94\xab\x58\xe5\xc4\x9d\xac\x10\xcf\x56\x94\xdb\xba\x2d\x36\xd2\x4d\xf5\x72\x4e\xc9\x23\x10\x2c\xce\x5c\x03\x50\xac\x12\x29\x37\x04\xfd\xd5\x0a\x91\xfc\x9b\xf8\xd9\x94\x5d\x86\xec\x57\xe3\x70\xe7\x28\x27\xba\x55\x2a\xa6\x4a\xde\xcd\x3b\x5a\x9d\x28\xa6\xe6\x38\x5c\xb7\xb7\xb8\xd4\xde\xc6\xf6\xe4\x89\xa2\x8d\x23\x11\xc7\x2d\x8a\xb7\xb9\x7e\x57\x99\x49\x47\x6f\xa3\x45\xbe\xc5\xaf\xf3\x85\x47\x8a\xe7\x4a\x1c\x15\xc1\x88\x0b\x90\x39\xb5\x15\x59\xf0\xe0\x9f\xf8\x7a\xe0\x1a\x2f\x89\x22\xc9\xdd\xfc\xdc\xd6\x2a\xcc\x77\x19\xed\xdf\x2a\xff\x0f\x00\x00\xff\xff\xd3\x06\x0c\xa1\xdc\x01\x00\x00")

func dataFiltersStringParametersTxtBytes() ([]byte, error) {
	return bindataRead(
		_dataFiltersStringParametersTxt,
		"data/filters/string-parameters.txt",
	)
}

func dataFiltersStringParametersTxt() (*asset, error) {
	bytes, err := dataFiltersStringParametersTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/filters/string-parameters.txt", size: 476, mode: os.FileMode(420), modTime: time.Unix(1534152277, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataSqlDs_store = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x98\x31\x4a\xc5\x40\x10\x86\xff\x59\x53\x2c\x68\xb1\xa5\xe5\x5e\xc0\xc2\x1b\x2c\x21\x9e\x20\x17\x50\x54\x8c\x20\x51\x42\x4c\x9d\xca\x73\x79\x34\x09\xfb\x8b\x4a\x12\x9b\xc7\xe3\xe5\x3d\xfe\x0f\x96\xaf\xc8\x4c\x12\xb6\xd8\x9d\x19\x00\x56\xbe\x3f\x5c\x03\x01\x80\x47\x36\x3e\xb0\x88\xe7\x9a\xe1\x68\xcb\x2b\x00\x57\xb8\x43\x87\x7b\x34\x78\x5e\x7e\xd7\x8c\x29\xf7\x02\x0d\x7a\xf4\x78\xfb\x93\x3f\xe0\x71\xa8\xbb\xf6\xe5\xb5\x7d\x62\x9c\x10\x42\x08\x21\x76\x87\x77\xaa\x3f\x3f\xf4\x8f\x08\x21\x36\xc7\x74\x3e\x44\x3a\xd1\x63\xb6\xf1\xb9\xa3\x8b\x5f\x39\x81\x8e\x74\xa2\xc7\x6c\x63\x9c\xa3\x0b\xda\xd3\x81\x8e\x74\xa2\xc7\x6c\x1e\x5a\xc6\xe6\xc3\xf8\x65\x63\x87\x62\x81\x8e\x74\xda\xcf\xde\x08\x71\xec\x9c\x65\x85\xe9\xfe\xbf\x59\xef\xff\x85\x10\x27\x8c\x15\x55\x5d\x95\xff\x0c\xd9\x1c\x0b\x81\x5b\xc6\x7c\x7e\x27\xae\x14\x02\x2e\x0f\x0c\x2f\xf1\x13\xa7\x62\x40\x88\x0d\xf1\x15\x00\x00\xff\xff\xeb\xa0\x56\x2e\x04\x18\x00\x00")

func dataSqlDs_storeBytes() ([]byte, error) {
	return bindataRead(
		_dataSqlDs_store,
		"data/sql/.DS_Store",
	)
}

func dataSqlDs_store() (*asset, error) {
	bytes, err := dataSqlDs_storeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/sql/.DS_Store", size: 6148, mode: os.FileMode(420), modTime: time.Unix(1534152160, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataSqlGithubDeletedFilesTestSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\x4f\x6f\xa3\x30\x14\xc4\xef\xfe\x14\x23\xb4\x12\xb0\x4a\x22\xed\x79\x4f\x2c\x71\xd8\x48\x89\x1d\x19\xa3\x36\xca\x1f\x02\xc4\x4a\x88\x20\x50\xdb\x54\xad\xc4\x87\xaf\xa0\x69\x4f\x3d\xce\xd3\xfc\xde\xbc\x79\xa1\xa0\x81\xa4\x90\x74\xbd\xe1\x22\x10\x5b\x2c\x12\x16\xca\x25\x67\x04\x50\x6f\x56\x67\x85\x4d\xdb\xcc\x5e\x3d\xd3\xe5\x37\x55\x58\x18\xab\xcb\xfb\xc5\x27\x80\xa0\x32\x11\x2c\x46\x2c\xc5\x92\x45\x04\x58\x05\x2c\x4a\x82\x88\xe2\x66\x10\xc4\x70\x1c\x87\x00\x5a\xd9\x4e\xdf\xf1\xc0\x67\xa6\xad\x4a\xeb\xb9\x70\xfd\xdd\x9f\x03\x71\x1c\xe7\x2f\x89\xe9\x8a\x86\x92\x00\x43\xce\x84\x00\x21\x4f\x98\xf4\x06\xe5\x0f\x7b\x8a\xa6\xbb\x5b\xb2\x10\x7c\x0d\x8f\x00\xdf\xf6\x9f\x0f\x1c\x91\x61\x40\x80\x81\x19\x8d\xa7\xbc\xbc\xbc\x74\x4a\xbf\x4f\xdb\x2e\xaf\xca\x62\x7a\xce\x6c\x36\xbb\x94\xf6\xda\xe5\xa9\x56\x6d\x63\x66\x26\xab\xdb\x4a\xa5\x45\x53\xd7\xa5\x35\x27\x02\x3c\xfd\xa7\x82\x8e\xb8\xa0\x11\x7d\xde\xa4\x21\x67\x32\x58\xb2\xf8\x2b\x6a\x02\xf7\xe8\x69\x55\x37\xaf\xfd\x59\x55\xca\xf6\x4a\x67\xc6\xf7\x54\xaf\xce\xbd\x32\xfd\xf0\x26\xec\x8e\xfb\xbd\x39\xfc\xfe\xe5\xfa\xf0\x49\x24\x78\xb2\xc1\xbf\xed\xa3\x2b\xe1\x62\x4e\xc5\xa7\x1e\x4b\x62\x4e\xe3\xf0\x23\x00\x00\xff\xff\x81\xbd\xa6\xc4\x92\x01\x00\x00")

func dataSqlGithubDeletedFilesTestSqlBytes() ([]byte, error) {
	return bindataRead(
		_dataSqlGithubDeletedFilesTestSql,
		"data/sql/github/deleted-files-test.sql",
	)
}

func dataSqlGithubDeletedFilesTestSql() (*asset, error) {
	bytes, err := dataSqlGithubDeletedFilesTestSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/sql/github/deleted-files-test.sql", size: 402, mode: os.FileMode(420), modTime: time.Unix(1535099018, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataSqlGithubDeletedFilesSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\xcf\x6e\xb2\x40\x14\xc5\xf7\xf7\x29\x4e\xc8\x97\x00\x5f\xd4\xa4\xeb\xae\x28\x8e\xd4\x44\x07\x33\x0c\x69\x8d\x7f\x01\x27\x8a\x51\xa1\x33\x43\xd3\x26\x3c\x7c\x03\xb5\x5d\x75\x79\x6e\xce\xef\x9e\x7b\x6e\x28\x58\x20\x19\x24\x9b\x2f\x62\x11\x88\x25\x26\x29\x0f\xe5\x34\xe6\x04\xa8\x0f\xab\xb3\xc2\xee\xea\xcc\x9e\x3c\xd3\xe4\x67\x55\x58\x18\xab\xcb\xdb\xd1\x27\x40\x30\x99\x0a\x9e\x20\x91\x62\xca\x23\x02\x66\x01\x8f\xd2\x20\x62\x38\x1b\x04\x09\x1c\xc7\x21\x40\x2b\xdb\xe8\x1b\xee\xf8\xc8\xd4\x97\xd2\x7a\x2e\x5c\x7f\xf5\xb0\x21\xc7\x71\x1e\x29\x61\x33\x16\x4a\x02\xba\x9c\x01\x01\x61\x9c\x72\xe9\x75\xca\xef\xf6\x14\x55\x73\xb3\x34\x11\xf1\x1c\x1e\x01\xbf\xf6\xbf\x0f\xec\x91\x6e\x40\x40\xc7\xf4\xc6\x7d\x5e\x1e\xdf\x1a\xa5\x3f\x87\x75\x93\x5f\xca\x62\x78\xc8\x6c\x36\x3a\x96\xf6\xd4\xe4\x3b\xad\xea\xca\x8c\x8a\xea\x7a\x2d\xad\xd9\x13\xf0\xf2\xcc\x04\xeb\x39\xc1\x22\xf6\xba\xd8\x85\x31\x97\xc1\x94\x27\x3f\x19\x03\xb8\x5b\x4f\xab\x6b\xf5\xde\x1e\xd4\x45\xd9\x56\xe9\xcc\xf8\x9e\x6a\xd5\xa1\x55\xa6\xed\xfe\x83\xd5\x76\xbd\x36\x9b\xff\xff\x5c\x1f\x3e\x45\x22\x4e\x17\x78\x5a\xde\x4b\x52\x2c\xc6\x4c\x7c\xeb\xbe\x1d\xc6\x2c\x09\xbf\x02\x00\x00\xff\xff\x24\xcb\x8a\x5b\x8b\x01\x00\x00")

func dataSqlGithubDeletedFilesSqlBytes() ([]byte, error) {
	return bindataRead(
		_dataSqlGithubDeletedFilesSql,
		"data/sql/github/deleted-files.sql",
	)
}

func dataSqlGithubDeletedFilesSql() (*asset, error) {
	bytes, err := dataSqlGithubDeletedFilesSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/sql/github/deleted-files.sql", size: 395, mode: os.FileMode(420), modTime: time.Unix(1535099018, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataSqlGithubNodejsRoutesTestSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x51\x51\x6f\xd3\x30\x18\x7c\xf7\xaf\x38\x2c\x60\xc9\xa0\x29\xaf\x30\x8a\x08\x99\x57\x82\x5a\x67\x72\x12\xa1\xb2\x8e\x36\x4d\xdd\xd4\x28\x4d\x42\xe2\x20\xa0\xeb\x7f\x47\x71\xda\xb1\x49\xbc\xf0\xf4\xd9\xdf\x7d\x77\xe7\xef\xec\x09\xe6\x46\x0c\x11\x9b\x5e\x07\xc2\x15\x33\x5c\xc5\xdc\x8b\xfc\x80\xa3\x4a\xea\x46\x8a\xb2\xd5\xb2\xb1\x6a\x53\xbc\x72\x2d\x11\x46\xc2\xe7\x63\x9b\x08\x16\xc5\x82\x87\x70\x85\x70\x67\x6f\xfb\xee\x3b\x4c\x5c\x3e\x8e\xdd\x31\xc3\xb7\x06\x6e\x08\x4a\x29\x21\xc0\x8f\xa4\x46\x2d\x33\xf9\x13\x23\x0c\x2d\x7a\x77\x66\xdf\x24\x83\xdf\xee\xe0\xcb\xe2\xd5\xe0\xf5\x7c\x3e\x9c\xcf\xdf\x3c\x79\xff\xf4\xd9\xd7\xe7\xe7\x96\xfd\x62\x74\x7b\x6e\x9d\xdd\x51\x7b\xb8\xcb\xd4\x05\x01\xd4\x06\x0f\xfc\x6d\xec\x09\x70\x92\x6c\xda\x5c\x63\x84\xbf\xb0\xb3\x4b\x74\xba\xb5\x8c\x99\xdd\xb1\x0f\x90\x79\x23\xff\x45\xba\xb9\x35\x38\x01\x6a\xa9\xdb\xba\x38\x22\x5d\x93\x52\x7a\x41\x48\xc8\x26\xcc\x8b\x3a\xbc\x93\x7f\x49\x00\x2f\x88\x79\xd4\x3f\xc6\xee\xd6\x4b\xcb\xb6\xd0\xe4\x4a\x04\x53\x58\x04\xb8\x27\xe0\x51\x78\xa9\x93\x96\x85\x96\x85\x36\x1c\xc3\x26\x40\xc7\x32\xa3\xcb\x95\xca\xbe\xb7\xb2\xfe\x35\xa8\xda\x55\xae\xd2\xc1\x3a\xd1\x89\x93\x29\xbd\x6d\x57\x8b\x5a\x56\x65\xe3\x34\xc9\xae\xca\xe5\xe2\x28\xd3\x2c\x91\x12\xe0\x53\xe0\x73\x63\xfb\xc8\x18\x50\x6b\x73\xb8\xd7\xff\x1f\x87\x8d\xca\x65\xb3\x34\xb4\xcf\x1f\x99\x60\x47\x81\x2a\xd1\x5b\x8c\x40\xd3\xb2\xd8\xa8\x6c\xd8\xc7\xed\xd4\x2b\x0a\x1b\x1b\x02\x04\xdc\x0c\x6e\x1c\xb5\xc6\x08\x69\x57\xec\xe3\xa7\x74\xb1\xc5\x9c\xb3\xf0\x61\x6e\x7d\x06\x27\x07\x73\x83\x1f\x82\x07\x11\x78\x3c\x99\x90\xb1\x08\xe2\x6b\x7c\x98\x9d\x40\x12\x88\x4b\x26\xfa\x86\xc9\x1c\x97\x2c\xf4\xc8\xc4\x9f\xfa\xdd\xd6\xfb\x7d\xae\x76\x4a\x1f\x0e\x7f\x02\x00\x00\xff\xff\x08\xac\xb3\x1f\xcd\x02\x00\x00")

func dataSqlGithubNodejsRoutesTestSqlBytes() ([]byte, error) {
	return bindataRead(
		_dataSqlGithubNodejsRoutesTestSql,
		"data/sql/github/nodejs-routes-test.sql",
	)
}

func dataSqlGithubNodejsRoutesTestSql() (*asset, error) {
	bytes, err := dataSqlGithubNodejsRoutesTestSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/sql/github/nodejs-routes-test.sql", size: 717, mode: os.FileMode(420), modTime: time.Unix(1534148741, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataSqlGithubNodejsRoutesSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func dataSqlGithubNodejsRoutesSqlBytes() ([]byte, error) {
	return bindataRead(
		_dataSqlGithubNodejsRoutesSql,
		"data/sql/github/nodejs-routes.sql",
	)
}

func dataSqlGithubNodejsRoutesSql() (*asset, error) {
	bytes, err := dataSqlGithubNodejsRoutesSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/sql/github/nodejs-routes.sql", size: 0, mode: os.FileMode(420), modTime: time.Unix(1534039257, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataSqlGithubRailsRoutesTestSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x51\x51\x6f\xd3\x30\x18\x7c\xf7\xaf\x38\x2c\x60\xc9\xa0\x29\xaf\x30\x8a\x08\x99\x57\x82\x5a\x67\x72\x12\xa1\xb2\x8e\x36\x4d\xdd\xd4\x28\x4d\x42\xe2\x20\xa0\xeb\x7f\x47\x71\xda\xb1\x49\xbc\xf0\xf4\xd9\xdf\x7d\x77\xe7\xef\xec\x09\xe6\x46\x0c\x11\x9b\x5e\x07\xc2\x15\x33\x5c\xc5\xdc\x8b\xfc\x80\xa3\x4a\xea\x46\x8a\xb2\xd5\xb2\xb1\x6a\x53\xbc\x72\x2d\x11\x46\xc2\xe7\x63\x9b\x08\x16\xc5\x82\x87\x70\x85\x70\x67\x6f\xfb\xee\x3b\x4c\x5c\x3e\x8e\xdd\x31\xc3\xb7\x06\x6e\x08\x4a\x29\x21\xc0\x8f\xa4\x46\x2d\x33\xf9\x13\x23\x0c\x2d\x7a\x77\x66\xdf\x24\x83\xdf\xee\xe0\xcb\xe2\xd5\xe0\xf5\x7c\x3e\x9c\xcf\xdf\x3c\x79\xff\xf4\xd9\xd7\xe7\xe7\x96\xfd\x62\x74\x7b\x6e\x9d\xdd\x51\x7b\xb8\xcb\xd4\x05\x01\xd4\x06\x0f\xfc\x6d\xec\x09\x70\x92\x6c\xda\x5c\x63\x84\xbf\xb0\xb3\x4b\x74\xba\xb5\x8c\x99\xdd\xb1\x0f\x90\x79\x23\xff\x45\xba\xb9\x35\x38\x01\x6a\xa9\xdb\xba\x38\x22\x5d\x93\x52\x7a\x41\x48\xc8\x26\xcc\x8b\x3a\xbc\x93\x7f\x49\x00\x2f\x88\x79\xd4\x3f\xc6\xee\xd6\x4b\xcb\xb6\xd0\xe4\x4a\x04\x53\x58\x04\xb8\x27\xe0\x51\x78\xa9\x93\x96\x85\x96\x85\x36\x1c\xc3\x26\x40\xc7\x32\xa3\xcb\x95\xca\xbe\xb7\xb2\xfe\x35\xa8\xda\x55\xae\xd2\xc1\x3a\xd1\x89\x93\x29\xbd\x6d\x57\x8b\x5a\x56\x65\xe3\x34\xc9\xae\xca\xe5\xe2\x28\xd3\x2c\x91\x12\xe0\x53\xe0\x73\x63\xfb\xc8\x18\x50\x6b\x73\xb8\xd7\xff\x1f\x87\x8d\xca\x65\xb3\x34\xb4\xcf\x1f\x99\x60\x47\x81\x2a\xd1\x5b\x8c\x40\xd3\xb2\xd8\xa8\x6c\xd8\xc7\xed\xd4\x2b\x0a\x1b\x1b\x02\x04\xdc\x0c\x6e\x1c\xb5\xc6\x08\x69\x57\xec\xe3\xa7\x74\xb1\xc5\x9c\xb3\xf0\x61\x6e\x7d\x06\x27\x07\x73\x83\x1f\x82\x07\x11\x78\x3c\x99\x90\xb1\x08\xe2\x6b\x7c\x98\x9d\x40\x12\x88\x4b\x26\xfa\x86\xc9\x1c\x97\x2c\xf4\xc8\xc4\x9f\xfa\xdd\xd6\xfb\x7d\xae\x76\x4a\x1f\x0e\x7f\x02\x00\x00\xff\xff\x08\xac\xb3\x1f\xcd\x02\x00\x00")

func dataSqlGithubRailsRoutesTestSqlBytes() ([]byte, error) {
	return bindataRead(
		_dataSqlGithubRailsRoutesTestSql,
		"data/sql/github/rails-routes-test.sql",
	)
}

func dataSqlGithubRailsRoutesTestSql() (*asset, error) {
	bytes, err := dataSqlGithubRailsRoutesTestSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/sql/github/rails-routes-test.sql", size: 717, mode: os.FileMode(420), modTime: time.Unix(1534136427, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataSqlGithubRailsRoutesSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\x51\x73\x93\x40\x14\x85\xdf\xf7\x57\x1c\x77\xd4\x42\x35\x89\xaf\x5a\x71\x44\xba\x8d\x38\x74\xe9\x2c\x30\x4e\x6c\x6a\x43\xc8\x42\xd6\x21\x10\x61\x71\xd4\x34\xff\xdd\x61\x49\x6a\x3b\xe3\x83\x4f\x17\xee\xb9\xdf\x3d\xdc\x83\x27\x98\x1b\x33\xc4\xec\xf2\x2a\x14\xae\x98\xe1\x22\xe1\x5e\xec\x87\x1c\xdb\xb4\x69\xa5\xa8\x3b\x2d\x5b\xab\x31\xc5\xab\x57\x12\x51\x2c\x7c\x3e\xb5\x89\x60\x71\x22\x78\x04\x57\x08\x77\xf6\x76\xe8\xbe\x43\xe0\xf2\x69\xe2\x4e\x19\xbe\xb5\x70\x23\x50\x4a\x09\x01\x7e\xa4\x0d\x1a\x59\xc8\x9f\x70\x30\xb1\xe8\xdd\x89\x7d\x9d\x8e\x7e\xbb\xa3\x2f\xb7\xaf\x46\xaf\xe7\xf3\xc9\x7c\xfe\xe6\xc9\xfb\xa7\xcf\xbe\x3e\x3f\xb5\xec\x17\xce\xcd\xa9\x75\x72\x47\xed\xc9\xa6\x50\x67\x04\x50\x39\x1e\xf8\xdb\xd8\x11\xe0\xb8\xb2\xed\x4a\x0d\x07\x7f\xe5\xf1\x26\xd5\xd9\xda\x32\x66\x76\x4f\xef\x21\xcb\x56\xfe\x0b\xba\xbe\x31\x3a\x01\x1a\xa9\xbb\xa6\x3a\x28\x7d\x93\x52\x7a\x46\x48\xc4\x02\xe6\xc5\xbd\xde\xaf\x7f\x49\x00\x2f\x4c\x78\x3c\x7c\x8c\xdd\x9f\x97\xd5\x5d\xa5\xc9\x85\x08\x2f\x61\x11\xe0\x1e\xc0\xa3\xf0\xb2\x71\x56\x57\x5a\x56\xda\x30\x86\x26\x40\x4f\x99\xd1\xc5\x52\x15\xdf\x3b\xd9\xfc\x1a\x6d\xbb\x65\xa9\xb2\xd1\x2a\xd5\xe9\xb8\x50\x7a\xdd\x2d\x6f\x1b\xb9\xad\xdb\x23\xdf\x2e\x90\x11\xe0\x53\xe8\x73\xe3\xf7\xc8\x11\x50\x2b\xf3\x70\xbf\xf8\xbf\x56\xe7\xaa\x94\xed\xc2\xcc\x7f\xfe\xc8\x04\x3b\x90\xdb\x54\xaf\xe1\x80\x66\x75\x95\xab\x62\x32\x04\x3c\x6e\x96\x14\x36\x72\x02\x84\xdc\x0c\xe6\x63\xb5\x82\x83\xac\x2f\xf6\xe1\x37\xf4\x41\x25\x9c\xb3\xe8\x61\x52\xc3\xd5\x47\x07\xf3\x06\x3f\x02\x0f\x63\xf0\x24\x08\xc8\x54\x84\xc9\x15\x3e\xcc\x8e\x22\x09\xc5\x39\x13\x43\xc3\xa4\x8c\x73\x16\x79\x24\xf0\x2f\xfd\xfe\xdc\xdd\xae\x54\x1b\xa5\xf7\xfb\x3f\x01\x00\x00\xff\xff\x01\xe3\x7c\x25\xbf\x02\x00\x00")

func dataSqlGithubRailsRoutesSqlBytes() ([]byte, error) {
	return bindataRead(
		_dataSqlGithubRailsRoutesSql,
		"data/sql/github/rails-routes.sql",
	)
}

func dataSqlGithubRailsRoutesSql() (*asset, error) {
	bytes, err := dataSqlGithubRailsRoutesSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/sql/github/rails-routes.sql", size: 703, mode: os.FileMode(420), modTime: time.Unix(1534136417, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataSqlGithubTomcatRoutesTestSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x51\x51\x6f\xd3\x30\x18\x7c\xf7\xaf\x38\x2c\x60\xc9\xa0\x29\xaf\x30\x8a\x08\x99\x57\x82\x5a\x67\x72\x12\xa1\xb2\x8e\x36\x4d\xdd\xd4\x28\x4d\x42\xe2\x20\xa0\xeb\x7f\x47\x71\xda\xb1\x49\xbc\xf0\xf4\xd9\xdf\x7d\x77\xe7\xef\xec\x09\xe6\x46\x0c\x11\x9b\x5e\x07\xc2\x15\x33\x5c\xc5\xdc\x8b\xfc\x80\xa3\x4a\xea\x46\x8a\xb2\xd5\xb2\xb1\x6a\x53\xbc\x72\x2d\x11\x46\xc2\xe7\x63\x9b\x08\x16\xc5\x82\x87\x70\x85\x70\x67\x6f\xfb\xee\x3b\x4c\x5c\x3e\x8e\xdd\x31\xc3\xb7\x06\x6e\x08\x4a\x29\x21\xc0\x8f\xa4\x46\x2d\x33\xf9\x13\x23\x0c\x2d\x7a\x77\x66\xdf\x24\x83\xdf\xee\xe0\xcb\xe2\xd5\xe0\xf5\x7c\x3e\x9c\xcf\xdf\x3c\x79\xff\xf4\xd9\xd7\xe7\xe7\x96\xfd\x62\x74\x7b\x6e\x9d\xdd\x51\x7b\xb8\xcb\xd4\x05\x01\xd4\x06\x0f\xfc\x6d\xec\x09\x70\x92\x6c\xda\x5c\x63\x84\xbf\xb0\xb3\x4b\x74\xba\xb5\x8c\x99\xdd\xb1\x0f\x90\x79\x23\xff\x45\xba\xb9\x35\x38\x01\x6a\xa9\xdb\xba\x38\x22\x5d\x93\x52\x7a\x41\x48\xc8\x26\xcc\x8b\x3a\xbc\x93\x7f\x49\x00\x2f\x88\x79\xd4\x3f\xc6\xee\xd6\x4b\xcb\xb6\xd0\xe4\x4a\x04\x53\x58\x04\xb8\x27\xe0\x51\x78\xa9\x93\x96\x85\x96\x85\x36\x1c\xc3\x26\x40\xc7\x32\xa3\xcb\x95\xca\xbe\xb7\xb2\xfe\x35\xa8\xda\x55\xae\xd2\xc1\x3a\xd1\x89\x93\x29\xbd\x6d\x57\x8b\x5a\x56\x65\xe3\x34\xc9\xae\xca\xe5\xe2\x28\xd3\x2c\x91\x12\xe0\x53\xe0\x73\x63\xfb\xc8\x18\x50\x6b\x73\xb8\xd7\xff\x1f\x87\x8d\xca\x65\xb3\x34\xb4\xcf\x1f\x99\x60\x47\x81\x2a\xd1\x5b\x8c\x40\xd3\xb2\xd8\xa8\x6c\xd8\xc7\xed\xd4\x2b\x0a\x1b\x1b\x02\x04\xdc\x0c\x6e\x1c\xb5\xc6\x08\x69\x57\xec\xe3\xa7\x74\xb1\xc5\x9c\xb3\xf0\x61\x6e\x7d\x06\x27\x07\x73\x83\x1f\x82\x07\x11\x78\x3c\x99\x90\xb1\x08\xe2\x6b\x7c\x98\x9d\x40\x12\x88\x4b\x26\xfa\x86\xc9\x1c\x97\x2c\xf4\xc8\xc4\x9f\xfa\xdd\xd6\xfb\x7d\xae\x76\x4a\x1f\x0e\x7f\x02\x00\x00\xff\xff\x08\xac\xb3\x1f\xcd\x02\x00\x00")

func dataSqlGithubTomcatRoutesTestSqlBytes() ([]byte, error) {
	return bindataRead(
		_dataSqlGithubTomcatRoutesTestSql,
		"data/sql/github/tomcat-routes-test.sql",
	)
}

func dataSqlGithubTomcatRoutesTestSql() (*asset, error) {
	bytes, err := dataSqlGithubTomcatRoutesTestSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/sql/github/tomcat-routes-test.sql", size: 717, mode: os.FileMode(420), modTime: time.Unix(1534149328, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataSqlGithubTomcatRoutesSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func dataSqlGithubTomcatRoutesSqlBytes() ([]byte, error) {
	return bindataRead(
		_dataSqlGithubTomcatRoutesSql,
		"data/sql/github/tomcat-routes.sql",
	)
}

func dataSqlGithubTomcatRoutesSql() (*asset, error) {
	bytes, err := dataSqlGithubTomcatRoutesSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/sql/github/tomcat-routes.sql", size: 0, mode: os.FileMode(420), modTime: time.Unix(1534039265, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataSqlGithubWordsWithExtSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xce\xcd\x4e\x84\x30\x14\x40\xe1\xfd\x7d\x8a\xbb\x30\x19\x30\xce\xbc\x80\x2b\x2d\x57\x24\x01\x4a\x0a\x44\x5d\xf1\x67\x85\x26\x08\x48\xdb\x44\xd3\xf4\xdd\x8d\x38\xcb\xb3\x3a\x5f\x49\x29\xb1\x0a\x10\xb7\xce\x4c\x77\x80\xc8\x78\x9d\x57\xc1\x6d\x88\xc3\x6a\x17\x03\x4f\x82\x67\x80\xd8\xf6\x6a\xfc\xb2\x72\xff\x39\x6f\xb6\x9f\xd5\x70\x7e\xef\x4c\x77\x19\x95\x99\x6c\xdf\xec\x72\x5b\xf5\xe5\x43\xcd\x52\xb7\xf0\xf2\x4c\x82\x00\x51\x50\x4c\xaf\x45\xc3\x78\x5e\x3d\x24\x79\x89\xc1\x31\xc0\xfd\x14\x38\x27\xbf\x8d\x5c\xb4\x5a\x17\xed\x7d\x78\x73\x0a\x21\x16\xbc\x2e\xf0\xf1\xed\xea\x00\x2e\x22\x12\xff\x7d\x30\x30\xa2\x92\x41\x9a\x64\xc9\x1f\xd5\xb9\x59\x7d\x2a\xe3\xfd\xfd\x6f\x00\x00\x00\xff\xff\xdb\xe0\xde\xd1\xbe\x00\x00\x00")

func dataSqlGithubWordsWithExtSqlBytes() ([]byte, error) {
	return bindataRead(
		_dataSqlGithubWordsWithExtSql,
		"data/sql/github/words-with-ext.sql",
	)
}

func dataSqlGithubWordsWithExtSql() (*asset, error) {
	bytes, err := dataSqlGithubWordsWithExtSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/sql/github/words-with-ext.sql", size: 190, mode: os.FileMode(420), modTime: time.Unix(1532021038, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataSqlHackernewsSubdomainsSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8f\x4f\x8f\x9b\x30\x14\xc4\xef\xef\x53\x8c\x38\xc5\x52\x43\xee\x45\x3d\xb0\xac\x17\x21\x25\x10\x19\x38\xe4\xd4\x35\x84\x64\xbd\x0b\x66\xeb\x3f\xea\x46\x51\xbe\x7b\x15\x42\x53\xa5\x07\x4b\xf6\xf3\xcc\xbc\xdf\x24\x82\xc7\x15\x47\xc5\x37\xdb\x42\xc4\x62\x87\x97\x3a\x4f\xaa\xac\xc8\x71\xec\x5c\xe9\x9b\xfd\x38\x48\xa5\x17\x5f\x28\x2b\x91\xe5\x29\x03\x48\xf0\xaa\x16\x79\x39\x4f\x00\x02\xd6\x71\x9e\xd6\x71\xca\xf1\x6e\x11\x97\x08\x82\x80\x80\x83\xd7\xad\x53\xa3\x7e\x4c\xb2\x0c\x67\x02\x00\x67\x4e\xf3\x0d\x30\x9d\xf3\x46\xa3\x16\xd9\xc2\xb2\xd0\xde\xc5\x2c\x9a\x04\x17\xb4\xd2\xb5\x6f\x58\x74\x5f\xec\x7f\x8f\x9d\x25\x74\x3b\xf3\xf4\x11\x9e\x45\x74\x45\x2a\xb6\xd7\x62\x25\x16\x13\x73\xaf\x1a\x23\xcd\xe9\x47\x70\xb4\xdf\x57\xab\x76\x1c\x86\x51\xdb\xcf\x4e\x7e\x2c\xfd\xfe\xb0\xaa\x45\x16\x0e\x4a\x87\xef\x36\x20\x16\x11\x95\x7c\xcd\x93\x6a\x32\x3e\x64\x7b\xd3\xb3\x6b\xe5\x3b\xf3\x37\x02\xda\xd1\x6b\x77\xfb\x92\xf6\xf6\xa2\x17\x51\x6c\x26\xfb\x6b\xa3\x8e\xbf\x7c\x67\x4e\xcb\x4f\xdf\xf4\xaa\x5d\xee\xa5\x93\xe1\x9b\x6c\x3f\x3a\xf3\x53\x77\xbf\x6d\x78\xf0\x7d\xff\x4a\xa9\x28\xea\x2d\x9e\x76\x84\x7f\xe1\x54\x88\x67\x2e\xf0\xb4\xc3\xdf\x2d\x78\xe6\x65\x42\xeb\x6c\x93\x55\x04\x9c\xcf\xbd\x1a\x94\xbb\x5c\xa2\x3f\x01\x00\x00\xff\xff\x59\xeb\x00\xbf\xd8\x01\x00\x00")

func dataSqlHackernewsSubdomainsSqlBytes() ([]byte, error) {
	return bindataRead(
		_dataSqlHackernewsSubdomainsSql,
		"data/sql/hackernews/subdomains.sql",
	)
}

func dataSqlHackernewsSubdomainsSql() (*asset, error) {
	bytes, err := dataSqlHackernewsSubdomainsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/sql/hackernews/subdomains.sql", size: 472, mode: os.FileMode(420), modTime: time.Unix(1534034400, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataSqlHttpArchiveDirectoriesSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8f\x4f\x8f\x9b\x30\x10\xc5\xef\xf3\x29\x9e\x38\x61\xa9\x25\xf7\xa2\x1e\x28\xf1\x22\xa4\x04\x22\x63\x0e\x7b\xea\x52\xf0\x6e\x9c\xf0\x27\xb5\x4d\x95\x08\xf1\xdd\x2b\x12\xd2\x4a\x7b\x18\x69\xe6\xe9\xcd\xd3\xfb\xc5\x82\x47\x92\x43\xf2\xfd\x21\x17\x91\x78\xc5\x4b\x99\xc5\x32\xcd\x33\x02\x3e\x94\xdb\x6a\xe3\x5f\x51\x48\x91\x66\x09\x23\x40\x70\x59\x8a\xac\x58\x15\x02\x76\x51\x96\x94\x51\xc2\x71\xb2\x88\x0a\x78\x9e\x47\xc0\xfb\xd8\xd7\x4e\x0f\xfd\x33\xc1\x32\x4c\x04\x00\xce\xdc\xd6\x0d\x30\xca\x8d\xa6\x47\x29\x52\xdf\xb2\xa0\xd1\x46\xd5\x6e\x30\x37\x9f\x85\x77\xc3\x8c\xba\x72\xf5\x11\xbe\xba\xb2\xcf\x3f\x76\xb5\xd0\x63\x56\xf5\x59\x97\x85\xb4\xd4\xc8\x0f\x0b\x46\x41\x80\x8f\x56\xff\x32\x95\xb9\x7d\xf7\x3e\xec\xb7\xcd\xa6\x1e\xba\x6e\xe8\xed\x45\x55\xe7\xaf\x63\xf3\xbe\x29\x45\x1a\x74\xba\x0f\x4e\xd6\x03\x0b\xa9\xe0\x3b\x1e\xcb\xff\xfc\xa3\x69\xd9\xc2\x36\x9a\xf6\x0b\x01\x71\x5e\x66\xf2\x9f\x58\x0f\x63\xef\xe8\x45\xe4\x7b\x02\xde\x8e\xce\x5d\x2a\x53\x1f\xf5\x1f\x15\x18\xf5\x7b\x54\xd6\xd9\x60\x9a\x9a\xca\xa9\x79\xfe\xd9\x28\x7b\x76\xc3\xe5\x8d\x12\x91\x97\x07\xfc\x78\x25\x2c\xa1\x94\x8b\x2d\x17\x8f\xf3\x1e\x87\x2d\x2f\x62\xda\xa5\xfb\x54\x62\x9a\x5a\xdd\x69\x37\xcf\xe1\xdf\x00\x00\x00\xff\xff\xa0\xad\x0d\x72\xaa\x01\x00\x00")

func dataSqlHttpArchiveDirectoriesSqlBytes() ([]byte, error) {
	return bindataRead(
		_dataSqlHttpArchiveDirectoriesSql,
		"data/sql/http-archive/directories.sql",
	)
}

func dataSqlHttpArchiveDirectoriesSql() (*asset, error) {
	bytes, err := dataSqlHttpArchiveDirectoriesSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/sql/http-archive/directories.sql", size: 426, mode: os.FileMode(420), modTime: time.Unix(1605580993, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataSqlHttpArchiveSubdomainsSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x90\x41\xef\x9a\x40\x14\xc4\xef\xef\x53\x4c\x38\xb1\x49\x8b\xf7\x92\x1e\x28\xff\x95\x90\x28\x98\x65\x39\x78\xaa\x08\xab\xa2\x02\x76\x77\x69\x34\x84\xef\xde\xa8\xc4\x46\x0f\x9b\x6c\xe6\xcd\x9b\xfc\xe6\x85\x82\x07\x92\x43\xf2\xe5\x2a\x15\x81\x58\x63\x9e\x27\xa1\x8c\xd3\x84\x80\xbd\xb2\x59\xbf\xad\xba\xa6\xa8\x5b\xf7\x8a\x4c\x8a\x38\x89\x18\x01\x82\xcb\x5c\x24\xd9\xa4\x10\xb0\x08\x92\x28\x0f\x22\x8e\xa3\x41\x90\xc1\x71\x1c\x02\x76\x7d\x5b\xda\xba\x6b\xdf\x73\x0c\xc3\x40\x00\x60\xf5\x6d\xfa\x01\x5a\xd9\x5e\xb7\xc8\x45\xec\x1a\xe6\x99\x97\x99\xf9\x0f\xc3\x88\xb2\xb0\xe5\x01\xae\xba\xb2\xcf\x1d\x33\x59\xe8\xf9\x26\xf5\x1d\x9d\xf9\x74\x47\x4a\x57\xf7\x62\x19\x01\x2e\xce\xf5\x56\x17\xfa\xf6\xd3\xd9\x9b\x1f\xb3\x59\xd9\x35\x4d\xd7\x9a\x8b\x2a\x4e\xdf\xfb\x6a\x37\xcb\x45\xec\x35\x75\xeb\x1d\x8d\x03\xe6\x53\xc6\x17\x3c\x94\x9f\x17\xe9\xf5\x99\xdd\xdb\xbe\x70\xbf\x11\x10\xa6\x79\x22\x5f\xa3\xb2\xeb\x5b\x4b\x73\x91\x2e\x09\xd8\x1c\xac\xbd\x14\xba\x3c\xd4\x7f\x95\xa7\xd5\x9f\x5e\x19\x6b\xbc\x61\xa8\x0a\xab\xc6\xf1\x77\xa5\xcc\xc9\x76\x97\x0d\x45\x22\xcd\x57\xf8\xb5\x26\xfc\x8f\xa6\x54\x7c\x71\xf1\x14\x1f\xa1\xf8\xe2\x59\x48\x8b\x78\x19\x4b\x0c\xc3\xb9\x6e\x6a\x3b\x8e\xfe\xbf\x00\x00\x00\xff\xff\xc5\x85\x14\xbd\xce\x01\x00\x00")

func dataSqlHttpArchiveSubdomainsSqlBytes() ([]byte, error) {
	return bindataRead(
		_dataSqlHttpArchiveSubdomainsSql,
		"data/sql/http-archive/subdomains.sql",
	)
}

func dataSqlHttpArchiveSubdomainsSql() (*asset, error) {
	bytes, err := dataSqlHttpArchiveSubdomainsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/sql/http-archive/subdomains.sql", size: 462, mode: os.FileMode(420), modTime: time.Unix(1605607736, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataSqlHttpArchiveTechnologiesSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x0a\x76\xf5\x71\x75\x0e\xe1\x52\x50\x28\x2d\xca\xe1\x72\x0b\xf2\xf7\xe5\x52\x50\x48\xc8\x28\x29\x29\x48\x2c\x4a\xce\xc8\x2c\x4b\xd5\x2b\x49\x4d\xce\xc8\xcb\xcf\xc9\x4f\xcf\x4c\x2d\xd6\xab\xae\x4e\x49\x2c\x49\xad\xad\x8d\x4f\x49\x2d\xce\x2e\xc9\x2f\x48\xe0\x0a\xf7\x70\x0d\x72\xe5\x52\x50\x48\x2c\x28\x50\xb0\x55\x50\xaf\xae\x86\xab\xaf\xac\xad\x55\xe7\xf2\xf1\xf4\xf5\x04\x99\x5e\x5d\x9d\x93\x99\x9b\x59\x52\x5b\x0b\x08\x00\x00\xff\xff\x9f\x3f\x79\xb3\x70\x00\x00\x00")

func dataSqlHttpArchiveTechnologiesSqlBytes() ([]byte, error) {
	return bindataRead(
		_dataSqlHttpArchiveTechnologiesSql,
		"data/sql/http-archive/technologies.sql",
	)
}

func dataSqlHttpArchiveTechnologiesSql() (*asset, error) {
	bytes, err := dataSqlHttpArchiveTechnologiesSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/sql/http-archive/technologies.sql", size: 112, mode: os.FileMode(420), modTime: time.Unix(1605610287, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataSqlHttpArchiveWordsWithExtSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x90\x51\x6f\x9b\x30\x14\x85\xdf\xef\xaf\x38\x42\x93\x8a\xa5\x8e\xbc\x2f\xda\x03\xa3\x0e\x43\x4a\x20\x32\x46\x5b\x9f\x5a\x46\x9c\xc6\x2d\x98\xcc\x36\x53\x2a\xc4\x7f\x9f\x68\xd9\xa4\xe6\xc1\x92\x7d\x7c\xee\xd1\x77\x6e\x22\x78\x2c\x39\x24\xdf\xed\x0b\x11\x8b\x7b\x6c\xaa\x3c\x91\x59\x91\x13\xf0\xa4\xfc\x46\xb7\xca\xd4\x9d\x0a\x2f\x28\xa5\xc8\xf2\x94\x11\x20\xb8\xac\x44\x5e\x2e\x0a\x01\xdb\x38\x4f\xab\x38\xe5\x78\x76\x88\x4b\x04\x41\x40\xc0\x71\x30\x8d\xd7\xbd\xf9\x10\xe3\x18\x46\x02\x00\x6f\x5f\x97\x1b\x60\x95\x1f\xac\x41\x25\xb2\xd0\xb1\xe8\xf8\xcf\xcb\xd6\x6f\xff\x13\x9a\xda\x37\x27\x84\xea\xc2\xae\x47\xdc\x62\xa1\xf7\xb3\xa8\x1f\xb8\xd9\x9a\x66\x9e\x62\x3f\x97\x2a\x09\x08\xd1\xea\x5f\xb6\xb6\xaf\x5f\x83\x27\xf7\x65\xb5\x6a\xfa\xae\xeb\x8d\x3b\xab\xfa\xe5\xf3\x70\x38\xae\x2a\x91\x45\x9d\x36\xd1\xb3\x0b\xc0\xd6\x54\xf2\x2d\x4f\xe4\xd5\x36\x06\xdb\xb2\xb9\xe9\x60\xdb\x5b\x02\x92\xa2\xca\xe5\x7f\xb1\xe9\x07\xe3\x69\x23\x8a\x1d\x01\x8f\x27\xef\xcf\xb5\x6d\x4e\xfa\x8f\x8a\xac\xfa\x3d\x28\xe7\x5d\x34\x8e\x87\xda\xab\x69\x7a\x38\x28\xf7\xe2\xfb\xf3\x23\xfd\xf8\xce\x05\x87\xe0\x29\xff\xb9\x7f\x48\x8a\x5c\xc6\x59\x5e\x62\xce\xbc\x85\xbd\x09\xc7\x51\x5d\xbc\x32\x4e\xf7\xc6\x4d\x13\xfb\x74\xc3\x28\x15\x45\xb5\xc7\xb7\x7b\xc2\x8c\x41\x85\xb8\xe3\xe2\xfd\xf9\x06\x80\x3b\x5e\x26\xb4\xcd\x76\x99\xc4\x38\xb6\xba\xd3\x7e\x9a\xd6\x7f\x03\x00\x00\xff\xff\x2f\xdd\x00\x30\xef\x01\x00\x00")

func dataSqlHttpArchiveWordsWithExtSqlBytes() ([]byte, error) {
	return bindataRead(
		_dataSqlHttpArchiveWordsWithExtSql,
		"data/sql/http-archive/words-with-ext.sql",
	)
}

func dataSqlHttpArchiveWordsWithExtSql() (*asset, error) {
	bytes, err := dataSqlHttpArchiveWordsWithExtSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/sql/http-archive/words-with-ext.sql", size: 495, mode: os.FileMode(420), modTime: time.Unix(1605527689, 0)}
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
	"data/.DS_Store":                           dataDs_store,
	"data/filters/numerical-parameters.txt":    dataFiltersNumericalParametersTxt,
	"data/filters/string-parameters.txt":       dataFiltersStringParametersTxt,
	"data/sql/.DS_Store":                       dataSqlDs_store,
	"data/sql/github/deleted-files-test.sql":   dataSqlGithubDeletedFilesTestSql,
	"data/sql/github/deleted-files.sql":        dataSqlGithubDeletedFilesSql,
	"data/sql/github/nodejs-routes-test.sql":   dataSqlGithubNodejsRoutesTestSql,
	"data/sql/github/nodejs-routes.sql":        dataSqlGithubNodejsRoutesSql,
	"data/sql/github/rails-routes-test.sql":    dataSqlGithubRailsRoutesTestSql,
	"data/sql/github/rails-routes.sql":         dataSqlGithubRailsRoutesSql,
	"data/sql/github/tomcat-routes-test.sql":   dataSqlGithubTomcatRoutesTestSql,
	"data/sql/github/tomcat-routes.sql":        dataSqlGithubTomcatRoutesSql,
	"data/sql/github/words-with-ext.sql":       dataSqlGithubWordsWithExtSql,
	"data/sql/hackernews/subdomains.sql":       dataSqlHackernewsSubdomainsSql,
	"data/sql/http-archive/directories.sql":    dataSqlHttpArchiveDirectoriesSql,
	"data/sql/http-archive/subdomains.sql":     dataSqlHttpArchiveSubdomainsSql,
	"data/sql/http-archive/technologies.sql":   dataSqlHttpArchiveTechnologiesSql,
	"data/sql/http-archive/words-with-ext.sql": dataSqlHttpArchiveWordsWithExtSql,
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
	"data": &bintree{nil, map[string]*bintree{
		".DS_Store": &bintree{dataDs_store, map[string]*bintree{}},
		"filters": &bintree{nil, map[string]*bintree{
			"numerical-parameters.txt": &bintree{dataFiltersNumericalParametersTxt, map[string]*bintree{}},
			"string-parameters.txt":    &bintree{dataFiltersStringParametersTxt, map[string]*bintree{}},
		}},
		"sql": &bintree{nil, map[string]*bintree{
			".DS_Store": &bintree{dataSqlDs_store, map[string]*bintree{}},
			"github": &bintree{nil, map[string]*bintree{
				"deleted-files-test.sql": &bintree{dataSqlGithubDeletedFilesTestSql, map[string]*bintree{}},
				"deleted-files.sql":      &bintree{dataSqlGithubDeletedFilesSql, map[string]*bintree{}},
				"nodejs-routes-test.sql": &bintree{dataSqlGithubNodejsRoutesTestSql, map[string]*bintree{}},
				"nodejs-routes.sql":      &bintree{dataSqlGithubNodejsRoutesSql, map[string]*bintree{}},
				"rails-routes-test.sql":  &bintree{dataSqlGithubRailsRoutesTestSql, map[string]*bintree{}},
				"rails-routes.sql":       &bintree{dataSqlGithubRailsRoutesSql, map[string]*bintree{}},
				"tomcat-routes-test.sql": &bintree{dataSqlGithubTomcatRoutesTestSql, map[string]*bintree{}},
				"tomcat-routes.sql":      &bintree{dataSqlGithubTomcatRoutesSql, map[string]*bintree{}},
				"words-with-ext.sql":     &bintree{dataSqlGithubWordsWithExtSql, map[string]*bintree{}},
			}},
			"hackernews": &bintree{nil, map[string]*bintree{
				"subdomains.sql": &bintree{dataSqlHackernewsSubdomainsSql, map[string]*bintree{}},
			}},
			"http-archive": &bintree{nil, map[string]*bintree{
				"directories.sql":    &bintree{dataSqlHttpArchiveDirectoriesSql, map[string]*bintree{}},
				"subdomains.sql":     &bintree{dataSqlHttpArchiveSubdomainsSql, map[string]*bintree{}},
				"technologies.sql":   &bintree{dataSqlHttpArchiveTechnologiesSql, map[string]*bintree{}},
				"words-with-ext.sql": &bintree{dataSqlHttpArchiveWordsWithExtSql, map[string]*bintree{}},
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
