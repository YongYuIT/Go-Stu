package testHandler

import (
	"fmt"
	"net/http"
)

type MyHandler struct {
}

func (this *MyHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprint(writer, "this is test for MyHandler")
}

func DoTestFunc001() {
	hander := MyHandler{}
	ser := http.Server{
		Addr: "0.0.0.0:8080",
		//指定处理器，这样所有发到localhost:8888的请求都归&hander处理
		Handler: &hander,
	}
	ser.ListenAndServe()
}
