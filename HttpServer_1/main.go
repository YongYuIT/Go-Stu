package main

import (
	"fmt"
	"./testHandler"
)

func main() {
	fmt.Println("enc_test_success")

	if(1==0){
		//样例1：指定http.Server的Handler
		testHandler.DoTestFunc001()
	}
	if(1==0){
		//样例2：不指定http.Server的Handler
		testHandler.DoTestFunc002()
	}
	if(1==0){
		//样例3，使用HandlerFunc
		testHandler.DoTestFunc003()
	}
	if(1==1){
		//样例4，使用串联HandlerFunc
		testHandler.DoTestFunc004()
	}
}
