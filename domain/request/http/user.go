package http

import "github.com/device-server/domain/base"

type UserListRequest struct {
	base.PageInfo
}

// 增加用户代理
type AgentRequest struct {
	UserId int64 `json:"UserId"`
}
