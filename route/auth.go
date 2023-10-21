package route

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"gameapp/repository"
	"gameapp/request"
	"gameapp/request/userrequest"
	"gameapp/service"
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
	var reqData = make([]byte, 1024)
	var validationErrors *request.ValidationErrors

	n, _ := req.Body.Read(reqData)
	var registerReq userrequest.RegisterRequest
	err := json.Unmarshal(reqData[:n], &registerReq)
	if err != nil {
		fmt.Println(err)
	}
	row, err := userService.Repo.FirstWhere("phone_number", registerReq.PhoneNumber)

	if err != sql.ErrNoRows && row.ID != 0 {
		message := fmt.Sprint("this mobile already registered!")
		validationErrors = registerReq.Validate(&message)
		handleValidation(res, validationErrors)
		return
	}

	if validationErrors = registerReq.Validate(nil); validationErrors != nil {
		fmt.Println(validationErrors)
		handleValidation(res, validationErrors)
		return
	}
	user, err := userService.Repo.Store(registerReq)

	if err != nil {
		fmt.Println(err)
		http.Error(res, err.Error(), http.StatusServiceUnavailable)
		return
	}

	userBytes, err := json.Marshal(user)
	if err != nil {
		return
	}

	fmt.Fprint(res, string(userBytes))
}

func handleValidation(res http.ResponseWriter, validationErrors *request.ValidationErrors) {
	marshal, _ := json.Marshal(validationErrors.Errors)
	http.Error(res, string(marshal), http.StatusUnprocessableEntity)
}
