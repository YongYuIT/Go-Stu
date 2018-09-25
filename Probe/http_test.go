package Probe

import (
	"testing"
	"net/http"
	"fmt"
)

//w表示response对象
//r表示request对象
func do_test(w http.ResponseWriter, r *http.Request) {
	test_key := r.PostFormValue("test_key")
	//返回内容
	fmt.Fprintf(w, "Hello golang http "+test_key)
}

func TestHttp(t *testing.T) {
	http.HandleFunc("/do_test", do_test)
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Printf("Start Error: ", err)
	}
}
