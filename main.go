package main

import (
	"fmt"
	_ "gameapp/repository/mysql"
	"gameapp/route"
	_ "gameapp/service"
	"net/http"
)

func main() {

	route.Route()

	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		fmt.Println(err)

		return
	}
}
