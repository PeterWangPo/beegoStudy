package routers

import (
	"github.com/astaxie/beego"
	"webApp/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/article/add", &controllers.AddArticleController{}, "post:Add")
	beego.Router("/article/del", &controllers.DelArticleController{})
	beego.Router("/article/list", &controllers.GetArticleController{})
	beego.Router("/article/search", &controllers.SearchArticleController{})
}
