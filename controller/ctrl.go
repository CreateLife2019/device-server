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
}

var instance *Service
var once sync.Once

func GetInstance() *Service {
	once.Do(func() {
		instance = &Service{verifyCode: impl.NewUserVerifyCodeService(global.Db), login: impl.NewAccountService(global.Db), user: impl.NewUserService(global.Db), tcpServer: tcp_server.New(fmt.Sprintf("0.0.0.0:%d", global.Cfg.ServerCfg.TcpPort))}
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
func (s *Service) StartTcpServer(callback func(c *tcp_server.Client, message []byte)) {
	s.tcpServer.OnNewMessage(callback)
	go s.tcpServer.Listen()
}
