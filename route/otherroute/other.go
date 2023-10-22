package otherroute

import (
	"gameapp/handler"
	"gameapp/route"
)

func init() {
	route.HttpHandlers.Push(route.HttpRoute{
		"GET@/other": handler.OtherHandler,
	})
}
