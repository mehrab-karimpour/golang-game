package handler

import (
	"database/sql"
	"fmt"
	"gameapp/app/entity"
	"gameapp/app/http/request"
	"gameapp/app/http/request/userrequest"
	"gameapp/app/http/response"
	"gameapp/app/http/response/userresponse"
	"gameapp/app/repository"
	"gameapp/app/service"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

type AppContext struct {
	echo.Context
}

var userService service.User
var userAuthService service.Auth

func init() {
	userService.Repo = repository.DataBaseService()
	userService.UserAuth = userAuthService.New()
}

func Login(c echo.Context) error {
	var validationError *request.ValidationErrors
	var req userrequest.LoginRequest
	var lRes userresponse.Login
	var res response.HttpResponse
	var user *entity.User
	var err error
	var httpCode = http.StatusOK

	_ = c.Bind(&req)
	validationError = req.Validate(nil)
	if validationError != nil {
		res.Message = *validationError.Error
		httpCode = http.StatusUnprocessableEntity
		goto sendResponse
	}

	user, err = userService.Repo.FirstWhere("phone_number", req.PhoneNumber)

	if (err != nil && err == sql.ErrNoRows) || comparePass(user.Password, req.Password) {
		httpCode = http.StatusBadRequest
		res.Message = err.Error()
		goto sendResponse
	}
	lRes.RefreshToken = *userService.UserAuth.GenerateRefreshToken(int(user.ID))
	lRes.AccessToken = *userService.UserAuth.GenerateAccessToken(int(user.ID))

	res.Message = "you logged in successfully"
	res.Status = true
	res.Data = lRes

sendResponse:
	_ = c.JSON(httpCode, res)
	return nil
}

func Register(c echo.Context) error {
	var req userrequest.RegisterRequest
	var res userresponse.Register
	var httpRes response.HttpResponse
	var validationError *request.ValidationErrors
	var httpCode = http.StatusCreated
	var err error

	_ = c.Bind(&req)
	validationError = req.Validate(nil)
	if validationError != nil {
		httpRes.Message = *validationError.Error
		httpCode = http.StatusUnprocessableEntity
		goto response
	}

	req.Password = bcryptPass(req.Password)
	res.User, err = userService.Repo.Store(req)
	httpRes.Data = res
response:
	err = c.JSON(httpCode, httpRes)

	return err
}

func Profile(c echo.Context) error {
	var auth *any
	var user *entity.User
	var res response.HttpResponse
	var statusCode int
	var err error

	authorization := c.Request().Header.Get("Authorization")
	token := strings.Split(authorization, "Bearer ")
	tokenString := token[1]

	if userService.UserAuth.TokenIsValid(tokenString) {
		auth = userService.UserAuth.GetAuth(tokenString)
	} else {
		res.Status = false
		statusCode = http.StatusUnauthorized
		goto returnResponse
	}

	user, err = userService.Repo.FirstWhere("id", *auth)
	if err != nil {
		fmt.Println(err)
	}
	res.Data = user
	statusCode = http.StatusOK
returnResponse:
	err = c.JSON(statusCode, res)

	return err
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
