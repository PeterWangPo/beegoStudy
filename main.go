package main

import (
	"github.com/astaxie/beego"
	_ "webApp/initial"
	_ "webApp/routers"
)

func main() {
	beego.Run()
}
