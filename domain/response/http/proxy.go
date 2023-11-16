package http

import (
	"github.com/device-server/domain/base"
	"time"
)

type ProxyInfo struct {
	ProxyId     int64     `json:"proxyId,omitempty"`
	ProxyHost   string    `json:"proxyHost"`
	ProxyPort   int       `json:"proxyPort"`
	ProxySecret string    `json:"proxySecret"`
	SetTime     time.Time `json:"setTime"`
}
type ProxyInfoListData struct {
	base.PageInfo
	Proxies []ProxyInfo `json:"proxies"`
}
type ProxyDetailResponse struct {
	base.BaseResponse
	Data ProxyInfo `json:"data"`
}
type ProxyInfoListResponse struct {
	base.BaseResponse
	Data ProxyInfoListData `json:"data"`
}
type CreateProxyData struct {
	ProxyId int64 `json:"proxyId"`
}
type CreateProxyResponse struct {
	base.BaseResponse
	Data CreateProxyData `json:"data"`
}
type UpdateProxyResponse struct {
	base.BaseResponse
}

type DeleteProxyResponse struct {
	base.BaseResponse
}
