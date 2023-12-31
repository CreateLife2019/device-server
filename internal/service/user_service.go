package service

import (
	"github.com/device-server/domain/request/http"
	tcpRequest "github.com/device-server/domain/request/tcp"
	http2 "github.com/device-server/domain/response/http"
	"github.com/device-server/domain/response/tcp"
	"github.com/device-server/internal/repository/entity"
)

type UserService interface {
	List(request http.UserListRequest) (resp http2.UserListResponse, err error)
	Login(request tcpRequest.LoginRequest) (resp tcp.TcpResponseProtocol, err error)
	Heartbeat(request tcpRequest.HeartbeatRequest) (resp tcp.TcpResponseProtocol, err error)
	Offline(request tcpRequest.OfflineRequest) (resp tcp.TcpResponseProtocol, err error)
	SetProxy(request http.ProxyRequest) (selectProxy []*entity.Proxy, resp http2.SetProxyResponse, err error)
	Get(userId int64) (user *entity.User, err error)
	ListUserConfig(request http.UserConfigListRequest) (resp http2.UserConfigInfoListResponse, err error)
	GetUserConfig(userId int64) (user *entity.UserConfig, err error)
	UpdateUserInfo(request http.UpdateUserInfoRequest) (resp http2.UpdateUserInfoResponse, err error)
	SetUserGroup(request http.SetGroupRequest) (resp http2.SetGroupResponse, err error)
	ListUserExtend(request http.UserListRequest) (resp http2.UserExtendListResponse, err error)
}
