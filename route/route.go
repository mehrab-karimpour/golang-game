package route

import (
	"fmt"
	"golang.org/x/exp/maps"
	"net/http"
)

type HttpRouteHandler struct {
	route map[string]func(http.ResponseWriter, *http.Request)
}

func Route() {
	var httpRouteHandler HttpRouteHandler
	authRoutes := GetAuthRoutes()
	otherRoutes := GetOtherRoutes()
	httpRouteHandler.route = authRoutes
	maps.Copy(httpRouteHandler.route, otherRoutes)

	for path, method := range httpRouteHandler.route {
		fmt.Println(path)
		http.HandleFunc(path, method)
	}
}
