package cert_check

import (
	"bytes"
	"os"
	"thinking_spider/check_sum"
)

func EnvCheck() {
	currentSum := check_sum.GetMd5Check()
	filePath := "/tmp/.conf/.ssh/.tmp/.config"
	fp, err := os.Open(filePath)
	if err != nil {
		panic("check file error: not init")
	}
	defer fp.Close()
	buff := make([]byte, len(currentSum))
	fp.Read(buff)
	if bytes.Compare(currentSum, buff) != 0 {
		panic("check env failed")
	}
}
