package req

import "../../do"
import "../../../utils"

type RegisterReq struct {
	baseReq
	Name string `form:"user_name"`
	Pwd  string `form:"user_pwd"`
}

func (thiz *RegisterReq) GetUserDO() *do.UserDo {
	user := do.UserDo{}
	user.UserId = utils.Uuid()
	user.UserName = thiz.Name
	user.PassWd = utils.Md5(thiz.Pwd)
	//变量逃逸，C++里面这样写就直接回收了（未验证）
	return &user
}
