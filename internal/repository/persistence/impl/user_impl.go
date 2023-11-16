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
		DoUpdates: clause.AssignmentColumns([]string{"f_online", "f_client_ip", "f_login_time", "f_app_version", "f_proxy_type", "f_heartbeat_time"}),
	}).Create(user).Error
}
func (u *UserIerImpl) Get(db *gorm.DB, scopes ...func(*gorm.DB) *gorm.DB) (account *entity.User, err error) {
	account = &entity.User{}
	db = db.Model(&entity.User{}).Scopes(scopes...)
	err = db.First(account).Error
	return
}
func (u *UserIerImpl) UpdateUserExtend(tx *gorm.DB, in *entity.UserExtend, scopes ...func(db *gorm.DB) *gorm.DB) (err error) {
	err = tx.Model(&entity.UserExtend{}).Scopes(scopes...).Updates(&in).Error
	return
}
func (u *UserIerImpl) UpdateUserConfig(tx *gorm.DB, in *entity.UserConfig, scopes ...func(db *gorm.DB) *gorm.DB) (err error) {
	err = tx.Model(&entity.UserConfig{}).Scopes(scopes...).Updates(&in).Error
	return
}
func (u *UserIerImpl) Update(tx *gorm.DB, in *entity.User, scopes ...func(db *gorm.DB) *gorm.DB) (err error) {
	err = tx.Model(&entity.User{}).Scopes(scopes...).Updates(&in).Error
	return
}
func (u *UserIerImpl) SearchUserExtend(db *gorm.DB, page *entity.Page, scopes ...func(*gorm.DB) *gorm.DB) (users []*entity.UserExtend, err error) {
	db = db.Model(&entity.UserExtend{}).Scopes(scopes...)
	if page != nil {
		db.Count(&page.Total)
		db = db.Scopes(filter.Page(page))
	}
	err = db.Order("f_created_at desc").Find(&users).Error
	return
}
func (u *UserIerImpl) GetUserConfig(db *gorm.DB, scopes ...func(*gorm.DB) *gorm.DB) (account *entity.UserConfig, err error) {
	account = &entity.UserConfig{}
	db = db.Model(&entity.UserConfig{}).Scopes(scopes...)
	err = db.First(account).Error
	return
}
func (u *UserIerImpl) GetOrCreateUserConfig(db *gorm.DB, in *entity.UserConfig, scopes ...func(*gorm.DB) *gorm.DB) (out *entity.UserConfig, err error) {
	out, err = u.GetUserConfig(db, scopes...)
	if err != nil && err == gorm.ErrRecordNotFound {
		in.Id = 0
		if err = db.Create(&in).Error; err != nil {
			return
		}
		out = in
		return
	}
	return
}
func (u *UserIerImpl) SearchUserConfig(db *gorm.DB, page *entity.Page, scopes ...func(*gorm.DB) *gorm.DB) (users []*entity.UserConfig, err error) {
	db = db.Model(&entity.UserConfig{}).Scopes(scopes...)
	if page != nil {
		db.Count(&page.Total)
		db = db.Scopes(filter.Page(page))
	}
	err = db.Order("f_created_at desc").Find(&users).Error
	return
}

func (u *UserIerImpl) OfflineUsers(db *gorm.DB, userIds []int64) error {
	return db.Model(&entity.UserExtend{}).Where("f_id in ?", userIds).Update("f_online", 2).Error
}
