package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	"strconv"
	"strings"
)

func main() {
	config := viper.New()
	config.SetConfigName("conf")
	config.SetConfigType("yaml")
	config.AddConfigPath("./")
	if err := config.ReadInConfig(); err != nil {
		fmt.Println(err)
		return
	}
	host_name := config.Get("host_name").(string)
	port := config.Get("port").(string)
	user_name := config.Get("user_name").(string)
	passwd := config.Get("passwd").(string)
	db_name := config.Get("db_name").(string)
	sslmode := config.Get("sslmode").(string)
	query_sql := config.Get("query_sql").(string)
	excel_file := config.Get("excel_file").(string)

	excelFile, err := excelize.OpenFile(excel_file)
	if err != nil {
		fmt.Println("read excel file err:", err)
		return
	}

	sheetName := "Sheet1"
	maxX := getMaxX(excelFile, sheetName)
	startX := (maxX + 1)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host_name, port, user_name, passwd, db_name, sslmode)
	fmt.Println("get conn str-->" + psqlInfo)
	conn, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for i := 1; true; i++ {
		val, _ := excelFile.GetCellValue(sheetName, "A"+strconv.Itoa(i))
		if strings.EqualFold(val, "") {
			break
		}
		p_code := readFirstCol(excelFile, sheetName, i)
		readAddrFromDB(query_sql, conn, p_code, func(info AddrInfo) {
			excelFile.SetCellValue(sheetName, string(startX)+strconv.Itoa(i), info.DetialAddress)
			excelFile.SetCellValue(sheetName, string(startX+1)+strconv.Itoa(i), info.RegionName)
			excelFile.SetCellValue(sheetName, string(startX+2)+strconv.Itoa(i), info.PName)
		})
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

func readFirstCol(excelFile *excelize.File, sheetName string, row int) string {
	val, err := excelFile.GetCellValue(sheetName, "A"+strconv.Itoa(row));
	if err != nil {
		return ""
	}
	return val
}

type AddrInfo struct {
	DetialAddress string `gorm:"column:detial_address"`
	RegionName    string `gorm:"column:region_name"`
	PName         string `gorm:"column:ps_name"`
}

func readAddrFromDB(query_sql string, conn *gorm.DB, code string, onReadCallback func(AddrInfo)) {
	var result_item AddrInfo
	conn.Raw(query_sql, code).Scan(&result_item)
	if strings.EqualFold(result_item.PName, "") {
		fmt.Println("no info for-->", query_sql, code)
	} else {
		onReadCallback(result_item)
	}
}
