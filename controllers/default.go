package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	a := c.GetString("a")
	c.Data["Website"] = beego.VERSION
	c.Data["Email"] = a
	c.TplName = "logincontroller/get.tpl"
}

func (c *LoginController) Post() {
	a := c.Input().Get("a")
	c.Data["Website"] = "beego.me111"
	c.Data["Email"] = a
	c.TplName = "logincontroller/get.tpl"
}

func (c *MainController) Get() {
	c.Data["Website"] = beego.VERSION
	c.Data["Email"] = "122211@123.com"
	c.TplName = "index.tpl"
}
