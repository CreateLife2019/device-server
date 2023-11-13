package service

import (
	"github.com/device-server/domain/request"
	"github.com/device-server/domain/response"
)

type UserService interface {
	List(request request.UserListRequest) (resp response.UserListResponse, err error)
}
