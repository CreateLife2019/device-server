package persistence

import (
	"github.com/device-server/internal/repository/entity"
	"gorm.io/gorm"
)

type VerifyCodeIer interface {
	BatchSave(db *gorm.DB, admins []*entity.VerifyCode) error
	Get(db *gorm.DB, scopes ...func(*gorm.DB) *gorm.DB) (account *entity.VerifyCode, err error)
}
