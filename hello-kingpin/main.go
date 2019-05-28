package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"reflect"
	"strings"
)

var (
	func_name = kingpin.Flag("func_name", "plus or sub").Required().String()
	first_num = kingpin.Arg("first_num", "a num").Default("0").Int()
	next_num  = kingpin.Arg("next_name", "a num").Default("0").Int()
)

func main() {
	kingpin.Parse()
	fmt.Println(reflect.TypeOf(func_name).String())
	fmt.Println(reflect.TypeOf(first_num).String())
	fmt.Println(*func_name)
	if strings.Compare(*func_name, "plus") == 0 {
		fmt.Printf("reslut is %+d \n", (*first_num)+(*next_num))
	} else if strings.Compare(*func_name, "sub") == 0 {
		fmt.Printf("reslut is %+d \n", (*first_num)-(*next_num))
	} else {
		fmt.Println("error \n")
	}
	/*
		$ go run main.go --func_name sub 1 2
		*string
		*int
		sub
		reslut is -1
	*/
}
