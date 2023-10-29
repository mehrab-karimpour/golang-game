package handler

import (
	"fmt"
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
	var req userrequest.LoginRequest
	var lRes userresponse.Login
	var res response.HttpResponse

	_ = c.Bind(&req)

	user, err := userService.Repo.FirstWhere("phone_number", req.PhoneNumber)
	if err != nil {
		return err
	}
	lRes.RefreshToken = *userService.UserAuth.GenerateRefreshToken(int(user.ID))
	lRes.AccessToken = *userService.UserAuth.GenerateAccessToken(int(user.ID))

	res.Message = "you logged in success fully"
	res.Status = true
	res.Data = lRes
	_ = c.JSON(http.StatusOK, res)
	return nil
}

func Register(c echo.Context) error {
	var req userrequest.RegisterRequest
	var res userresponse.Register
	var httpRes response.HttpResponse
	var err error

	_ = c.Bind(&req)
	req.Password = bcryptPass(req.Password)
	res.User, err = userService.Repo.Store(req)
	httpRes.Data = res
	err = c.JSON(http.StatusCreated, httpRes)

	return err
}

func Profile(res http.ResponseWriter, req *http.Request) {

	var auth any = 0
	var token = ""
	authHeader := req.Header.Get("Authorization")
	authHeaders := strings.Split(authHeader, "Bearer ")
	if len(authHeaders) > 1 {
		token = authHeaders[1]
	}

	if userService.UserAuth.TokenIsValid(token) {
		auth = userService.UserAuth.GetAuth(token)
	}

	user, err := userService.Repo.FirstWhere("id", auth)
	if err != nil {
		fmt.Println(err)
	}

	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	message := response.Messages{}
	message["status"] = "user profile "
	prepareResponse := response.Prepare(true, user, message)
	_, err = res.Write([]byte(*prepareResponse))
	if err != nil {
		fmt.Println(err)
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
