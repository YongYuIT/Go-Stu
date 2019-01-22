package testHandler

import (
	"fmt"
	"net/http"
)

func func_1(func_x http.HandlerFunc) http.HandlerFunc{
	return func(writer http.ResponseWriter,req *http.Request){
		fmt.Println("this is func_1")
		func_x(writer,req)
	}
}

func func_2(func_x http.HandlerFunc) http.HandlerFunc{
	return func(writer http.ResponseWriter,req *http.Request){
		fmt.Println("this is func_2")
		func_x(writer,req)
	}
}

func hello(writer http.ResponseWriter, req *http.Request) {
	fmt.Println( "this is hello")
	fmt.Fprint(writer,"hello")
}

func DoTestFunc004() {
	ser := http.Server{
		Addr:    "localhost:8888",
		//不指定处理器，这样就会使用默认处理器DefaultServeMux
		Handler: nil,
	}
	http.HandleFunc("/test",func_1(func_2(hello)))
	ser.ListenAndServe()
}