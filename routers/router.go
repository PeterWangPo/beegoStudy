package routers

import (
	"github.com/astaxie/beego"
	"webApp/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/article/add", &controllers.ArticleController{}, "post:Add")
	beego.Router("/article/del", &controllers.ArticleController{}, "post:DelOne")
	beego.Router("/article/detail", &controllers.ArticleController{}, "get:Detail")
}
