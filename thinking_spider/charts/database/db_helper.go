package database

import (
	"charts/config"
	_ "charts/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var CurrentDB *gorm.DB

func init() {
	db, err := gorm.Open(mysql.Open(config.DBConn.Conn), &gorm.Config{})
	if err != nil {
		fmt.Println("conn db err-->", db)
		os.Exit(1)
	}
	CurrentDB = db
}
