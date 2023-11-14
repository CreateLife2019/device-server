package service

import (
	"github.com/device-server/domain/request/http"
	http2 "github.com/device-server/domain/response/http"
)

type LoginService interface {
	Login(request http.LoginRequest) (resp http2.LoginResponse, err error)
}
