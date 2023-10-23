package route

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

type HttpRoute map[string]func(http.ResponseWriter, *http.Request)

type HttpRouteHandler struct {
	Route HttpRoute
}

var HttpHandlers HttpRouteHandler

func init() {
	routes := HttpRoute{}
	HttpHandlers = HttpRouteHandler{
		Route: routes,
	}
}

func (httpRouter HttpRouteHandler) Push(routes HttpRoute) {
	for index, method := range routes {
		httpRouter.Route[index] = method
	}
}

func Route() {
	for fullPath, handler := range HttpHandlers.Route {
		fullPathSlice := strings.Split(fullPath, "@")
		method := fullPathSlice[0]
		path := fullPathSlice[1]
		switch method {
		case http.MethodPost:
			http.Handle(path, HttpHandlers.post(handler))
			break
		case http.MethodGet:
			http.Handle(path, HttpHandlers.get(handler))
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
