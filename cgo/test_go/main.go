package main

/*
#cgo CFLAGS : -I./include
#cgo LDFLAGS: -L./lib -ltest
#include "test.h"
*/
import "C"

import "fmt"

func main() {
	fmt.Println(C.GoString(C.test_hello(C.CString("yuyong"))));
}

//export LD_LIBRARY_PATH=/home/yong/Desktop/cgo-test20180723001/test_001/lib
