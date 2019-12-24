package db_try

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

//读写已建好的表
/*
create schema f_test;
grant all on schema f_test to yuyong;
create table f_test.f_stu_info(
    f_stu_name varchar(255),
    f_stu_age varchar(255)
);
grant all on f_test.f_stu_info to yuyong;

insert into f_test.f_stu_info(f_stu_name,f_stu_age)
values('aaa','111'),
      ('bbb','222'),
      ('ccc','333');

select * from f_test.f_stu_info
*/

type StuInfo struct {
	FStuName string `gorm:"column:f_stu_name"`
	FStuAge  string `gorm:"column:f_stu_age"`
}

func (StuInfo) TableName() string {
	return "f_test.f_stu_info"
}

func TestGP() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "0.0.0.0", 5432, "yuyong", "123456", "testDB")
	fmt.Println("get conn str-->" + psqlInfo)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	var stu StuInfo
	db.First(&stu)
	fmt.Println(stu)

	var stus []StuInfo
	db.Find(&stus)
	fmt.Println(stus)
}
