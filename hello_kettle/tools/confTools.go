package tools

import (
	"fmt"
	"github.com/beevik/etree"
	"github.com/jinzhu/gorm"
	"strings"
)

func FillConnWithConf(elm *etree.Element, conf map[string]interface{}) {
	elm.FindElement("server").SetText(conf["host"].(string))
	elm.FindElement("username").SetText(conf["user_name"].(string))
	elm.FindElement("password").SetText(conf["pwd"].(string))
	elm.FindElement("database").SetText(conf["db_name"].(string))
	elm.FindElement("port").SetText(conf["port"].(string))
}

func HandleTab(elm *etree.Element, handle_str string, mysqlConn *gorm.DB, gpConn *gorm.DB) {
	//获取表所有字段
	tabs := strings.Split(handle_str, "-->")
	mysql_tab := tabs[0]
	pg_tab := tabs[1]
	fmt.Println("mysql,", mysql_tab, "gp,", pg_tab)
	mysql_tab_infos := strings.Split(mysql_tab, ".")
	get_mysql_tab := fmt.Sprintf("select COLUMN_NAME from INFORMATION_SCHEMA.COLUMNS Where table_name = '%s' AND table_schema = '%s'", mysql_tab_infos[1], mysql_tab_infos[0])
	fmt.Println(get_mysql_tab)
	rows, err := mysqlConn.Raw(get_mysql_tab).Rows()
	if err != nil {
		fmt.Println("read mysql err-->", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		item := MySQLColumInfo{}
		mysqlConn.ScanRows(rows, &item)
		fmt.Println(item)
	}
}

type MySQLColumInfo struct {
	COLUMN_NAME string `gorm:"column:COLUMN_NAME"`
}
