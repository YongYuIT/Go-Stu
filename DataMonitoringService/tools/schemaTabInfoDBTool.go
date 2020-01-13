package tools

import (
	"../model"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type SchemaTabInfoDBTool struct {
	conn *gorm.DB
}

func (thiz *SchemaTabInfoDBTool) InitTool(id string) *SchemaTabInfoDBTool {
	thiz.conn = GetConn(id)
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

func (thiz *SchemaTabInfoDBTool) GetTabDataCount(tab *TabMonItem) *model.TabDataRecord {
	if thiz.conn == nil {
		return nil
	}
	var count int64 = -1
	thiz.conn.Table(tab.ScheName + "." + tab.Tabname).Count(&count)
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

func (thiz *SchemaTabInfoDBTool) SaveTabCountRecode(record *model.TabDataRecord) error {
	if thiz.conn == nil {
		return fmt.Errorf("DB conn error")
	}
	thiz.conn.Create(record)
	if thiz.conn.NewRecord(*record) {
		return fmt.Errorf("create record error")
	}
	return nil
}
