package main

import (
	"./excel"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
	"strings"
)

func main() {

	exFile := "AAA.xlsx"

	sheetName := "sheet1"
	startPoint := &excel.Point{"DY", "4"}
	endPoint := &excel.Point{"EO", "83"}
	var orgs []OrgDataItem

	readFromExcel(exFile,
		sheetName,
		startPoint,
		endPoint,
		&orgs)

	sheetName = "sheet2"
	endPoint = &excel.Point{"EO", "63"}

	readFromExcel(exFile,
		sheetName,
		startPoint,
		endPoint,
		&orgs)

	fmt.Println("all or:=g items-->", len(orgs))
	err := WriteToExcel(orgs)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("out success")
	}
}

type OrgDataItem struct {
	DataItemName string
	Date         string
	OrgName      string
	OrgType      string
	Qu           string
	Value        string
	Mark         string
	EXPosition   string
}

func readFromExcel(fileName string, sheetName string, startPoint *excel.Point, endPoint *excel.Point, orgs *[]OrgDataItem) {
	eTool := excel.ExcelTool{}
	eTool.InitTool(fileName)
	thizPoint := startPoint
	for { //for X
		for { //for Y
			fmt.Println(thizPoint.GetPosition())
			var data OrgDataItem
			data.DataItemName = eTool.GetDataItemName(sheetName, thizPoint)
			data.Date = eTool.GetDate(sheetName, thizPoint)
			data.OrgName = eTool.GetOrgName(sheetName, thizPoint)
			data.OrgType = eTool.GetOrgType(sheetName, thizPoint)
			data.Qu = eTool.GetQu(sheetName, thizPoint)
			data.Value = eTool.ReadValue(sheetName, thizPoint)
			data.Mark = sheetName
			data.EXPosition = thizPoint.GetPosition()
			if strings.TrimSpace(data.Value) == "" {
				fmt.Println("igo")
			} else {
				fmt.Println(data)
				*orgs = append(*orgs, data)
			}
			fmt.Println("-------------------------------")

			next_y := thizPoint.GetNextY()
			if next_y.Y == endPoint.Y {
				break
			}
			thizPoint = next_y
		}
		next_x := thizPoint.GetNextX()
		next_x.Y = startPoint.Y
		if next_x.X == endPoint.X {
			break
		}
		thizPoint = next_x
	}
}

func WriteToExcel(orgs []OrgDataItem) error {
	f := excelize.NewFile()
	index := f.NewSheet("data_items")

	f.SetCellValue("data_items", "A"+strconv.Itoa(1), "DataItemName")
	f.SetCellValue("data_items", "B"+strconv.Itoa(1), "Date")
	f.SetCellValue("data_items", "C"+strconv.Itoa(1), "OrgName")
	f.SetCellValue("data_items", "D"+strconv.Itoa(1), "OrgType")
	f.SetCellValue("data_items", "E"+strconv.Itoa(1), "Qu")
	f.SetCellValue("data_items", "F"+strconv.Itoa(1), "Value")
	f.SetCellValue("data_items", "G"+strconv.Itoa(1), "Mark")
	f.SetCellValue("data_items", "H"+strconv.Itoa(1), "EXPosition")

	for i := 0; i < len(orgs); i++ {
		f.SetCellValue("data_items", "A"+strconv.Itoa(i+2), orgs[i].DataItemName)
		f.SetCellValue("data_items", "B"+strconv.Itoa(i+2), orgs[i].Date)
		f.SetCellValue("data_items", "C"+strconv.Itoa(i+2), orgs[i].OrgName)
		f.SetCellValue("data_items", "D"+strconv.Itoa(i+2), orgs[i].OrgType)
		f.SetCellValue("data_items", "E"+strconv.Itoa(i+2), orgs[i].Qu)
		f.SetCellValue("data_items", "F"+strconv.Itoa(i+2), orgs[i].Value)
		f.SetCellValue("data_items", "G"+strconv.Itoa(i+2), orgs[i].Mark)
		f.SetCellValue("data_items", "H"+strconv.Itoa(i+2), orgs[i].EXPosition)
	}
	f.SetActiveSheet(index)
	if err := f.SaveAs("output20200227002.xlsx"); err != nil {
		return err
	} else {
		return nil
	}
}
