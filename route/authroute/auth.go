package authroute

import (
	"gameapp/app/http/handler"
	"gameapp/route"
)

func init() {
	route.HttpHandlers.Push(route.HttpRoute{
		"GET@/auth/profile": handler.Profile,
	})
}
