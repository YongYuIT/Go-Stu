package main

import (
	_ "./routers"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(0.0.0.0:3306)/test_bee?charset=utf8")
	//设置数据库的最大空闲连接
	orm.SetMaxIdleConns("default", 30)
	//设置数据库的最大数据库连接
	orm.SetMaxOpenConns("default", 30)
	orm.DefaultTimeLoc = time.UTC
}

func main() {
	beego.Run()
}
