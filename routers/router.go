package routers

import (
	"github.com/astaxie/beego"
	"webApp/controllers"
	"webApp/controllers/account"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	// beego.Router("/login", &controllers.LoginController{})
	// beego.Router("/article/add", &controllers.ArticleController{}, "post:Add")
	// beego.Router("/article/del", &controllers.ArticleController{}, "post:DelOne")
	// beego.Router("/article/delBatch", &controllers.ArticleController{}, "post:DelBatch")
	// beego.Router("/article/detail", &controllers.ArticleController{}, "get:Detail")
	// beego.Router("/article/queryArticle", &controllers.ArticleController{}, "get:QueryArticle")

	//用户
	beego.Router("/account/login", &account.AccountController{}, "post:AjaxLogin")
	beego.Router("/account/index", &account.AccountController{}, "get:Index")
}
