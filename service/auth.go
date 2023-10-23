package service

import (
	"fmt"
	"gameapp/entity"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/golang-jwt/jwt/v5"
	"time"
)

type Auth struct{}

func (auth Auth) GenerateToken(user entity.User) *string {
	secretKey := []byte("my-secret-key")
	mapClaims := jwt.MapClaims{
		"username": user.ID,
		"exp":      time.Now().Add(time.Second + 2).Unix(),
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
	secretKey := []byte("my-secret-key")
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		fmt.Println("pars token error:", err)
	}

	return token.Valid
}

func (auth Auth) GetAuth(tokenStr string) *any {
	secretKey := []byte("my-secret-key")
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
