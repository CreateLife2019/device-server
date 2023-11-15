package impl

import (
	"github.com/device-server/internal/repository/entity"
	"github.com/device-server/internal/repository/filter"
	"gorm.io/gorm"
)

type AccountImpl struct {
}

func (a *AccountImpl) Get(db *gorm.DB, scopes ...func(*gorm.DB) *gorm.DB) (account *entity.Account, err error) {
	account = &entity.Account{}
	db = db.Model(&entity.Account{}).Scopes(scopes...)
	err = db.First(account).Error
	return
}
func (a *AccountImpl) Save(db *gorm.DB, in *entity.Account) (account *entity.Account, err error) {
	err = db.Model(&entity.Account{}).Create(&in).Error
	return in, err
}
func (a *AccountImpl) Update(tx *gorm.DB, in *entity.Account, scopes ...func(db *gorm.DB) *gorm.DB) (err error) {
	err = tx.Model(&entity.Account{}).Scopes(scopes...).Updates(&in).Error
	return
}
func (u *AccountImpl) SearchAccount(db *gorm.DB, page *entity.Page, scopes ...func(*gorm.DB) *gorm.DB) (users []*entity.Account, err error) {
	db = db.Model(&entity.Account{}).Scopes(scopes...)
	if page != nil {
		db.Count(&page.Total)
		db = db.Scopes(filter.Page(page))
	}
	err = db.Order("f_created_at desc").Find(&users).Error
	return
}
func (u *AccountImpl) Delete(db *gorm.DB, in *entity.Account) (err error) {
	err = db.Model(&entity.Account{}).Unscoped().Where(&in).Delete(&in).Error
	return err
}
