package impl

import (
	"github.com/device-server/internal/repository/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type VerifyCodeImpl struct {
}

func (v *VerifyCodeImpl) BatchSave(db *gorm.DB, admins []*entity.VerifyCode) error {
	err := db.CreateInBatches(admins, 1000).Error
	if err != nil {
		logrus.Errorf("批量创建验证码失败:%s", err.Error())
		return err
	}
	return nil
}

func (v *VerifyCodeImpl) Get(db *gorm.DB, scopes ...func(*gorm.DB) *gorm.DB) (account *entity.VerifyCode, err error) {
	model := entity.VerifyCode{}
	db = db.Model(&entity.VerifyCode{}).Scopes(scopes...)
	err = db.First(&model).Error
	return
}
