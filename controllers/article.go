package controllers

import (
	"encoding/json"
	"fmt"
	"go_demo/models"
	"go_demo/services"
	"go_demo/utils"
	"strconv"

	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

//添加文章
func (this *ArticleController) AddArticle() {
	token := this.Ctx.Input.Header("token")
	ifToken := utils.CheckJwt(token)
	if !ifToken {
		return
	}
	response := make(map[string]interface{})
	var requestBody map[string]string
	json.Unmarshal(this.Ctx.Input.RequestBody, &requestBody)
	title := requestBody["title"]
	content := requestBody["content"]
	ok, id := services.AddArticle(content, title)
	if !ok {
		response["code"] = 400
		response["id"] = id
		response["msg"] = "文章新建错误"
	} else {
		response["code"] = 200
		response["id"] = id
		response["msg"] = "文章新建成功"
	}
	this.Data["json"] = response
	this.ServeJSON()
	return
}

//获取所有文章
func (this *ArticleController) GetAllArticle() {
	response := make(map[string]interface{})
	token := this.Ctx.Input.Header("token")
	fmt.Println(token)
	if len(token) == 0 {
		response["code"] = 401
		response["msg"] = "没有权限"

	} else {
		ok := utils.CheckJwt(token)
		if !ok {
			response["code"] = 401
			response["msg"] = "权限已过期，请重新登录"
		}
		article := services.GetAllArticle()
		response["code"] = 200
		response["msg"] = article
	}
	this.Data["json"] = response
	this.ServeJSON()
}

//根据ID获取单个文章
func (this *ArticleController) GetOneArticle() {
	id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
	article, err := models.GetOne(int64(id))
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
	id, err := strconv.Atoi(requestBody["id"])
	if err != nil {
		return
	}
	content := requestBody["content"]
	article := models.UpdateOne(int64(id), title, content)
	this.Data["json"] = map[string]interface{}{"status": 200, "msg": article}
	this.ServeJSON()
	return
}

//根据ID删除文章
func (this *ArticleController) DeleteArticle() {
	token := this.Ctx.Input.Header("token")
	ifToken := utils.CheckJwt(token)
	if !ifToken {
		this.Data["json"] = map[string]interface{}{"msg": token}
	}
	response := make(map[string]interface{})
	var requestBody map[string]string
	json.Unmarshal(this.Ctx.Input.RequestBody, &requestBody)
	idInt, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
	if err != nil {
		response["code"] = 200
		response["msg"] = "删除成功"
	}
	id := int64(idInt)
	article := models.DeleteOne(id)
	fmt.Println(article)
	response["code"] = 200
	response["msg"] = "删除成功"
	this.Data["json"] = response
	this.ServeJSON()
	return
}
