package route

import (
	"encoding/json"
	"fmt"
	"gameapp/request/userrequest"
	"net/http"
)

var authRoute HttpRouteHandler

func init() {
	authRoute.route = map[string]func(http.ResponseWriter, *http.Request){
		"POST@/auth/register": registerHandler,
	}
}
func GetAuthRoutes() map[string]func(http.ResponseWriter, *http.Request) {
	return authRoute.route
}

func registerHandler(res http.ResponseWriter, req *http.Request) {
	var reqData = make([]byte, 1024)
	n, _ := req.Body.Read(reqData)
	var registerReq userrequest.RegisterRequest
	err := json.Unmarshal(reqData[:n], &registerReq)
	if err != nil {
		fmt.Println(err)
	}

	// 1.6.0

	http.Error(res, "implement register handler ...."+string(reqData[:n]), http.StatusOK)
}
