package tasks

import (
	"charts/database"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

const (
	startHtml = "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <title>Title</title>\n</head>\n<body>\n\n<table border=\"1\">\n"
	endHtml   = "</table>\n</body>\n</html>"
)

func DoListTask(sql_file string) {
	if strings.EqualFold("", sql_file) {
		fmt.Println("you need to set a sql file by -f")
		return
	}
	sqlFile, err := os.Open(sql_file)
	if err != nil {
		fmt.Println("sql file open error!")
		return
	}
	bSQL, err := ioutil.ReadAll(sqlFile)
	if err != nil {
		fmt.Println("sql file read error!")
		return
	}
	queString := string(bSQL)
	outFilePath := "./lists/"
	os.MkdirAll(outFilePath, os.ModePerm)
	file, err := os.Create(fmt.Sprintf(outFilePath+"%d.html", time.Now().Unix()))
	if err != nil {
		return
	}
	defer file.Close()
	fmt.Fprintf(file, startHtml)
	//write data start
	var infos []map[string]interface{}
	database.CurrentDB.Raw(queString).Scan(&infos)
	keys := []string{}
	tabTitle := "<tr>"
	for s := range infos[0] {
		keys = append(keys, s)
		tabTitle += fmt.Sprintf("<td>%s</td>", s)
	}
	tabTitle += "</tr>"
	fmt.Fprintf(file, tabTitle)
	for i, info := range infos {
		fmt.Println("print-->", i)
		itemStrFormat := "<tr name=\"items\"> "
		for i2 := 0; i2 < len(keys); i2++ {
			key := keys[i2]
			value := info[key]
			_, isStr := value.(string)
			if isStr {
				if strings.Contains(key, "pic_url") {
					itemStrFormat += fmt.Sprintf("<td><img height=100 src=\"%s\"/></td>", value)
				} else {
					itemStrFormat += fmt.Sprintf("<td>%s</td>", value)
				}
			} else {
				_, isInt := value.(int64)
				if isInt {
					itemStrFormat += fmt.Sprintf("<td>%d</td>", value)
				} else {
					itemStrFormat += fmt.Sprintf("<td>%f</td>", value)
				}
			}
		}
		itemStrFormat += "</tr>"
		fmt.Fprintf(file, itemStrFormat)
	}
	//write data end
	fmt.Fprintf(file, endHtml)
}
