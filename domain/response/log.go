package response

import "time"

type LoginLogInfo struct {
	Ip        string    `json:"Ip"`
	LoginTime time.Time `json:"LoginTime"`
}

type LoginLogResponse struct {
	BaseResponse
	PageInfo
	LoginLogs []LoginLogInfo `json:"LoginLogs"`
}
