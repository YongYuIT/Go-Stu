package controllers

import "github.com/astaxie/beego"
import "../models/dto/req"
import "../models/dto/resp"

type LoginOutController struct {
	beego.Controller
}

func (thiz *LoginOutController) Post() {
	resp := resp.UserLoginResp{}
	resp.Init()

	user := req.UserLoginReq{}
	//这里需要增加表单验证
	thiz.ParseForm(&user)
	resp.RequestCode = user.RequestCode
	beego.Informational("login info: " + user.Username + "-->" + user.Pwd)

	thiz.Data["json"] = &resp
	thiz.ServeJSON()
}
