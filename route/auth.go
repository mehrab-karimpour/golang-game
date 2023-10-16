package route

import "net/http"

var authRoute HttpRouteHandler

func init() {
	authRoute.route = map[string]func(http.ResponseWriter, *http.Request){
		"/auth/register": registerHandler,
	}
}
func GetAuthRoutes() map[string]func(http.ResponseWriter, *http.Request) {
	return authRoute.route
}

func registerHandler(res http.ResponseWriter, req *http.Request) {
	http.Error(res, "implement register handler ....", http.StatusOK)
}
