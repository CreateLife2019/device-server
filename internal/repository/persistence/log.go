package persistence

import (
	"github.com/device-server/internal/repository/entity"
	"gorm.io/gorm"
)

type LogIer interface {
	Save(db *gorm.DB, in *entity.LoginLog) (account *entity.LoginLog, err error)
	SearchLog(db *gorm.DB, page *entity.Page, scopes ...func(*gorm.DB) *gorm.DB) (logs []*entity.LoginLog, err error)
}
