package impl

import (
	"github.com/device-server/internal/repository/entity"
	"github.com/device-server/internal/repository/filter"
	"gorm.io/gorm"
)

type ProxyImpl struct {
}

func (l *ProxyImpl) Save(db *gorm.DB, in *entity.Proxy) (out *entity.Proxy, err error) {
	err = db.Model(&entity.Proxy{}).Create(&in).Error
	return in, err
}
func (l *ProxyImpl) Update(tx *gorm.DB, in *entity.Proxy, scopes ...func(db *gorm.DB) *gorm.DB) (err error) {
	err = tx.Model(&entity.Proxy{}).Scopes(scopes...).Updates(&in).Error
	return
}
func (l *ProxyImpl) Delete(db *gorm.DB, in *entity.Proxy) (err error) {
	err = db.Model(&entity.Proxy{}).Unscoped().Where(&in).Delete(&in).Error
	return err
}
func (l *ProxyImpl) SearchProxy(db *gorm.DB, page *entity.Page, scopes ...func(*gorm.DB) *gorm.DB) (proxies []*entity.Proxy, err error) {
	db = db.Model(&entity.Proxy{}).Scopes(scopes...)
	if page != nil {
		db.Count(&page.Total)
		db = db.Scopes(filter.Page(page))
	}
	err = db.Order("f_created_at desc").Find(&proxies).Error
	return
}
func (l *ProxyImpl) Get(db *gorm.DB, scopes ...func(*gorm.DB) *gorm.DB) (proxy *entity.Proxy, err error) {
	proxy = &entity.Proxy{}
	db = db.Model(&entity.Proxy{}).Scopes(scopes...)
	err = db.First(proxy).Error
	return
}
