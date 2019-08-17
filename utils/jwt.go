package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func SetJwt(name string) (token string, err error) {
	claims := make(jwt.MapClaims)
	claims["username"] = name
	claims["exp"] = time.Now().Add(time.Hour * 480).Unix() //20天有效期，过期需要重新登录获取token
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用自定义字符串加密
	tokenString, err := tokens.SignedString([]byte("mykey"))
	return tokenString, err
}
func CheckJwt(token string) bool {
	tokenString, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Errorf("Unexpected signing method: %v", token.Header["mykey"])
		}
		return token, nil
	})
	fmt.Println(err)
	// if err != nil {
	// 	return false
	// }

	if claims, ok := tokenString.Claims.(jwt.MapClaims); ok {
		nowTime := time.Now().Unix()
		lastTime := claims["exp"].(float64)
		if nowTime < int64(lastTime) {
			return true
		}
		return false
	} else {
		return false
	}
}
