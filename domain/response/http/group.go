package http

import (
	"github.com/device-server/domain/base"
)

type GroupInfo struct {
	GroupId int64  `json:"groupId,omitempty"`
	Name    string `json:"name"`
}
type GroupInfoListData struct {
	base.PageInfo
	Groups []GroupInfo `json:"groups"`
}
type GroupDetailResponse struct {
	base.BaseResponse
	Data GroupInfo `json:"data"`
}
type GroupInfoListResponse struct {
	base.BaseResponse
	Data GroupInfoListData `json:"data"`
}
type CreateGroupData struct {
	GroupId int64 `json:"groupId,string"`
}
type CreateGroupResponse struct {
	base.BaseResponse
	Data CreateGroupData `json:"data"`
}
type UpdateGroupResponse struct {
	base.BaseResponse
}

type DeleteGroupResponse struct {
	base.BaseResponse
}
