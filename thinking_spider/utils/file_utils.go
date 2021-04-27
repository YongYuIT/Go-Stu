package utils

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func DownloadFile(url string, filepath string, fileName string) bool {
	client := http.Client{
		Timeout: 60 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("get resource failed")
		return false
	}
	defer resp.Body.Close()

	os.MkdirAll(filepath, os.ModePerm)
	file, err := os.Create(filepath + fileName)
	if err != nil {
		return false
	}
	defer file.Close()

	reader := &MyReader{
		Reader: resp.Body,
		Total:  resp.ContentLength,
	}

	_, err = io.Copy(file, reader)
	if err != nil {
		return false
	}
	return true
}

type MyReader struct {
	io.Reader
	Total   int64
	Current int64
}

func (r *MyReader) Read(p []byte) (n int, err error) {
	n, err = r.Reader.Read(p)
	r.Current += int64(n)
	fmt.Printf("\rdownload... %.2f%%", float64(r.Current*10000/r.Total)/100)
	return
}

func CreateSImgFromPatentsPDF(pdfPath string, pdfFile string) string {
	jpgFile := pdfPath + strings.ReplaceAll(pdfFile, ".pdf", ".jpg")
	sJpgFile := pdfPath + strings.ReplaceAll(pdfFile, ".pdf", ".s.jpg")
	pdfFile = pdfPath + pdfFile

	_, err := os.Stat(sJpgFile)
	if err == nil {
		return sJpgFile
	}

	err = RunCommand("convert", "-density", "300", pdfFile, jpgFile)
	if err != nil {
		fmt.Println("covert err-->", err)
		return ""
	}

	fIn, err := os.Open(jpgFile)
	defer fIn.Close()
	if err != nil {
		fmt.Println("read img err-->", err)
		return ""
	}

	fOut, _ := os.Create(sJpgFile)
	defer fOut.Close()

	origin, fm, err := image.Decode(fIn)
	if err != nil {
		fmt.Println("read img err-->", err)
		return ""
	}
	fmt.Println("get img frm-->", fm)

	img := origin.(*image.Gray)
	subImg := img.SubImage(image.Rect(200, 1800, 2200, 3300)).(*image.Gray)
	jpeg.Encode(fOut, subImg, &jpeg.Options{100})
	return sJpgFile
}
