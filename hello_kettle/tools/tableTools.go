package tools

import (
	"fmt"
	"github.com/beevik/etree"
	"github.com/jinzhu/gorm"
	"strings"
)

func HandleTab(ktr *etree.Element, handle_str string, mysqlConn *gorm.DB, gpConn *gorm.DB) string {
	//处理 in step
	steps := ktr.FindElements("step")
	var step_in *etree.Element
	var step_out *etree.Element
	for i := 0; i < len(steps); i++ {
		if strings.EqualFold("tmp_tab_in", steps[i].FindElement("name").Text()) {
			step_in = steps[i]
		}
		if strings.EqualFold("tmp_tab_out", steps[i].FindElement("name").Text()) {
			step_out = steps[i]
		}
	}
	fmt.Println("step_in is found-->", step_in == nil)
	fmt.Println("step_out is found-->", step_out == nil)
	tabs := strings.Split(handle_str, "-->")
	mysql_tab := tabs[0]
	pg_tab := tabs[1]
	fmt.Println("mysql,", mysql_tab, "gp,", pg_tab)
	mysql_tab_infos := strings.Split(mysql_tab, ".")
	get_mysql_tab := fmt.Sprintf("select COLUMN_NAME from INFORMATION_SCHEMA.COLUMNS Where table_name = '%s' AND table_schema = '%s'", mysql_tab_infos[1], mysql_tab_infos[0])
	fmt.Println(get_mysql_tab)
	step_in.FindElement("sql").SetText(fmt.Sprintf("select * from %s", mysql_tab_infos[1]))
	step_in_name := mysql_tab_infos[1] + "_full_input"
	step_in.FindElement("name").SetText(step_in_name)
	//处理 out step
	gp_tab_infos := strings.Split(pg_tab, ".")
	step_out.FindElement("schema").SetText(gp_tab_infos[0])
	step_out.FindElement("table").SetText(gp_tab_infos[1])
	//获取表所有字段
	fields_elm := step_out.FindElement("fields")
	remove_els := fields_elm.ChildElements()
	for i := 0; i < len(remove_els); i++ {
		fields_elm.RemoveChild(remove_els[i])
	}
	rows, err := mysqlConn.Raw(get_mysql_tab).Rows()
	if err != nil {
		fmt.Println("read mysql err-->", err)
		return ""
	}
	defer rows.Close()
	for rows.Next() {
		item := MySQLColumInfo{}
		mysqlConn.ScanRows(rows, &item)
		fmt.Println(item)
		field := fields_elm.CreateElement("field")
		field_yuan := field.CreateElement("stream_name")
		field_yuan.SetText(item.COLUMN_NAME)
		field_mubiao := field.CreateElement("column_name")
		field_mubiao.SetText(item.COLUMN_NAME)
		field.AddChild(field_yuan)
		field.AddChild(field_mubiao)
		fields_elm.AddChild(field)
	}
	step_out_name := mysql_tab_infos[1] + "_full_output"
	step_out.FindElement("name").SetText(step_out_name)
	line := ktr.FindElement("order").FindElement("hop")
	line.FindElement("from").SetText(step_in_name)
	line.FindElement("to").SetText(step_out_name)
	return mysql_tab_infos[1]
}

type MySQLColumInfo struct {
	COLUMN_NAME string `gorm:"column:COLUMN_NAME"`
}
