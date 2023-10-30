package main

import (
	"gameapp/app/http/handler"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	auth := e.Group("/auth/")
	auth.POST("login", handler.Login)
	auth.POST("register", handler.Register)
	auth.GET("profile", handler.Profile)

	_ = e.Start(":8000")
}
