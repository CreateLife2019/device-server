package filter

import (
	"github.com/device-server/internal/repository/entity"
	"gorm.io/gorm"
)

func Page(page *entity.Page) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page != nil {
			offset := int((page.Page - 1) * page.PageSize)
			return db.Offset(offset).Limit(int(page.PageSize))
		}
		return db
	}
}
