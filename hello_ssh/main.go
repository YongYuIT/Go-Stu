package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

var (
	task       string
	hostname   string
	port       int
	cert       string
	user       string
	cmdStr     string
	sourcePath string
	remotePath string
)

func init() {
	flag.StringVar(&task, "t", "upload", "name of task: upload/cmd")
	flag.StringVar(&hostname, "h", "localhost", "remote ssh host name")
	flag.IntVar(&port, "p", 22, "remote ssh port")
	flag.StringVar(&cert, "i", "~/.ssh/id_rsa", "remote ssh cert")
	flag.StringVar(&user, "u", "root", "remote ssh username")
	flag.StringVar(&cmdStr, "c", "hostname", "cmd of task")
	flag.StringVar(&sourcePath, "S", "./", "path of source")
	flag.StringVar(&remotePath, "D", "./", "path of remote")

}
func main() {
	flag.Parse()

	if strings.EqualFold("upload", task) {
		upload()
	}
	if strings.EqualFold("cmd", task) {
		cmd()
	}

}
func cmd() {
	Exec_Cmd(hostname, user, cmdStr, port, cert)
}

func upload() {
	sourcePath = sourcePath + "/"
	_sourcePath := sourcePath
	handleDir(_sourcePath, "", func(path string, fileName string) {
		fmt.Println("get path-->", path, "-->", fileName)
		fullPath := path + "/" + fileName
		oppoPath := path[len(sourcePath)+1:]
		_remotePath := remotePath + "/" + oppoPath
		fmt.Println("remote-->", _remotePath)
		Exec_Cmd(hostname, user, "mkdir -p "+_remotePath+"/", port, cert)
		UploadFile(hostname, user, fullPath, _remotePath, port, cert)
	})
}

func handleDir(path string, filename string, handle func(string, string)) {
	fmt.Println("getpathXXX-->", path, "-->", filename)
	path = path + "/" + filename
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("read path err-->", err)
		return
	}
	for _, fi := range dir {
		if fi.IsDir() {
			handleDir(path, fi.Name(), handle)
		} else {
			handle(path, fi.Name())
		}
	}
}
