package main

import (
	"github.com/astaxie/beego"
	_ "webApp/initial"
	_ "webApp/routers"
)

func main() {
	// beego.BConfig.WebConfig.Session.SessionProvider = "file"
	// beego.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"
	beego.Run()
}
