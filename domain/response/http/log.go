package http

import (
	"github.com/device-server/domain/base"
	"time"
)

type LoginLogInfo struct {
	Account   string    `json:"Account"`
	Ip        string    `json:"Ip"`
	LoginTime time.Time `json:"LoginTime"`
}
type LoginLogData struct {
	base.PageInfo
	LoginLogs []LoginLogInfo `json:"LoginLogs"`
}
type LoginLogResponse struct {
	base.BaseResponse
	Data LoginLogData `json:"data"`
}
