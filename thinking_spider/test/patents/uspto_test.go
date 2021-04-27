package patents

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"image"
	"image/jpeg"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"
	"thinking_spider/utils"
	"time"
)

func Test_get_pages_list(test *testing.T) {
	c := colly.NewCollector(
		//不限深度
		func(collector *colly.Collector) {
			extensions.RandomUserAgent(collector)
			extensions.Referer(collector)
		},
	)
	c.SetRequestTimeout(time.Second * 60)
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
		//降低爬取频率
		time.Sleep(2 * time.Second)
	})

	c.OnResponse(func(response *colly.Response) {
		fmt.Println("resp-code-->", response.Request.URL, "-->", response.StatusCode)
		logPath := "./logs/tmp_html/"
		os.MkdirAll(logPath, os.ModePerm)
		file, err := os.Create(fmt.Sprintf(logPath+"%d.html", time.Now().Unix()))
		if err != nil {
			return
		}
		defer file.Close()
		fmt.Fprintf(file, "<!-- "+response.Request.URL.String()+" -->\n")
		//fmt.Fprintf(file, string(response.Body))
		file.Write(response.Body)
	})

	c.OnHTML("TABLE", func(element1 *colly.HTMLElement) {
		if strings.Contains(element1.ChildText("TH[scope='col']"), "PAT. NO.") {
			element1.ForEach("TR", func(i int, element2 *colly.HTMLElement) {
				if strings.Contains(element2.ChildAttr("TD", "valign"), "top") {
					fmt.Println("------------------start")
					element2.ForEach("TD", func(i int, element *colly.HTMLElement) {
						if i == 0 {
							fmt.Println("index-->", element.Text)
						}
						if i == 1 {
							fmt.Println("p.id-->", element.Text)
						}
						if i == 3 {
							fmt.Println("title-->", strings.TrimSpace(element.Text))
							detailUrl := element.ChildAttr("a", "href")
							fmt.Println("url-->", detailUrl)
							startNewCtrl("http://patft.uspto.gov" + detailUrl)
						}
					})
					fmt.Println("------------------end")
				}
			})
		}
	})

	isVisit := make(map[string]bool)

	c.OnHTML("a[href]", func(element *colly.HTMLElement) {
		if strings.Contains(element.ChildAttr("img[align='MIDDLE']", "src"), "nextlist") {
			url := element.Attr("href")
			if !isVisit[url] {
				fmt.Println("next page-->", "http://patft.uspto.gov"+url)
				isVisit[url] = true
				element.Request.Visit("http://patft.uspto.gov" + url)
			}
		}
	})

	startUrl := "http://patft.uspto.gov/netacgi/nph-Parser?Sect1=PTO2&Sect2=HITOFF&p=1&u=%2Fnetahtml%2FPTO%2Fsearch-bool.html&r=0&f=S&l=50&TERM1=Soap+box&FIELD1=&co1=OR&TERM2=Soap+dish&FIELD2=&d=PTXT"
	c.Visit(startUrl)
}

func Test_get_img_pdf(test *testing.T) {
	startNewCtrl("http://patft.uspto.gov" + "/netacgi/nph-Parser?Sect1=PTO2&Sect2=HITOFF&p=1&u=%2Fnetahtml%2FPTO%2Fsearch-bool.html&r=26&f=G&l=50&co1=OR&d=PTXT&s1=%22Soap+box%22&s2=%22Soap+dish%22&OS=\"Soap+box\"+OR+\"Soap+dish\"&RS=\"Soap+box\"+OR+\"Soap+dish\"")
}

func startNewCtrl(startUrl string) {
	c := colly.NewCollector(
		colly.MaxDepth(2),
		func(collector *colly.Collector) {
			extensions.RandomUserAgent(collector)
			extensions.Referer(collector)
		},
	)
	c.SetRequestTimeout(time.Second * 60)
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
		//降低爬取频率
		time.Sleep(2 * time.Second)
	})

	c.OnResponse(func(response *colly.Response) {
		fmt.Println("resp-code-->", response.Request.URL, "-->", response.StatusCode)
		logPath := "./logs/tmp_html/"
		os.MkdirAll(logPath, os.ModePerm)
		file, err := os.Create(fmt.Sprintf(logPath+"%d.html", time.Now().Unix()))
		if err != nil {
			return
		}
		defer file.Close()
		fmt.Fprintf(file, "<!-- "+response.Request.URL.String()+" -->\n")
		//fmt.Fprintf(file, string(response.Body))
		file.Write(response.Body)
	})

	isImaged := false

	c.OnHTML("a[href]", func(element *colly.HTMLElement) {
		if strings.Contains(element.ChildAttr("img", "src"), "image.gif") {
			if !isImaged {
				imgSrc := element.Attr("href")
				fmt.Println("get img src-->", imgSrc)
				element.Request.Visit(imgSrc)
			}
			isImaged = true
		}
	})

	c.OnHTML("embed[type='application/pdf']", func(element *colly.HTMLElement) {
		pdfUrl := element.Attr("src")[2:]
		fmt.Println("get pdf-->", pdfUrl)
		pdfName := strconv.FormatInt(time.Now().Unix(), 10) + ".pdf"
		fmt.Println("save to-->", pdfName)
		downloadFile("https://"+pdfUrl, "./logs/tmp_pdf/", pdfName)
	})

	c.Visit(startUrl)
}

func downloadFile(url string, filepath string, fileName string) {
	client := http.Client{
		Timeout: 60 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("get resource failed")
		return
	}
	defer resp.Body.Close()

	os.MkdirAll(filepath, os.ModePerm)
	file, err := os.Create(filepath + fileName)
	if err != nil {
		return
	}
	defer file.Close()

	reader := &MyReader{
		Reader: resp.Body,
		Total:  resp.ContentLength,
	}

	io.Copy(file, reader)
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

func Test_get_pic_from_PDF(test *testing.T) {
	// sudo apt-get install imagemagick
	// convert -density 300 1619433676.pdf 1619433676.jpg
	// if error "attempt to perform an operation not allowed by the security policy `PDF'"
	// fix as: https://www.jianshu.com/p/48b4c3ff52e2
	// sudo gedit /etc/ImageMagick-6/policy.xml
	// <!-- <policy domain="coder" rights="none" pattern="PDF" /> -->

	path := "./logs/tmp_pdf/"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("read dir error-->", err)
		return
	}
	for i := range files {
		fmt.Println(files[i].Name())
		doConvert(path, files[i].Name())
	}

}

func doConvert(path string, pdfFile string) {

	jpgFile := path + strings.ReplaceAll(pdfFile, "pdf", "jpg")
	sJpgFile := path + strings.ReplaceAll(pdfFile, "pdf", "s.jpg")
	pdfFile = path + pdfFile

	err := utils.RunCommand("convert", "-density", "300", pdfFile, jpgFile)
	if err != nil {
		fmt.Println("covert err-->", err)
		return
	}

	fIn, err := os.Open(jpgFile)
	defer fIn.Close()
	if err != nil {
		fmt.Println("read img err-->", err)
		return
	}

	fOut, _ := os.Create(sJpgFile)
	defer fOut.Close()

	origin, fm, err := image.Decode(fIn)
	if err != nil {
		fmt.Println("read img err-->", err)
		return
	}
	fmt.Println("get img frm-->", fm)

	img := origin.(*image.Gray)
	subImg := img.SubImage(image.Rect(200, 1800, 2200, 3300)).(*image.Gray)
	jpeg.Encode(fOut, subImg, &jpeg.Options{100})
}
