package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
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
		if strings.EqualFold(str, "") {
			continue
		}
		fmt.Print("doing job-->", str)

		runCommand("./thinking_spider", "-k", str[:len(str)-1], "-t", task)

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("run cmd err-->", err)
			break
		}
	}
}

func runCommand(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	// 命令的错误输出和标准输出都连接到同一个管道
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout

	if err != nil {
		return err
	}

	if err = cmd.Start(); err != nil {
		return err
	}
	// 从管道中实时获取输出并打印到终端
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		fmt.Print(string(tmp))
		if err != nil {
			break
		}
	}

	if err = cmd.Wait(); err != nil {
		return err
	}
	return nil
}
