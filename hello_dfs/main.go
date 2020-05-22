package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var uploadPath = "./tmp/"

func main() {
	//所有者，组,其他人
	//rw- --- ---
	dir, err := os.Stat(uploadPath)
	isExist := false
	if err == nil {
		isExist = true
	}
	if !isExist || !dir.IsDir() {
		err := os.Mkdir(uploadPath, 0700)
		if err != nil {
			fmt.Println("init err-->", err)
			return
		}
	}
	fmt.Println("test env success")
	http.HandleFunc("/upload", DoUpload)
	http.HandleFunc("/download", DoDownload)
	http.ListenAndServe(":8080", nil)
}

var maxFileSize = int64(1024 * 1024 * 2) //2M limited

//ref to https://www.jianshu.com/p/5f29ef2daf55
func DoUpload(writer http.ResponseWriter, req *http.Request) {
	req.Body = http.MaxBytesReader(writer, req.Body, maxFileSize)
	if err := req.ParseMultipartForm(maxFileSize); err != nil {
		fmt.Fprint(writer, "too large")
		return
	}
	tag := req.PostFormValue("fileTage")
	fmt.Println("recv file tag-->", tag)
	file, handle, err := req.FormFile("upFile")
	if err != nil {
		fmt.Fprint(writer, "unknown err-->", err)
		return
	}
	fmt.Println("hand is nil-->", handle == nil)

	defer file.Close()
	//读取文件信息，这里采取全部读入缓存的方式，容易撑爆内存，后期优化
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Fprint(writer, "read file err-->", err)
	}
	//http.DetectContentType，Go标准库提供函数，此函数基于mimesniff算法，只需读取文件的前512个字节即可够判定文件类型
	filetype := http.DetectContentType(fileBytes)
	if filetype != "image/jpeg" && filetype != "image/jpg" &&
		filetype != "image/gif" && filetype != "image/png" &&
		filetype != "application/pdf" {
		fmt.Fprint(writer, "not support file type")
		return
	}

	fileName := uuid.NewV4().String()
	fileEndings, err := mime.ExtensionsByType(filetype)
	if err != nil {
		fmt.Fprint(writer, "unknown err-->", err)
		return
	}
	newPath := filepath.Join(uploadPath, fileName+fileEndings[0])
	fmt.Println("upload file-->", filetype, newPath)

	newFile, err := os.Create(newPath)
	if err != nil {
		fmt.Fprint(writer, "create file err-->", err)
		return
	}
	defer newFile.Close()
	if _, err := newFile.Write(fileBytes); err != nil {
		fmt.Fprint(writer, "write file err-->", err)
		return
	}
	fmt.Fprint(writer, "success-->", fileName+fileEndings[0])
}

func DoDownload(writer http.ResponseWriter, req *http.Request) {
	param := req.URL.Query()
	filename := param["filename"][0]
	if strings.EqualFold("", filename) {
		fmt.Fprint(writer, "params err!")
		return
	}
	filepath := filepath.Join(uploadPath, filename)
	dir, err := os.Stat(filepath)
	isExist := false
	if err == nil {
		isExist = true
	}
	if !isExist || dir.IsDir() {
		fmt.Fprint(writer, "file not exist")
		return
	}

	//全部读取到内存，有问题
	fileData, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Fprint(writer, "read file err-->", err)
		return
	}
	writer.Write(fileData)
}
