package authroute

import (
	"gameapp/http/handler"
	"gameapp/route"
)

func init() {
	route.HttpHandlers.Push(route.HttpRoute{
		"GET@/auth/profile": handler.Profile,
	})
}
