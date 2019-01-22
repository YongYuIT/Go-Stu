package testHandler

import (
	"fmt"
	"net/http"
)

type TestHandler1 struct {
}

func (this *TestHandler1) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprint(writer, "this is test for 1")
}

func Test2(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprint(writer, "this is test for 2")
}

func Test3(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprint(writer, "this is test for 3")
}


func DoTestFunc003() {
	ser := http.Server{
		Addr:    "localhost:8888",
		//不指定处理器，这样就会使用默认处理器DefaultServeMux
		Handler: nil,
	}
	//以下三种方式等价，方式1
	http.Handle("/test1",&TestHandler1{})
	//方式2
	testHandler2:=http.HandlerFunc(Test2)
	http.Handle("/test2",&testHandler2)
	//方式3
	http.HandleFunc("/test3",Test3)
	ser.ListenAndServe()
}