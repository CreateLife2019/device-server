package http

import (
	"github.com/device-server/domain/base"
)

type AccountData struct {
	AccountId int64  `json:"accountId"`
	Token     string `json:"token"`
}
type LoginResponse struct {
	base.BaseResponse
	Data AccountData `json:"data"`
}
type LogoutResponse struct {
	base.BaseResponse
}
