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

/*
create table f_test.f_stu_info(
    f_stu_name varchar(255),
    f_stu_age varchar(255)
);

*/
