package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var (
	app      = kingpin.New("stu_kingpin_v2", "this is stu_kingpin_v2 test")
	start    = app.Command("start", "start this fuck test").Default()
	stu_name = app.Command("stu_name", "show name of this fuck test")
)

func main() {
	fullCmd := kingpin.MustParse(app.Parse(os.Args[1:]))
	fmt.Println("cmd is --> " + fullCmd)
	if fullCmd == stu_name.FullCommand() {
		fmt.Println("this test name is aaaaaa")
	}
}

//go run main.go --help
//go run main.go stu_name
