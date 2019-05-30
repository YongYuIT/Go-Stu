package controllers

import (
	"github.com/astaxie/beego"
)
import "../models/dto/req"
import "../models/dto/resp"

type RegisterController struct {
	beego.Controller
}

func (thiz *RegisterController) Post() {
	registerResp := resp.RegisterResp{}
	registerResp.Init()

	user := req.RegisterReq{}
	//这里需要增加表单验证
	thiz.ParseForm(&user)
	registerResp.RequestCode = user.RequestCode

	userDo := user.GetUserDO()
	_, err := userDo.Save()
	if err != nil {
		beego.Error("insert user error:" + err.Error())
		thiz.Abort("500")
	}

	registerResp.ResultDesc = "success"
	registerResp.Result = userDo.UserId
	registerResp.ResultCode = resp.SUCCESS

	thiz.Data["json"] = &registerResp
	thiz.ServeJSON()
}
