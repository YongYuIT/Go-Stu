package db_try

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Profile   Profile
	ProfileID int
}

type Profile struct {
	gorm.Model
	Name string
}


func TryPkFk()  {
	db, err := gorm.Open("mysql", "root:19911214yu@(localhost:3306)/testdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败-->" + err.Error())
	}
	defer db.Close()
	db.AutoMigrate(&User{},&Profile{})
}