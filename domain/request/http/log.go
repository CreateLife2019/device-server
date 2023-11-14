package http

import "github.com/device-server/domain/base"

type LoginLogRequest struct {
	base.PageInfo
	UserId int64 `form:"UserId"`
}
