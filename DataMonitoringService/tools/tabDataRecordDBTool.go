package tools

import (
	"../model"
	"fmt"
	"time"
)

type TabDataRecordDBTool struct {
	dbConn
}

func (thiz *TabDataRecordDBTool) InitTool(id string) DBToolInterfce {
	thiz.dbConn.InitTool(id)
	return thiz
}

func (thiz *TabDataRecordDBTool) GetTabDataCount(tab *TabMonItem) *model.TabDataRecord {
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

func (thiz *TabDataRecordDBTool) SaveTabCountRecode(record *model.TabDataRecord) error {
	if thiz.conn == nil {
		return fmt.Errorf("DB conn error")
	}
	thiz.conn.Create(record)
	if thiz.conn.NewRecord(*record) {
		return fmt.Errorf("create record error")
	}
	return nil
}
