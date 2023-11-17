package http

import (
	"github.com/device-server/domain/base"
)

type UserInfo struct {
	Id         int64  `json:"id,string"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	NickName   string `json:"nickName"` // 用户名
	Avatar     string `json:"avatar"`   // 头像
	Remark     string `json:"remark"`   // 备注信息
	DeviceName string `json:"deviceName"`
	Ip         string `json:"ip"`
	Online     int    `json:"online"`
	Agent      int    `json:"agent"`
	GroupId    int64  `json:"groupId"`
	GroupName  string `json:"groupName"`
	ProxyIp    string `json:"proxyIp"`
}

type UserListData struct {
	base.PageInfo
	Users []UserInfo `json:"Users"`
}
type UserListResponse struct {
	base.BaseResponse
	Data UserListData `json:"data"`
}

type SetProxyResponse struct {
	base.BaseResponse
}

type SendProxyResponse struct {
	base.BaseResponse
}

type UserConfigInfo struct {
	Id       int64  `json:"id,string"`
	UserId   int64  `json:"userId"`
	UserName string `json:"userName"`

	Proxies []ProxyInfo `json:"proxies"`
}
type UserConfigInfoListData struct {
	base.PageInfo
	UserConfigs []UserConfigInfo `json:"userConfigs"`
}
type UserConfigInfoListResponse struct {
	base.BaseResponse
	Data UserConfigInfoListData `json:"data"`
}
type UpdateUserInfoResponse struct {
	base.BaseResponse
}
type SetGroupResponse struct {
	base.BaseResponse
}
