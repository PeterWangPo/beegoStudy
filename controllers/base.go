package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

type BaseController struct {
	beego.Controller
	IsLogin  bool
	UserId   int64
	Username string
}

func (this *BaseController) Prepare() {
	sessionStr := this.GetSession("userInfo")
	fmt.Println(sessionStr)
	if sessionStr != nil {
		sessStr := strings.Split((sessionStr).(string), "|")
		this.IsLogin = true
		this.UserId, _ = strconv.ParseInt(sessStr[0], 10, 64)
		this.Username = sessStr[1]
	}
	fmt.Println(this.UserId)
	fmt.Println(this.Username)
	fmt.Println("base prepare done...")
}
