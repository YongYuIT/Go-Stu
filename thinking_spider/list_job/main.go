package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"list_job/utils"
	"os"
	"strings"
)

const (
	KEYWORD_TASK = "KEYWORD_TASK"
	DETAIL_TASK  = "DETAIL_TASK"
)

func init() {
	flag.StringVar(&task, "t", KEYWORD_TASK, "set task")
}

var (
	task string
)

func main() {

	file, err := os.Open("key_word.list")
	if err != nil {
		fmt.Printf("list open failed --> ", err)
		return
	}

	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		fmt.Print("doing job-->", str)
		str = strings.ReplaceAll(strings.TrimSpace(str), " ", "+")
		if !strings.EqualFold(str, "") {
			utils.RunCommand("./thinking_spider", "-k", str, "-t", task)
		}
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("run cmd err-->", err)
			break
		}
	}
}
