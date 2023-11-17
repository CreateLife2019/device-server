package impl

import (
	"github.com/device-server/domain/base"
	"github.com/device-server/domain/constants"
	"github.com/device-server/domain/request/http"
	tcpRequest "github.com/device-server/domain/request/tcp"
	http2 "github.com/device-server/domain/response/http"
	"github.com/device-server/domain/response/tcp"
	"github.com/device-server/global"
	"github.com/device-server/internal/repository/entity"
	"github.com/device-server/internal/repository/filter"
	"github.com/device-server/internal/repository/persistence"
	"github.com/device-server/internal/repository/persistence/impl"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strings"
	"time"
)

type UserServiceImpl struct {
	db       *gorm.DB
	user     persistence.UserIer
	proxy    persistence.ProxyIer
	group    persistence.GroupIer
	stopChan chan struct{}
}

func NewUserService(db *gorm.DB) *UserServiceImpl {

	s := &UserServiceImpl{db: db, user: &impl.UserIerImpl{}, proxy: &impl.ProxyImpl{}, group: &impl.GroupImpl{}}
	s.checkHeartbeat()
	return s
}
func (u *UserServiceImpl) List(request http.UserListRequest) (resp http2.UserListResponse, err error) {
	resp.Data.Users = make([]http2.UserInfo, 0)
	users := make([]*entity.User, 0)
	page := &entity.Page{
		Page:     request.Page,
		PageSize: request.PageSize,
	}
	resp.Data.Page = page.Page
	resp.Data.PageSize = page.PageSize
	userIds := make([]int64, 0)
	if request.UserId != 0 {
		userIds = append(userIds, request.UserId)
	}
	if request.GroupId != 0 {
		var userGroups []*entity.UserGroup
		userGroups, err = u.user.SearchUserGroup(u.db, page, filter.WithGroupId([]int64{request.GroupId}))
		if err != nil {
			return
		}
		if len(userGroups) == 0 {
			resp.Code = constants.Status200
			resp.Msg = constants.MessageSuc
			return
		}
		for _, v := range userGroups {
			userIds = append(userIds, v.UserId)
		}
	}
	users, err = u.user.SearchUser(u.db, page, filter.WithInId(userIds))
	if err != nil {
		return
	}
	resp.Data.Total = page.Total

	userIds = make([]int64, 0)
	for _, v := range users {
		userIds = append(userIds, v.Id)
	}
	userExtends, err := u.user.SearchUserExtend(u.db, nil, filter.WithInUserId(userIds))
	if err != nil {
		return
	}
	userExtendMap := map[int64]*entity.UserExtend{}
	for _, v := range userExtends {
		userExtendMap[v.UserId] = v
	}
	userGroups, err := u.user.SearchUserGroup(u.db, nil, filter.WithInUserId(userIds))
	if err != nil {
		return
	}
	userGroupMap := map[int64]*entity.UserGroup{}
	groupIds := make([]int64, 0)
	for _, v := range userGroups {
		userGroupMap[v.UserId] = v
		groupIds = append(groupIds, v.GroupId)
	}
	groups, err := u.group.SearchGroup(u.db, nil, filter.WithInId(groupIds))
	if err != nil {
		return
	}
	groupMap := map[int64]*entity.Group{}
	for _, v := range groups {
		groupMap[v.Id] = v
	}
	for _, v := range users {
		item := http2.UserInfo{
			Name:     v.Name,
			Phone:    v.Phone,
			NickName: v.NickName,
			Avatar:   v.Avatar,
			Remark:   v.Remark,
			Id:       v.Id,
		}
		if find, ok := userExtendMap[v.Id]; ok {
			item.Ip = find.ClientIp
			item.Online = find.Online
			if strings.ToLower(find.ProxyType) != "none" {
				item.Agent = 1
			}
			item.DeviceName = find.AppVersion
			item.ProxyIp = find.ProxyIp
		}
		if find, ok := userGroupMap[v.Id]; ok {
			item.GroupId = find.GroupId
			if findGroup, ok2 := groupMap[find.GroupId]; ok2 {
				item.GroupName = findGroup.Name
			}
		}
		resp.Data.Users = append(resp.Data.Users, item)

	}
	resp.Code = constants.Status200
	resp.Msg = constants.MessageSuc
	return
}
func (u *UserServiceImpl) Login(request tcpRequest.LoginRequest) (resp tcp.TcpResponseProtocol, err error) {
	user := &entity.User{
		Name:     request.UserName,
		Phone:    request.Phone,
		NickName: request.NickName,
	}
	tcpResp := tcp.LoginResponse{}
	loginTime, err := time.ParseInLocation(constants.TimeFormat, request.LoginTime, time.Local)
	if err != nil {
		logrus.Errorf("时间解析错误:%s,手机号:%s", request.LoginTime, request.Phone)
		loginTime = time.Now()
	}
	userExtend := entity.UserExtend{
		Base:          entity.Base{},
		Online:        1,
		ClientIp:      request.ClientIP,
		LoginTime:     loginTime,
		HeartbeatTime: loginTime,
		AppVersion:    request.AppVersion,
		ProxyType:     request.ProxyType,
	}

	err = u.db.Transaction(func(tx *gorm.DB) error {
		err = u.user.SaveOrUpdate(u.db, user)
		if err != nil {
			return err
		}
		user, err = u.user.Get(u.db, filter.WithPhone(user.Phone))
		if err != nil {
			return err
		}
		userExtend.UserId = user.Id
		err = u.user.SaveOrUpdateExtend(u.db, &userExtend)
		if err != nil {
			return err
		}
		tcpResp.UserId = userExtend.UserId
		return err
	})
	resp = &tcpResp
	return
}
func (u *UserServiceImpl) Heartbeat(request tcpRequest.HeartbeatRequest) (resp tcp.TcpResponseProtocol, err error) {
	tcpResp := tcp.HeartbeatResponse{}
	user, err := u.user.Get(u.db, filter.WithPhone(request.Phone))
	if err != nil {
		return
	} else {
		tcpResp.Phone = request.Phone
		err = u.user.UpdateUserExtend(u.db, &entity.UserExtend{
			HeartbeatTime: time.Now(),
		}, filter.WithUserId(user.Id))
		if err != nil {
			return
		}
	}
	resp = &tcpResp
	return
}
func (u *UserServiceImpl) Offline(request tcpRequest.OfflineRequest) (resp tcp.TcpResponseProtocol, err error) {
	tcpResp := tcp.OfflineResponse{}
	user, err := u.user.Get(u.db, filter.WithPhone(request.Phone))
	if err != nil {
		return
	} else {
		tcpResp.Phone = request.Phone
		err = u.user.UpdateUserExtend(u.db, &entity.UserExtend{
			Online: 2,
		}, filter.WithUserId(user.Id))
		if err != nil {
			return
		}
	}
	resp = &tcpResp
	return
}
func (u *UserServiceImpl) Get(userId int64) (user *entity.User, err error) {
	return u.user.Get(u.db, filter.WithId(userId))
}
func (u *UserServiceImpl) SetProxy(request http.ProxyRequest) (selectProxy []*entity.Proxy, resp http2.SetProxyResponse, err error) {
	resp = http2.SetProxyResponse{BaseResponse: base.BaseResponse{Code: constants.Status200, Msg: constants.MessageSuc}}
	_, err = u.user.Get(u.db, filter.WithId(request.UserId))
	if err != nil {
		resp.Code = constants.Status500
		resp.Msg = constants.MessageFailedNotFound
		return
	}
	var proxies []*entity.Proxy
	proxies, err = u.proxy.SearchProxy(u.db, nil)
	if err != nil {
		resp.Code = constants.Status500
		resp.Msg = constants.MessageFailedNoProxy
		return
	}
	if len(proxies) == 0 {
		resp.Code = constants.Status500
		resp.Msg = constants.MessageFailedNoProxy
		return
	}

	selectProxy = proxies
	userConfig := &entity.UserConfig{}
	for _, v := range proxies {
		v.Time = time.Now()
		userConfig.Proxies = append(userConfig.Proxies, *v)
	}
	userConfig.UserId = request.UserId
	_, err = u.user.GetOrCreateUserConfig(u.db, userConfig, filter.WithUserId(request.UserId))
	if err != nil {
		resp.Code = constants.Status500
		resp.Msg = constants.MessageFailedNotFound
		return
	}
	err = u.user.UpdateUserConfig(u.db, userConfig, filter.WithUserId(request.UserId))
	if err != nil {
		resp.Code = constants.Status500
		resp.Msg = constants.MessageFailedNotFound
		return
	}
	resp.Code = constants.Status200
	resp.Msg = constants.MessageSuc
	return
}

func (u *UserServiceImpl) ListUserConfig(request http.UserConfigListRequest) (resp http2.UserConfigInfoListResponse, err error) {
	configs := make([]*entity.UserConfig, 0)
	page := &entity.Page{
		Page:     request.Page,
		PageSize: request.PageSize,
	}
	configs, err = u.user.SearchUserConfig(u.db, page)
	if err != nil {
		resp.Code = constants.Status500
		resp.Msg = constants.MessageFailedNotFound
		return
	}
	resp.Data.Total = page.Total
	resp.Data.Page = page.Page
	resp.Data.PageSize = page.PageSize
	resp.Data.UserConfigs = make([]http2.UserConfigInfo, 0)
	userIds := make([]int64, 0)
	for _, v := range configs {
		userIds = append(userIds, v.UserId)
	}
	var users []*entity.User
	users, err = u.user.SearchUser(u.db, nil, filter.WithInId(userIds))
	if err != nil {
		resp.Code = constants.Status500
		resp.Msg = constants.MessageFailedNotFound
		return
	}
	userMap := map[int64]*entity.User{}
	for _, v := range users {
		userMap[v.Id] = v
	}
	for _, v := range configs {
		item := v.ToResponse()
		item.UserName = userMap[v.UserId].Name
		resp.Data.UserConfigs = append(resp.Data.UserConfigs, item)
	}
	resp.Code = constants.Status200
	resp.Msg = constants.MessageSuc
	return
}
func (u *UserServiceImpl) GetUserConfig(userId int64) (user *entity.UserConfig, err error) {
	return u.user.GetUserConfig(u.db, filter.WithUserId(userId))
}
func (u *UserServiceImpl) UpdateUserInfo(request http.UpdateUserInfoRequest) (resp http2.UpdateUserInfoResponse, err error) {
	err = u.user.Update(u.db, &entity.User{Remark: request.Remark}, filter.WithId(request.UserId))
	if err != nil {
		resp.Code = constants.Status500
		resp.Msg = err.Error()
		return
	}
	resp.Code = constants.Status200
	resp.Msg = constants.MessageSuc
	return
}

func (u *UserServiceImpl) SetUserGroup(request http.SetGroupRequest) (resp http2.SetGroupResponse, err error) {
	// 一个用户只能一个分组
	_, err = u.group.Get(u.db, filter.WithId(request.GroupId))
	if err != nil {
		resp.Code = constants.Status500
		resp.Msg = constants.MessageFailedGroupNotFound
		return
	}
	err = u.user.UpdateUserGroup(u.db, &entity.UserGroup{
		UserId:  request.UserId,
		GroupId: request.GroupId,
	}, filter.WithUserId(request.UserId))
	if err != nil {
		resp.Code = constants.Status500
		resp.Msg = err.Error()
		return
	}
	resp.Code = constants.Status200
	resp.Msg = constants.MessageSuc
	return
}
func (u *UserServiceImpl) checkHeartbeat() {
	go func() {

		ticker := time.NewTicker(time.Duration(global.Cfg.ServerCfg.HeartbeatTime))
		defer func() {
			ticker.Stop()
		}()
		for {
			select {
			case <-u.stopChan:
				return
			case <-ticker.C:
				logrus.Infof("检测客户端上线下线")
				users, err := u.user.SearchUserExtend(u.db, nil, filter.WithOnline())
				if err != nil {
					continue
				}
				ids := make([]int64, 0)
				for _, v := range users {
					if time.Now().Sub(v.HeartbeatTime).Seconds() > float64(global.Cfg.ServerCfg.HeartbeatTime/1000000000) {
						logrus.Infof("未检测到心跳:%v", *v)
						//
						ids = append(ids, v.UserId)
					}
				}
				if len(ids) != 0 {
					err = u.user.OfflineUsers(u.db, ids)
					if err != nil {
						logrus.Errorf("更新用户状态失败:%v", ids)
						continue
					}
				}

			}
		}
	}()
}
