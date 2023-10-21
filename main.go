package main

import (
	"fmt"
	"gameapp/handler"
	_ "gameapp/repository/mysql"
	_ "gameapp/service"
	"net/http"
)

func main() {

	handler.Route()

	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		fmt.Println(err)

		return
	}
}
