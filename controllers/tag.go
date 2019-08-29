package controllers

import (
	"go_demo/models"

	"github.com/astaxie/beego"
)

type TagController struct {
	beego.Controller
}

func (this *TagController) GetTag() {
	token := this.Ctx.Input.Header("token")
	response := make(map[string]interface{})
	if len(token) == 0 {
		response["msg"] = "没有权限"
	}
	content, err := models.GetTag()
	if err != nil {
		response["msg"] = err
	}
	response["msg"] = content
	response["code"] = "200"
	this.Data["json"] = response
	this.ServeJSON()
}
