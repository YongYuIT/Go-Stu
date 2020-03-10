package test

import (
	"fmt"
	"github.com/hpcloud/tail"
	"testing"
	"time"
)

func TestTailFileChanges(t *testing.T) {

	//单元测试不支持输入，所以需要拷贝的主函数测试、运行

	filename := "/home/yong/Go-Stu20200302001/ETLLogMonitoringService/test_path/test1.txt"
	tails, err := tail.TailFile(filename, tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})

	if err != nil {
		fmt.Println("tail file err:", err)
		return
	}

	var msg *tail.Line
	var ok bool
	for true {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen,filenam:%s\n", tails, filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		fmt.Println("msg:", msg.Text)
	}
}
