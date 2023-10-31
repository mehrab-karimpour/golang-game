package response

import (
	"gameapp/pkg/richerror"
	"net/http"
)

type HttpResponse struct {
	Message string `json:"message"`
	Status  `json:"status"`
	Data    any `json:"data"`
}
type Messages map[string]string
type Response map[string]any
type Status bool

func (res HttpResponse) MapKindErrorToHttpCodeError(kind richerror.KindType) int {
	var httpErrors = make(map[richerror.KindType]int)
	httpErrors[richerror.KindUnexpected] = http.StatusInternalServerError
	httpErrors[richerror.KindNotFound] = http.StatusNotFound
	httpErrors[richerror.KindForbidden] = http.StatusForbidden
	httpErrors[richerror.KindUnauthorized] = http.StatusUnauthorized
	httpErrors[richerror.KindUnprocessable] = http.StatusUnprocessableEntity
	return httpErrors[kind]
}
