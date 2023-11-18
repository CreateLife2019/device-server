package impl

import (
	"github.com/device-server/internal/repository/entity"
	"github.com/device-server/internal/repository/filter"
	"gorm.io/gorm"
)

type LogImpl struct {
}

func (l *LogImpl) Save(db *gorm.DB, in *entity.LoginLog) (account *entity.LoginLog, err error) {
	err = db.Model(&entity.LoginLog{}).Create(&in).Error
	return in, err
}
func (l *LogImpl) SearchLog(db *gorm.DB, page *entity.Page, scopes ...func(*gorm.DB) *gorm.DB) (users []*entity.LoginLog, err error) {
	db = db.Model(&entity.LoginLog{}).Scopes(scopes...)
	if page != nil {
		db.Count(&page.Total)
		db = db.Scopes(filter.Page(page))
	}
	err = db.Order("f_created_at desc").Find(&users).Error
	return
}
