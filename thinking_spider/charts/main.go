package main

import (
	"charts/database"
	_ "charts/database"
	"flag"
	"fmt"
	v2charts "github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"time"
)

var (
	table string
)

func init() {
	flag.StringVar(&table, "t", "table_example", "table name to export")
}

type LineItem struct {
	DataTime  time.Time `gorm:"column:data_time"`
	DataValue float32   `gorm:"column:data_value" sql:"type:decimal(10,2);"`
	DataKey   string    `gorm:"column:data_key"`
}

func (LineItem) TableName() string {
	return table
}

type Key struct {
	Key string `gorm:"column:data_key"`
}

type TimeShow struct {
	TimeValue time.Time `gorm:"column:data_time"`
}

func main() {
	flag.Parse()
	defer database.CloseDB()

	keys := []Key{}
	rows, err := database.CurrentDB.Raw("select data_key from table_example group by (data_key)").Rows()
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		for rows.Next() {
			item := Key{}
			database.CurrentDB.ScanRows(rows, &item)
			keys = append(keys, item)
		}
	}

	times := []TimeShow{}
	rows, err = database.CurrentDB.Raw("select data_time from table_example group by data_time order by data_time").Rows()
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		for rows.Next() {
			item := TimeShow{}
			database.CurrentDB.ScanRows(rows, &item)
			times = append(times, item)
		}
	}

	// create a new line instance
	line := v2charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		v2charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		v2charts.WithTitleOpts(opts.Title{
			Title:    "Line from amazon",
			Subtitle: "from tab: " + table,
		}))

	// Put data into instance
	var timesStr []string
	for _, show := range times {
		timesStr = append(timesStr, show.TimeValue.String())
	}
	line = line.SetXAxis(timesStr)
	for _, key := range keys {
		line = line.AddSeries(key.Key, generateLineItems(key.Key, times), func(s *v2charts.SingleSeries) {
			s.Name = key.Key
		})
	}

	line.SetSeriesOptions(v2charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

	file_name := "./" + table + ".html"
	f, err := os.Create(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	line.Render(f)
}

// generate random data for line chart
func generateLineItems(key string, times []TimeShow) []opts.LineData {

	var lineItems []LineItem
	database.CurrentDB.Find(&lineItems, "data_key = ?", key).Order("data_time")
	fmt.Println("get-->", key, "-->", len(lineItems))

	var lastValue float32 = 0
	items := make([]opts.LineData, 0)
	for _, item := range lineItems {
		if findTime(times, item.DataTime) {
			items = append(items, opts.LineData{Value: item.DataValue})
		} else {
			items = append(items, opts.LineData{Value: lastValue})
		}
		lastValue = item.DataValue
	}

	return items
}

func findTime(times []TimeShow, time time.Time) bool {
	for _, show := range times {
		if show.TimeValue.Equal(time) {
			return true
		}
	}
	return false
}
