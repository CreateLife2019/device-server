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
	"time"
)

type AccountServiceImpl struct {
	db      *gorm.DB
	account persistence.AccountIer
	log     persistence.LogIer
}

func NewAccountService(db *gorm.DB) *AccountServiceImpl {
	return &AccountServiceImpl{db: db, account: &impl.AccountImpl{}, log: &impl.LogImpl{}}
}
func (a *AccountServiceImpl) Login(request http.LoginRequest) (resp http2.LoginResponse, err error) {
	var account *entity.Account
	account, err = a.account.Get(a.db, filter.WithAccountPassword(request.Account, request.Password))
	if err != nil {
		resp.Code = "400"
		resp.Msg = err.Error()
		return
	}
	account.LoginTime = time.Now()
	err = a.account.Update(a.db, account, filter.WithAccountId(account.Id))
	if err != nil {
		resp.Code = "400"
		resp.Msg = err.Error()
		return
	}
	resp.Code = "200"
	resp.Msg = constants.MessageSuc
	resp.Data.AccountId = account.Id
	_, err = a.log.Save(a.db, &entity.LoginLog{
		AccountId: account.Id,
		Account:   account.UserName,
		Ip:        request.ClientIp,
	})
	if err != nil {
		resp.Code = "400"
		resp.Msg = err.Error()
		return
	}
	return
}
func (a *AccountServiceImpl) CreateAccount(request http.CreateAccountRequest) (resp http2.CreateAccountResponse, err error) {
	if request.ConfirmPassword != request.Password {
		err = fmt.Errorf(constants.MessageFailedNotSamePassword)
		resp.Code = "400"
		resp.Msg = constants.MessageFailedNotSamePassword
		return
	}
	var account *entity.Account
	account, err = a.account.Get(a.db, filter.WithAccount(request.Account))
	if err == nil {
		err = fmt.Errorf(constants.MessageFailedConflictAccount)
		resp.Code = "500"
		resp.Msg = constants.MessageFailedConflictAccount
		return
	} else {
		account.Password = request.ConfirmPassword
		account.UserName = request.Account
		account.LoginTime = time.Now()
		account, err = a.account.Save(a.db, account)
		if err != nil {
			resp.Code = "500"
			resp.Msg = err.Error()
			return
		}
		resp.Code = "200"
		resp.Msg = constants.MessageSuc
		resp.Data.AccountId = account.Id
	}
	return
}
func (a *AccountServiceImpl) UpdateAccount(request http.UpdateAccountRequest) (resp http2.UpdateAccountResponse, err error) {
	if request.NewPassword == request.OldPassword {
		err = fmt.Errorf(constants.MessageFailedSamePassword)
		resp.Code = "400"
		resp.Msg = constants.MessageFailedSamePassword
		return
	}
	account := &entity.Account{}
	account, err = a.account.Get(a.db, filter.WithAccountId(request.AccountId))
	if err != nil {
		err = fmt.Errorf(constants.MessageFailedNotFound)
		resp.Code = "400"
		resp.Msg = constants.MessageFailedNotFound
		return
	}
	if account.Password != request.OldPassword {
		err = fmt.Errorf(constants.MessageFailedNotFound)
		resp.Code = "400"
		resp.Msg = constants.MessageFailedWrongPassword
		return
	}
	account.Password = request.NewPassword
	err = a.account.Update(a.db, account, filter.WithAccountId(request.AccountId))
	if err != nil {
		resp.Code = "500"
		resp.Msg = err.Error()
		return
	}
	resp.Code = "200"
	resp.Msg = constants.MessageSuc
	resp.Data.AccountId = account.Id
	return
}
func (a *AccountServiceImpl) AccountList(request http.AccountListRequest) (resp http2.AccountListResponse, err error) {
	accounts := make([]*entity.Account, 0)
	page := &entity.Page{
		Page:     request.Page,
		PageSize: request.PageSize,
	}
	resp.Data.Page = page.Page
	resp.Data.PageSize = page.PageSize
	accounts, err = a.account.SearchAccount(a.db, page)
	if err != nil {
		return
	}
	resp.Code = "200"
	resp.Msg = constants.MessageSuc
	resp.Data.Total = page.Total
	resp.Data.Page = page.Page
	resp.Data.PageSize = page.PageSize
	resp.Data.Accounts = make([]http2.AccountInfo, 0)
	for _, v := range accounts {
		resp.Data.Accounts = append(resp.Data.Accounts, http2.AccountInfo{
			Account:   v.UserName,
			LoginTime: v.LoginTime,
			CreatedAt: *v.CreatedAt,
			Id:        v.Id,
		})
	}
	return
}
func (a *AccountServiceImpl) DeleteAccount(accountId int64) (resp http2.DeleteAccountResponse, err error) {
	account := &entity.Account{
		Base: entity.Base{Id: accountId},
	}
	err = a.account.Delete(a.db, account)
	if err != nil {
		err = fmt.Errorf(constants.MessageFailedNotFound)
		resp.Code = "400"
		resp.Msg = constants.MessageFailedNotFound
		return
	}
	resp.Code = "200"
	resp.Msg = constants.MessageSuc
	return
}
func (a *AccountServiceImpl) CheckUser(userName, password string) bool {
	_, err := a.account.Get(a.db, filter.WithAccountPassword(userName, password))
	if err != nil {
		return false
	}
	return true
}
func (a *AccountServiceImpl) SearchLoginLog(request http.LoginLogRequest) (resp http2.LoginLogResponse, err error) {
	logs := make([]*entity.LoginLog, 0)
	page := &entity.Page{
		Page:     request.Page,
		PageSize: request.PageSize,
	}
	resp.Data.Page = page.Page
	resp.Data.PageSize = page.PageSize
	logs, err = a.log.SearchLog(a.db, page, filter.WithLickAccount(request.Account))
	if err != nil {
		return
	}
	resp.Code = "200"
	resp.Msg = constants.MessageSuc
	resp.Data.Total = page.Total
	resp.Data.Page = page.Page
	resp.Data.PageSize = page.PageSize
	resp.Data.LoginLogs = make([]http2.LoginLogInfo, 0)
	for _, v := range logs {
		resp.Data.LoginLogs = append(resp.Data.LoginLogs, http2.LoginLogInfo{
			Ip:        v.Ip,
			LoginTime: *v.CreatedAt,
			Account:   v.Account,
		})
	}
	return
}
