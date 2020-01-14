package charts

import (
	"../model"
	"github.com/go-echarts/go-echarts/charts"
	"log"
	"os"
)

type LinePrintTool struct {
}

func (thiz *LinePrintTool) PrintTabDataRecords(records []model.TabDataRecord) {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: records[0].SchemaName})
	times := []string{}
	counts := []int64{}
	for _, v := range records {
		times = append(times, v.CkechTime.Format("2006-01-02 15:04:05"))
		counts = append(counts, v.Count)
	}
	line.AddXAxis(times).AddYAxis(records[0].TabName, counts)
	file_name := "./charts_out/" + records[0].DBIPPort + ":" + records[0].DBName + "." + records[0].SchemaName + "." + records[0].TabName + ".html"
	f, err := os.Create(file_name)
	if err != nil {
		log.Println(err)
	}
	line.Render(f)
}
