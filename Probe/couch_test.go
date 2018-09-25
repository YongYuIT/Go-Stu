package Probe

import (
	"testing"
	"net/http"
	"io/ioutil"
	"fmt"
)

/*
$ echo "deb https://apache.bintray.com/couchdb-deb xenial main" \
    | sudo tee -a /etc/apt/sources.list
$ curl -L https://couchdb.apache.org/repo/bintray-pubkey.asc \
    | sudo apt-key add -
$ sudo apt-get update && sudo apt-get install couchdb
*/
func readResp(resp *http.Response) {
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Printf("Read Error: ", err)
		return
	}
	fmt.Println(string(body))
}

func TestCouch(t *testing.T) {

	resp, err := http.Get("http://127.0.0.1:5984")
	if err != nil {
		fmt.Printf("Get Error: ", err)
		return
	}
	readResp(resp)

	resp, err = http.Get("http://127.0.0.1:5984/_all_dbs")
	if err != nil {
		fmt.Printf("Get Error: ", err)
		return
	}
	readResp(resp)

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPut, "http://admin:123456@127.0.0.1:5984/test_db", nil)
	resp, err = client.Do(req)
	if err != nil {
		fmt.Printf("Get Error: ", err)
		return
	}
	readResp(resp)

}
