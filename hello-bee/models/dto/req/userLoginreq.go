package req

type UserLoginReq struct {
	baseReq
	Username string `form:"username"`
	Pwd      string `form:"userPassWord"`
}
