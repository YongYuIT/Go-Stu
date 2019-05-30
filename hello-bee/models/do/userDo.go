package do

import (
	"database/sql"
	"github.com/astaxie/beego/orm"
)

type UserDo struct {
	UserId   string
	UserName string
	PassWd   string
	Token    string
}

func (thiz *UserDo) Save() (sql.Result, error) {
	o := orm.NewOrm()
	return o.Raw("insert into user set userId=? , userName=? , passwd=?", thiz.UserId, thiz.UserName, thiz.PassWd).Exec()
}
