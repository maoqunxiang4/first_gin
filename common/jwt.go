package common

import (
	"github.com/dgrijalva/jwt-go"
	"goPro/model"
	"time"
)

//jwt:  github.com/dgrijalva/jwt-go

var jwtKey = []byte("a_secret_cret")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// 生成并发放token
func ReleaseToken(user model.User) (string, error) {
	//确定token失效时间
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := Claims{
		//ID要用 当前用户的id
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			//token失效时间
			ExpiresAt: expirationTime.Unix(),
			//确定发布时间
			IssuedAt: time.Now().Unix(),
			//发布人
			Issuer: "ocanlearn.tech",
			//确定主题
			Subject: "user token",
		},
	}

	//只能够使用SigningMethodHS256方法，使用秘钥生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// 解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, claims, err

}
