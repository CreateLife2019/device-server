package service

import (
	"github.com/device-server/domain/request/http"
	tcpRequest "github.com/device-server/domain/request/tcp"
	http2 "github.com/device-server/domain/response/http"
)

type UserService interface {
	List(request http.UserListRequest) (resp http2.UserListResponse, err error)
	Login(request tcpRequest.LoginRequest) (resp []byte)
}
