package response

import (
	"encoding/json"
	"fmt"
)

type HttpResponse struct {
	Message string `json:"message"`
	Status  `json:"status"`
	Data    any `json:"data"`
}
type Messages map[string]string
type Response map[string]any
type Status bool

func Prepare(status Status, data any, messages Messages) *string {
	var res = Response{}
	res["status"] = status
	res["data"] = data
	res["messages"] = messages
	bytes, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	response := string(bytes)

	return &response
}

func (m Messages) ToJson() string {
	messageByte, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}

	return string(messageByte)
}
