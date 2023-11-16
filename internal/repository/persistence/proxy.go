package persistence

import (
	"github.com/device-server/internal/repository/entity"
	"gorm.io/gorm"
)

type ProxyIer interface {
	Save(db *gorm.DB, in *entity.Proxy) (out *entity.Proxy, err error)
	Update(tx *gorm.DB, in *entity.Proxy, scopes ...func(db *gorm.DB) *gorm.DB) (err error)
	Delete(db *gorm.DB, in *entity.Proxy) (err error)
	SearchProxy(db *gorm.DB, page *entity.Page, scopes ...func(*gorm.DB) *gorm.DB) (proxies []*entity.Proxy, err error)
	Get(db *gorm.DB, scopes ...func(*gorm.DB) *gorm.DB) (proxy *entity.Proxy, err error)
}
