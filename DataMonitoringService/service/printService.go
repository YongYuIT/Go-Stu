package service

import "strings"
import "../tools"
import "../charts"

type PrintService struct {
}

func (thiz *PrintService) PrintTabPolyline(tabInfo string) {
	info := strings.Split(tabInfo, ".")
	tDbTool := (&tools.TabDataRecordDBTool{}).InitTool(info[0]).(*tools.TabDataRecordDBTool)
	records := tDbTool.GetTabDataRecordByTabInfo(info[0], info[1], info[2])
	//records := tDbTool.GetTabDataRecordBySchaInfo(info[0], info[1])
	printer := &charts.LinePrintTool{}
	if records != nil {
		printer.PrintTabDataRecords(records)
	}
}
