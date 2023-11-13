package impl

import (
	"github.com/device-server/domain/request"
	"github.com/device-server/domain/response"
	"github.com/device-server/internal/repository/filter"
	"github.com/device-server/internal/repository/persistence"
	"github.com/device-server/internal/repository/persistence/impl"
	"gorm.io/gorm"
)

type LoginServiceImpl struct {
	db      *gorm.DB
	account persistence.AccountIer
}

func NewLoginService(db *gorm.DB) *LoginServiceImpl {
	return &LoginServiceImpl{db: db, account: &impl.AccountImpl{}}
}
func (l *LoginServiceImpl) Login(request request.LoginRequest) (resp response.LoginResponse, err error) {
	_, err = l.account.Get(l.db, filter.WithAccount(request.Account, request.Password))
	if err != nil {
		return
	}
	return
}
func (l *LoginServiceImpl) VerifyCode() (resp response.VerifyCodeResponse, err error) {
	return
}
