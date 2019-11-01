package controllers

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)

type BaseController struct {
	beego.Controller
}

// ParseToken parse JWT token in http header.
func (c *BaseController) ParseToken() (t *jwt.Token, e error) {
	var authString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjUxNDUzMDQsInVzZXJuYW1lIjoibHBmIn0.MN2sLfOAGB_odNuiVneKTqVG9drEJR-fHdQfR-sa6DU"
	//authString := c.Ctx.Input.Header("token")
	beego.Debug("AuthString:", authString)

	kv := strings.Split(authString, ".")
	if len(kv) != 3 {
		beego.Error("AuthString invalid:", kv)
		return nil, nil
	}
	// tokenString := kv[2]
	token, err := jwt.Parse(authString, func(token *jwt.Token) (interface{}, error) {
		return []byte("mykey"), nil
	})

	if token.Valid {
		fmt.Println("You look nice today")
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
		} else {
			fmt.Println("Couldn't handle this token:", err)
		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}

	// Parse token
	// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	// 	return []byte("mykey"), nil
	// })
	// if err != nil {
	// 	beego.Error("Parse token:", err)
	// 	if ve, ok := err.(*jwt.ValidationError); ok {
	// 		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
	// 			// That's not even a token
	// 			return nil, nil
	// 		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
	// 			// Token is either expired or not active yet
	// 			return nil, nil
	// 		} else {
	// 			// Couldn't handle this token
	// 			return nil, nil
	// 		}
	// 	} else {
	// 		// Couldn't handle this token
	// 		return nil, nil
	// 	}
	// }
	// if !token.Valid {
	// 	beego.Error("Token invalid:", tokenString)
	// 	return nil, nil
	// }
	beego.Debug("Token:", token)

	return token, nil
}
