package impl

import (
	"fmt"
	"github.com/device-server/domain/constants"
	"github.com/device-server/domain/request/http"
	http2 "github.com/device-server/domain/response/http"
	"github.com/device-server/internal/repository/entity"
	"github.com/device-server/internal/repository/filter"
	"github.com/device-server/internal/repository/persistence"
	"github.com/device-server/internal/repository/persistence/impl"
	"gorm.io/gorm"
)

type GroupServiceImpl struct {
	db    *gorm.DB
	group persistence.GroupIer
	user  persistence.UserIer
}

func NewGroupService(db *gorm.DB) *GroupServiceImpl {
	s := &GroupServiceImpl{db: db, group: &impl.GroupImpl{}, user: &impl.UserIerImpl{}}
	return s
}

func (p *GroupServiceImpl) CreateGroup(request http.CreateGroupRequest) (resp http2.CreateGroupResponse, err error) {
	var group *entity.Group
	group, err = p.group.Get(p.db, filter.WithName(request.Name))
	if err == nil {
		resp.Code = constants.Status500
		resp.Msg = constants.MessageFailedConflictGroup
		return
	} else {
		group, err = p.group.Save(p.db, &entity.Group{Name: request.Name})
		if err != nil {
			resp.Code = constants.Status500
			resp.Msg = constants.MessageFailedServer
			return
		}
		resp.Data.GroupId = group.Id
	}
	resp.Code = constants.Status200
	resp.Msg = constants.MessageSuc
	return
}
func (p *GroupServiceImpl) UpdateGroup(request http.UpdateGroupRequest) (resp http2.UpdateGroupResponse, err error) {
	var group *entity.Group
	group, err = p.group.Get(p.db, filter.WithName(request.Name))
	if err == nil {
		err = fmt.Errorf(constants.MessageFailedConflictGroup)
		resp.Code = constants.Status500
		resp.Msg = err.Error()
		return
	}

	group, err = p.group.Get(p.db, filter.WithId(request.GroupId))
	if err != nil {
		resp.Code = constants.Status500
		resp.Msg = constants.MessageFailedServer
		return
	}
	group.Name = request.Name
	err = p.group.Update(p.db, group, filter.WithId(group.Id))
	if err != nil {
		resp.Code = constants.Status500
		resp.Msg = err.Error()
		return
	}
	resp.Code = constants.Status200
	resp.Msg = constants.MessageSuc
	return
}
func (p *GroupServiceImpl) GroupList(request http.GroupListRequest) (resp http2.GroupInfoListResponse, err error) {
	groups := make([]*entity.Group, 0)
	page := &entity.Page{
		Page:     request.Page,
		PageSize: request.PageSize,
	}
	resp.Data.Page = page.Page
	resp.Data.PageSize = page.PageSize
	resp.Data.Groups = make([]http2.GroupInfo, 0)
	groups, err = p.group.SearchGroup(p.db, page, filter.WithLickName(request.Name))
	if err != nil {
		resp.Code = constants.Status500
		resp.Msg = err.Error()
		return
	}
	resp.Code = "200"
	resp.Msg = constants.MessageSuc
	resp.Data.Total = page.Total
	resp.Data.Page = page.Page
	resp.Data.PageSize = page.PageSize
	for _, v := range groups {
		resp.Data.Groups = append(resp.Data.Groups, http2.GroupInfo{
			GroupId: v.Id,
			Name:    v.Name,
		})
	}
	return
}
func (p *GroupServiceImpl) DeleteGroup(groupId int64) (resp http2.DeleteGroupResponse, err error) {
	group := &entity.Group{
		Base: entity.Base{Id: groupId},
	}
	var userGroups []*entity.UserGroup
	userGroups, err = p.user.SearchUserGroup(p.db, nil, filter.WithGroupId([]int64{groupId}))
	if err != nil {
		return
	}
	if len(userGroups) > 0 {
		err = fmt.Errorf(constants.MessageFailedNotAllowedDeleteGroup)
		resp.Code = constants.Status500
		resp.Msg = constants.MessageFailedNotFound
		return
	}
	err = p.group.Delete(p.db, group)
	if err != nil {
		resp.Code = constants.Status500
		resp.Msg = constants.MessageFailedNotFound
		return
	}
	resp.Code = constants.Status200
	resp.Msg = constants.MessageSuc
	return
}
func (p *GroupServiceImpl) GroupDetail(groupId int64) (resp http2.GroupDetailResponse, err error) {
	var group *entity.Group
	group, err = p.group.Get(p.db, filter.WithId(groupId))
	if err != nil {
		resp.Code = constants.Status400
		resp.Msg = constants.MessageFailedNoProxy
		return
	}
	resp.Data.GroupId = group.Id
	resp.Data.Name = group.Name
	resp.Code = constants.Status200
	resp.Msg = constants.MessageSuc
	return
}
