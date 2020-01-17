package test

import "testing"
import "../service"

func TestSendEmail(t *testing.T) {
	emailService := service.EmailService{}
	//emailService.SendSimpleSchaInfoEmail("db1_id.sch1", "m13145882354@126.com")
	//emailService.SendSimpleSchaInfoEmail("db1_id.sch1", "13145882354@163.com")
	emailService.SendSimpleSchaInfoEmail("db1_id.sch1", "yuyong_thinking@outlook.com")
}
