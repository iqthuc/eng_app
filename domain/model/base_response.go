package model

import (
	"encoding/json"
)

type BaseData[T any] struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    T      `json:"data"`
	Status  int    `json:"status"`
}

func (item BaseData[T]) ToResponseData() []byte {
	result, _ := json.Marshal(item)
	return result
}
