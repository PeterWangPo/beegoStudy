package controllers

import (
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/utils/pagination"
	"fmt"
	// "strconv"
	. "webApp/common"
	m "webApp/models"
)

type AddArticleController struct {
	beego.Controller
}

func (this *AddArticleController) Add() {
	// if !this.isLogin {
	// 	this.Data["json"] = map[string]interface{}{"code": 0, "message": "请先登陆"}
	// 	this.ServeJSON()
	// 	return
	// }
	title := this.GetString("title")
	content := this.GetString("content")
	keywords := this.GetString("keywords")
	uri := this.GetString("uri")
	summary := this.GetString("summary")
	author := this.GetString("author")

	if "" == title {
		this.Data["json"] = GetMsg(0, "请填写标题")
		this.ServeJSON()
		return
	}

	if "" == content {
		this.Data["json"] = GetMsg(0, "请填写内容")
		this.ServeJSON()
		return
	}

	var art m.Article
	art.Title = title
	art.Keywords = keywords
	art.Uri = uri
	art.Summary = summary
	art.Content = content
	art.Author = author

	if _, err := m.AddArticle(art); err == nil {
		this.Data["json"] = GetMsg(1, "博客添加成功")
	} else {
		this.Data["json"] = GetMsg(0, "博客添加出错")
	}
	this.ServeJSON()
}

type DelArticleController struct {
	beego.Controller
}

func (this *DelArticleController) Post() {
	id, err1 := this.GetInt("id")
	fmt.Println(id)
	if err1 != nil {
		this.Data["json"] = GetMsg(0, "参数错误")
		this.ServeJSON()
		return
	}
	err := m.DelArticle(id)
	if err == nil {
		this.Data["json"] = GetMsg(1, "删除成功")
	} else {
		this.Data["json"] = GetMsg(0, "删除失败")
	}
	this.ServeJSON()
}

type GetArticleController struct {
	beego.Controller
}

func (this *GetArticleController) Get() {
	if id, err1 := this.GetInt("id"); err1 != nil {
		fmt.Println(err1)
		this.Data["json"] = GetMsg(0, "参数错误")
	} else {
		if art, err := m.GetArticle(id); err != nil {
			this.Data["json"] = GetMsg(0, "获取失败")
		} else {
			this.Data["json"] = GetMsg(1, "获取成功", art)
		}
	}
	this.ServeJSON()
}

type SearchArticleController struct {
	beego.Controller
}

func (this *SearchArticleController) Get() {
	this.Data["json"] = map[string]interface{}{"code": 0, "message": "请先登陆"}
	this.ServeJSON()
}
