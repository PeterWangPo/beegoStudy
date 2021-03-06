package controllers

import (
	// "github.com/astaxie/beego"
	// "github.com/astaxie/beego/utils/pagination"
	"fmt"
	"strconv"
	"strings"
	. "webApp/common"
	m "webApp/models"
)

type ArticleController struct {
	BaseController
}

/**ArticleController这个控制器所有方法的前置方法,如果没有就会查找父级是否存在该方法，存在就执行父级的该方法**/
// func (this *ArticleController) Prepare() {
// 	fmt.Println("prepare done...")
// }
func (this *ArticleController) Finish() {
	fmt.Println("finish done...")
}

/**添加一条文章*/
func (this *ArticleController) Add() {
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

/**删除一条文章*/
func (this *ArticleController) DelOne() {
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

/**批量删除文章
***param ids string id用逗号分隔
**/
func (this *ArticleController) DelBatch() {
	ids := this.GetString("ids")
	if ids == "" {
		this.Data["json"] = GetMsg(0, "参数错误")
		this.ServeJSON()
		return
	}
	_ids := strings.Split(ids, ",")
	success := 0
	failed := 0
	for _, v := range _ids {
		if id, err2 := strconv.Atoi(v); err2 == nil {
			if err := m.DelArticle(id); err != nil {
				failed++
			} else {
				success++
			}
		} else {
			failed++
		}
	}
	if failed == 0 {
		this.Data["json"] = GetMsg(1, "操作成功")
	} else {
		this.Data["json"] = GetMsg(1, "成功删除:"+strconv.Itoa(success)+"条 删除失败:"+strconv.Itoa(failed)+"条")
	}
	this.ServeJSON()
}

/**获取一篇文章详情**/
func (this *ArticleController) Detail() {
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

/**搜索文章**/
func (this *ArticleController) QueryArticle() {
	args := map[string]string{}
	if id, err := this.GetInt("id"); err == nil {
		args["id"] = strconv.Itoa(id)
	}
	if title := this.GetString("title"); title != "" {
		args["title"] = title
	}
	// fmt.Println(args)
	articles, err := m.SearchArticle(args)
	if err == nil {
		this.Data["json"] = GetMsg(1, "获取成功", articles)
	} else {
		this.Data["json"] = GetMsg(0, "获取失败")
	}
	this.ServeJSON()
}
