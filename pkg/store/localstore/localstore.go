// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package localstore

import (
	"nctr/pkg/alog"
	"nctr/pkg/constant"
	"path/filepath"
	"strings"

	"github.com/syndtr/goleveldb/leveldb"
)

var levelDB *leveldb.DB

func InitDb(path string) {
	dataPath := filepath.Join(path, "localstore")
	db, err := leveldb.OpenFile(dataPath, nil)
	levelDB = db
	if err != nil {
		alog.Error("leveldb read failed")
	}
	// defer levelDB.Close()

}

func Put(key string, value string) {

	levelDB.Put([]byte(key), []byte(value), nil)

}

func Get(key string) string {
	data, err := levelDB.Get([]byte(key), nil)
	if err != nil {
		return ""
	}
	result := string(data[:])
	return result
}

func FindAll() map[string]string {
	iter := levelDB.NewIterator(nil, nil)
	var dict map[string]string = make(map[string]string)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		dict[string(key)] = string(value)
	}
	defer iter.Release()
	return dict
}

func FindAllHeart(lastTime string) map[string]string {
	iter := levelDB.NewIterator(nil, nil)
	var dict map[string]string = make(map[string]string)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()

		if strings.Contains(string(key), constant.HeartKeyPrefix) {

			time := strings.Replace(string(key), constant.HeartKeyPrefix, "", -1)

			if time > lastTime {
				dict[string(key)] = string(value)
				if len(dict) >= constant.UploadNum {
					return dict
				}
			}

		}
	}
	return dict
}
