package main

import (
	"fmt"
	"gameapp/repository"
	_ "gameapp/repository/mysql"
	"gameapp/request/userrequest"
	"gameapp/service"
)

func main() {

	myDb := repository.DataBaseService()

	var userService = service.User{
		Repo: myDb,
	}

	registerResponse, err := userService.Register(userrequest.RegisterRequest{
		FirstName:   "mehrab",
		LastName:    "karimpour",
		Password:    "1212",
		PhoneNumber: "09180131109",
	})
	if err != nil {
		return
	}

	fmt.Println(registerResponse)

}
