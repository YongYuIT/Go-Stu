package main

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
)

var db *leveldb.DB = nil

func main() {
	var err error
	db, err = leveldb.OpenFile("./test_db", nil)
	defer func() {
		if (db != nil) {
			db.Close()
		}
	}()
	if (err != nil) {
		fmt.Println("conn failed-->" + err.Error())
		return
	}
	db.Put([]byte("name"), []byte("yuyong"), nil)
	name, err := db.Get([]byte("name"), nil)
	if (err != nil) {
		fmt.Println("conn failed-->" + err.Error())
		return
	}
	fmt.Println(string(name))
}
