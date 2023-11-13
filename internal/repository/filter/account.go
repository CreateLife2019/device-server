package filter

import "gorm.io/gorm"

func WithAccount(userName, password string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Where("f_user_name = ? and f_password = ? ", userName, password)
		return db
	}
}
