package filter

import "gorm.io/gorm"

func WithAccountPassword(userName, password string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Where("f_user_name = ? and f_password = ? ", userName, password)
		return db
	}
}
func WithAccount(userName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Where("f_user_name = ? ", userName)
		return db
	}
}
func WithAccountId(accountId int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Where("f_id = ? ", accountId)
		return db
	}
}
