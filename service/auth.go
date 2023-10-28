package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/golang-jwt/jwt/v5"
	"time"
)

type Auth struct {
	signKey               []byte
	accessExpirationTime  int64
	refreshExpirationTime int64
}

func (auth Auth) Prepare() Auth {
	day := time.Hour * 24
	auth.signKey = []byte("my-secret-key")
	auth.accessExpirationTime = time.Now().Add(day * 7).Unix()
	auth.refreshExpirationTime = time.Now().Add(day * 30).Unix()
	return auth
}

func (auth Auth) GenerateAccessToken(ID int) *string {
	return auth.generateToken(ID, auth.accessExpirationTime, "accessToken")
}

func (auth Auth) GenerateRefreshToken(ID int) *string {
	return auth.generateToken(ID, auth.refreshExpirationTime, "refreshToken")
}

func (auth Auth) generateToken(ID int, expirationTime int64, subject string) *string {
	mapClaims := jwt.MapClaims{
		"username": ID,
		"exp":      expirationTime,
		"sub":      subject,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	tokenString, err := token.SignedString(auth.signKey)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &tokenString
}

func (auth Auth) TokenIsValid(tokenStr string) bool {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return auth.signKey, nil
	})

	if err != nil {
		fmt.Println("pars token error:", err)
		return false
	}

	return token.Valid
}

func (auth Auth) GetAuth(tokenStr string) *any {

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return auth.signKey, nil
	})
	if err != nil {
		fmt.Println("pars token error:", err)
		return nil
	}

	userId := token.Claims.(jwt.MapClaims)["username"]
	return &userId
}
