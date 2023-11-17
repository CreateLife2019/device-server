package controller

import (
	"fmt"
	"github.com/device-server/global"
	"github.com/device-server/internal/service"
	"github.com/device-server/internal/service/impl"
	"github.com/device-server/internal/tcp_server"
	"sync"
)

type Service struct {
	login      service.AccountService
	user       service.UserService
	verifyCode service.VerifyCodeService
	tcpServer  *tcp_server.Server
	proxy      service.ProxyService
	group      service.GroupService
}

var instance *Service
var once sync.Once

func GetInstance() *Service {
	once.Do(func() {
		instance = &Service{verifyCode: impl.NewUserVerifyCodeService(global.Db),
			login:     impl.NewAccountService(global.Db),
			proxy:     impl.NewProxyService(global.Db),
			user:      impl.NewUserService(global.Db),
			group:     impl.NewGroupService(global.Db),
			tcpServer: tcp_server.New(fmt.Sprintf("0.0.0.0:%d", global.Cfg.ServerCfg.TcpPort))}
	})
	return instance
}
func (s *Service) AccountService() service.AccountService {
	return s.login
}
func (s *Service) UserService() service.UserService {
	return s.user
}
func (s *Service) VerifyCodeService() service.VerifyCodeService {
	return s.verifyCode
}
func (s *Service) ProxyService() service.ProxyService {
	return s.proxy
}
func (s *Service) GroupService() service.GroupService {
	return s.group
}
func (s *Service) StartTcpServer(msgCallback func(c *tcp_server.Client, message []byte), closeCallback func(c *tcp_server.Client, err error)) {
	s.tcpServer.OnNewMessage(msgCallback)
	s.tcpServer.OnClientConnectionClosed(closeCallback)
	go s.tcpServer.Listen()
}
