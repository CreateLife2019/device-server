package http

import "github.com/device-server/domain/base"

type LoginLogRequest struct {
	base.PageInfo
	Account string `form:"account,omitempty" json:"account"`
}
