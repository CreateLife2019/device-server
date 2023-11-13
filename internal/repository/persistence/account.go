package persistence

import (
	"github.com/device-server/internal/repository/entity"
	"gorm.io/gorm"
)

type AccountIer interface {
	Get(db *gorm.DB, scopes ...func(*gorm.DB) *gorm.DB) (account *entity.Account, err error)
}
