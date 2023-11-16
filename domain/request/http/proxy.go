package http

import "github.com/device-server/domain/base"

type ProxyInfo struct {
	ProxyHost   string `json:"proxyHost" form:"proxyHost" binding:"required"`
	ProxyPort   int    `json:"proxyPort" form:"ProxyPort" binding:"required"`
	ProxySecret string `json:"proxySecret" form:"ProxySecret" binding:"required"`
}
type CreateProxyRequest struct {
	ProxyInfo
}
type UpdateProxyRequest struct {
	ProxyInfo
	ProxyId int64 `json:"-"`
}
type DeleteProxyRequest struct {
	ProxyId int64 `form:"proxyId" json:"-" binding:"required"`
}

type ProxyListRequest struct {
	base.PageInfo
}
