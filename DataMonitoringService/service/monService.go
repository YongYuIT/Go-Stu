package service

import (
	"../tools"
	"fmt"
)

type MonService struct {
}

func (thiz *MonService) StartService() {
	dbItems, err := tools.GetDBConfig()
	if err != nil {
		fmt.Println("err-->", err)
	} else {
		fmt.Println("dbItems-->", dbItems)
	}
	tabIntems, err := tools.GetZeroTabMonConf()
	if err != nil {
		fmt.Println("err-->", err)
	} else {
		fmt.Println("tabIntems-->", tabIntems)
	}
}
