package handler

import "net/http"

var otherRoute HttpRouteHandler

func init() {
	otherRoute.route = map[string]func(http.ResponseWriter, *http.Request){
		"POST@/other": otherHandler,
	}
}
func GetOtherRoutes() map[string]func(http.ResponseWriter, *http.Request) {
	return otherRoute.route
}

func otherHandler(res http.ResponseWriter, req *http.Request) {
	http.Error(res, "implement other handler ....", http.StatusOK)

}
