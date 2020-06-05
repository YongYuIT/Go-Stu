package service

import (
	"encoding/json"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
)

var ThisDB *leveldb.DB = nil

const db_cache_path = "/db_cache/"

func InitDb(root_path string) error {
	var err error = nil
	ThisDB, err = leveldb.OpenFile(root_path+db_cache_path, nil)
	if err != nil {
		ThisDB = nil
		fmt.Println("init db err-->", err)
		return err
	}
	return nil
}

func RecordUpload(fid string, fdis FileDiscribe) error {
	if ThisDB == nil {
		return fmt.Errorf("DB not inited!")
	}
	json, err := json.Marshal(fdis)
	if err != nil {
		fmt.Println("get file json err-->", err)
		return err
	}
	return ThisDB.Put([]byte(fid), json, nil)
}

func CloseDB() {
	if ThisDB != nil {
		ThisDB.Close()
	}
}
