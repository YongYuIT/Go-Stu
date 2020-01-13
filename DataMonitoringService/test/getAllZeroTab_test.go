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
	saveDbTool := (&tools.SchemaTabInfoDBTool{}).InitTool("db1_id")
	for _, v := range tabIntems {
		fmt.Println(v.DBConf.ID + "-->" + v.ScheName + "-->" + v.Tabname)
		sDbTool := (&tools.SchemaTabInfoDBTool{}).InitTool(v.DBConf.ID)
		count := sDbTool.GetTabDataCount(&v)
		fmt.Println(count)
		saveDbTool.SaveTabCountRecode(count)
	}
}
