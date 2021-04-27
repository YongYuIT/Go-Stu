package patents

import (
	"fmt"
	"net/url"
	"testing"
)

func Test_get_url_params(test *testing.T) {
	thisUrl, err := url.Parse("http://patft.uspto.gov/netacgi/nph-Parser?Sect1=PTO2&Sect2=HITOFF&p=1&u=%2Fnetahtml%2FPTO%2Fsearch-bool.html&r=0&f=S&l=50&TERM1=Soap+box&FIELD1=&co1=OR&TERM2=Soap+dish&FIELD2=&d=PTXT")
	if err != nil {
		fmt.Println("error1", err)
		return
	}
	kvs, err := url.ParseQuery(thisUrl.RawQuery)
	if err != nil {
		fmt.Println("error2", err)
		return
	}
	for s := range kvs {
		val := kvs[s]
		fmt.Println(s, "-->", val)
	}

	//change
	kvs.Set("TERM1", "Silicone  gloves")
	kvs.Set("TERM2", "Anti-scald gloves")
	changed := kvs.Encode()
	fmt.Println(changed)
}
