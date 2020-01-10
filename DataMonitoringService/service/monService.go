package service

import (
	"../tools"
	"fmt"
)

type MonService struct {
}

func (thiz *MonService) StartService() {
	tabIntems, err := tools.GetZeroTabMonConf()
	if err != nil {
		fmt.Println("err-->", err)
	} else {
		fmt.Println("tabIntems-->", tabIntems)
	}
	for _, v := range tabIntems {
		count := tools.GetTabDataCount(&v)
		fmt.Println(count)
	}
}
