package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"thinking_spider/config"
	_ "thinking_spider/config"
)

var CurrentDB *gorm.DB

func init() {
	db, err := gorm.Open(config.DBConn.Type, config.DBConn.Conn)
	if err != nil {
		fmt.Println("conn db err-->", db)
		os.Exit(1)
	}
	CurrentDB = db
}

func CloseDB() {
	CurrentDB.Close()
}
