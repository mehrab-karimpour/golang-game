package route

import (
	"fmt"
	"golang.org/x/exp/maps"
	"net/http"
	"strings"
	"time"
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

	for fullPath, handler := range httpRouteHandler.route {
		fullPathSlice := strings.Split(fullPath, "@")
		method := fullPathSlice[0]
		path := fullPathSlice[1]
		switch method {
		case http.MethodPost:
			http.Handle(path, httpRouteHandler.post(handler))
			break
		case http.MethodGet:
			http.Handle(path, httpRouteHandler.get(handler))
			break
		default:
			panic(fmt.Sprintf("the method %s not suported!", method))
		}
	}
}

func (httpRouter HttpRouteHandler) post(handler func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method, r.URL, " ", time.Now().Format("15:4:5"))
		if r.Method != http.MethodPost {
			http.Error(w, `{"message":"method not allowed!"}`, http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	})
}
func (httpRouter HttpRouteHandler) get(handler func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method, r.URL, " ", time.Now().Format("15:4:5"))
		if r.Method != http.MethodGet {
			http.Error(w, `{"message":"method not allowed!"}`, http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	})
}
