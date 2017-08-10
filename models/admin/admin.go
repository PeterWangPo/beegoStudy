package admin

import (
	"github.com/astaxie/beego/orm"
	"webApp/models"
)

type AdminInfo struct {
	Id       int
	Username string
	Password string
	BaseModel
}

func (this *AdminInfo) GetAdminInfo() {

}

func (this *AdminInfo) Login(username, password, verifyCode string) {

}

func (this *AdminInfo) Logout() {

}
