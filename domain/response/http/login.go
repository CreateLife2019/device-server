package http

import (
	"github.com/device-server/domain/base"
)

type AccountData struct {
	AccountId int64 `json:"accountId"`
}
type LoginResponse struct {
	base.BaseResponse
	Data  AccountData `json:"data"`
	Token string      `json:"token"`
}
type LogoutResponse struct {
	base.BaseResponse
}
