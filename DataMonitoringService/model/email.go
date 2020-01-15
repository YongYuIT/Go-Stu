package model

type Email struct {
	ID        int    `gorm:"primary_key;column:id"`
	From      string `gorm:"column:e_from"`
	To        string `gorm:"column:e_to"`
	Cc        string `gorm:"column:cc"`
	Bcc       string `gorm:"column:bcc"`
	Title     string `gorm:"column:title"`
	NeedReply bool   `gorm:"column:need_reply"`
	Content   string `gorm:"column:content"`
	SendTime  string `gorm:"column:send_time"`
}

func (Email) TableName() string {
	return "etl_tem.tab_mon_email"
}

/*
create table etl_tem.tab_mon_email(
id serial PRIMARY KEY,
e_from varchar(255),
e_to varchar(255),
cc varchar(255),
bcc varchar(255),
title varchar(255),
need_reply int2,
content text,
send_time timestamp(0)
);
grant all on etl_tem.tab_mon_email to yuyong;
*/
