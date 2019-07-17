package controllers

import (
	"encoding/json"
	"go_demo/models"
	"time"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

//登录
func (user *UserController) Login() {

	//从前端获取数据信息
	var requestBody map[string]string
	json.Unmarshal(user.Ctx.Input.RequestBody, &requestBody)
	username := requestBody["username"]
	password := requestBody["password"]

	if username == "" || password == "" {
		user.Data["json"] = map[string]interface{}{"status": 400, "msg": "账号或密码不可为空", "time": time.Now().Format("2006-12-12 12:12:12")}
		user.ServeJSON()
		return
	}
	users, err := models.GetUserByName(username)
	if err != nil {
		beego.Info(err)
	}
	if users != nil {
		if users.Password == password {
			user.Data["json"] = map[string]interface{}{"status": 200, "user": users, "time": time.Now().Format("2006-01-02 15:04:05")}
			user.ServeJSON()
			return
		}
	}

}

//注册
func (this *UserController) Register() {
	var requestBody map[string]string
	json.Unmarshal(this.Ctx.Input.RequestBody, &requestBody)
	tel := requestBody["tel"]
	Password := requestBody["password"]
	Name := requestBody["username"]
	ID := time.Now().Format("2006-01-02 15:04:05")

	if tel == "" || Password == "" || Name == "" { //如果手机号为空
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "账号或密码不为空！", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}
	//判断该手机号是否已经注册
	// user, err := models.GetUserByTel(tel)
	// if err != nil { //如果有错误
	// 	beego.Info(err)
	// 	this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
	// 	this.ServeJSON()
	// 	return
	// }
	// if user != nil { //如果有该用户
	// 	this.Data["json"] = map[string]interface{}{"status": 400, "msg": "该手机号已注册！", "time": time.Now().Format("2006-01-02 15:04:05")}
	// 	this.ServeJSON()
	// 	return
	// }
	id, errs := models.InsertUser(tel, Name, Password, ID)
	if errs == nil {
		this.Data["json"] = map[string]interface{}{"status": 200, "msg": "注册成功", "user": id}
		this.ServeJSON()
		return
	}

}
