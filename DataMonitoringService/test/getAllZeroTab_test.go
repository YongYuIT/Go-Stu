package test

import (
	"../tools"
	"fmt"
	"testing"
)

func TestGetAllZeroTab(t *testing.T) {
	tabIntems, err := tools.GetZeroTabMonConf()
	if err != nil {
		fmt.Println("err-->", err)
		return
	}
	saveDbTool := (&tools.TabDataRecordDBTool{}).InitTool("db1_id").(*tools.TabDataRecordDBTool)
	for _, v := range tabIntems {
		fmt.Println(v.DBConf.ID + "-->" + v.ScheName + "-->" + v.Tabname)
		tDbTool := (&tools.TabDataRecordDBTool{}).InitTool(v.DBConf.ID).(*tools.TabDataRecordDBTool)
		count := tDbTool.GetTabDataCount(&v)
		fmt.Println(count)
		saveDbTool.SaveTabCountRecode(count)
	}
}
