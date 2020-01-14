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

func (thiz *TabDataRecordDBTool) CalculateTabDataRecord(tab *TabMonItem) *model.TabDataRecord {
	if thiz.conn == nil {
		return nil
	}
	var count int64 = -1
	thiz.conn.Table(tab.ScheName + "." + tab.Tabname).Count(&count)
	record := &model.TabDataRecord{}
	record.DBName = tab.DBConf.DBName
	record.TabName = tab.Tabname
	record.CkechTime = time.Now()
	record.Condition = "1=1"
	record.Count = count
	record.DBIPPort = tab.DBConf.IPPort
	record.SchemaName = tab.ScheName
	return record
}

func (thiz *TabDataRecordDBTool) GetTabDataRecordByTabInfo(id string, schName string, tabName string) []model.TabDataRecord {
	conf := GetDBConfigByID(id)
	qurystr := fmt.Sprintf("db_ip_port = '%s' and db_name = '%s' and schema_name = '%s' and tab_name = '%s'", conf.IPPort, conf.DBName, schName, tabName)
	return thiz.getTabDataRecordCondition(qurystr)
}

func (thiz *TabDataRecordDBTool) GetTabDataRecordBySchaInfo(id string, schName string) []model.TabDataRecord {
	conf := GetDBConfigByID(id)
	qurystr := fmt.Sprintf("db_ip_port = '%s' and db_name = '%s' and schema_name = '%s'", conf.IPPort, conf.DBName, schName)
	return thiz.getTabDataRecordCondition(qurystr)
}

func (thiz *TabDataRecordDBTool) getTabDataRecordCondition(condition string) []model.TabDataRecord {
	result := []model.TabDataRecord{}
	thiz.conn.Where(condition).Order("check_time asc").Find(&result)
	return result
}

func (thiz *TabDataRecordDBTool) SaveTabDataRecode(record *model.TabDataRecord) error {
	if thiz.conn == nil {
		return fmt.Errorf("DB conn error")
	}
	thiz.conn.Create(record)
	if thiz.conn.NewRecord(*record) {
		return fmt.Errorf("create record error")
	}
	return nil
}
