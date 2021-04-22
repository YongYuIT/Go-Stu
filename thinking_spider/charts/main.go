package main

import (
	_ "charts/database"
	"charts/tasks"
	"flag"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strings"
)

var (
	table   string
	task    string
	sqlFile string
)

func init() {
	flag.StringVar(&table, "t", "table_example", "table name to export")
	flag.StringVar(&sqlFile, "f", "", "set sql file")
	flag.StringVar(&task, "T", LINE_TASK, "set task")
}

const (
	LINE_TASK = "LINE_TASK"
	LIST_TASK = "LIST_TASK"
)

func main() {

	flag.Parse()

	if strings.EqualFold(LINE_TASK, task) {
		tasks.TableName = table
		tasks.DoLineTask()
	} else if strings.EqualFold(LIST_TASK, task) {
		tasks.DoListTask(sqlFile)
	} else {
		fmt.Println("you need to set a task")
	}
}
