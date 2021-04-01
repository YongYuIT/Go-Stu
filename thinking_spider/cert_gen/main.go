package main

import (
	"bytes"
	"cert_gen/check_sum"
	"encoding/binary"
	"fmt"
	"os"
)

func main() {
	thisCryBs := check_sum.GetMd5Check()
	keyPath := "/tmp/.conf/.ssh/.tmp/"
	err := os.MkdirAll(keyPath, 0700)
	if err != nil {
		panic("cannot create path" + err.Error())
	}
	file, err := os.Create(keyPath + ".config")
	if err != nil {
		panic("cannot create file" + err.Error())
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, thisCryBs)
	file.Write(buf.Bytes())

	fmt.Println("init success!")
}
