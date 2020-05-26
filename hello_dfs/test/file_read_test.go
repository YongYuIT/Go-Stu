package test

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestBufferRead(test *testing.T) {
	file, err := os.Open("./test_file.png")
	if err != nil {
		fmt.Println("open file err-->", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buf := make([]byte, 10)
	totalSize := 0
	for {
		//虽然缓冲区长度为10，但是每次不一定会读取10个字节到缓冲区，而且读取之前缓冲区不清空。所以要特别注意那些没有读满缓冲区的情况。
		getsize, err := reader.Read(buf)
		if err != nil {
			fmt.Println("read err-->", err)
			break
		}
		totalSize += getsize
		fmt.Println(printHexStr(buf), "-->", getsize, "-->", totalSize)
	}
}

func TestCopy(test *testing.T) {
	file, err := os.Open("./test_file.png")
	if err != nil {
		fmt.Println("open file err-->", err)
		return
	}
	defer file.Close()

	fileCpy, err := os.Create("./test_file_cpy.png")
	if err != nil {
		fmt.Println("create file err-->", err)
		return
	}
	defer fileCpy.Close()

	writer := bufio.NewWriter(fileCpy)
	reader := bufio.NewReader(file)
	buf := make([]byte, 10)
	for {
		getsize, err := reader.Read(buf)
		if err != nil {
			fmt.Println("read err-->", err)
			break
		}
		writesize, err := writer.Write(buf[:getsize])
		//writesize, err := writer.Write(buf)
		if err != nil {
			fmt.Println("write err-->", err)
			break
		}
		fmt.Println("read-->", getsize, "-->,write-->", writesize)
	}
}

func printHexStr(bytes []byte) string {
	reslut := ""
	for i := 0; i < len(bytes); i++ {
		reslut += strconv.FormatInt(int64(bytes[i]&0xff), 16) + " "
	}
	return reslut
}
