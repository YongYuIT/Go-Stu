package service

import (
	"../model"
	"strings"
)
import "../tools"
import "../charts"

type PrintService struct {
}

func (thiz *PrintService) PrintTabPolyline(tabInfo string) {
	info := strings.Split(tabInfo, ".")
	tDbTool := (&tools.TabDataRecordDBTool{}).InitTool("db1_id").(*tools.TabDataRecordDBTool)
	records := tDbTool.GetTabDataRecordByTabInfo(info[0], info[1], info[2])
	printer := &charts.LinePrintTool{}
	if records != nil {
		printer.PrintTabDataTabRecords(records)
	}
}

func (thiz *PrintService) PrintSchaPolylines(scha string) {
	info := strings.Split(scha, ".")
	tDbTool := (&tools.TabDataRecordDBTool{}).InitTool("db1_id").(*tools.TabDataRecordDBTool)
	records := tDbTool.GetTabDataRecordBySchaInfo(info[0], info[1])
	recordsTabs := [][]*model.TabDataRecord{}
	for i := 0; i < len(records); i++ {
		isSave := false
		for j := 0; j < len(recordsTabs); j++ {
			if recordsTabs[j][0] != nil && strings.EqualFold(recordsTabs[j][0].TabName, records[i].TabName) {
				recordsTabs[j] = append(recordsTabs[j], &records[i])
				isSave = true
				break
			}
		}
		if !isSave {
			new_records := []*model.TabDataRecord{}
			new_records = append(new_records, &records[i])
			recordsTabs = append(recordsTabs, new_records)
		}
	}
	printer := &charts.LinePrintTool{}
	if records != nil {
		printer.PrintTabDataSchaRecords(recordsTabs)
	}
}
