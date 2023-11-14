package http

import "github.com/device-server/domain/base"

type UserInfo struct {
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	NickName   string `json:"nickName"` // 用户名
	Avatar     string `json:"avatar"`   // 头像
	Remark     string `json:"remark"`   // 备注信息
	DeviceName string `json:"deviceName"`
	Ip         string `json:"ip"`
	Online     int    `json:"online"`
	Agent      int    `json:"agent"`
}

type UserListData struct {
	base.PageInfo
	Users []UserInfo `json:"Users"`
}
type UserListResponse struct {
	base.BaseResponse
	Data UserListData `json:"data"`
}
