package authroute

import (
	"gameapp/route"
)

func init() {
	route.HttpHandlers.Push(route.HttpRoute{})
}
