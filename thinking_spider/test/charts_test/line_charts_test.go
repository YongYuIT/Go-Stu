package charts_test

import (
	"fmt"
	"github.com/go-echarts/go-echarts/charts"
	v2charts "github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"math/rand"
	"os"
	"testing"
)

func Test_hello_line(test *testing.T) {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: "by ratings"})
	times := []string{}
	counts := []int64{}

	times = append(times, "2006-01-02 15:04:05")
	times = append(times, "2006-01-03 15:04:05")
	times = append(times, "2006-01-04 15:04:05")
	times = append(times, "2006-01-05 15:04:05")

	counts = append(counts, 1)
	counts = append(counts, 2)
	counts = append(counts, 1)
	counts = append(counts, 3)

	line.AddXAxis(times).AddYAxis("ratings", counts)
	file_name := "./test.html"
	f, err := os.Create(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	line.Render(f)
}

// generate random data for line chart
func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}
	return items
}

func Test_mut(test *testing.T) {
	// create a new line instance
	line := v2charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		v2charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		v2charts.WithTitleOpts(opts.Title{
			Title:    "Line example in Westeros theme",
			Subtitle: "Line chart rendered by the http server this time",
		}))

	// Put data into instance
	line.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
		AddSeries("Category A", generateLineItems()).
		AddSeries("Category B", generateLineItems()).
		SetSeriesOptions(v2charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

	file_name := "./test.html"
	f, err := os.Create(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	line.Render(f)
}
