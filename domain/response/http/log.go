package http

import (
	"github.com/device-server/domain/base"
	"time"
)

type LoginLogInfo struct {
	Ip        string    `json:"Ip"`
	LoginTime time.Time `json:"LoginTime"`
}

type LoginLogResponse struct {
	base.BaseResponse
	base.PageInfo
	LoginLogs []LoginLogInfo `json:"LoginLogs"`
}
