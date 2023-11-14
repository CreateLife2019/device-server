package http

import (
	"github.com/device-server/domain/base"
)

type VerifyCodeData struct {
	RequestId  int64  `json:"requestId,string"`
	VerifyCode string `json:"verifyCode"`
}
type VerifyCodeResponse struct {
	base.BaseResponse
	Data VerifyCodeData `json:"data"`
}
