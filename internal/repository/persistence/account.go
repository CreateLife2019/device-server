package persistence

import (
	"github.com/device-server/internal/repository/entity"
	"gorm.io/gorm"
)

type AccountIer interface {
	Get(db *gorm.DB, scopes ...func(*gorm.DB) *gorm.DB) (account *entity.Account, err error)
	Save(db *gorm.DB, in *entity.Account) (account *entity.Account, err error)
	Update(tx *gorm.DB, in *entity.Account, scopes ...func(db *gorm.DB) *gorm.DB) (err error)
	SearchAccount(db *gorm.DB, page *entity.Page, scopes ...func(*gorm.DB) *gorm.DB) (accounts []*entity.Account, err error)
	Delete(db *gorm.DB, in *entity.Account) (err error)
}
