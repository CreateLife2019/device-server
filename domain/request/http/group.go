package http

import "github.com/device-server/domain/base"

type GroupInfo struct {
	Name string `json:"name" form:"name" binding:"required"`
}
type CreateGroupRequest struct {
	GroupInfo
}
type UpdateGroupRequest struct {
	GroupInfo
	GroupId int64 `json:"-"`
}
type DeleteGroupRequest struct {
	GroupId int64 `form:"groupId" json:"-" binding:"required"`
}

type GroupListRequest struct {
	base.PageInfo
	Name string `form:"name,omitempty" json:"name"`
}
