package controllers

import (
	"encoding/json"
	"go_demo/models"
	"time"

	"go_demo/services"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)

type UserController struct {
	beego.Controller
	BaseController
}

//登录
func (this *UserController) Post() {
	response := make(map[string]interface{})
	var requestBody map[string]string
	json.Unmarshal(this.Ctx.Input.RequestBody, &requestBody)
	username := requestBody["username"]
	password := requestBody["password"]
	if username == "" || password == "" {
		response["code"] = 201
		response["msg"] = "账号或密码不可为空"
		response["time"] = time.Now().Format("2006-01-02 15:04:05")
	}
	if tokenString, ok := services.Login(username, password); ok {
		response["code"] = 200
		response["msg"] = "登录成功"
		response["token"] = tokenString
		response["time"] = time.Now().Format("2006-01-02 15:04:05")
	} else {
		response["code"] = 201
		response["msg"] = tokenString
		response["time"] = time.Now().Format("2006-01-02 15:04:05")
	}
	this.Data["json"] = response
	this.ServeJSON()
}

//注册
func (this *UserController) Register() {
	var requestBody map[string]string
	json.Unmarshal(this.Ctx.Input.RequestBody, &requestBody)
	tel := requestBody["tel"]
	Password := requestBody["password"]
	Name := requestBody["username"]
	ID := 1

	if tel == "" || Password == "" || Name == "" { //如果手机号为空
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

	id, errs := models.InsertUser(tel, Name, Password, ID)
	if errs == nil {
		this.Data["json"] = map[string]interface{}{"status": 200, "msg": "注册成功", "user": id}
		this.ServeJSON()
		return
	}

}
func (this *UserController) GetAllUser() {
	token, e := this.ParseToken()

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {

		return
	}
	if e == nil {
		var user string = claims["username"].(string)
		this.Data["json"] = map[string]interface{}{"message": "xas", "user": user}
		this.ServeJSON()
		return
	}
}
