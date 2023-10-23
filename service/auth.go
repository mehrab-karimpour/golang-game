package service

import (
	"fmt"
	"gameapp/entity"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/golang-jwt/jwt/v5"
	"time"
)

type Auth struct{}

var secretKey = []byte("my-secret-key")

func (auth Auth) GenerateToken(user entity.User) *string {
	mapClaims := jwt.MapClaims{
		"username": user.ID,
		"exp":      time.Now().Add(time.Hour + 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &tokenString
}

func (auth Auth) TokenIsValid(tokenStr string) bool {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		fmt.Println("pars token error:", err)
	}

	return token.Valid
}

func (auth Auth) GetAuth(tokenStr string) *any {

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		fmt.Println("pars token error:", err)
		return nil
	}

	userId := token.Claims.(jwt.MapClaims)["username"]
	return &userId
}
