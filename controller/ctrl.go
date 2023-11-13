package controller

import (
	"github.com/device-server/global"
	"github.com/device-server/internal/service"
	"github.com/device-server/internal/service/impl"
	"sync"
)

type Service struct {
	login service.LoginService
	user  service.UserService
}

var instance *Service
var once sync.Once

func GetInstance() *Service {
	once.Do(func() {
		instance = &Service{login: impl.NewLoginService(global.Db), user: impl.NewUserService(global.Db)}
	})
	return instance
}
func (s *Service) LoginService() service.LoginService {
	return s.login
}
func (s *Service) UserService() service.UserService {
	return s.user
}
