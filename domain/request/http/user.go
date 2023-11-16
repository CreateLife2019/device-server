package http

import "github.com/device-server/domain/base"

type UserListRequest struct {
	base.PageInfo
	UserId int64 `form:"userId,omitempty" json:"userId"`
}
type UserConfigListRequest struct {
	base.PageInfo
	UserId int64 `form:"userId,omitempty" json:"userId"`
}
type ProxyInfo struct {
	ProxyHost   string `json:"proxyHost"`
	ProxyPort   int    `json:"proxyPort"`
	ProxySecret string `json:"proxySecret"`
}

// 增加用户代理
type ProxyRequest struct {
	UserId      int64       `json:"userId"` // 用户列表返回，唯一
	Proxies     []ProxyInfo `json:"proxies"`
	Immediately bool        `json:"immediately"` //是否立即生效
}
