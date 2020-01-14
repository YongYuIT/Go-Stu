package tools

import (
	"../model"
	"fmt"
)

type SchemaTabInfoDBTool struct {
	dbConn
}

func (thiz *SchemaTabInfoDBTool) InitTool(id string) DBToolInterfce {
	thiz.dbConn.InitTool(id)
	return thiz
}

func (thiz *SchemaTabInfoDBTool) ReadAllTabsUnderSchema(schameName string) []model.SchemaTabInfo {
	reslut := []model.SchemaTabInfo{}
	if thiz.conn == nil {
		return reslut
	}
	rows, err := thiz.conn.Raw("select table_catalog,table_schema,table_name,table_type from information_schema.tables where table_schema = ?", schameName).Rows()
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		for rows.Next() {
			item := model.SchemaTabInfo{}
			thiz.conn.ScanRows(rows, &item)
			reslut = append(reslut, item)
		}
	}
	return reslut
}
