package testHandler

import (
	"fmt"
	"net/http"
)

type HelloHandler struct {
}

func (this *HelloHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprint(writer, "this is test for HelloHandler")
}

type WordHandler struct {

}

func (this *WordHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprint(writer, "this is test for WordHandler")
}

func DoTestFunc002() {
	ser := http.Server{
		Addr:    "localhost:8888",
		//不指定处理器，这样就会使用默认处理器DefaultServeMux
		Handler: nil,
	}
	http.Handle("/hello",&HelloHandler{})
	http.Handle("/word",&WordHandler{})
	ser.ListenAndServe()
}