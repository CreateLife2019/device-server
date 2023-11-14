package http

import (
	"github.com/device-server/domain/base"
)

type AdminIdData struct {
	Id int64 `json:"id"`
}
type LoginResponse struct {
	base.BaseResponse
	Data AdminIdData `json:"data"`
}
