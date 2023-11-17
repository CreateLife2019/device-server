package impl

import (
	"github.com/device-server/internal/repository/entity"
	"github.com/device-server/internal/repository/filter"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type GroupImpl struct {
}

func (g *GroupImpl) SearchGroup(db *gorm.DB, page *entity.Page, scopes ...func(*gorm.DB) *gorm.DB) (groups []*entity.Group, err error) {
	db = db.Model(&entity.Group{}).Scopes(scopes...)
	if page != nil {
		db.Count(&page.Total)
		db = db.Scopes(filter.Page(page))
	}
	err = db.Order("f_created_at desc").Find(&groups).Error
	return
}
func (g *GroupImpl) Save(db *gorm.DB, in *entity.Group) (out *entity.Group, err error) {
	err = db.Model(&entity.Group{}).Create(&in).Error
	return in, err
}
func (g *GroupImpl) BatchSave(db *gorm.DB, in []*entity.Group) (err error) {
	err = db.CreateInBatches(in, 1000).Error
	if err != nil {
		logrus.Errorf("批量创建验证码失败:%s", err.Error())
		return err
	}
	return nil
}
func (g *GroupImpl) Update(tx *gorm.DB, in *entity.Group, scopes ...func(db *gorm.DB) *gorm.DB) (err error) {
	err = tx.Model(&entity.Group{}).Scopes(scopes...).Updates(&in).Error
	return
}
func (g *GroupImpl) Delete(db *gorm.DB, in *entity.Group) (err error) {
	err = db.Model(&entity.Group{}).Unscoped().Where(&in).Delete(&in).Error
	return err
}
func (g *GroupImpl) Get(db *gorm.DB, scopes ...func(*gorm.DB) *gorm.DB) (group *entity.Group, err error) {
	group = &entity.Group{}
	db = db.Model(&entity.Group{}).Scopes(scopes...)
	err = db.First(group).Error
	return
}
