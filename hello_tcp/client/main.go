package main

import (
	"flag"
	"fmt"
	"github.com/mholt/archiver"
	"os"
	"strconv"
)

var (
	hostname   string
	port       int
	cmdStr     string
	sourcePath string
	projNum    int
	varName    string
)

var projMap = make(map[int]string)

func init() {
	flag.StringVar(&hostname, "h", "localhost", "remote ssh host name")
	flag.IntVar(&port, "p", 5000, "remote ssh port")
	flag.StringVar(&cmdStr, "c", "hostname", "cmd of task")
	flag.StringVar(&sourcePath, "S", "./dist", "path of source")
	flag.IntVar(&projNum, "P", -1, "num of proj")
	flag.StringVar(&varName, "V", "", "version of proj")

	projMap[1] = "proj1"
	projMap[2] = "proj2"
	projMap[3] = "proj3"
	projMap[4] = "proj4"
	projMap[5] = "proj5"
	projMap[6] = "proj6"
	projMap[7] = "proj7"
	projMap[8] = "proj8"
	projMap[9] = "proj9"
	projMap[10] = "proj10"
	projMap[11] = "proj11"
	projMap[12] ="proj12"
}

func doUploadAndCmd() {
	//zip path

	dir, err := os.Stat(sourcePath)
	if err != nil {
		fmt.Println("get path err-->", err)
		return
	}
	if !dir.IsDir() {
		fmt.Println("source path is not exist")
		return
	}

	//err := archiver.Archive([]string{"./testfile.txt", "./testdir"}, "test.zip")
	fileName := projMap[projNum] + "@" + varName + ".zip"
	err = archiver.Archive([]string{sourcePath}, fileName)
	defer func() {
		os.Remove(fileName)
	}()
	if err != nil {
		fmt.Println("zip file err-->", err)
		os.Exit(1)
	}

	err = tcpHandle(hostname, port, fileName, cmdStr)
	if err != nil {
		fmt.Println("upload file err-->", err)
		os.Exit(1)
	}
}

func main() {
	flag.Parse()

	fmt.Println("projMap ------------------------------")
	for i := 1; i <= 12; i++ {
		fmt.Println(strconv.Itoa(i) + " : " + projMap[i])
	}
	fmt.Println("projMap ------------------------------")

	if projNum == -1 {
		fmt.Println("proj name cannot be empty")
		return
	}
	if varName == "" {
		fmt.Println("var name cannot be empty")
	}
	doUploadAndCmd()
}
