package tools

import (
	"fmt"
	"github.com/beevik/etree"
	"github.com/jinzhu/gorm"
	"strings"
)

func GetConnStrForMySql(mysql_conf_info map[string]interface{}) string {
	return fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		mysql_conf_info["user_name"].(string),
		mysql_conf_info["pwd"].(string),
		mysql_conf_info["host"].(string),
		mysql_conf_info["port"].(string),
		mysql_conf_info["db_name"].(string))
}

func GetConnStrForPG(pg_conf_info map[string]interface{}) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		pg_conf_info["host"].(string),
		pg_conf_info["port"].(string),
		pg_conf_info["user_name"].(string),
		pg_conf_info["pwd"].(string),
		pg_conf_info["db_name"].(string))
}

func GetConn(dbType string, connStr string) *gorm.DB {
	conn, err := gorm.Open(dbType, connStr)
	if err != nil {
		fmt.Println("mysql conn failed-->", err)
		return nil
	}
	fmt.Println("conn mysql success")
	return conn
}

func HandleTabIncreMsg(ktr *etree.Element, tabInfo map[interface{}]interface{}, mysqlConn *gorm.DB) string {

	//处理 in step
	steps := ktr.FindElements("step")
	var step_in *etree.Element
	var step_out *etree.Element
	for i := 0; i < len(steps); i++ {
		if strings.EqualFold("TabInput", steps[i].FindElement("name").Text()) {
			step_in = steps[i]
		}
		if strings.EqualFold("InsertUpdate", steps[i].FindElement("name").Text()) {
			step_out = steps[i].FindElement("lookup")
		}
	}
	fmt.Println("step_in is found-->", step_in == nil)
	fmt.Println("step_out is found-->", step_out == nil)

	mysql_tab := strings.Split(tabInfo["source_tab"].(string), ".")[1]
	step_in.FindElement("sql").SetText(fmt.Sprintf("select * from %s ", mysql_tab) + "where update_time>STR_TO_DATE('${MAX_MODIFYRQ}', '%Y-%m-%d %H:%i:%s')")

	pg_tab := tabInfo["tag_tab"].(string)
	gp_tab_infos := strings.Split(pg_tab, ".")
	step_out.FindElement("schema").SetText(gp_tab_infos[0])
	step_out.FindElement("table").SetText(gp_tab_infos[1])

	return mysql_tab
}

func HandleTabFullMsg(ktr *etree.Element, tabInfo map[interface{}]interface{}, mysqlConn *gorm.DB) string {
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
	mysql_tab := tabInfo["source_tab"].(string)
	pg_tab := tabInfo["tag_tab"].(string)
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
