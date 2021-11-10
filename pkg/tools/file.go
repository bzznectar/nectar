// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tools

import (
	"nctr/pkg/alog"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
)

type FileInfo struct {
	name string
	size float64
}

func CreateRandomFile(dir string, minSize int64, maxSize int64, name string) {
	path := filepath.Join(dir)
	fullPath := filepath.Join(path, name)
	if FileExist(fullPath) {
		return
	}
	if !FileExist(dir) {
		os.Mkdir(dir, 0755)
	}
	f, err := os.Create(fullPath)
	if err != nil {
		alog.Error(err)
	}
	defer f.Close()
	size := GetRandomInt(minSize, maxSize)
	//f.Truncate(size)

	f.WriteString(GetRandomString(int(size)))

}

func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func FileSize(path string) uint64 {
	f, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return uint64(f.Size())
}
func DirSize(dir string) uint64 {
	files, err := DirInfo(dir)
	if err != nil {
		return 0
	}
	size := 0.0
	for _, item := range files {
		size += item.size
	}
	return uint64(size)
}
func DirInfo(path string) ([]FileInfo, error) {
	op, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	files, err := ioutil.ReadDir(op)
	if err != nil {
		return nil, err
	}
	var fileInfos []FileInfo
	for _, f := range files {
		if f.IsDir() {
			fs, err := DirInfo(op + "/" + f.Name())
			if err != nil {
				return nil, err
			}
			fileInfos = append(fileInfos, fs...)
		} else {
			fi := FileInfo{op + "/" + f.Name(), float64(f.Size())}
			fileInfos = append(fileInfos, fi)
		}
	}
	return fileInfos, nil
}
