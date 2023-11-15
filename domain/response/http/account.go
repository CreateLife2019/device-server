package http

import (
	"github.com/device-server/domain/base"
	"time"
)

type CreateAccountResponse struct {
	base.BaseResponse
	Data AccountData `json:"data"`
}
type UpdateAccountResponse struct {
	base.BaseResponse
	Data AccountData `json:"data"`
}
type DeleteAccountResponse struct {
	base.BaseResponse
}
type AccountInfo struct {
	Account   string    `json:"account"`
	LoginTime time.Time `json:"loginTime"`
	CreatedAt time.Time `json:"createdAt"`
	Id        int64     `json:"id"`
}
type AccountListData struct {
	base.PageInfo
	Accounts []AccountInfo `json:"accounts"`
}
type AccountListResponse struct {
	base.BaseResponse
	Data AccountListData `json:"data"`
}
