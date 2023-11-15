package impl

import (
	"github.com/device-server/internal/repository/entity"
	"github.com/device-server/internal/repository/filter"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserIerImpl struct {
}

func (u *UserIerImpl) SearchUser(db *gorm.DB, page *entity.Page, scopes ...func(*gorm.DB) *gorm.DB) (users []*entity.User, err error) {
	db = db.Model(&entity.User{}).Scopes(scopes...)
	if page != nil {
		db.Count(&page.Total)
		db = db.Scopes(filter.Page(page))
	}
	err = db.Order("f_created_at desc").Find(&users).Error
	return
}
func (u *UserIerImpl) SaveOrUpdate(db *gorm.DB, user *entity.User) error {
	return db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "f_phone"}},
		DoUpdates: clause.AssignmentColumns([]string{"f_nick_name", "f_name"}),
	}).Create(user).Error
}
func (u *UserIerImpl) SaveOrUpdateExtend(db *gorm.DB, user *entity.UserExtend) error {
	return db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "f_user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"f_online", "f_client_ip", "f_login_time", "f_app_version", "f_proxy_type", "f_heartbeat@"}),
	}).Create(user).Error
}
func (u *UserIerImpl) Get(db *gorm.DB, scopes ...func(*gorm.DB) *gorm.DB) (account *entity.User, err error) {
	account = &entity.User{}
	db = db.Model(&entity.User{}).Scopes(scopes...)
	err = db.First(account).Error
	return
}
