package account

import (
	"fmt"
	// "webApp/models/admin"
	"webApp/controllers"
)

type AccountController struct {
	controllers.BaseController
}

//登陆
func (this *AccountController) Login() {

}

//退出
func (this *AccountController) Logout() {

}

//验证码
func (this *AccountController) VarifiedCode() (code, pic string) {

}
