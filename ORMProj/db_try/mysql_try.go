package db_try

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Student struct {
	gorm.Model
	Name   string
	Age    uint
	Gender string
}

func MysqlTry() {
	db, err := gorm.Open("mysql", "root:19911214yu@(localhost:3306)/testdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败-->" + err.Error())
	}
	defer db.Close()

	//自动迁移；也就是建表
	db.AutoMigrate(&Student{})
	//写数据
	db.Create(&Student{Name: "yuyong", Age: 27, Gender: "M"})
	//查数据
	var yong Student
	db.First(&yong, "Name = ?", "yuyong")
	//改数据
	db.Model(&yong).Update("Age", 28)
	//删数据
	db.Delete(&yong)
}
