package http

import "github.com/device-server/domain/base"

type CreateAccountRequest struct {
	Account         string `form:"account" json:"account" binding:"required"`
	Password        string `form:"password" json:"password" binding:"required"`
	ConfirmPassword string `form:"confirmPassword" json:"confirmPassword" binding:"required"`
}
type UpdateAccountRequest struct {
	AccountId   int64  `form:"account" json:"accountId" binding:"required"`
	OldPassword string `form:"oldPassword" json:"oldPassword" binding:"required"`
	NewPassword string `form:"newPassword" json:"newPassword" binding:"required"`
}
type DeleteAccountRequest struct {
	AccountId int64 `form:"account" json:"accountId" binding:"required"`
}

type AccountListRequest struct {
	base.PageInfo
}
