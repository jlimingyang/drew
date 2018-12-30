package utils

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

var (
	key []byte = []byte(beego.AppConfig.String("jwt_key"))
)

type MyCustomClaims struct {
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
	jwt.StandardClaims
}

// 产生json web token
func GenToken(userId int, userName string, exTime int64) string {
	claims := MyCustomClaims{
		strconv.Itoa(userId),
		userName,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix()),
			ExpiresAt: int64(time.Now().Unix() + exTime*1000),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)
	if err != nil {
		beego.Error(err)
		return ""
	}
	return ss
}

// 校验token是否有效
func CheckToken(tokenString string) (custom *MyCustomClaims) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims
	} else {
		fmt.Println(err)
		return nil
	}
}
