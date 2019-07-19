package controllers

import (
	"encoding/json"
	"go_demo/models"

	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

//添加文章
func (this *ArticleController) AddArticle() {
	var requestBody map[string]string
	json.Unmarshal(this.Ctx.Input.RequestBody, &requestBody)
	title := requestBody["title"]
	content := requestBody["content"]
	id, err := models.Add(content, title)
	if err == nil {
		this.Data["json"] = map[string]interface{}{"status": 200, "id": id, "msg": "插入成功"}
		this.ServeJSON()
		return
	}
}

//获取所有文章
func (this *ArticleController) GetAllArticle() {
	article, err := models.GetAll()
	if err == nil {
		this.Data["json"] = map[string]interface{}{"status": 200, "msg": article}
		this.ServeJSON()
		return
	}
}

//根据ID获取单个文章
func (this *ArticleController) GetOneArticle() {
	id := this.Ctx.Input.Param(":id")
	article, err := models.GetOne(id)
	if err == nil {
		this.Data["json"] = map[string]interface{}{"status": 200, "msg": article}
		this.ServeJSON()
		return
	}
}

//根据ID更新文章
func (this *ArticleController) UpdateArticle() {
	var requestBody map[string]string
	json.Unmarshal(this.Ctx.Input.RequestBody, &requestBody)
	title := requestBody["title"]
	id := requestBody["id"]
	content := requestBody["content"]
	article := models.UpdateOne(id, title, content)
	this.Data["json"] = map[string]interface{}{"status": 200, "msg": article}
	this.ServeJSON()
	return
}

//根据ID删除文章
func (this *ArticleController) DeleteArticle() {
	var requestBody map[string]string
	json.Unmarshal(this.Ctx.Input.RequestBody, &requestBody)
	id := requestBody["id"]
	article := models.DeleteOne(id)
	this.Data["json"] = map[string]interface{}{"status": 200, "msg": article, "tip": "删除成功"}
	this.ServeJSON()
	return
}
