package models

import (
	"fmt"
	"github.com/astaxie/beego"
)

type BaseModel struct {
}

func (this *BaseModel) TableName(tableName string) {
	return beego.AppConfig.String("mysqlprefix") + tableName
}
