package account

import (
	"fmt"
	"github.com/astaxie/beego"
	// . "webApp/common"
	// "html/template"
	// "strings"
	"webApp/controllers"
	. "webApp/models/admin"
)

type AccountController struct {
	controllers.BaseController
}

func (this *AccountController) Index() {
	isVarified, _ := beego.AppConfig.Bool("isOpenVarified")
	// fmt.Println(isVarified)
	if isVarified {
		//生成验证码
	}
	this.Data["isVarified"] = isVarified
	// this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	// fmt.Println("ok")
	this.TplName = "accountcontroller/index.tpl"
}

//ajax登陆
func (this *AccountController) AjaxLogin() {
	username := this.GetString("username")
	password := this.GetString("password")
	isVarified, _ := beego.AppConfig.Bool("isOpenVarified")
	captcha := this.GetString("captcha", "")
	var adminInfo AdminInfo
	fmt.Println(username, password)
	// fmt.Println(password, captcha)
	// fmt.Println(adminInfo)
	if code, Msg := adminInfo.Login(username, password, isVarified, captcha); code == 1 {
		//登陆成功,todo...
		// userInfo := strings.Split(Msg, "|")
		fmt.Println(">>>>>>>>>>>>>>>>>>>>", Msg)
		this.SetSession("userInfo", Msg)
		this.Data["json"] = map[string]interface{}{"code": 1, "msg": "登陆成功", "url": "/"}
	} else {
		//登陆失败,todo...
		this.Data["json"] = map[string]interface{}{"code": 0, "msg": Msg}
	}

	this.ServeJSON()
}

//退出
func (this *AccountController) Logout() {

}

//验证码
func (this *AccountController) VarifiedCode() (code, pic string) {
	return "123", "132"
}
