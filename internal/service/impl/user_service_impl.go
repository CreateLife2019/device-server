package impl

import (
	"github.com/device-server/domain/constants"
	"github.com/device-server/domain/request/http"
	tcpRequest "github.com/device-server/domain/request/tcp"
	http2 "github.com/device-server/domain/response/http"
	"github.com/device-server/domain/response/tcp"
	"github.com/device-server/internal/repository/entity"
	"github.com/device-server/internal/repository/filter"
	"github.com/device-server/internal/repository/persistence"
	"github.com/device-server/internal/repository/persistence/impl"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type UserServiceImpl struct {
	db   *gorm.DB
	user persistence.UserIer
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
	users, err = u.user.SearchUser(u.db, page)
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
		})
	}
	return
}
func (u *UserServiceImpl) Login(request tcpRequest.LoginRequest) (resp []byte) {
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

	if err != nil {
		return tcpResp.BuildFailed("500")
	} else {
		return tcpResp.BuildSuc()
	}
}
func (u *UserServiceImpl) Heartbeat(request tcpRequest.HeartbeatRequest) (resp []byte) {

	return
}
func (u *UserServiceImpl) checkHeartbeat() {
	go func() {
		//for {
		//
		//}
	}()
}
