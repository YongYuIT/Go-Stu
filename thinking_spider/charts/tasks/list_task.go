package tasks

import (
	"charts/database"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"time"
)

const (
	startHtml = "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <title>Title</title>\n</head>\n<script>\n    function fun() {\n        //获取所有的 checkbox 属性的 input标签\n        obj = document.getElementsByName(\"asin_select\");\n        check_val = [];\n        for (k in obj) {\n            //判断复选框是否被选中\n            if (obj[k].checked)\n                //获取被选中的复选框的值\n                check_val.push(\"'\" + obj[k].value + \"'\");\n        }\n        alert(check_val);\n    }\n</script>\n<body>\n<input type='button' value='提交' onclick=\"fun()\"/>\n<table border=\"1\">"
	endHtml   = "</table>\n</body>\n</html>"
)

func DoListTask(sql_file string, contentFunc func(file *os.File, infos []map[string]interface{})) {
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
	contentFunc(file, infos)
	//write data end
	fmt.Fprintf(file, endHtml)
}

const LineCount = 6

func GetGradeContent(file *os.File, infos []map[string]interface{}) {
	line := len(infos) / LineCount
	if line*LineCount < len(infos) {
		line += 1
	}

	keys := []string{}
	for s := range infos[0] {
		keys = append(keys, s)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i][0] < keys[j][0]
	})

	for l := 0; l < line; l++ {
		tabLine := "<tr>"
		for i := 0; i < LineCount; i++ {
			tableItem := "<td> <ul>"
			index := l*LineCount + i
			if index >= len(infos) {
				break
			}
			info := infos[index]
			for k := 0; k < len(keys); k++ {
				key := keys[k]
				value := info[key]

				_, isStr := value.(string)
				if isStr {
					if strings.Contains(key, "pic_url") {
						tableItem += "<img height=200 src=\"" + value.(string) + "\"/>"
					} else {
						tableItem += fmt.Sprintf("<li>%s</li>", value)
					}
				} else {
					_, isInt := value.(int64)
					if isInt {
						tableItem += fmt.Sprintf("<li>%d</li>", value)
					} else {
						tableItem += fmt.Sprintf("<li>%f</li>", value)
					}
				}

			}
			tableItem += "</ul> </td>"
			tabLine += tableItem
		}
		tabLine += "</tr>"
		file.Write([]byte(tabLine))
	}

}

func GetListContent(file *os.File, infos []map[string]interface{}) {
	keys := []string{}
	tabTitle := "<tr>"
	for s := range infos[0] {
		keys = append(keys, s)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i][0] < keys[j][0]
	})
	for i := range keys {
		tabTitle += fmt.Sprintf("<td>%s</td>", keys[i])
	}
	tabTitle += "<td>select</td>"

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
					itemStrFormat += "<td><img height=100 src=\"" + value.(string) + "\"/></td>"
				} else if strings.Contains(key, "_url") {
					itemStrFormat += "<td><a href=\"" + value.(string) + "\" target='_blank'>查看详情</a></td>"
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
		itemStrFormat += fmt.Sprintf("<td><input type=\"checkbox\" name=\"asin_select\" value=\"%s\"/></td>", info[keys[0]])
		itemStrFormat += "</tr>"
		file.Write([]byte(itemStrFormat))
	}
}
