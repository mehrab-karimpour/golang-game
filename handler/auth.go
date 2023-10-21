package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"gameapp/repository"
	"gameapp/request/userrequest"
	"gameapp/response"
	"gameapp/service"
	"io"
	"net/http"
)

var authRoute HttpRouteHandler
var userService service.User

func init() {
	userRepo := repository.DataBaseService()
	userService.Repo = userRepo

	authRoute.route = map[string]func(http.ResponseWriter, *http.Request){
		"POST@/auth/register": registerHandler,
	}
}
func GetAuthRoutes() map[string]func(http.ResponseWriter, *http.Request) {
	return authRoute.route
}

func registerHandler(res http.ResponseWriter, req *http.Request) {
	messages := response.Messages{}
	reqData, readErr := io.ReadAll(req.Body)
	if readErr != nil {
		fmt.Println(readErr)
		return
	}

	var registerReq userrequest.RegisterRequest
	err := json.Unmarshal(reqData, &registerReq)
	if err != nil {
		fmt.Println(err)
	}
	row, err := userService.Repo.FirstWhere("phone_number", registerReq.PhoneNumber)

	if err != sql.ErrNoRows && row.ID != 0 {
		message := fmt.Sprint("this mobile already registered!")
		validationErrors := registerReq.Validate(&message)
		preparedResponse := response.Prepare(false, nil, validationErrors.Errors)
		http.Error(res, *preparedResponse, http.StatusUnprocessableEntity)
		return
	}

	if validationErrors := registerReq.Validate(nil); validationErrors != nil {
		preparedResponse := response.Prepare(false, nil, validationErrors.Errors)
		http.Error(res, *preparedResponse, http.StatusUnprocessableEntity)
		return
	}
	user, err := userService.Repo.Store(registerReq)

	if err != nil {
		messages["storeError"] = err.Error()
		preparedResponse := response.Prepare(false, nil, messages)
		http.Error(res, *preparedResponse, http.StatusServiceUnavailable)
		return
	}

	preparedResponse := response.Prepare(false, user, messages)
	_, err = fmt.Fprint(res, *preparedResponse)
	if err != nil {
		return
	}
}
