package models

import (
	// "fmt"
	"github.com/astaxie/beego/orm"
	// "github.com/gogather/com"
)

type User struct {
	Id          int
	Phone       string
	UserProfile *UserProfile `orm:"rel(one)"`
	Password    string
	Status      int
	Created     int64
	Changed     int64
}
type UserProfile struct {
	Id       int
	Realname string
	Sex      int
	Birth    string
	Email    string
	Phone    string
	Address  string
	Hobby    string
	Intro    string
	User     *User `orm:"reverse(one)"`
}

func (this *User) TableName() string {
	return "user"
}
func init() {
	orm.RegisterModel(new(User), new(UserProfile))
}
