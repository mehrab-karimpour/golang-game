package main

import (
	"gameapp/http/handler"
	"gameapp/repository"
	_ "gameapp/repository/mysql"
	_ "gameapp/route/authroute"
	_ "gameapp/route/otherroute"
	"gameapp/service"
	_ "gameapp/service"
	"github.com/labstack/echo/v4"
)

var userService service.User
var userAuthService service.Auth

func init() {
	userService.Repo = repository.DataBaseService()
	userService.UserAuth = userAuthService.New()
}

func main() {

	e := echo.New()

	e.POST("/auth/login", handler.Login)
	e.POST("/auth/register", handler.Register)

	_ = e.Start(":8000")
}
