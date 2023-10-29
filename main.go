package main

import (
	"gameapp/app/http/handler"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	e.POST("/auth/login", handler.Login)
	e.POST("/auth/register", handler.Register)

	_ = e.Start(":8000")
}
