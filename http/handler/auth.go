package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"gameapp/http/request/userrequest"
	"gameapp/http/response"
	"gameapp/repository"
	"gameapp/service"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
)

var userService service.User

func init() {
	userService.Repo = repository.DataBaseService()
}

func LoginHandler(res http.ResponseWriter, req *http.Request) {
	var loginRequest userrequest.LoginRequest
	var err error

	bodyData, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = json.Unmarshal(bodyData, &loginRequest)

	if validationErrors := loginRequest.Validate(nil); validationErrors != nil {
		responsePrepared := response.Prepare(false, nil, validationErrors.Errors)
		http.Error(res, *responsePrepared, http.StatusUnprocessableEntity)
		return
	}

	failedMessage := "mobile or password is not correct"
	user, err := userService.Repo.FirstWhere("phone_number", loginRequest.PhoneNumber)

	userDataIsUnprocessable := err == sql.ErrNoRows || !comparePass(user.Password, loginRequest.Password) || user.PhoneNumber != loginRequest.PhoneNumber

	if userDataIsUnprocessable {
		validationErrors := loginRequest.Validate(&failedMessage)
		responsePrepared := response.Prepare(false, nil, validationErrors.Errors)
		http.Error(res, *responsePrepared, http.StatusUnprocessableEntity)
		return
	}

	var successMessage = response.Messages{
		"msg": "you logged in successfully...",
	}
	res.WriteHeader(http.StatusCreated)

	authToken := userService.Auth.GenerateToken(*user)
	responsePrepared := response.Prepare(true, *authToken, successMessage)
	_, err = fmt.Fprint(res, *responsePrepared)
	if err != nil {
		return
	}
}

func RegisterHandler(res http.ResponseWriter, req *http.Request) {
	messages := response.Messages{}
	reqData, readErr := io.ReadAll(req.Body)
	if readErr != nil {
		fmt.Println(readErr)
		return
	}

	var registerReq userrequest.RegisterRequest
	err := json.Unmarshal(reqData, &registerReq)
	if err != nil {
		fmt.Println("Unmarshal error registerHandler", err)
	}

	if validationErrors := registerReq.Validate(nil); validationErrors != nil {
		preparedResponse := response.Prepare(false, nil, validationErrors.Errors)
		http.Error(res, *preparedResponse, http.StatusUnprocessableEntity)
		return
	}

	row, err := userService.Repo.FirstWhere("phone_number", registerReq.PhoneNumber)

	if err != sql.ErrNoRows && row.ID != 0 {
		message := fmt.Sprint("this mobile already registered!")
		validationErrors := registerReq.Validate(&message)
		preparedResponse := response.Prepare(false, nil, validationErrors.Errors)
		http.Error(res, *preparedResponse, http.StatusUnprocessableEntity)
		return
	}

	registerReq.Password = bcryptPass(registerReq.Password)
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

func Profile(res http.ResponseWriter, req *http.Request) {

	user, err := userService.Repo.FirstWhere("phone_number ", "09180131109")
	if err != nil {
		fmt.Println(err)
	}

	res.WriteHeader(http.StatusOK)
	message := response.Messages{}
	message["status"] = "user profile "
	prepareResponse := response.Prepare(true, user, message)
	_, err = fmt.Fprint(res, *prepareResponse)
	if err != nil {
		return
	}

}

func bcryptPass(password string) string {
	var err error
	var passHashedByte []byte
	passByte := []byte(password)
	passHashedByte, err = bcrypt.GenerateFromPassword(passByte, 6)
	if err != nil {
		fmt.Println("bcryptPass error", err)
	}
	return string(passHashedByte)
}

func comparePass(hashedPass string, currentPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(currentPass))
	if err != nil {
		return false
	}
	return true
}
