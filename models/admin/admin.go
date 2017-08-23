package admin

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	. "webApp/common"
)

type AdminInfo struct {
	Id       int
	Username string
	Password string
	Status   string
}

func init() {
	orm.RegisterModel(new(AdminInfo))
}

func (this *AdminInfo) GetAdminInfo() {

}

func (this *AdminInfo) Login(username string, password string, isverifed bool, verifyCode string) string {
	if username == "" {
		return "密码为空"
	}
	if password == "" {
		return "密码为空"
	}
	qb, db_err := orm.NewQueryBuilder("mysql")
	if db_err != nil {
		return "登陆失败"
	}
	var adminInfo AdminInfo
	qb.Select("id", "username", "password", "status").From(Common_getTableName("user")).Where("username = ? and status = 1").Limit(1)
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql, username).QueryRow(&adminInfo)
	if adminInfo.Id <= 0 {
		return "账号不存在"
	} else {
		if Common_md5(password) == adminInfo.Password {
			return ""
		} else {
			return "密码错误"
		}
	}
}

func (this *AdminInfo) Logout() {

}
