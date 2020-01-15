package test

import (
	"../service"
	"testing"
)

func TestPrintTab(t *testing.T) {
	tabInfo := "db2_id.sch2.test2"
	pService := service.PrintService{}
	pService.PrintTabPolyline(tabInfo)
}

func TestPrintTabLines(t *testing.T) {
	tabInfo := "db1_id.sch1"
	pService := service.PrintService{}
	pService.PrintSchaPolylines(tabInfo)
}
