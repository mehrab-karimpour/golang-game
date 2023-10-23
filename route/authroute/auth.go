package authroute

import (
	"gameapp/handler"
	"gameapp/route"
)

func init() {
	route.HttpHandlers.Push(route.HttpRoute{
		"POST@/auth/register": handler.RegisterHandler,
		"POST@/auth/login":    handler.LoginHandler,
		"GET@/auth/profile":   handler.Profile,
	})
}
