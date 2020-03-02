package excel

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
	"time"
)

var READ_EXCEL_ERROR = "READ_EXCEL_ERROR"

type ExcelTool struct {
	excelFile *excelize.File
}

func (thiz *ExcelTool) InitTool(filePath string) *ExcelTool {
	eFile, err := excelize.OpenFile(filePath)
	if err == nil {
		thiz.excelFile = eFile
	} else {
		thiz.excelFile = nil
	}
	return thiz
}

func (thiz *ExcelTool) ReadValue(sheet string, p *Point) string {
	val, err := thiz.excelFile.GetCellValue(sheet, p.GetPosition())
	if err == nil {
		return val
	} else {
		return READ_EXCEL_ERROR
	}
}

func (thiz *ExcelTool) GetOrgName(sheet string, p *Point) string {
	return thiz.ReadValue(sheet, p.GetOrgName())
}

func (thiz *ExcelTool) GetDate(sheet string, p *Point) string {
	return convertToFormatDay(thiz.ReadValue(sheet, p.GetDate()))
}

func (thiz *ExcelTool) GetDataItemName(sheet string, p *Point) string {
	return thiz.ReadValue(sheet, p.GetDataItemName())
}

func (thiz *ExcelTool) GetOrgType(sheet string, p *Point) string {
	return thiz.ReadValue(sheet, p.GetOrgType())
}

func (thiz *ExcelTool) GetQu(sheet string, p *Point) string {
	return thiz.ReadValue(sheet, p.GetQu())
}

// excel日期字段格式化 yyyy-mm-dd
func convertToFormatDay(excelDaysString string) string {
	// 2006-01-02 距离 1900-01-01的天数
	baseDiffDay := 38719 - 1 //在网上工具计算的天数需要加2天，什么原因没弄清楚
	curDiffDay := excelDaysString
	b, _ := strconv.Atoi(curDiffDay)
	// 获取excel的日期距离2006-01-02的天数
	realDiffDay := b - baseDiffDay
	//fmt.Println("realDiffDay:",realDiffDay)
	// 距离2006-01-02 秒数
	realDiffSecond := realDiffDay * 24 * 3600
	//fmt.Println("realDiffSecond:",realDiffSecond)
	// 2006-01-02 15:04:05距离1970-01-01 08:00:00的秒数 网上工具可查出
	baseOriginSecond := 1136185445
	resultTime := time.Unix(int64(baseOriginSecond+realDiffSecond), 0).Format("2006-01-02")
	return resultTime
}
