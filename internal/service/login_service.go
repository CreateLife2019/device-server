package service

import (
	"github.com/device-server/domain/request"
	"github.com/device-server/domain/response"
)

type LoginService interface {
	Login(request request.LoginRequest) (resp response.LoginResponse, err error)
}
