package persistence

import (
	"github.com/device-server/internal/repository/entity"
	"gorm.io/gorm"
)

type GroupIer interface {
	SearchGroup(db *gorm.DB, page *entity.Page, scopes ...func(*gorm.DB) *gorm.DB) (users []*entity.Group, err error)
	Save(db *gorm.DB, in *entity.Group) (out *entity.Group, err error)
	Update(tx *gorm.DB, in *entity.Group, scopes ...func(db *gorm.DB) *gorm.DB) (err error)
	Delete(db *gorm.DB, in *entity.Group) (err error)
	Get(db *gorm.DB, scopes ...func(*gorm.DB) *gorm.DB) (proxy *entity.Group, err error)
}
