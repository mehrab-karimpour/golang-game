package otherroute

import (
	"gameapp/app/http/handler"
	"gameapp/route"
)

func init() {
	route.HttpHandlers.Push(route.HttpRoute{
		"GET@/other": handler.OtherHandler,
	})
}
