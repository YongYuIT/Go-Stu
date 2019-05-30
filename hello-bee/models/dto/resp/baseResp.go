package resp

import "time"

type baseResp struct {
	Result      string `json:"reslut"`
	ResultCode  int    `json:"recult_code"`
	RequestCode int64  `json:"request_code"`
	ResultDesc  string `json:"result_desc"`
	Timestamp   int64  `json:"reslut_time"`
}

func (this *baseResp) Init() {
	this.Result = "known error"
	this.ResultCode = ERROR
	this.ResultDesc = "known error"
	this.Timestamp = time.Now().UnixNano()
}
