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
	startHtml = "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <title>Title</title>\n</head>\n<body>\n\n<table border=\"1\">\n    <tr>\n        <td>KeyWord</td>\n        <td>Page</td>\n        <td>PageIndex</td>\n        <td>Asin</td>\n        <td>Titles</td>\n        <td>Ratings</td>\n        <td>Starts</td>\n        <td>Price</td>\n        <td>MainPicUrl</td>\n        <td>Select</td>\n    </tr>"
	endHtml   = "</table>\n</body>\n</html>"
)

type InfoWithPic struct {
	KeyWord    string
	Page       int
	PageIndex  int
	Asin       string
	Titles     string
	Ratings    int
	Starts     float32 `sql:"type:decimal(10,2);"`
	Price      float32 `sql:"type:decimal(10,2);"`
	MainPicUrl string  `sql:"type:text;"`
}

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
	var infos []InfoWithPic
	database.CurrentDB.Raw(queString).Scan(&infos)
	for i, info := range infos {
		fmt.Println("print-->", i)
		itemStr := fmt.Sprintf("<tr name=\"items\"> "+
			"<td>%s</td>"+
			"<td>%d</td>"+
			"<td>%d</td>"+
			"<td>%s</td>"+
			"<td>%s</td>"+
			"<td>%d</td>"+
			"<td>%f</td>"+
			"<td>%f</td>"+
			"<td><img height=100 src=\"%s\"/></td>"+
			"<td><input type='checkbox' value=%d name=\"prod\" /></td>"+
			"</tr>",
			info.KeyWord,
			info.Page,
			info.PageIndex,
			info.Asin,
			info.Titles,
			info.Ratings,
			info.Starts,
			info.Price,
			info.MainPicUrl,
			i,
		)
		fmt.Fprintf(file, itemStr)
	}
	//write data end
	fmt.Fprintf(file, endHtml)
}
