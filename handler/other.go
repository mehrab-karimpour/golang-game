package handler

import "net/http"

func OtherHandler(res http.ResponseWriter, req *http.Request) {
	http.Error(res, "implement other handler ....", http.StatusOK)

}
