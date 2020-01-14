package test

import (
	"../charts"
	"../tools"
	"strings"
	"testing"
)

func TestPrintTab(t *testing.T) {
	tabInfo := "db1_id.sch1.test2"

	info := strings.Split(tabInfo, ".")
	tDbTool := (&tools.TabDataRecordDBTool{}).InitTool("db1_id").(*tools.TabDataRecordDBTool)
	//records := tDbTool.GetTabDataRecordByTabInfo(info[0], info[1], info[2])
	records := tDbTool.GetTabDataRecordBySchaInfo(info[0], info[1])
	printer := &charts.LinePrintTool{}
	if records != nil {
		printer.PrintTabDataRecords(records)
	}
}
