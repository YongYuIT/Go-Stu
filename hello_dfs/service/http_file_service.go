package service

import (
	"bufio"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var UploadPath = "./tmp/"

type TKHHttpService struct {
	Port int
}

func (this *TKHHttpService) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	var pfunc func(http.ResponseWriter, *http.Request)
	switch path {
	case "/upload":
		pfunc = this.DoUpload
	case "/upload_big":
		pfunc = this.DoUploadBig
	case "/download":
		pfunc = this.DoDownload
	default:
		pfunc = nil
	}
	if pfunc == nil {
		fmt.Fprint(writer, "unknown path")
		return
	}
	pfunc(writer, req)
}

//需改造成单例
func (this *TKHHttpService) StartService() error {
	ser := http.Server{
		Addr:    "0.0.0.0:" + strconv.Itoa(this.Port),
		Handler: this,
	}
	mime.AddExtensionType(".rar", "application/x-rar-compressed")
	mime.AddExtensionType(".iso", "application/octet-stream")
	return ser.ListenAndServe()
}

var maxFileSize = int64(1024 * 1024 * 2) //2M limited
var bigMaxFileSize = int64(1024 * 1024 * 1024 * 8)

//ref to https://www.jianshu.com/p/5f29ef2daf55
func (this *TKHHttpService) DoUpload(writer http.ResponseWriter, req *http.Request) {
	req.Body = http.MaxBytesReader(writer, req.Body, maxFileSize)
	if err := req.ParseMultipartForm(maxFileSize); err != nil {
		fmt.Fprint(writer, "too large")
		return
	}
	tag := req.PostFormValue("fileTage")
	fmt.Println("recv file tag-->", tag)
	file, header, err := req.FormFile("upFile")
	if err != nil {
		fmt.Fprint(writer, "unknown err-->", err)
		return
	}
	fmt.Println("header is nil-->", header == nil)
	if header != nil {
		fmt.Println("header size-->", header.Size)
	}
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
	newPath := filepath.Join(UploadPath, fileName+fileEndings[0])
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

func (this *TKHHttpService) DoUploadBig(writer http.ResponseWriter, req *http.Request) {
	//采取分片读写的方法，避免撑爆内存
	req.Body = http.MaxBytesReader(writer, req.Body, bigMaxFileSize)
	if err := req.ParseMultipartForm(bigMaxFileSize); err != nil {
		fmt.Fprint(writer, "too large")
		return
	}

	tag := req.PostFormValue("fileTage")
	fmt.Println("recv file tag-->", tag)
	//断点调试可知，此处文件已经加入内存（或许在内存缓冲区），故此这种处理大文件上传的方法也不好
	file, header, err := req.FormFile("upFile")
	if err != nil {
		fmt.Fprint(writer, "unknown err-->", err)
		return
	}
	fmt.Println("header is nil-->", header == nil)
	if header != nil {
		fmt.Println("header size-->", header.Size)
	}
	defer file.Close()

	fileName := uuid.NewV4().String()
	newPath := filepath.Join(UploadPath, fileName)
	fmt.Println("upload file-->", newPath)
	newFile, err := os.Create(newPath)
	if err != nil {
		fmt.Fprint(writer, "create file err-->", err)
		return
	}
	defer newFile.Close()

	reader := bufio.NewReader(file)
	fwriter := bufio.NewWriter(newFile)
	fileTypeName := ""
	buf := make([]byte, 1024)
	for i := 0; true; i++ {
		getsize, err := reader.Read(buf)
		if err != nil {
			fmt.Println("read err-->", err)
			break
		}

		if i == 0 {
			filetype := http.DetectContentType(buf)
			if filetype != "image/jpeg" && filetype != "image/jpg" &&
				filetype != "image/gif" && filetype != "image/png" &&
				filetype != "application/pdf" && filetype != "application/x-rar-compressed" &&
				filetype != "application/octet-stream" {
				fmt.Fprint(writer, "not support file type")
				return
			}
			fileTypeNames, err := mime.ExtensionsByType(filetype)
			if err != nil {
				fmt.Fprint(writer, "unknown err-->", err)
				return
			}
			fileTypeName = fileTypeNames[0]
		}

		writesize, err := fwriter.Write(buf[:getsize])
		if err != nil {
			fmt.Fprint(writer, "write err-->", err)
			break
		}
		fmt.Println("read-->", getsize, "-->write-->", writesize, "-->", i)
	}
	fnewPath := newPath + fileTypeName
	os.Rename(newPath, fnewPath)
	fmt.Fprint(writer, "success-->", fnewPath, "-->", newPath)
}

func (this *TKHHttpService) DoDownload(writer http.ResponseWriter, req *http.Request) {
	param := req.URL.Query()
	filename := param["filename"][0]
	if strings.EqualFold("", filename) {
		fmt.Fprint(writer, "params err!")
		return
	}
	filepath := filepath.Join(UploadPath, filename)
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
