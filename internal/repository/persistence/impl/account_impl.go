package impl

import (
	"github.com/device-server/internal/repository/entity"
	"gorm.io/gorm"
)

type AccountImpl struct {
}

func (a *AccountImpl) Get(db *gorm.DB, scopes ...func(*gorm.DB) *gorm.DB) (account *entity.Account, err error) {
	model := entity.Account{}
	db = db.Model(&entity.Account{}).Scopes(scopes...)
	err = db.First(&model).Error
	return
}
