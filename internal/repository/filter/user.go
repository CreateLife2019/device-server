package filter

import "gorm.io/gorm"

func WithPhone(phone string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Where("f_phone = ?  ", phone)
		return db
	}
}
