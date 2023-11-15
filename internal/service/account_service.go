package service

import (
	"github.com/device-server/domain/request/http"
	http2 "github.com/device-server/domain/response/http"
)

type AccountService interface {
	Login(request http.LoginRequest) (resp http2.LoginResponse, err error)
	CreateAccount(request http.CreateAccountRequest) (resp http2.CreateAccountResponse, err error)
	UpdateAccount(request http.UpdateAccountRequest) (resp http2.UpdateAccountResponse, err error)
	AccountList(request http.AccountListRequest) (resp http2.AccountListResponse, err error)
	DeleteAccount(accountId int64) (resp http2.DeleteAccountResponse, err error)
}
