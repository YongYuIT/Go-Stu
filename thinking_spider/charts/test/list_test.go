package test

import (
	"charts/database"
	"fmt"
	"os"
	"testing"
	"time"
)

const (
	startHtml = "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <title>Title</title>\n</head>\n<body>\n\n<table border=\"1\">"
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

func Test_printLit(test *testing.T) {

	queString := "select * from key_word_prod_records where uuid != ''"

	logPath := "./logs/tmp_html/"
	os.MkdirAll(logPath, os.ModePerm)
	file, err := os.Create(fmt.Sprintf(logPath+"%d.html", time.Now().Unix()))
	if err != nil {
		return
	}
	defer file.Close()
	fmt.Fprintf(file, startHtml)
	fmt.Fprintf(file, "    <tr>\n        <td>KeyWord</td>\n        <td>Page</td>\n        <td>PageIndex</td>\n        <td>Asin</td>\n        <td>Titles</td>\n        <td>Ratings</td>\n        <td>Starts</td>\n        <td>Price</td>\n        <td>MainPicUrl</td>\n        <td>Select</td>\n    </tr>")
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
