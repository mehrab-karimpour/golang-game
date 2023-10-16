package main

import (
	"fmt"
	"gameapp/repository"
	_ "gameapp/repository/mysql"
	"gameapp/request/userrequest"
	"gameapp/service"
	_ "gameapp/service"
)

func main() {

	myDb := repository.DataBaseService()

	var userService = service.User{
		Repo: myDb,
	}

	registerResponse, registerError := userService.Register(userrequest.RegisterRequest{
		FirstName:   "mehrab",
		LastName:    "karimpour",
		Password:    "1212",
		PhoneNumber: "09180131105",
	})
	if registerError != nil {
		fmt.Println("registerError", registerError)
		return
	}

	fmt.Println(registerResponse)

}
