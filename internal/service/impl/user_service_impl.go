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
	"time"
)

type UserServiceImpl struct {
	db       *gorm.DB
	user     persistence.UserIer
	stopChan chan struct{}
}

func NewUserService(db *gorm.DB) *UserServiceImpl {

	s := &UserServiceImpl{db: db, user: &impl.UserIerImpl{}}
	s.checkHeartbeat()
	return s
}
func (u *UserServiceImpl) List(request http.UserListRequest) (resp http2.UserListResponse, err error) {
	users := make([]*entity.User, 0)
	page := &entity.Page{
		Page:     request.Page,
		PageSize: request.PageSize,
	}
	users, err = u.user.SearchUser(u.db, page, filter.WithId(request.UserId))
	if err != nil {
		return
	}
	resp.Data.Total = page.Total
	resp.Data.Page = page.Page
	resp.Data.PageSize = page.PageSize
	resp.Data.Users = make([]http2.UserInfo, 0)
	for _, v := range users {
		resp.Data.Users = append(resp.Data.Users, http2.UserInfo{
			Name:       v.Name,
			Phone:      v.Phone,
			NickName:   v.NickName,
			Avatar:     v.Avatar,
			Remark:     v.Remark,
			DeviceName: "",
			Ip:         "",
			Online:     0,
			Agent:      0,
			Id:         v.Id,
		})
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
func (u *UserServiceImpl) SetProxy(request http.ProxyRequest) (resp http2.SetProxyResponse, err error) {
	resp = http2.SetProxyResponse{BaseResponse: base.BaseResponse{Code: constants.Status200, Msg: constants.MessageSuc}}
	_, err = u.user.Get(u.db, filter.WithId(request.UserId))
	if err != nil {
		resp.Code = constants.Status500
		resp.Msg = constants.MessageFailedNotFound
		return
	}
	userConfig := &entity.UserConfig{}
	userConfig.ReadProxy(request)
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
				for _, v := range users {
					if time.Now().Sub(v.HeartbeatTime).Seconds() > float64(global.Cfg.ServerCfg.HeartbeatTime/1000000000) {
						logrus.Infof("未检测到心跳:%v", *v)
						//
					}
				}
			}
		}
	}()
}
