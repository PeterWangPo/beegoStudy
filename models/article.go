package models

import (
	"github.com/astaxie/beego/orm"
	// _ "github.com/go-sql-driver/mysql"
	// "fmt"
	"strings"
	"time"
	//_ "webApp/initial"
)

type Article struct {
	Id       int
	Title    string
	Uri      string
	Keywords string
	Summary  string
	Content  string
	Author   string
	Created  int64
	Viewnum  int
	Status   int
}

var getTableName = func() string {
	return "article"
}

func init() {
	//orm.RegisterDriver("mysql", orm.DRMySQL)
	//orm.RegisterDataBase("default", "mysql", "root:@/blog?charset=utf8", 30)
	orm.RegisterModel(new(Article))
	//orm.RunSyncdb("default", false, true)
}
func AddArticle(updArt Article) (int64, error) {
	// fmt.Println(time.Now().Unix())
	o := orm.NewOrm()
	o.Using("default")
	art := new(Article)
	art.Title = updArt.Title
	art.Uri = updArt.Uri
	art.Keywords = updArt.Keywords
	art.Summary = updArt.Summary
	art.Content = updArt.Content
	art.Author = updArt.Author
	art.Created = time.Now().Unix()
	art.Viewnum = 1
	art.Status = updArt.Status

	id, err := o.Insert(art)
	return id, err
}
func UpdateArticle(id int, updArt Article) error {
	o := orm.NewOrm()
	o.Using("default")
	art := Article{Id: id}

	art.Title = updArt.Title
	art.Uri = updArt.Uri
	art.Keywords = updArt.Keywords
	art.Summary = updArt.Summary
	art.Content = updArt.Content
	art.Author = updArt.Author
	art.Status = updArt.Status

	_, err := o.Update(&art)
	return err
}
func GetArticle(id int) (Article, error) {
	o := orm.NewOrm()
	o.Using("default")
	art := Article{Id: id}
	err := o.Read(&art)
	return art, err
}

func SearchArticle(arg map[string]string) ([]Article, error) {
	o := orm.NewOrm()
	var sql = "select * from " + getTableName()
	conditions := []string{}
	for k, v := range arg {
		if k == "title" {
			conditions = append(conditions, " title like '%"+v+"%'")
		} else {
			conditions = append(conditions, k+" = "+"'"+v+"'")
		}
	}
	if len(conditions) > 0 {
		sql = sql + " where " + strings.Join(conditions, " and ")
	}
	var arts []Article
	_, err := o.Raw(sql).QueryRows(&arts)
	if err == nil {
		return arts, nil
	} else {
		return nil, err
	}
}

// func SearchArticle(args ...interface{}) (Article, error) {
// 	p :=map[string]interface{}{}

// 	o := orm.NewOrm()
// 	o.Using("default")
// 	//....
// }

func DelArticle(id int) error {
	o := orm.NewOrm()
	o.Using("default")
	_, error := o.Delete(&Article{Id: id})
	return error
}
