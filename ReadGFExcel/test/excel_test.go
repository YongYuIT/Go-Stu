package test

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"testing"
)

func TestReadFromExcel(t *testing.T) {
	f, err := excelize.OpenFile("AAA.xlsx")
	if err != nil {
		println(err.Error())
		return
	}
	// Get value from cell by given worksheet name and axis.
	//cell, err := f.GetCellValue("sheet1", "AW4")
	//cell, err := f.GetCellValue("sheet1", "A4")
	//cell, err := f.GetCellValue("sheet1", "A7")
	cell, err := f.GetCellValue("sheet1", "A8")
	if err != nil {
		println(err.Error())
		return
	}
	println(cell)
}
