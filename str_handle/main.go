package main

import (
	"flag"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"str_handle/tools"
	"strconv"
	"strings"
)

func main() {

	var debug string
	var s1_s2 string
	var s2_s1 string
	flag.StringVar(&s1_s2, "s1_s2", "0.5", "s1/s2 默认0.5")
	flag.StringVar(&s2_s1, "s2_s1", "0.5", "s2/s1 默认0.5")
	flag.StringVar(&debug, "debug", "false", "s2/s1 默认false")
	flag.Parse()

	rates1_s2, err := strconv.ParseFloat(s1_s2, 64)
	if err != nil {
		fmt.Println("s1/s2 rate err-->", err)
		return
	}

	rates2_s1, err := strconv.ParseFloat(s2_s1, 64)
	if err != nil {
		fmt.Println("s2/s1 rate err-->", err)
		return
	}

	fmt.Println("s1/s2", rates1_s2, "s2/s1", rates2_s1, "debug", debug)

	excelFile, err := excelize.OpenFile("comp_datas.xlsx")
	if err != nil {
		fmt.Println("read excel file err:", err)
		return
	}

	s1_row_start := 2
	s3_row_start := 2

	//优化：读取sheet2，避免重复读取、比较
	s2_row_start := 2
	data_Sheet2 := [][2]string{}
	for {
		s2_str := excelFile.GetCellValue("Sheet2", "A"+strconv.Itoa(s2_row_start))
		s2_id := excelFile.GetCellValue("Sheet2", "B"+strconv.Itoa(s2_row_start))
		if strings.EqualFold("", strings.Trim(s2_id, " ")) {
			break
		}
		if strings.EqualFold("", strings.Trim(s2_str, " ")) {
			continue
		}
		data_Sheet2 = append(data_Sheet2, [2]string{s2_str, s2_id})
		s2_row_start += 1
	}

	for {
		s1_str := excelFile.GetCellValue("Sheet1", "A"+strconv.Itoa(s1_row_start))
		s1_id := excelFile.GetCellValue("Sheet1", "B"+strconv.Itoa(s1_row_start))
		if strings.EqualFold("", strings.Trim(s1_id, " ")) {
			break
		}
		if strings.EqualFold("", strings.Trim(s1_str, " ")) {
			continue
		}
		fmt.Println("handling--->", s1_str)

		for sheet2_index := 0; sheet2_index < len(data_Sheet2); sheet2_index++ {
			s2_str := data_Sheet2[sheet2_index][0]
			s2_id := data_Sheet2[sheet2_index][1]
			s1_s2 := tools.Compstr(s2_str, s1_str)
			s2_s1 := tools.Compstr(s1_str, s2_str)
			if s1_s2 > rates1_s2 || s2_s1 > rates2_s1 {
				excelFile.SetCellValue("Sheet3", "A"+strconv.Itoa(s3_row_start), s1_str)
				excelFile.SetCellValue("Sheet3", "B"+strconv.Itoa(s3_row_start), s1_id)
				excelFile.SetCellValue("Sheet3", "C"+strconv.Itoa(s3_row_start), s2_str)
				excelFile.SetCellValue("Sheet3", "D"+strconv.Itoa(s3_row_start), s2_id)
				excelFile.SetCellValue("Sheet3", "E"+strconv.Itoa(s3_row_start), s1_s2)
				excelFile.SetCellValue("Sheet3", "F"+strconv.Itoa(s3_row_start), s2_s1)
				s3_row_start += 1
			}
		}

		excelFile.Save()
		s1_row_start += 1
	}
}
