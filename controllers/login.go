package controllers

import (
	"encoding/json"
	"fmt"
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
	//	response := make(map[string]interface{})
	var requestBody map[string]string
	json.Unmarshal(user.Ctx.Input.RequestBody, &requestBody)
	username := requestBody["username"]
	password := requestBody["password"]
	// beego.Info(pwd)
	// beego.Info(tel)

	if username == "" || password == "" {
		user.Data["json"] = map[string]interface{}{"status": 400, "msg": "账号或密码不可为空", "time": time.Now().Format("2006-12-12 12:12:12")}
		user.ServeJSON()
		return
	}

	//判断该手机号是否注册
	// user, err := models.GetUserByTel(tel)

	// if err != nil { //如果有错误
	// 	beego.Info(err)
	// }
	// if user != nil { //如果有该用户
	// 	if user.password == password {
	// 		user.password = ""
	// 		user.Data["json"] = map[string]interface{}{"status": 200, "user": user, "time": time.Now().Format("2006-01-02 15:04:05")}
	// 		user.ServeJSON()
	// 		return

	// 	}
	// }
	if username == "lpf" && password == "211314lpf" {
		user.Data["json"] = map[string]interface{}{"status": 200, "msg": "登录成功", "time": time.Now().Format("2006-01-02 15:04:05")}
	}
	user.ServeJSON()
	return
}

//注册
func (this *UserController) Register() {
	tel := this.Input().Get("Tel")
	Password := this.Input().Get("password")
	fmt.Println(tel, "a", Password)
	beego.Info(tel)
	if tel == "" { //如果手机号为空
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "账号或密码不为空！", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}
	//判断该手机号是否已经注册
	user, err := models.GetUserByTel(tel)
	if err != nil { //如果有错误
		beego.Info(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}
	if user != nil { //如果有该用户
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "该手机号已注册！", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

}
