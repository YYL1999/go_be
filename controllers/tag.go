package controllers

import (
	"encoding/json"
	"fmt"
	"go_demo/services"

	"github.com/astaxie/beego"
)

type TagController struct {
	beego.Controller
}

func (this *TagController) GetTag() {
	token := this.Ctx.Input.Header("token")
	response := make(map[string]interface{})
	fmt.Println(token)
	if len(token) == 0 {
		response["msg"] = "没有权限"
	}
	content, err := services.GetTag()
	if err != nil {
		response["msg"] = err
	}
	response["msg"] = content
	response["code"] = "200"
	this.Data["json"] = response
	this.ServeJSON()
	return
}
func (this *TagController) SetTag() {
	var requestBody map[string]string
	json.Unmarshal(this.Ctx.Input.RequestBody, &requestBody)
	response := make(map[string]interface{})
	content := requestBody["content"]
	link := requestBody["link"]
	token := this.Ctx.Input.Header("token")
	fmt.Println(token)
	if len(token) == 0 {
		response["msg"] = "没有权限"
	}
	ok := services.SetTag(content, link)

	response["msg"] = ok
	this.Data["json"] = response
	this.ServeJSON()

}
