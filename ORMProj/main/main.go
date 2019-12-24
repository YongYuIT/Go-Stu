package main

import (
	"../db_try"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	fmt.Println("env_test")
	//db_try.MysqlTry()
	//db_try.TryPkFk()
	db_try.TestGP()
}
