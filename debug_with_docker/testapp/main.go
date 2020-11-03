package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type MyHandler struct {
}

func (this *MyHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	for i := 0; i < 10; i++ {
		fmt.Println("print -->" + strconv.Itoa(i))
	}
	fmt.Fprint(writer, "this is test for MyHandler")
}

func main() {
	hander := MyHandler{}
	ser := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: &hander,
	}
	ser.ListenAndServe()
}
