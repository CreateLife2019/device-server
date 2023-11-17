package service

import (
	"github.com/device-server/domain/request/http"
	http2 "github.com/device-server/domain/response/http"
)

type GroupService interface {
	CreateGroup(request http.CreateGroupRequest) (resp http2.CreateGroupResponse, err error)
	UpdateGroup(request http.UpdateGroupRequest) (resp http2.UpdateGroupResponse, err error)
	GroupList(request http.GroupListRequest) (resp http2.GroupInfoListResponse, err error)
	DeleteGroup(groupId int64) (resp http2.DeleteGroupResponse, err error)
	GroupDetail(groupId int64) (resp http2.GroupDetailResponse, err error)
}
