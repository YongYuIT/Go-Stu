package tools

import (
	"../model"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"strings"
	"time"
)

var dbConnContext = make(map[string]*gorm.DB)

func init() {
	dbItems, err := GetDBConfig()
	if err != nil {
		fmt.Println("err-->", err)
		return
	} else {
		fmt.Println("dbItems-->", dbItems)
	}
	for i := 0; i < len(dbItems); i++ {
		host := strings.Split(dbItems[i].IPPort, ":")[0]
		port := strings.Split(dbItems[i].IPPort, ":")[1]
		psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, dbItems[i].UserName, dbItems[i].Passwd, dbItems[i].DBName)
		conn, err := gorm.Open("postgres", psqlInfo)
		if err != nil {
			fmt.Println("dberror-->", err)
		} else {
			dbConnContext[dbItems[i].ID] = conn
		}
	}
}

func GetConn(id string) *gorm.DB {
	return dbConnContext[id]
}

func ReadAllTabsUnderSchema(id string, schameName string) []model.SchemaTabInfo {
	reslut := []model.SchemaTabInfo{}
	conn := GetConn(id)
	if conn == nil {
		return reslut
	}
	rows, err := conn.Raw("select table_catalog,table_schema,table_name,table_type from information_schema.tables where table_schema = ?", schameName).Rows()
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		for rows.Next() {
			item := model.SchemaTabInfo{}
			conn.ScanRows(rows, &item)
			reslut = append(reslut, item)
		}
	}
	return reslut
}

func GetTabDataCount(tab *TabMonItem) *model.TabDataRecord {
	conn := GetConn(tab.DBConf.ID)
	if conn == nil {
		return nil
	}
	var count int64 = -1
	conn.Table(tab.ScheName + "." + tab.Tabname).Count(&count)
	record := &model.TabDataRecord{}
	record.DBName = tab.Tabname
	record.TabName = tab.Tabname
	record.CkechTime = time.Now()
	record.Condition = "1=1"
	record.Count = count
	record.DBIPPort = tab.DBConf.IPPort
	record.SchemaName = tab.ScheName
	return record
}

func SaveTabCountRecode(record *model.TabDataRecord, tab *TabMonItem) error {
	conn := GetConn("db1_id")
	if conn == nil {
		return fmt.Errorf("DB conn error")
	}
	conn.Create(record)
	if conn.NewRecord(*record) {
		return fmt.Errorf("create record error")
	}
	return nil
}

/*
docker run --name YongPG1 -e POSTGRES_PASSWORD=123456 -e POSTGRES_USER=yuyong -p 5432:5432 -d postgres
docker run --name YongPG2 -e POSTGRES_PASSWORD=123456 -e POSTGRES_USER=yuyong -p 5433:5432 -d postgres
docker ps -a | grep post

create database db1;
create schema sch1;
grant all on schema sch1 to yuyong;
create table sch1.test1(
    f_stu_name varchar(255),
    f_stu_age varchar(255)
);
grant all on sch1.test1 to yuyong;
create table sch1.test2(
    f_stu_name varchar(255),
    f_stu_age varchar(255)
);
grant all on sch1.test2 to yuyong;
create table sch1.test3(
    f_stu_name varchar(255),
    f_stu_age varchar(255)
);
grant all on sch1.test3 to yuyong;

insert into sch1.test1(f_stu_name,f_stu_age)
values('aaa','111'),
      ('bbb','222'),
      ('ccc','333');

insert into sch1.test2(f_stu_name,f_stu_age)
values('aaa','111'),
      ('bbb','222'),
      ('bbb','444'),
      ('ccc','333');

insert into sch1.test3(f_stu_name,f_stu_age)
values('aaa','111'),
      ('bbb','222'),
      ('ccc','333');

----------------------------------------------------------------

create database db2;
create schema sch2;
grant all on schema sch2 to yuyong;
create table sch2.test1(
    f_stu_name varchar(255),
    f_stu_age varchar(255)
);
grant all on sch2.test1 to yuyong;
create table sch2.test2(
    f_stu_name varchar(255),
    f_stu_age varchar(255)
);
grant all on sch2.test2 to yuyong;

insert into sch2.test1(f_stu_name,f_stu_age)
values('aaa','111'),
      ('bbb','222'),
      ('ccc','333');

insert into sch2.test2(f_stu_name,f_stu_age)
values('aaa','111'),
      ('bbb','222'),
      ('bbb','444'),
      ('ccc','333');
*/
