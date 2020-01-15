package charts

import (
	"../model"
	"github.com/go-echarts/go-echarts/charts"
	"log"
	"os"
)

type LinePrintTool struct {
}

func (thiz *LinePrintTool) PrintTabDataTabRecords(records []model.TabDataRecord) {
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

func (thiz *LinePrintTool) PrintTabDataSchaRecords(scha_records [][]*model.TabDataRecord) string {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: scha_records[0][0].SchemaName}, charts.InitOpts{Theme: "shine"})
	for i := 0; i < len(scha_records); i++ {
		records := scha_records[i]
		times := []string{}
		counts := []int64{}
		for _, v := range records {
			times = append(times, v.CkechTime.Format("2006-01-02 15:04:05"))
			counts = append(counts, v.Count)
		}
		if i == 0 {
			line = line.AddXAxis(times).AddYAxis(records[0].TabName, counts)
		} else {
			line = line.AddYAxis(records[0].TabName, counts)
		}
	}
	file_name := "./charts_out/" + scha_records[0][0].DBIPPort + ":" + scha_records[0][0].DBName + "." + scha_records[0][0].SchemaName + ".html"
	f, err := os.Create(file_name)
	if err != nil {
		log.Println(err)
	}
	line.Render(f)
	return file_name
}
