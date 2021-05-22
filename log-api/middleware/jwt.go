package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log-api/model"
	"time"
)

var SecretKey string = "helloworld"

type LoginClaims struct {
	StandardClaims jwt.StandardClaims
	User model.User
}

func (this LoginClaims) Valid() error{
	return this.StandardClaims.Valid()
}

func GenerateToken(user model.User, expireDuration time.Duration) (string, error) {
	expire := time.Now().Add(expireDuration)
	// 将 uid，用户角色， 过期时间作为数据写入 token 中
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, LoginClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
		},
	})
	// SecretKey 用于对用户数据进行签名，不能暴露
	return token.SignedString([]byte(SecretKey))
}

func VerifyToken(tokenStr string) (LoginClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &LoginClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return LoginClaims{}, errors.New("Error Unauthorized")
	}
	if claims, ok := token.Claims.(*LoginClaims); ok && token.Valid {
		return *claims, nil
	} else {
		return LoginClaims{}, errors.New("Error Unauthorized")
	}
}

