package services

import (
	"go_demo/models"

	"go_demo/utils"

	"github.com/astaxie/beego"
)

func Login(username string, password string) (string, bool) {
	users, err := models.GetUserByName(username)
	if err != nil {
		beego.Info(err)
	}
	if users != nil {
		if users.Password == password {
			tokenString, err := utils.SetJwt(username)
			if err != nil {
				beego.Error("jwt.SignedString:", err)
				return utils.FailedCode, false
			}
			return tokenString, true
		}
		errorMsg := "密码错误"
		return errorMsg, false
	}
	return utils.FailedMsg, false
}
