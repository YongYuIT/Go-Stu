package main

import (
	"bufio"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"os"
	"strconv"
	"strings"
)

func main() {

	excelFile, err := excelize.OpenFile("test_file.xlsx")
	if err != nil {
		fmt.Println("read file err:", err)
	}
	f, err := os.Create("output_file.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("txt：将excel中的内容导成TXT文本")
	fmt.Println("comp：比较sheet1和sheet2第一列不同的数据")
	fmt.Println("he：比较sheet1和sheet2第一列的数据，一样的话将sheet2的其他列数据合并在sheet1")
	fmt.Println("concat；将所有列全部拼接在一起")

	inputReader := bufio.NewReader(os.Stdin)
	input_str, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("input error-->", err)
		return
	}
	if strings.EqualFold(input_str, "txt\n") {
		readText(excelFile, f)
	} else if strings.EqualFold(input_str, "comp\n") {
		comp(excelFile, f)
	} else if strings.EqualFold(input_str, "in\n") {
		get_in(excelFile, f)
	} else if strings.EqualFold(input_str, "he\n") {
		hebing(excelFile)
	} else if strings.EqualFold(input_str, "concat\n") {
		concat(excelFile, f)
	}

}

func concat(excelFile *excelize.File, f *os.File) {
	maxSheet1X := getMaxX(excelFile, "Sheet1")
	for i := 1; true; i++ {
		str := ""
		for x := 'A'; x < maxSheet1X; x++ {
			val, _ := excelFile.GetCellValue("Sheet1", string(x)+strconv.Itoa(i))
			str += val
		}
		if strings.EqualFold(str, "") {
			break
		}
		f.Write([]byte(str + "\n"))
	}
}

func hebing(excelFile *excelize.File) {
	maxSheet1X := getMaxX(excelFile, "Sheet1")
	maxSheet2X := getMaxX(excelFile, "Sheet2")

	sheet1 := []string{}
	sheet2 := []string{}
	readSheetToArray(excelFile, &sheet1, "Sheet1")
	readSheetToArray(excelFile, &sheet2, "Sheet2")

	for i := 1; i <= len(sheet1); i++ {
		valSheet1 := sheet1[i-1]
		for j := 1; j <= len(sheet2); j++ {
			valSheet2 := sheet2[j-1]
			if strings.EqualFold(valSheet1, valSheet2) {
				for x := 'A'; x < maxSheet2X; x++ {
					val, _ := excelFile.GetCellValue("Sheet2", string(x)+strconv.Itoa(j))
					excelFile.SetCellValue("Sheet1", string(maxSheet1X+x-'A')+strconv.Itoa(i), val)
				}
				break
			}
		}
	}
	excelFile.Save()
}

func getMaxX(excelFile *excelize.File, sheetName string) int32 {
	maxX := 'A'
	for {
		val, _ := excelFile.GetCellValue(sheetName, string(maxX)+"1")
		if strings.TrimSpace(val) == "" {
			break
		}
		maxX++
	}
	return maxX
}

func get_in(excelFile *excelize.File, f *os.File) {
	sheet1 := []string{}
	readSheetToArray(excelFile, &sheet1, "Sheet1")
	f.Write([]byte("("))
	for i := 0; i < len(sheet1); i++ {
		val := sheet1[i]
		fmt.Println("val-->s", val)
		f.Write([]byte("'" + val + "'"))
		if i < len(sheet1)-1 {
			f.Write([]byte(","))
		}
	}
	f.Write([]byte(")"))
}

func comp(excelFile *excelize.File, f *os.File) {
	sheet1 := []string{}
	sheet2 := []string{}
	readSheetToArray(excelFile, &sheet1, "Sheet1")
	readSheetToArray(excelFile, &sheet2, "Sheet2")
	f.Write([]byte("sheet1 has, sheet2 hasn't:\n"))
	for i := 0; i < len(sheet1); i++ {
		val := sheet1[i]
		has := false
		for _, v := range sheet2 {
			if strings.EqualFold(v, val) {
				has = true
				break
			}
		}
		if !has {
			f.Write([]byte(val + "\n"))
		}
	}

	f.Write([]byte("\n\nsheet2 has, sheet1 hasn't:\n"))
	for i := 0; i < len(sheet2); i++ {
		val := sheet2[i]
		has := false
		for _, v := range sheet1 {
			if strings.EqualFold(v, val) {
				has = true
				break
			}
		}
		if !has {
			f.Write([]byte(val + "\n"))
		}
	}
}

func readSheetToArray(excelFile *excelize.File, array *[]string, sheetName string) {
	for i := 1; true; i++ {
		val, err := excelFile.GetCellValue(sheetName, "A"+strconv.Itoa(i))
		if err != nil {
			fmt.Println("read file err:", err)
		} else {
			if strings.TrimSpace(val) == "" {
				break
			}
			*array = append(*array, val)
		}
	}
}

func readText(excelFile *excelize.File, f *os.File) {

	for i := 1; true; i++ {
		val, err := excelFile.GetCellValue("Sheet1", "A"+strconv.Itoa(i))
		if err != nil {
			fmt.Println("read file err:", err)
		} else {
			fmt.Println("val-->s", val)
			f.Write([]byte(val + "\n"))
			if strings.TrimSpace(val) == "" {
				break
			}
		}
	}
}
