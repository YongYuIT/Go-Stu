package main

import (
	"fmt"
	"../db_try"
)

func main() {
	fmt.Println("env_test")
	//db_try.MysqlTry()
	db_try.TryPkFk()
}
