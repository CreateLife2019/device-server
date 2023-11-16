package impl

import (
	"github.com/device-server/internal/repository/entity"
	"gorm.io/gorm"
)

type LogImpl struct {
}

func (l *LogImpl) Save(db *gorm.DB, in *entity.LoginLog) (account *entity.LoginLog, err error) {
	err = db.Model(&entity.LoginLog{}).Create(&in).Error
	return in, err
}
