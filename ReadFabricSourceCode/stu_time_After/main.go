package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var prt_str string

func main() {
	channel_test := make(chan string)

	go output(channel_test)

	for {
		inputReader := bufio.NewReader(os.Stdin)
		str, err := inputReader.ReadString('\n')
		if err == nil {
			channel_test <- str
			if strings.Compare(str, "stop\n") == 0 {
				break
			}
		}
	}

}

//receive only
func output(ch <-chan string) {
	var timer <-chan time.Time
	for {
		select {
		case x := <-ch:
			prt_str = x
			fmt.Print("receive:" + prt_str)
			timer = time.After(time.Second * 3) //延迟3秒向timer发送信号，触发case <-timer
		case <-timer:
			fmt.Print("doing print:" + prt_str)
		}

	}
}
