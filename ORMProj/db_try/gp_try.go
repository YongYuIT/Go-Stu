package db_try

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

//读写已建好的表
/*
create schema f_test;
grant all on schema f_test to yuyong;

DROP TABLE IF EXISTS f_test.f_stu_info;
create table f_test.f_stu_info(
    f_stu_name varchar(255),
    f_stu_age varchar(255),
	f_stu_add varchar(255),
	f_marks varchar(255)
);
grant all on f_test.f_stu_info to yuyong;

insert into f_test.f_stu_info(f_stu_name,f_stu_age,f_stu_add,f_marks)
values('aaa','111','aaaaaa','111'),
      ('bbb','222','bbbbbbb','121212'),
      ('ccc','333','dddddd','23232323'),
	  ('cccaaa','333','dddddd','23232323');

select * from f_test.f_stu_info
*/

type StuInfo struct {
	FuckStuName string `gorm:"column:f_stu_name"` //这里属性必须大写开头，否则外部不可见，导致映射失败
	FuckStuAge  string `gorm:"column:f_stu_age"`
	FuckStuAdd  string `gorm:"column:f_stu_add`
	FuckMarks   string `gorm:"column:f_marks`
}

func (StuInfo) TableName() string {
	return "f_test.f_stu_info"
}

func TestGP() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "0.0.0.0", 5432, "yuyong", "123456", "testdb")
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

//执行原生SQL查询，以进行复杂查询
//定义结果集
type SqlResult struct {
	SQLStuName string `gorm:"column:f_stu_name"` //这里属性必须大写开头，否则外部不可见，导致映射失败
	SQLStuAge  string `gorm:"column:f_stu_age"`
}

func TestGPSQL() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "0.0.0.0", 5432, "yuyong", "123456", "testdb")
	fmt.Println("get conn str-->" + psqlInfo)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	result := []SqlResult{}
	rows, err := db.Raw("select f_stu_name,f_stu_age from f_test.f_stu_info").Rows()
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		for rows.Next() {
			item := SqlResult{}
			db.ScanRows(rows, &item)
			result = append(result, item)
		}
	}
	fmt.Println("--------------------------------------------------------1")
	fmt.Println(result)
	fmt.Println("--------------------------------------------------------2")
}
