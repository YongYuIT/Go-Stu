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
	for _, v := range tabIntems {
		fmt.Println(v.DBConf.ID + "-->" + v.ScheName + "-->" + v.Tabname)
	}
}
