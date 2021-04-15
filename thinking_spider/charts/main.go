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
	Key
	TimeValue
}

type TimeValue struct {
	DataValue float32 `gorm:"column:data_value" sql:"type:decimal(10,2);"`
	DTime
}

type Key struct {
	DataKey string `gorm:"column:data_key"`
}

type DTime struct {
	DataTime time.Time `gorm:"column:data_time"`
}

func (LineItem) TableName() string {
	return table
}

func main() {
	flag.Parse()
	defer database.CloseDB()

	//1. 统计出应画几条线
	var keys []Key
	database.CurrentDB.Raw("select data_key from table_example group by (data_key)").Scan(&keys)
	//2. 统计完全时间轴
	times := []DTime{}
	database.CurrentDB.Raw("select data_time from table_example group by data_time order by data_time").Scan(&times)
	//3. 根据实际数据，参照完全时间轴，补齐缺失数据
	kvs := make(map[string]*[]TimeValue)
	for i := range keys {
		kvs[keys[i].DataKey] = &[]TimeValue{}
		thisKVS := kvs[keys[i].DataKey]
		var lineItems []LineItem
		database.CurrentDB.Find(&lineItems, "data_key = ?", keys[i].DataKey).Order("data_time")
		for i2 := range times {
			*thisKVS = append(*thisKVS, TimeValue{})
			(*thisKVS)[i2].DataTime = times[i2].DataTime
			(*thisKVS)[i2].DataValue = -1
			for i3 := range lineItems {
				if lineItems[i3].DataTime.Equal((*thisKVS)[i2].DataTime) {
					(*thisKVS)[i2].DataValue = lineItems[i3].DataValue
				}
			}
			if (*thisKVS)[i2].DataValue == -1 {
				if i2 > 0 {
					(*thisKVS)[i2].DataValue = (*thisKVS)[i2-1].DataValue
				}
			}
		}

	}
	//画图
	// create a new line instance
	line := v2charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		v2charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		v2charts.WithTitleOpts(opts.Title{
			Title:    "Line from amazon",
			Subtitle: "from tab: " + table,
		}),
		v2charts.WithLegendOpts(opts.Legend{Show: true}),
	)

	// Put data into instance
	var timesStr []string
	for _, show := range times {
		timesStr = append(timesStr, show.DataTime.String())
	}
	line = line.SetXAxis(timesStr)
	for i := range keys {
		line = line.AddSeries(keys[i].DataKey, generateLineItems(*kvs[keys[i].DataKey]))
	}
	line.SetSeriesOptions(v2charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

	//output to html file
	file_name := "./" + table + ".html"
	f, err := os.Create(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	line.Render(f)

}

// generate random data for line chart
func generateLineItems(tValues []TimeValue) []opts.LineData {
	items := make([]opts.LineData, 0)
	for _, item := range tValues {
		items = append(items, opts.LineData{Value: item.DataValue})
	}
	return items
}
