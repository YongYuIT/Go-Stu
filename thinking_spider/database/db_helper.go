package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"thinking_spider/config"
	_ "thinking_spider/config"
	"thinking_spider/model"
)

var thizDB *gorm.DB

func init() {
	db, err := gorm.Open(config.DBConn.Type, config.DBConn.Conn)
	if err != nil {
		fmt.Println("conn db err-->", db)
		return
	}
	thizDB = db
}

func SaveKeyWordProdRecord(record *model.KeyWordProdRecord) {
	if !thizDB.HasTable(record) {
		thizDB.AutoMigrate(record)
	}
	thizDB.Create(record)
}

func CloseDB() {
	thizDB.Close()
}
