package model

import (
	"time"
)

type TabDataRecord struct {
	ID         int       `gorm:"primary_key;column:id"`
	DBIPPort   string    `gorm:"column:db_ip_port"`
	DBName     string    `gorm:"column:db_name"`
	SchemaName string    `gorm:"column:schema_name"`
	TabName    string    `gorm:"column:tab_name"`
	Count      int64     `gorm:"column:check_count"`
	Condition  string    `gorm:"column:condition"`
	CkechTime  time.Time `gorm:"column:check_time"`
}

func (TabDataRecord) TableName() string {
	return "etl_tem.tab_mon_record"
}

/*
create schema etl_tem;
grant all on schema etl_tem to yuyong;
create table etl_tem.tab_mon_record(
id serial PRIMARY KEY,
db_ip_port varchar(255),
db_name varchar(255),
schema_name varchar(255),
tab_name varchar(255),
condition varchar(255),
check_count bigint,
check_time timestamp(0)
);
grant all on etl_tem.tab_mon_record to yuyong;
*/
